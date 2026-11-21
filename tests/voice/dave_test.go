package voice_test

import (
	"bytes"
	"crypto/rand"
	"slices"
	"testing"

	"github.com/prettylittlelies/cry/mls"
	"github.com/prettylittlelies/cry/snowflake"
	"github.com/prettylittlelies/cry/voice/dave"
)

const channel = snowflake.ID(1234567890123456789)

type gateway struct {
	t        *testing.T
	suite    *mls.CipherSuite
	signer   *mls.SignaturePrivateKey
	sender   mls.ExternalSender
	members  map[snowflake.ID]*dave.Session
	inGroup  []snowflake.ID
	packages map[snowflake.ID]*mls.KeyPackage
	epoch    uint64
}

func newGateway(t *testing.T) *gateway {
	t.Helper()
	s, _ := mls.Suite(mls.P256AES128GCMSHA256P256)
	signer, err := s.GenerateSignatureKey()
	if err != nil {
		t.Fatal(err)
	}
	return &gateway{
		t: t, suite: s, signer: signer,
		sender:   mls.ExternalSender{SignatureKey: signer.Public(), Credential: mls.BasicCredential([]byte("voice"))},
		members:  map[snowflake.ID]*dave.Session{},
		packages: map[snowflake.ID]*mls.KeyPackage{},
	}
}

func (g *gateway) recognized(user snowflake.ID) bool {
	_, ok := g.members[user]
	return ok
}

func (g *gateway) connect(user snowflake.ID) {
	g.t.Helper()
	s, err := dave.NewSession(user, nil)
	if err != nil {
		g.t.Fatal(err)
	}
	s.Reset(dave.ProtocolVersion, channel)
	raw, err := dave.ParseExternalSender(mls.Marshal(&g.sender))
	if err != nil {
		g.t.Fatal(err)
	}
	if err := s.SetExternalSender(raw); err != nil {
		g.t.Fatal(err)
	}
	encoded, err := s.KeyPackage()
	if err != nil {
		g.t.Fatal(err)
	}
	kp := new(mls.KeyPackage)
	if err := mls.Unmarshal(encoded, kp); err != nil {
		g.t.Fatal(err)
	}
	if err := kp.Verify(g.suite); err != nil {
		g.t.Fatal(err)
	}
	if got, ok := dave.UserOf(kp.LeafNode.Credential); !ok || got != user {
		g.t.Fatal("credential is not the user id")
	}
	g.members[user] = s
	g.packages[user] = kp
}

func (g *gateway) leafOf(user snowflake.ID, group *mls.Group) mls.LeafIndex {
	l, ok := group.Tree().Find(func(leaf *mls.LeafNode) bool {
		u, _ := dave.UserOf(leaf.Credential)
		return u == user
	})
	if !ok {
		g.t.Fatalf("user %d not in tree", user)
	}
	return l
}

func (g *gateway) proposal(p mls.Proposal) *mls.Message {
	m, err := mls.ExternalProposal(g.suite, g.signer, 0, channelID(), g.epoch, p)
	if err != nil {
		g.t.Fatal(err)
	}
	return m
}

func channelID() []byte {
	return []byte{0x11, 0x22, 0x10, 0xf4, 0x7d, 0xe9, 0x81, 0x15}
}

func (g *gateway) round(proposals []*mls.Message, recipients []snowflake.ID, joiners []snowflake.ID, winner snowflake.ID) {
	g.t.Helper()
	payload := dave.EncodeProposals(&dave.Proposals{Operation: dave.ProposalsAppend, Messages: proposals})
	parsed, err := dave.ParseProposals(payload)
	if err != nil {
		g.t.Fatal(err)
	}
	commits := map[snowflake.ID][]byte{}
	for _, u := range recipients {
		out, err := g.members[u].ProcessProposals(parsed, g.recognized)
		if err != nil {
			g.t.Fatalf("user %d proposals: %v", u, err)
		}
		if out == nil {
			g.t.Fatalf("user %d produced no commit", u)
		}
		commits[u] = out
	}
	commit, welcome, err := dave.ParseCommitWelcome(commits[winner])
	if err != nil {
		g.t.Fatal(err)
	}
	if (welcome != nil) != (len(joiners) > 0) {
		g.t.Fatalf("welcome present=%v joiners=%v", welcome != nil, joiners)
	}
	raw := mls.Marshal(commit)
	for _, u := range recipients {
		if !slices.Contains(g.inGroup, u) && u != winner {
			continue
		}
		if err := g.members[u].ProcessCommit(raw); err != nil {
			g.t.Fatalf("user %d commit: %v", u, err)
		}
	}
	for _, u := range joiners {
		if err := g.members[u].ProcessWelcome(mls.Marshal(welcome), g.recognized); err != nil {
			g.t.Fatalf("user %d welcome: %v", u, err)
		}
	}
	g.epoch++
}

func (g *gateway) verifyMedia(users ...snowflake.ID) {
	g.t.Helper()
	for _, sender := range users {
		for _, receiver := range users {
			sr, err := g.members[sender].Ratchet(sender)
			if err != nil {
				g.t.Fatal(err)
			}
			rr, err := g.members[receiver].Ratchet(sender)
			if err != nil {
				g.t.Fatal(err)
			}
			var enc dave.Encryptor
			enc.SetRatchet(sr)
			var dec dave.Decryptor
			dec.SetRatchet(rr)
			for range 3 {
				frame := make([]byte, 60)
				rand.Read(frame)
				protected, err := enc.Encrypt(frame)
				if err != nil {
					g.t.Fatal(err)
				}
				plain, err := dec.Decrypt(protected)
				if err != nil || !bytes.Equal(plain, frame) {
					g.t.Fatalf("%d -> %d: %v", sender, receiver, err)
				}
			}
		}
		if !bytes.Equal(g.members[sender].EpochAuthenticator(), g.members[users[0]].EpochAuthenticator()) {
			g.t.Fatalf("user %d epoch authenticator differs", sender)
		}
	}
}

func TestDAVEGroupLifecycle(t *testing.T) {
	g := newGateway(t)
	alice, bob, carol, dan := snowflake.ID(111111111111111111), snowflake.ID(222222222222222222), snowflake.ID(333333333333333333), snowflake.ID(444444444444444444)
	for _, u := range []snowflake.ID{alice, bob, carol} {
		g.connect(u)
	}
	initial := []*mls.Message{
		g.proposal(mls.Proposal{Type: mls.ProposalAdd, Add: g.packages[bob]}),
		g.proposal(mls.Proposal{Type: mls.ProposalAdd, Add: g.packages[carol]}),
	}
	g.round(initial, []snowflake.ID{alice}, []snowflake.ID{bob, carol}, alice)
	g.inGroup = []snowflake.ID{alice, bob, carol}
	g.verifyMedia(alice, bob, carol)

	g.connect(dan)
	g.round([]*mls.Message{g.proposal(mls.Proposal{Type: mls.ProposalAdd, Add: g.packages[dan]})}, []snowflake.ID{alice, bob, carol}, []snowflake.ID{dan}, bob)
	g.inGroup = append(g.inGroup, dan)
	g.verifyMedia(alice, bob, carol, dan)

	remove := g.proposal(mls.Proposal{Type: mls.ProposalRemove, Remove: g.leafOf(alice, mustGroup(t, g, bob))})
	delete(g.members, alice)
	g.inGroup = []snowflake.ID{bob, carol, dan}
	g.round([]*mls.Message{remove}, []snowflake.ID{bob, carol, dan}, nil, dan)
	g.verifyMedia(bob, carol, dan)
	if got := g.members[carol].Members(); len(got) != 3 || slices.Contains(got, alice) {
		t.Fatalf("roster %v", got)
	}
	if g.members[bob].Epoch() != 3 {
		t.Fatalf("epoch %d", g.members[bob].Epoch())
	}
}

func mustGroup(t *testing.T, g *gateway, user snowflake.ID) *mls.Group {
	t.Helper()
	return g.members[user].Group()
}

func TestDAVERejectsUnexpectedUsers(t *testing.T) {
	g := newGateway(t)
	alice, mallory := snowflake.ID(111111111111111111), snowflake.ID(999999999999999999)
	g.connect(alice)
	g.connect(mallory)
	delete(g.members, mallory)
	add := g.proposal(mls.Proposal{Type: mls.ProposalAdd, Add: g.packages[mallory]})
	parsed, _ := dave.ParseProposals(dave.EncodeProposals(&dave.Proposals{Operation: dave.ProposalsAppend, Messages: []*mls.Message{add}}))
	if _, err := g.members[alice].ProcessProposals(parsed, g.recognized); err == nil {
		t.Fatal("add for an unrecognized user must be refused")
	}
}
