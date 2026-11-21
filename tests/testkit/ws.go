package testkit

import (
	"bufio"
	"crypto/sha1"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"io"
	"net"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gobwas/ws"
)

type WSPeer struct {
	Conn    net.Conn
	Request []string
}

func (p *WSPeer) Send(f ws.Frame) error {
	return ws.WriteFrame(p.Conn, f)
}

func (p *WSPeer) Receive() (ws.Frame, error) {
	f, err := ws.ReadFrame(p.Conn)
	if err != nil {
		return f, err
	}
	if f.Header.Masked {
		ws.Cipher(f.Payload, f.Header.Mask, 0)
	}
	return f, nil
}

type WSServer struct {
	URL   string
	Pool  *x509.CertPool
	Peers chan *WSPeer
	Extra []string
}

func NewWSServer(t testing.TB) *WSServer {
	t.Helper()
	cert := httptest.NewUnstartedServer(nil)
	cert.StartTLS()
	tlsCert := cert.TLS.Certificates[0]
	pool := x509.NewCertPool()
	pool.AddCert(cert.Certificate())
	cert.Close()
	ln, err := tls.Listen("tcp", "127.0.0.1:0", &tls.Config{Certificates: []tls.Certificate{tlsCert}, NextProtos: []string{"http/1.1"}})
	if err != nil {
		t.Fatal(err)
	}
	s := &WSServer{URL: "wss://127.0.0.1:" + strings.Split(ln.Addr().String(), ":")[1] + "/", Pool: pool, Peers: make(chan *WSPeer, 8)}
	t.Cleanup(func() { ln.Close() })
	go func() {
		for {
			c, err := ln.Accept()
			if err != nil {
				return
			}
			go s.upgrade(c)
		}
	}()
	return s
}

func (s *WSServer) upgrade(c net.Conn) {
	br := bufio.NewReader(c)
	var lines []string
	key := ""
	for {
		l, err := br.ReadString('\n')
		if err != nil {
			c.Close()
			return
		}
		l = strings.TrimRight(l, "\r\n")
		if l == "" {
			break
		}
		lines = append(lines, l)
		if k, v, ok := strings.Cut(l, ": "); ok && strings.EqualFold(k, "Sec-WebSocket-Key") {
			key = v
		}
	}
	sum := sha1.Sum([]byte(key + "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"))
	resp := "HTTP/1.1 101 Switching Protocols\r\nUpgrade: websocket\r\nConnection: upgrade\r\nSec-WebSocket-Accept: " + base64.StdEncoding.EncodeToString(sum[:]) + "\r\n"
	for _, e := range s.Extra {
		resp += e + "\r\n"
	}
	io.WriteString(c, resp+"\r\n")
	s.Peers <- &WSPeer{Conn: c, Request: lines}
}
