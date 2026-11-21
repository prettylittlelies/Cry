package voice_test

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/prettylittlelies/cry/mls"
	"github.com/prettylittlelies/cry/snowflake"
	"github.com/prettylittlelies/cry/voice/dave"
)

type libdave struct {
	t   *testing.T
	in  io.WriteCloser
	out *bufio.Scanner
}

func startLibdave(t *testing.T) *libdave {
	t.Helper()
	path := os.Getenv("CRY_LIBDAVE_HARNESS")
	if path == "" {
		t.Skip("CRY_LIBDAVE_HARNESS not set")
	}
	cmd := exec.Command(path)
	in, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	out, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { in.Close(); cmd.Wait() })
	sc := bufio.NewScanner(out)
	sc.Buffer(make([]byte, 1<<20), 1<<24)
	return &libdave{t: t, in: in, out: sc}
}

func (l *libdave) call(args ...any) ([]byte, error) {
	l.t.Helper()
	parts := make([]string, len(args))
	for i, a := range args {
		if b, ok := a.([]byte); ok {
			parts[i] = hex.EncodeToString(b)
			continue
		}
		parts[i] = fmt.Sprint(a)
	}
	fmt.Fprintln(l.in, strings.Join(parts, " "))
	for l.out.Scan() {
		line := l.out.Text()
		switch {
		case line == "ok":
			return nil, nil
		case strings.HasPrefix(line, "ok "):
			return hex.DecodeString(line[3:])
		case strings.HasPrefix(line, "err"):
			return nil, fmt.Errorf("libdave: %s", line)
		}
	}
	l.t.Fatal("libdave harness exited")
	return nil, nil
}

func (l *libdave) must(args ...any) []byte {
	l.t.Helper()
	out, err := l.call(args...)
	if err != nil {
		l.t.Fatal(err)
	}
	return out
}

func proposalsPayload(msgs ...*mls.Message) []byte {
	return dave.EncodeProposals(&dave.Proposals{Operation: dave.ProposalsAppend, Messages: msgs})
}

func keyPackageOf(t *testing.T, raw []byte) *mls.KeyPackage {
	t.Helper()
	kp := new(mls.KeyPackage)
	if err := mls.Unmarshal(raw, kp); err != nil {
		t.Fatal(err)
	}
	return kp
}

func TestLibdaveInterop(t *testing.T) {
	for _, goWins := range []bool{true, false} {
		t.Run(fmt.Sprintf("go_commits_first=%v", goWins), func(t *testing.T) { runInterop(t, goWins) })
	}
}

func runInterop(t *testing.T, goWins bool) {
	native := startLibdave(t)
	g := newGateway(t)
	goUser, nativeUser, third := snowflake.ID(111111111111111111), snowflake.ID(222222222222222222), snowflake.ID(333333333333333333)
	g.connect(goUser)
	g.members[nativeUser] = nil

	native.must("init", nativeUser, channel)
	native.must("recognize", goUser)
	native.must("recognize", third)
	native.must("extsender", mls.Marshal(&g.sender))
	nativeKP := keyPackageOf(t, native.must("keypackage"))
	if err := nativeKP.Verify(g.suite); err != nil {
		t.Fatalf("libdave key package: %v", err)
	}

	goCommit, err := g.members[goUser].ProcessProposals(mustParse(t, proposalsPayload(g.proposal(mls.Proposal{Type: mls.ProposalAdd, Add: nativeKP}))), g.recognized)
	if err != nil {
		t.Fatal(err)
	}
	nativeCommit := native.must("proposals", proposalsPayload(g.proposal(mls.Proposal{Type: mls.ProposalAdd, Add: g.packages[goUser]})))
	if goWins {
		commit, welcome := split(t, goCommit)
		if err := g.members[goUser].ProcessCommit(mls.Marshal(commit)); err != nil {
			t.Fatal(err)
		}
		native.must("welcome", mls.Marshal(welcome))
	} else {
		commit, welcome := split(t, nativeCommit)
		native.must("commit", mls.Marshal(commit))
		if err := g.members[goUser].ProcessWelcome(mls.Marshal(welcome), g.recognized); err != nil {
			t.Fatalf("go joining libdave group: %v", err)
		}
	}
	g.epoch = 1
	if !bytes.Equal(native.must("auth"), g.members[goUser].EpochAuthenticator()) {
		t.Fatal("epoch authenticators differ after group creation")
	}
	exchangeMedia(t, native, g, goUser, nativeUser)

	g.connect(third)
	add := g.proposal(mls.Proposal{Type: mls.ProposalAdd, Add: g.packages[third]})
	goCommit, err = g.members[goUser].ProcessProposals(mustParse(t, proposalsPayload(add)), g.recognized)
	if err != nil {
		t.Fatal(err)
	}
	nativeCommit = native.must("proposals", proposalsPayload(add))
	winner := nativeCommit
	if goWins {
		winner = goCommit
	}
	commit, welcome := split(t, winner)
	if err := g.members[goUser].ProcessCommit(mls.Marshal(commit)); err != nil {
		t.Fatalf("go processing commit: %v", err)
	}
	native.must("commit", mls.Marshal(commit))
	if err := g.members[third].ProcessWelcome(mls.Marshal(welcome), g.recognized); err != nil {
		t.Fatal(err)
	}
	g.epoch = 2
	for _, u := range []snowflake.ID{goUser, third} {
		if !bytes.Equal(native.must("auth"), g.members[u].EpochAuthenticator()) {
			t.Fatalf("user %d authenticator differs at epoch 2", u)
		}
	}
	exchangeMedia(t, native, g, third, nativeUser)

	remove := g.proposal(mls.Proposal{Type: mls.ProposalRemove, Remove: g.leafOf(goUser, g.members[third].Group())})
	thirdCommit, err := g.members[third].ProcessProposals(mustParse(t, proposalsPayload(remove)), g.recognized)
	if err != nil {
		t.Fatal(err)
	}
	nativeCommit = native.must("proposals", proposalsPayload(remove))
	winner = thirdCommit
	if !goWins {
		winner = nativeCommit
	}
	commit, _ = split(t, winner)
	if err := g.members[third].ProcessCommit(mls.Marshal(commit)); err != nil {
		t.Fatalf("go processing removal: %v", err)
	}
	native.must("commit", mls.Marshal(commit))
	if !bytes.Equal(native.must("auth"), g.members[third].EpochAuthenticator()) {
		t.Fatal("authenticators differ after removal")
	}
	exchangeMedia(t, native, g, third, nativeUser)
}

func mustParse(t *testing.T, payload []byte) *dave.Proposals {
	t.Helper()
	p, err := dave.ParseProposals(payload)
	if err != nil {
		t.Fatal(err)
	}
	return p
}

func split(t *testing.T, payload []byte) (*mls.Message, *mls.Welcome) {
	t.Helper()
	commit, welcome, err := dave.ParseCommitWelcome(payload)
	if err != nil {
		t.Fatal(err)
	}
	return commit, welcome
}

func exchangeMedia(t *testing.T, native *libdave, g *gateway, goUser, nativeUser snowflake.ID) {
	t.Helper()
	frame := make([]byte, 80)
	rand.Read(frame)

	fromNative := native.must("encrypt", frame)
	r, err := g.members[goUser].Ratchet(nativeUser)
	if err != nil {
		t.Fatal(err)
	}
	var dec dave.Decryptor
	dec.SetRatchet(r)
	got, err := dec.Decrypt(fromNative)
	if err != nil || !bytes.Equal(got, frame) {
		t.Fatalf("go decrypting libdave frame: %v", err)
	}

	own, err := g.members[goUser].Ratchet(goUser)
	if err != nil {
		t.Fatal(err)
	}
	var enc dave.Encryptor
	enc.SetRatchet(own)
	protected, err := enc.Encrypt(frame)
	if err != nil {
		t.Fatal(err)
	}
	if back := native.must("decrypt", goUser, protected); !bytes.Equal(back, frame) {
		t.Fatal("libdave decrypted a different frame")
	}
}
