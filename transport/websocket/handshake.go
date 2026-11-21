package websocket

import (
	"bufio"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const acceptGUID = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"

const offeredExtensions = "permessage-deflate; client_max_window_bits"

type Browser struct {
	UserAgent      string
	Origin         string
	AcceptLanguage string
}

func newKey() string {
	var b [16]byte
	rand.Read(b[:])
	return base64.StdEncoding.EncodeToString(b[:])
}

func acceptFor(key string) string {
	sum := sha1.Sum([]byte(key + acceptGUID))
	return base64.StdEncoding.EncodeToString(sum[:])
}

func cookieHeader(jar http.CookieJar, u *url.URL) string {
	if jar == nil {
		return ""
	}
	cookies := jar.Cookies(httpsURL(u))
	parts := make([]string, len(cookies))
	for i, c := range cookies {
		parts[i] = c.Name + "=" + c.Value
	}
	return strings.Join(parts, "; ")
}

func httpsURL(u *url.URL) *url.URL {
	out := *u
	out.Scheme = "https"
	return &out
}

func writeRequest(w io.Writer, u *url.URL, b Browser, cookie, key string) error {
	var sb strings.Builder
	line := func(k, v string) {
		sb.WriteString(k)
		sb.WriteString(": ")
		sb.WriteString(v)
		sb.WriteString("\r\n")
	}
	sb.WriteString("GET " + u.RequestURI() + " HTTP/1.1\r\n")
	line("Host", u.Host)
	line("Connection", "Upgrade")
	line("Pragma", "no-cache")
	line("Cache-Control", "no-cache")
	line("User-Agent", b.UserAgent)
	line("Upgrade", "websocket")
	line("Origin", b.Origin)
	line("Sec-WebSocket-Version", "13")
	line("Accept-Encoding", "gzip, deflate, br, zstd")
	line("Accept-Language", b.AcceptLanguage)
	if cookie != "" {
		line("Cookie", cookie)
	}
	line("Sec-WebSocket-Key", key)
	line("Sec-WebSocket-Extensions", offeredExtensions)
	sb.WriteString("\r\n")
	_, err := io.WriteString(w, sb.String())
	return err
}

func readResponse(br *bufio.Reader, u *url.URL, key string) (*http.Response, error) {
	req := &http.Request{Method: http.MethodGet, URL: u}
	resp, err := http.ReadResponse(br, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusSwitchingProtocols {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		resp.Body.Close()
		return nil, &HandshakeError{Status: resp.StatusCode, Body: body}
	}
	switch {
	case !strings.EqualFold(resp.Header.Get("Upgrade"), "websocket"):
		return nil, &HandshakeError{Status: resp.StatusCode, Reason: "missing upgrade"}
	case !strings.EqualFold(resp.Header.Get("Connection"), "upgrade"):
		return nil, &HandshakeError{Status: resp.StatusCode, Reason: "missing connection upgrade"}
	case resp.Header.Get("Sec-WebSocket-Accept") != acceptFor(key):
		return nil, &HandshakeError{Status: resp.StatusCode, Reason: "bad accept key"}
	case resp.Header.Get("Sec-WebSocket-Extensions") != "":
		return nil, &HandshakeError{Status: resp.StatusCode, Reason: "unsupported extension accepted: " + resp.Header.Get("Sec-WebSocket-Extensions")}
	}
	return resp, nil
}
