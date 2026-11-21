package fetch

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var ErrRangeUnsupported = errors.New("fetch: server does not honor range requests")

type Ranged struct {
	ctx    context.Context
	client Client
	url    string
	size   int64
}

func OpenRanged(ctx context.Context, c Client, url string) (*Ranged, error) {
	resp, err := c.Get(ctx, url, byteRange(0, 0))
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusPartialContent {
		return nil, ErrRangeUnsupported
	}
	cr := resp.Header.Get("Content-Range")
	total, err := strconv.ParseInt(cr[strings.LastIndexByte(cr, '/')+1:], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("fetch: content-range %q: %w", cr, err)
	}
	return &Ranged{ctx: ctx, client: c, url: url, size: total}, nil
}

func (r *Ranged) Size() int64 {
	return r.size
}

func (r *Ranged) ReadAt(p []byte, off int64) (int, error) {
	if off >= r.size {
		return 0, io.EOF
	}
	end := min(off+int64(len(p)), r.size) - 1
	resp, err := r.client.Get(r.ctx, r.url, byteRange(off, end))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusPartialContent {
		return 0, ErrRangeUnsupported
	}
	n, err := io.ReadFull(resp.Body, p[:end-off+1])
	if err == nil && n < len(p) {
		err = io.EOF
	}
	return n, err
}

func byteRange(from, to int64) http.Header {
	return http.Header{"Range": {fmt.Sprintf("bytes=%d-%d", from, to)}}
}
