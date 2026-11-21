package main

import (
	"bufio"
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"
)

type Record struct {
	ALPN     string   `json:"alpn"`
	Hello    []byte   `json:"client_hello"`
	Method   string   `json:"method,omitempty"`
	Headers  []string `json:"headers"`
	Priority *struct {
		Weight    uint8  `json:"weight"`
		DependsOn uint32 `json:"depends_on"`
		Exclusive bool   `json:"exclusive"`
	} `json:"priority,omitempty"`
	BodyBytes int `json:"body_bytes"`
}

type recorder struct {
	mu  sync.Mutex
	out *json.Encoder
}

func (r *recorder) write(rec Record) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.out.Encode(rec)
}

type teeConn struct {
	net.Conn
	buf bytes.Buffer
	on  bool
}

func (c *teeConn) Read(p []byte) (int, error) {
	n, err := c.Conn.Read(p)
	if c.on {
		c.buf.Write(p[:n])
	}
	return n, err
}

func firstRecord(b []byte) []byte {
	if len(b) < 5 {
		return b
	}
	n := 5 + (int(b[3])<<8 | int(b[4]))
	return b[:min(n, len(b))]
}

func certificate() tls.Certificate {
	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	tmpl := &x509.Certificate{SerialNumber: big.NewInt(1), NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(24 * time.Hour), DNSNames: []string{"*"}}
	der, _ := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	return tls.Certificate{Certificate: [][]byte{der}, PrivateKey: key}
}

var corsHeaders = []hpack.HeaderField{
	{Name: "access-control-allow-origin", Value: "https://discord.com"},
	{Name: "access-control-allow-methods", Value: "GET, POST, PUT, PATCH, DELETE, OPTIONS"},
	{Name: "access-control-allow-headers", Value: "*"},
	{Name: "access-control-allow-credentials", Value: "true"},
}

func serveH2(conn *tls.Conn, rec *recorder, hello []byte) {
	br := bufio.NewReader(conn)
	preface := make([]byte, len(http2.ClientPreface))
	if _, err := io.ReadFull(br, preface); err != nil {
		return
	}
	fr := http2.NewFramer(conn, br)
	fr.ReadMetaHeaders = hpack.NewDecoder(65536, nil)
	fr.WriteSettings()
	var hbuf bytes.Buffer
	enc := hpack.NewEncoder(&hbuf)
	pending := map[uint32]*Record{}
	respond := func(id uint32, status string) {
		hbuf.Reset()
		enc.WriteField(hpack.HeaderField{Name: ":status", Value: status})
		for _, f := range corsHeaders {
			enc.WriteField(f)
		}
		enc.WriteField(hpack.HeaderField{Name: "content-type", Value: "application/json"})
		fr.WriteHeaders(http2.HeadersFrameParam{StreamID: id, BlockFragment: hbuf.Bytes(), EndHeaders: true})
		fr.WriteData(id, true, []byte("{}"))
	}
	for {
		f, err := fr.ReadFrame()
		if err != nil {
			return
		}
		switch f := f.(type) {
		case *http2.SettingsFrame:
			if !f.IsAck() {
				fr.WriteSettingsAck()
			}
		case *http2.MetaHeadersFrame:
			r := &Record{ALPN: "h2", Hello: hello, Method: f.PseudoValue("method")}
			for _, hf := range f.Fields {
				r.Headers = append(r.Headers, hf.Name+": "+hf.Value)
			}
			if f.HasPriority() {
				p := f.Priority
				r.Priority = &struct {
					Weight    uint8  `json:"weight"`
					DependsOn uint32 `json:"depends_on"`
					Exclusive bool   `json:"exclusive"`
				}{p.Weight, p.StreamDep, p.Exclusive}
			}
			pending[f.StreamID] = r
			if f.StreamEnded() {
				rec.write(*r)
				delete(pending, f.StreamID)
				respond(f.StreamID, map[bool]string{true: "204", false: "200"}[r.Method == http.MethodOptions])
			}
		case *http2.DataFrame:
			r, ok := pending[f.StreamID]
			if !ok {
				continue
			}
			r.BodyBytes += len(f.Data())
			if n := uint32(len(f.Data())); n > 0 {
				fr.WriteWindowUpdate(0, n)
				fr.WriteWindowUpdate(f.StreamID, n)
			}
			if f.StreamEnded() {
				rec.write(*r)
				delete(pending, f.StreamID)
				respond(f.StreamID, "200")
			}
		case *http2.PingFrame:
			if !f.IsAck() {
				fr.WritePing(true, f.Data)
			}
		}
	}
}

func serveHTTP1(conn *tls.Conn, rec *recorder, hello []byte) {
	br := bufio.NewReader(conn)
	line, err := br.ReadString('\n')
	if err != nil {
		return
	}
	r := Record{ALPN: "http/1.1", Hello: hello, Method: strings.Fields(line)[0], Headers: []string{strings.TrimRight(line, "\r\n")}}
	for {
		l, err := br.ReadString('\n')
		if err != nil {
			return
		}
		l = strings.TrimRight(l, "\r\n")
		if l == "" {
			break
		}
		r.Headers = append(r.Headers, l)
	}
	rec.write(r)
	io.WriteString(conn, "HTTP/1.1 400 Bad Request\r\nContent-Length: 0\r\nConnection: close\r\n\r\n")
}

func main() {
	addr := flag.String("addr", "127.0.0.1:8443", "listen address")
	out := flag.String("out", "capture.jsonl", "output file")
	flag.Parse()
	f, err := os.Create(*out)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	rec := &recorder{out: json.NewEncoder(f)}
	ln, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{certificate()}, NextProtos: []string{"h2", "http/1.1"}}
	log.Printf("capturing on %s -> %s", *addr, *out)
	for {
		raw, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			defer raw.Close()
			tee := &teeConn{Conn: raw, on: true}
			conn := tls.Server(tee, cfg)
			if err := conn.Handshake(); err != nil {
				return
			}
			tee.on = false
			hello := firstRecord(tee.buf.Bytes())
			if conn.ConnectionState().NegotiatedProtocol == "h2" {
				serveH2(conn, rec, hello)
				return
			}
			serveHTTP1(conn, rec, hello)
		}()
	}
}
