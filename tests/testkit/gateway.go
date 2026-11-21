package testkit

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/gobwas/ws"
	"github.com/klauspost/compress/zstd"
	"github.com/prettylittlelies/cry/gateway"
	"github.com/prettylittlelies/cry/internal/etf"
)

type GatewayPeer struct {
	t   testing.TB
	ws  *WSPeer
	buf bytes.Buffer
	zw  *zstd.Encoder
}

type Packet struct {
	Op   int             `json:"op"`
	D    json.RawMessage `json:"d"`
	Raw  json.RawMessage `json:"-"`
	Keys []string        `json:"-"`
}

func AcceptGatewayPeer(t testing.TB, srv *WSServer) *GatewayPeer {
	t.Helper()
	select {
	case p := <-srv.Peers:
		g := &GatewayPeer{t: t, ws: p}
		g.zw, _ = zstd.NewWriter(&g.buf)
		return g
	case <-time.After(5 * time.Second):
		t.Fatal("client never connected")
	}
	return nil
}

func (p *GatewayPeer) Request() []string {
	return p.ws.Request
}

func (p *GatewayPeer) Send(v any) {
	p.t.Helper()
	b, err := etf.Marshal(v)
	if err != nil {
		p.t.Fatal(err)
	}
	p.zw.Write(b)
	p.zw.Flush()
	msg := bytes.Clone(p.buf.Bytes())
	p.buf.Reset()
	if err := p.ws.Send(ws.NewBinaryFrame(msg)); err != nil {
		p.t.Fatal(err)
	}
}

func (p *GatewayPeer) Hello(interval int) {
	p.Send(map[string]any{"op": 10, "d": map[string]any{"heartbeat_interval": interval}})
}

func (p *GatewayPeer) Dispatch(event string, seq int, d any) {
	p.Send(map[string]any{"op": 0, "s": seq, "t": event, "d": d})
}

func KeysOf(t testing.TB, raw json.RawMessage) []string {
	t.Helper()
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.Token()
	var out []string
	for dec.More() {
		k, _ := dec.Token()
		out = append(out, k.(string))
		var skip json.RawMessage
		dec.Decode(&skip)
	}
	return out
}

func (p *GatewayPeer) Receive() Packet {
	p.t.Helper()
	p.ws.Conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	f, err := p.ws.Receive()
	if err != nil {
		p.t.Fatal(err)
	}
	if f.Header.OpCode != ws.OpBinary {
		p.t.Fatalf("expected binary frame, got %v", f.Header.OpCode)
	}
	raw, err := etf.ToJSON(f.Payload)
	if err != nil {
		p.t.Fatal(err)
	}
	var pk Packet
	json.Unmarshal(raw, &pk)
	pk.Raw, pk.Keys = raw, KeysOf(p.t, raw)
	return pk
}

func (p *GatewayPeer) Expect(op gateway.Opcode) Packet {
	p.t.Helper()
	for {
		pk := p.Receive()
		switch gateway.Opcode(pk.Op) {
		case op:
			return pk
		case gateway.OpQoSHeartbeat, gateway.OpUpdateTimeSpentSessionID:
			continue
		}
		p.t.Fatalf("expected op %d, got %s", op, pk.Raw)
	}
}

func (p *GatewayPeer) Close(code int) {
	p.ws.Send(ws.NewCloseFrame(ws.NewCloseFrameBody(ws.StatusCode(code), "")))
	p.ws.Conn.Close()
}
