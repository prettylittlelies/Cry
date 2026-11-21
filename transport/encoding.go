package transport

import (
	"compress/flate"
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/andybalholm/brotli"
	"github.com/klauspost/compress/zstd"
)

type decoded struct {
	io.Reader
	closers []io.Closer
}

func (d *decoded) Close() error {
	var first error
	for _, c := range d.closers {
		if err := c.Close(); err != nil && first == nil {
			first = err
		}
	}
	return first
}

func decode(resp *http.Response) error {
	enc := strings.ToLower(strings.TrimSpace(resp.Header.Get("Content-Encoding")))
	var r io.Reader
	closers := []io.Closer{resp.Body}
	switch enc {
	case "":
		return nil
	case "br":
		r = brotli.NewReader(resp.Body)
	case "gzip":
		gz, err := gzip.NewReader(resp.Body)
		if err != nil {
			return err
		}
		r, closers = gz, append(closers, gz)
	case "deflate":
		fl := flate.NewReader(resp.Body)
		r, closers = fl, append(closers, fl)
	case "zstd":
		zd, err := zstd.NewReader(resp.Body)
		if err != nil {
			return err
		}
		r, closers = zd, append(closers, zd.IOReadCloser())
	default:
		return nil
	}
	resp.Body = &decoded{Reader: r, closers: closers}
	resp.Header.Del("Content-Encoding")
	resp.Header.Del("Content-Length")
	resp.ContentLength = -1
	resp.Uncompressed = true
	return nil
}
