package fingerprint

import (
	"context"
	"net"
	"time"

	utls "github.com/refraction-networking/utls"
)

type ContextDialer interface {
	DialContext(ctx context.Context, network, addr string) (net.Conn, error)
}

type Dialer struct {
	Hello    ClientHello
	Sessions utls.ClientSessionCache
	Net      ContextDialer
}

func NewDialer(hello ClientHello) *Dialer {
	return &Dialer{
		Hello:    hello,
		Sessions: utls.NewLRUClientSessionCache(64),
		Net:      &net.Dialer{Timeout: 15 * time.Second, KeepAlive: 30 * time.Second},
	}
}

func (d *Dialer) DialTLS(ctx context.Context, network, addr string) (*utls.UConn, error) {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}
	raw, err := d.Net.DialContext(ctx, network, addr)
	if err != nil {
		return nil, err
	}
	conn := utls.UClient(raw, &utls.Config{
		ServerName:         host,
		ClientSessionCache: d.Sessions,
		NextProtos:         []string{"h2", "http/1.1"},
		OmitEmptyPsk:       true,
	}, utls.HelloCustom)
	if err := conn.ApplyPreset(d.Hello()); err != nil {
		raw.Close()
		return nil, err
	}
	if err := conn.HandshakeContext(ctx); err != nil {
		raw.Close()
		return nil, err
	}
	return conn, nil
}
