package transport

import (
	"io"
	"net/http"
	"strconv"

	"golang.org/x/net/http2"
)

func (c *conn) readLoop() {
	for {
		f, err := c.framer.ReadFrame()
		if err != nil {
			c.fail(err)
			return
		}
		switch f := f.(type) {
		case *http2.MetaHeadersFrame:
			c.onHeaders(f)
		case *http2.DataFrame:
			c.onData(f)
		case *http2.SettingsFrame:
			c.onSettings(f)
		case *http2.WindowUpdateFrame:
			c.onWindowUpdate(f)
		case *http2.PingFrame:
			if !f.IsAck() {
				data := f.Data
				c.writeControl(func(w *http2.Framer) error { return w.WritePing(true, data) })
			}
		case *http2.RSTStreamFrame:
			c.withStream(f.StreamID, func(s *stream) {
				c.closeStream(s, &StreamError{StreamID: f.StreamID, Code: f.ErrCode})
			})
		case *http2.GoAwayFrame:
			c.onGoAway(f)
		}
	}
}

func (c *conn) withStream(id uint32, fn func(*stream)) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if s, ok := c.streams[id]; ok {
		fn(s)
	}
}

func (c *conn) onHeaders(f *http2.MetaHeadersFrame) {
	c.withStream(f.StreamID, func(s *stream) {
		if s.responded {
			if f.StreamEnded() {
				c.finish(s)
			}
			return
		}
		status, err := strconv.Atoi(f.PseudoValue("status"))
		if err != nil {
			c.closeStream(s, ErrMissingStatus)
			return
		}
		if status >= 100 && status < 200 {
			return
		}
		h := make(http.Header, len(f.RegularFields()))
		for _, hf := range f.RegularFields() {
			h.Add(hf.Name, hf.Value)
		}
		cl := int64(-1)
		if v := h.Get("Content-Length"); v != "" {
			if n, err := strconv.ParseInt(v, 10, 64); err == nil {
				cl = n
			}
		}
		resp := &http.Response{
			Status:        strconv.Itoa(status) + " " + http.StatusText(status),
			StatusCode:    status,
			Proto:         "HTTP/2.0",
			ProtoMajor:    2,
			Header:        h,
			Body:          s.body,
			ContentLength: cl,
		}
		s.responded = true
		s.resp <- resp
		if f.StreamEnded() {
			c.finish(s)
		}
	})
}

func (c *conn) finish(s *stream) {
	s.body.fail(io.EOF)
	s.done = true
	delete(c.streams, s.id)
	c.cond.Broadcast()
}

func (c *conn) onData(f *http2.DataFrame) {
	size := f.Header().Length
	c.mu.Lock()
	s, ok := c.streams[f.StreamID]
	c.mu.Unlock()
	if !ok {
		if size > 0 {
			c.consumed(nil, int(size))
		}
		return
	}
	if pad := size - uint32(len(f.Data())); pad > 0 {
		c.consumed(s, int(pad))
	}
	if len(f.Data()) > 0 {
		s.body.write(append([]byte(nil), f.Data()...))
	}
	if f.StreamEnded() {
		c.mu.Lock()
		c.finish(s)
		c.mu.Unlock()
	}
}

func (c *conn) onSettings(f *http2.SettingsFrame) {
	if f.IsAck() {
		return
	}
	var table *uint32
	c.mu.Lock()
	f.ForeachSetting(func(st http2.Setting) error {
		switch st.ID {
		case http2.SettingInitialWindowSize:
			delta := int64(st.Val) - c.peerWindow
			c.peerWindow = int64(st.Val)
			for _, s := range c.streams {
				s.window += delta
			}
		case http2.SettingMaxFrameSize:
			c.maxFrame = st.Val
		case http2.SettingHeaderTableSize:
			table = &st.Val
		}
		return nil
	})
	c.cond.Broadcast()
	c.mu.Unlock()
	c.writeControl(func(w *http2.Framer) error {
		if table != nil {
			c.enc.SetMaxDynamicTableSizeLimit(*table)
		}
		return w.WriteSettingsAck()
	})
}

func (c *conn) onWindowUpdate(f *http2.WindowUpdateFrame) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if f.StreamID == 0 {
		c.window += int64(f.Increment)
	} else if s, ok := c.streams[f.StreamID]; ok {
		s.window += int64(f.Increment)
	}
	c.cond.Broadcast()
}

func (c *conn) onGoAway(f *http2.GoAwayFrame) {
	err := &GoAwayError{LastStreamID: f.LastStreamID, Code: f.ErrCode, Debug: string(f.DebugData())}
	c.mu.Lock()
	if c.err == nil {
		c.err = err
	}
	for id, s := range c.streams {
		if id > f.LastStreamID {
			c.closeStream(s, err)
		}
	}
	c.mu.Unlock()
}
