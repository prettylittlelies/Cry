package voice_test

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"encoding/json"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/gobwas/ws"
	"github.com/prettylittlelies/cry"
	"github.com/prettylittlelies/cry/build"
	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/mls"
	"github.com/prettylittlelies/cry/snowflake"
	"github.com/prettylittlelies/cry/tests/testkit"
	"github.com/prettylittlelies/cry/voice"
	"github.com/prettylittlelies/cry/voice/dave"
)

type voicePeer struct {
	t    *testing.T
	peer *testkit.WSPeer
	seq  uint16
}

func (p *voicePeer) sendJSON(op int, d any) {
	p.t.Helper()
	raw, _ := json.Marshal(map[string]any{"op": op, "d": d})
	if err := p.peer.Send(ws.NewTextFrame(raw)); err != nil {
		p.t.Fatal(err)
	}
}

func (p *voicePeer) sendBinary(op dave.Opcode, payload []byte) {
	p.t.Helper()
	p.seq++
	frame := append(binary.BigEndian.AppendUint16(nil, p.seq), byte(op))
	if err := p.peer.Send(ws.NewBinaryFrame(append(frame, payload...))); err != nil {
		p.t.Fatal(err)
	}
}

type inbound struct {
	op     int
	d      json.RawMessage
	binary []byte
}

func (p *voicePeer) next(want int) inbound {
	p.t.Helper()
	for {
		f, err := p.peer.Receive()
		if err != nil {
			p.t.Fatalf("waiting for op %d: %v", want, err)
		}
		var in inbound
		if f.Header.OpCode == ws.OpBinary {
			in = inbound{op: int(f.Payload[0]), binary: f.Payload[1:]}
		} else {
			var pk struct {
				Op int             `json:"op"`
				D  json.RawMessage `json:"d"`
			}
			json.Unmarshal(f.Payload, &pk)
			in = inbound{op: pk.Op, d: pk.D}
		}
		if in.op == want {
			return in
		}
		if in.op != int(voice.OpHeartbeat) && in.op != int(voice.OpSpeaking) {
			p.t.Fatalf("expected op %d, got %d: %s", want, in.op, in.d)
		}
	}
}

func voiceClient(t *testing.T, srv *testkit.WSServer, self snowflake.ID) *cry.Client {
	t.Helper()
	v := func(s string) build.Version { out, _ := build.ParseVersion(s); return out }
	snap := &build.Snapshot{Releases: map[build.Channel]build.Release{build.Stable: {Channel: build.Stable, Number: 1, MacOS: build.MacOS{Host: v("0.0.421"), Electron: v("44.10.0"), Chrome: v("152.0.7977.169")}}}}
	c, err := cry.New(context.Background(), "token", cry.WithSnapshot(snap))
	if err != nil {
		t.Fatal(err)
	}
	c.Gateway().Dialer().RootCAs = srv.Pool
	ready := new(discord.ReadyEvent)
	ready.User.ID = self
	c.State().Apply(ready)
	return c
}

func TestVoiceConnectionWithDAVE(t *testing.T) {
	self, remote := snowflake.ID(111111111111111111), snowflake.ID(222222222222222222)
	guild, ch := snowflake.ID(500000000000000005), channel
	srv := testkit.NewWSServer(t)
	c := voiceClient(t, srv, self)

	udp, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		t.Fatal(err)
	}
	defer udp.Close()
	clientAddr := make(chan *net.UDPAddr, 1)
	media := make(chan []byte, 64)
	go func() {
		buf := make([]byte, 2048)
		for {
			n, addr, err := udp.ReadFromUDP(buf)
			if err != nil {
				return
			}
			pkt := append([]byte(nil), buf[:n]...)
			if n == 74 && binary.BigEndian.Uint16(pkt) == 1 {
				resp := make([]byte, 74)
				binary.BigEndian.PutUint16(resp, 2)
				binary.BigEndian.PutUint16(resp[2:], 70)
				copy(resp[4:8], pkt[4:8])
				copy(resp[8:], "127.0.0.1")
				binary.BigEndian.PutUint16(resp[72:], uint16(addr.Port))
				udp.WriteToUDP(resp, addr)
				clientAddr <- addr
				continue
			}
			if pkt[0] == 0x80 {
				media <- pkt
			}
		}
	}()

	received := make(chan []byte, 8)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	joined := make(chan error, 1)
	var conn *voice.Conn
	go func() {
		var err error
		conn, err = voice.Join(ctx, c, guild, ch, voice.Options{OnOpus: func(user snowflake.ID, frame []byte) {
			if user == remote {
				received <- frame
			}
		}})
		joined <- err
	}()
	endpoint := strings.TrimPrefix(srv.URL, "wss://")
	stop := make(chan struct{})
	go func() {
		for range 40 {
			select {
			case <-stop:
				return
			case <-time.After(50 * time.Millisecond):
			}
			c.Dispatch(ctx, &discord.VoiceStateUpdateEvent{UserID: self, ChannelID: &ch, SessionID: "voice-session"})
			c.Dispatch(ctx, &discord.VoiceServerUpdateEvent{Token: "voice-token", GuildID: &guild, Endpoint: &endpoint})
		}
	}()

	p := &voicePeer{t: t, peer: <-srv.Peers}
	id := p.next(int(voice.OpIdentify))
	var ident struct {
		ServerID  string `json:"server_id"`
		SessionID string `json:"session_id"`
		Token     string `json:"token"`
		MaxDAVE   int    `json:"max_dave_protocol_version"`
	}
	json.Unmarshal(id.d, &ident)
	if ident.ServerID != guild.String() || ident.SessionID != "voice-session" || ident.Token != "voice-token" || ident.MaxDAVE != 1 {
		t.Fatalf("identify %s", id.d)
	}
	p.sendJSON(int(voice.OpHello), map[string]any{"v": 8, "heartbeat_interval": 5000})
	port := udp.LocalAddr().(*net.UDPAddr).Port
	p.sendJSON(int(voice.OpReady), map[string]any{"ssrc": 42, "ip": "127.0.0.1", "port": port, "modes": []string{voice.ModeAES256GCM, voice.ModeXChaCha20Poly1305}})
	sel := p.next(int(voice.OpSelectProtocol))
	var selected struct {
		Data struct {
			Mode string `json:"mode"`
			Port int    `json:"port"`
		} `json:"data"`
	}
	json.Unmarshal(sel.d, &selected)
	addr := <-clientAddr
	if selected.Data.Mode != voice.ModeAES256GCM || selected.Data.Port != addr.Port {
		t.Fatalf("select protocol %s", sel.d)
	}

	g := newGateway(t)
	g.connect(remote)
	key := bytes.Repeat([]byte{9}, 32)
	keyInts := make([]int, len(key))
	for i, b := range key {
		keyInts[i] = int(b)
	}
	p.sendJSON(int(voice.OpClientsConnect), map[string]any{"user_ids": []string{remote.String()}})
	p.sendJSON(int(voice.OpSessionDescription), map[string]any{"mode": voice.ModeAES256GCM, "secret_key": keyInts, "dave_protocol_version": 1, "audio_codec": "opus"})
	p.sendBinary(dave.OpExternalSender, mls.Marshal(&g.sender))
	err = <-joined
	close(stop)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close(context.Background())

	kpMsg := p.next(int(dave.OpKeyPackage))
	clientKP := keyPackageOf(t, kpMsg.binary)
	if u, _ := dave.UserOf(clientKP.LeafNode.Credential); u != self {
		t.Fatal("client key package credential")
	}
	g.members[self] = nil
	p.sendBinary(dave.OpProposals, proposalsPayload(g.proposal(mls.Proposal{Type: mls.ProposalAdd, Add: g.packages[remote]})))
	cw := p.next(int(dave.OpCommitWelcome))
	commit, welcome := split(t, cw.binary)
	if err := g.members[remote].ProcessWelcome(mls.Marshal(welcome), g.recognized); err != nil {
		t.Fatal(err)
	}
	p.sendBinary(dave.OpAnnounceCommit, append([]byte{0, 7}, mls.Marshal(commit)...))
	ready := p.next(int(dave.OpTransitionReady))
	if !strings.Contains(string(ready.d), `"transition_id":7`) {
		t.Fatalf("ready %s", ready.d)
	}
	p.sendJSON(int(dave.OpExecuteTransition), map[string]any{"transition_id": 7})
	deadline := time.Now().Add(2 * time.Second)
	for !conn.Encrypted() && time.Now().Before(deadline) {
		time.Sleep(10 * time.Millisecond)
	}
	if !conn.Encrypted() || !bytes.Equal(conn.EpochAuthenticator(), g.members[remote].EpochAuthenticator()) {
		t.Fatal("client did not activate e2ee on the shared group")
	}

	frame := []byte("twenty milliseconds of opus")
	if err := conn.WriteOpus(frame); err != nil {
		t.Fatal(err)
	}
	pkt := <-media
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, 12)
	copy(nonce, pkt[len(pkt)-4:])
	opened, err := gcm.Open(nil, nonce, pkt[12:len(pkt)-4], pkt[:12])
	if err != nil {
		t.Fatalf("transport: %v", err)
	}
	if binary.BigEndian.Uint32(pkt[8:12]) != 42 || pkt[1] != 120 {
		t.Fatalf("rtp header % x", pkt[:12])
	}
	r, _ := g.members[remote].Ratchet(self)
	var dec dave.Decryptor
	dec.SetRatchet(r)
	if plain, err := dec.Decrypt(opened); err != nil || !bytes.Equal(plain, frame) {
		t.Fatalf("dave: %v", err)
	}

	p.sendJSON(int(voice.OpSpeaking), map[string]any{"user_id": remote.String(), "ssrc": 77, "speaking": 1})
	time.Sleep(50 * time.Millisecond)
	own, _ := g.members[remote].Ratchet(remote)
	var enc dave.Encryptor
	enc.SetRatchet(own)
	protected, _ := enc.Encrypt([]byte("hello from the other side"))
	header := []byte{0x80, 120, 0, 1, 0, 0, 3, 192, 0, 0, 0, 77}
	sealed := append(append([]byte(nil), header...), gcm.Seal(nil, make([]byte, 12), protected, header)...)
	udp.WriteToUDP(append(sealed, 0, 0, 0, 0), addr)
	select {
	case got := <-received:
		if string(got) != "hello from the other side" {
			t.Fatalf("received %q", got)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("no inbound audio")
	}
}
