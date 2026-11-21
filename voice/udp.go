package voice

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"
)

const (
	discoveryRequest  = 0x1
	discoveryResponse = 0x2
	discoveryLength   = 70
	discoveryTimeout  = 5 * time.Second
)

var ErrDiscovery = errors.New("voice: ip discovery failed")

type udpLink struct {
	conn *net.UDPConn
}

func dialUDP(ctx context.Context, ip string, port int) (*udpLink, error) {
	var d net.Dialer
	c, err := d.DialContext(ctx, "udp", net.JoinHostPort(ip, strconv.Itoa(port)))
	if err != nil {
		return nil, err
	}
	return &udpLink{conn: c.(*net.UDPConn)}, nil
}

func (u *udpLink) discover(ssrc uint32) (string, int, error) {
	req := make([]byte, 4+discoveryLength)
	binary.BigEndian.PutUint16(req[0:], discoveryRequest)
	binary.BigEndian.PutUint16(req[2:], discoveryLength)
	binary.BigEndian.PutUint32(req[4:], ssrc)
	if _, err := u.conn.Write(req); err != nil {
		return "", 0, err
	}
	u.conn.SetReadDeadline(time.Now().Add(discoveryTimeout))
	defer u.conn.SetReadDeadline(time.Time{})
	resp := make([]byte, 128)
	for {
		n, err := u.conn.Read(resp)
		if err != nil {
			return "", 0, fmt.Errorf("%w: %w", ErrDiscovery, err)
		}
		if n < 4+discoveryLength || binary.BigEndian.Uint16(resp) != discoveryResponse {
			continue
		}
		addr := resp[8 : 8+64]
		if i := bytes.IndexByte(addr, 0); i >= 0 {
			addr = addr[:i]
		}
		return string(addr), int(binary.BigEndian.Uint16(resp[72:74])), nil
	}
}

func (u *udpLink) write(b []byte) error {
	_, err := u.conn.Write(b)
	return err
}

func (u *udpLink) read(buf []byte) (int, error) { return u.conn.Read(buf) }

func (u *udpLink) close() error { return u.conn.Close() }
