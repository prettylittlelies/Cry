package testkit

import (
	"crypto/x509"
	"net"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

type H2Server struct {
	*httptest.Server
	Pool  *x509.CertPool
	conns atomic.Int64
}

func NewH2Server(t testing.TB, h http.Handler) *H2Server {
	t.Helper()
	s := &H2Server{Server: httptest.NewUnstartedServer(h)}
	s.EnableHTTP2 = true
	s.Config.ConnState = func(_ net.Conn, st http.ConnState) {
		if st == http.StateNew {
			s.conns.Add(1)
		}
	}
	s.StartTLS()
	t.Cleanup(s.Close)
	s.Pool = x509.NewCertPool()
	s.Pool.AddCert(s.Certificate())
	return s
}

func (s *H2Server) Connections() int64 {
	return s.conns.Load()
}
