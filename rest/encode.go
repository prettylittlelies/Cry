package rest

import (
	"errors"
	"io"
	"net/url"
	"strconv"
	"strings"
	"syscall"

	"github.com/prettylittlelies/cry/transport"
)

type query struct {
	values url.Values
}

func newQuery(r *Request) *query {
	v := url.Values{}
	for k, vs := range r.Query {
		v[k] = append([]string(nil), vs...)
	}
	return &query{values: v}
}

func (q *query) bump() {
	n, _ := strconv.Atoi(q.values.Get("attempts"))
	q.values.Set("attempts", strconv.Itoa(n+1))
}

func (q *query) encode() string {
	if len(q.values) == 0 {
		return ""
	}
	return "?" + q.values.Encode()
}

const unreserved = "-_.!~*'()"

func encodeURIComponent(s string) string {
	var b strings.Builder
	const hex = "0123456789ABCDEF"
	for i := 0; i < len(s); i++ {
		c := s[i]
		if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9' || strings.IndexByte(unreserved, c) >= 0 {
			b.WriteByte(c)
			continue
		}
		b.WriteByte('%')
		b.WriteByte(hex[c>>4])
		b.WriteByte(hex[c&15])
	}
	return b.String()
}

func retryableNetErr(err error) bool {
	var goAway *transport.GoAwayError
	var stream *transport.StreamError
	return errors.Is(err, io.ErrUnexpectedEOF) || errors.Is(err, io.EOF) ||
		errors.Is(err, syscall.ECONNRESET) || errors.Is(err, transport.ErrConnClosed) ||
		errors.As(err, &goAway) || errors.As(err, &stream)
}
