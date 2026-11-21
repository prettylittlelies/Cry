package testkit

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type Server struct {
	*httptest.Server
	mux *http.ServeMux
}

func NewServer(t testing.TB) *Server {
	t.Helper()
	mux := http.NewServeMux()
	s := &Server{Server: httptest.NewServer(mux), mux: mux}
	t.Cleanup(s.Close)
	return s
}

func (s *Server) Serve(pattern, contentType string, body []byte) {
	s.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		http.ServeContent(w, r, "", time.Time{}, bytes.NewReader(body))
	})
}

func (s *Server) Handle(pattern string, h http.HandlerFunc) {
	s.mux.HandleFunc(pattern, h)
}
