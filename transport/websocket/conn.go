package websocket

import (
	"bufio"
	"bytes"
	"context"
	"crypto/rand"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gobwas/ws"
	"github.com/prettylittlelies/cry/fingerprint"
)

const (
	DefaultMaxMessage = 64 << 20
	closeWait         = 5 * time.Second
)

type Opcode = ws.OpCode

const (
	Text   = ws.OpText
	Binary = ws.OpBinary
)

type Config struct {
	Dialer     *fingerprint.Dialer
	Browser    Browser
	Jar        http.CookieJar
	MaxMessage int
}

type Conn struct {
	nc         net.Conn
	br         *bufio.Reader
	maxMessage int

	wmu       sync.Mutex
	closeOnce sync.Once
	closed    chan struct{}
	Response  *http.Response
}

func NewDialer(hello fingerprint.ClientHello) *fingerprint.Dialer {
	d := fingerprint.NewDialer(hello)
	d.ALPN = fingerprint.ALPNHTTP1
	return d
}

func Dial(ctx context.Context, rawURL string, cfg Config) (*Conn, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	addr := u.Host
	if u.Port() == "" {
		addr = net.JoinHostPort(u.Hostname(), "443")
	}
	tc, err := cfg.Dialer.DialTLS(ctx, "tcp", addr)
	if err != nil {
		return nil, err
	}
	if deadline, ok := ctx.Deadline(); ok {
		tc.SetDeadline(deadline)
	}
	key := newKey()
	if err := writeRequest(tc, u, cfg.Browser, cookieHeader(cfg.Jar, u), key); err != nil {
		tc.Close()
		return nil, err
	}
	br := bufio.NewReaderSize(tc, 64<<10)
	resp, err := readResponse(br, u, key)
	if err != nil {
		tc.Close()
		return nil, err
	}
	tc.SetDeadline(time.Time{})
	if cfg.Jar != nil {
		cfg.Jar.SetCookies(httpsURL(u), resp.Cookies())
	}
	limit := cfg.MaxMessage
	if limit <= 0 {
		limit = DefaultMaxMessage
	}
	return &Conn{nc: tc, br: br, maxMessage: limit, closed: make(chan struct{}), Response: resp}, nil
}

func mask() (m [4]byte) {
	rand.Read(m[:])
	return m
}

func (c *Conn) writeFrame(op ws.OpCode, payload []byte) error {
	c.wmu.Lock()
	defer c.wmu.Unlock()
	m := mask()
	masked := bytes.Clone(payload)
	ws.Cipher(masked, m, 0)
	h := ws.Header{Fin: true, OpCode: op, Masked: true, Mask: m, Length: int64(len(masked))}
	var buf bytes.Buffer
	if err := ws.WriteHeader(&buf, h); err != nil {
		return err
	}
	buf.Write(masked)
	_, err := c.nc.Write(buf.Bytes())
	return err
}

func (c *Conn) Write(op Opcode, payload []byte) error {
	select {
	case <-c.closed:
		return ErrClosed
	default:
	}
	return c.writeFrame(op, payload)
}

func (c *Conn) Read() (Opcode, []byte, error) {
	var (
		op  ws.OpCode
		msg []byte
	)
	for {
		h, err := ws.ReadHeader(c.br)
		if err != nil {
			return 0, nil, c.fail(err)
		}
		if h.Masked || h.Rsv != 0 {
			return 0, nil, c.fail(ErrProtocol)
		}
		if h.Length > int64(c.maxMessage) || int64(len(msg))+h.Length > int64(c.maxMessage) {
			return 0, nil, c.fail(ErrMessageTooLarge)
		}
		payload := make([]byte, h.Length)
		if _, err := io.ReadFull(c.br, payload); err != nil {
			return 0, nil, c.fail(err)
		}
		switch h.OpCode {
		case ws.OpPing:
			if err := c.writeFrame(ws.OpPong, payload); err != nil {
				return 0, nil, c.fail(err)
			}
			continue
		case ws.OpPong:
			continue
		case ws.OpClose:
			code, reason := ws.ParseCloseFrameData(payload)
			c.writeFrame(ws.OpClose, payload)
			c.shutdown()
			return 0, nil, &CloseError{Code: int(code), Reason: reason}
		case ws.OpContinuation:
			if op == 0 {
				return 0, nil, c.fail(ErrProtocol)
			}
		default:
			if op != 0 {
				return 0, nil, c.fail(ErrProtocol)
			}
			op = h.OpCode
		}
		msg = append(msg, payload...)
		if h.Fin {
			return op, msg, nil
		}
	}
}

func (c *Conn) Close(code int, reason string) error {
	err := c.writeFrame(ws.OpClose, ws.NewCloseFrameBody(ws.StatusCode(code), reason))
	c.nc.SetReadDeadline(time.Now().Add(closeWait))
	c.shutdown()
	return err
}

func (c *Conn) shutdown() {
	c.closeOnce.Do(func() {
		close(c.closed)
		c.nc.Close()
	})
}

func (c *Conn) fail(err error) error {
	c.shutdown()
	if errors.Is(err, net.ErrClosed) {
		return ErrClosed
	}
	return err
}
