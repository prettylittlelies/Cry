package transport

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/prettylittlelies/cry/fingerprint"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"
)

const (
	defaultWindow   = 65535
	defaultMaxFrame = 16384
)

type stream struct {
	id        uint32
	resp      chan *http.Response
	body      *body
	window    int64
	unacked   uint32
	done      bool
	responded bool
}

type conn struct {
	net     net.Conn
	profile fingerprint.Profile

	wmu    sync.Mutex
	bw     *bufio.Writer
	framer *http2.Framer
	hbuf   bytes.Buffer
	enc    *hpack.Encoder

	mu         sync.Mutex
	cond       *sync.Cond
	streams    map[uint32]*stream
	nextID     uint32
	window     int64
	unacked    uint32
	peerWindow int64
	maxFrame   uint32
	err        error

	recvWindow     uint32
	connRecvWindow uint32
}

func newConn(nc net.Conn, p fingerprint.Profile) (*conn, error) {
	c := &conn{
		net:            nc,
		profile:        p,
		bw:             bufio.NewWriterSize(nc, 32<<10),
		streams:        make(map[uint32]*stream),
		nextID:         1,
		window:         defaultWindow,
		peerWindow:     defaultWindow,
		maxFrame:       defaultMaxFrame,
		recvWindow:     defaultWindow,
		connRecvWindow: defaultWindow + p.H2.ConnectionFlow,
	}
	c.cond = sync.NewCond(&c.mu)
	c.framer = http2.NewFramer(c.bw, bufio.NewReaderSize(nc, 32<<10))
	c.enc = hpack.NewEncoder(&c.hbuf)
	if v, ok := p.H2.Setting(http2.SettingInitialWindowSize); ok {
		c.recvWindow = v
	}
	tableSize := uint32(4096)
	if v, ok := p.H2.Setting(http2.SettingHeaderTableSize); ok {
		tableSize = v
	}
	c.framer.ReadMetaHeaders = hpack.NewDecoder(tableSize, nil)
	if v, ok := p.H2.Setting(http2.SettingMaxHeaderListSize); ok {
		c.framer.MaxHeaderListSize = v
	}
	if err := c.preface(); err != nil {
		return nil, err
	}
	go c.readLoop()
	return c, nil
}

func (c *conn) preface() error {
	c.wmu.Lock()
	defer c.wmu.Unlock()
	if _, err := io.WriteString(c.bw, http2.ClientPreface); err != nil {
		return err
	}
	if err := c.framer.WriteSettings(c.profile.H2.Settings...); err != nil {
		return err
	}
	if inc := c.profile.H2.ConnectionFlow; inc > 0 {
		if err := c.framer.WriteWindowUpdate(0, inc); err != nil {
			return err
		}
	}
	return c.bw.Flush()
}

func (c *conn) usable() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.err == nil && c.nextID < 1<<31-1
}

func (c *conn) roundTrip(req *http.Request) (*http.Response, error) {
	fields := c.fields(req)
	urgency := fingerprint.ParseUrgency(req.Header.Get("Priority"))
	hasBody := req.Body != nil && req.Body != http.NoBody

	s, err := c.open(fields, !hasBody, c.profile.H2.Priority(urgency))
	if err != nil {
		return nil, err
	}
	if hasBody {
		go c.writeBody(s, req.Body)
	}

	select {
	case resp := <-s.resp:
		resp.Request = req
		if err := decode(resp); err != nil {
			resp.Body.Close()
			return nil, err
		}
		return resp, nil
	case <-req.Context().Done():
		c.cancel(s)
		return nil, req.Context().Err()
	}
}

func (c *conn) fields(req *http.Request) []hpack.HeaderField {
	path := req.URL.RequestURI()
	pseudo := map[fingerprint.Pseudo]string{
		fingerprint.Method:    req.Method,
		fingerprint.Authority: req.URL.Host,
		fingerprint.Scheme:    "https",
		fingerprint.Path:      path,
	}
	h := req.Header.Clone()
	if h == nil {
		h = http.Header{}
	}
	if req.ContentLength > 0 && h.Get("Content-Length") == "" {
		h.Set("Content-Length", strconv.FormatInt(req.ContentLength, 10))
	}
	out := make([]hpack.HeaderField, 0, len(h)+4)
	for _, p := range c.profile.H2.PseudoOrder {
		out = append(out, hpack.HeaderField{Name: string(p), Value: pseudo[p]})
	}
	return append(out, c.profile.Headers.Fields(h)...)
}

func (c *conn) open(fields []hpack.HeaderField, end bool, prio http2.PriorityParam) (*stream, error) {
	c.wmu.Lock()
	defer c.wmu.Unlock()
	c.mu.Lock()
	if c.err != nil {
		err := c.err
		c.mu.Unlock()
		return nil, err
	}
	s := &stream{id: c.nextID, resp: make(chan *http.Response, 1), window: c.peerWindow}
	s.body = newBody(func(n int) { c.consumed(s, n) }, func() { c.cancel(s) })
	c.nextID += 2
	c.streams[s.id] = s
	limit := int(c.maxFrame)
	c.mu.Unlock()
	if err := c.writeHeaders(s.id, fields, end, prio, limit); err != nil {
		c.fail(err)
		return nil, err
	}
	return s, nil
}

func (c *conn) writeHeaders(id uint32, fields []hpack.HeaderField, end bool, prio http2.PriorityParam, limit int) error {
	c.hbuf.Reset()
	for _, f := range fields {
		if err := c.enc.WriteField(f); err != nil {
			return err
		}
	}
	block := c.hbuf.Bytes()
	first := block[:min(len(block), limit)]
	block = block[len(first):]
	if err := c.framer.WriteHeaders(http2.HeadersFrameParam{
		StreamID:      id,
		BlockFragment: first,
		EndStream:     end,
		EndHeaders:    len(block) == 0,
		Priority:      prio,
	}); err != nil {
		return err
	}
	for len(block) > 0 {
		chunk := block[:min(len(block), limit)]
		block = block[len(chunk):]
		if err := c.framer.WriteContinuation(id, len(block) == 0, chunk); err != nil {
			return err
		}
	}
	return c.bw.Flush()
}

func (c *conn) writeBody(s *stream, r io.ReadCloser) {
	defer r.Close()
	buf := make([]byte, defaultMaxFrame)
	for {
		n, rerr := r.Read(buf)
		data := buf[:n]
		for len(data) > 0 {
			k, err := c.reserve(s, len(data))
			if err != nil {
				return
			}
			if err := c.writeData(s.id, data[:k], false); err != nil {
				c.fail(err)
				return
			}
			data = data[k:]
		}
		if rerr == io.EOF {
			if err := c.writeData(s.id, nil, true); err != nil {
				c.fail(err)
			}
			return
		}
		if rerr != nil {
			c.cancel(s)
			return
		}
	}
}

func (c *conn) reserve(s *stream, want int) (int, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for c.err == nil && !s.done && (c.window <= 0 || s.window <= 0) {
		c.cond.Wait()
	}
	if c.err != nil {
		return 0, c.err
	}
	if s.done {
		return 0, ErrConnClosed
	}
	n := int64(min(want, int(c.maxFrame)))
	n = min(n, c.window, s.window)
	c.window -= n
	s.window -= n
	return int(n), nil
}

func (c *conn) writeData(id uint32, p []byte, end bool) error {
	c.wmu.Lock()
	defer c.wmu.Unlock()
	if err := c.framer.WriteData(id, end, p); err != nil {
		return err
	}
	return c.bw.Flush()
}

func (c *conn) writeControl(fn func(*http2.Framer) error) {
	c.wmu.Lock()
	defer c.wmu.Unlock()
	if err := fn(c.framer); err == nil {
		c.bw.Flush()
	}
}

func (c *conn) consumed(s *stream, n int) {
	c.mu.Lock()
	c.unacked += uint32(n)
	streamInc, connInc := uint32(0), uint32(0)
	if s != nil {
		s.unacked += uint32(n)
		if !s.done && s.unacked >= c.recvWindow/2 {
			streamInc, s.unacked = s.unacked, 0
		}
	}
	if c.unacked >= c.connRecvWindow/2 {
		connInc, c.unacked = c.unacked, 0
	}
	c.mu.Unlock()
	if streamInc == 0 && connInc == 0 {
		return
	}
	c.writeControl(func(f *http2.Framer) error {
		if connInc > 0 {
			if err := f.WriteWindowUpdate(0, connInc); err != nil {
				return err
			}
		}
		if streamInc > 0 {
			return f.WriteWindowUpdate(s.id, streamInc)
		}
		return nil
	})
}

func (c *conn) cancel(s *stream) {
	c.mu.Lock()
	active := !s.done
	c.closeStream(s, context.Canceled)
	c.mu.Unlock()
	if active {
		c.writeControl(func(f *http2.Framer) error { return f.WriteRSTStream(s.id, http2.ErrCodeCancel) })
	}
}

func (c *conn) closeStream(s *stream, err error) {
	if s.done {
		return
	}
	s.done = true
	delete(c.streams, s.id)
	s.body.fail(err)
	c.cond.Broadcast()
}

func (c *conn) fail(err error) {
	c.mu.Lock()
	if c.err == nil {
		c.err = err
	}
	for _, s := range c.streams {
		c.closeStream(s, err)
	}
	c.mu.Unlock()
	c.net.Close()
}

func (c *conn) close() error {
	c.fail(ErrConnClosed)
	return nil
}
