package mls_test

import (
	"bytes"
	"testing"

	"github.com/prettylittlelies/cry/mls"
)

type member struct {
	name  string
	kp    *mls.KeyPackageSecrets
	group *mls.Group
}

func newMember(t *testing.T, s *mls.CipherSuite, name string) *member {
	t.Helper()
	signer, err := s.GenerateSignatureKey()
	if err != nil {
		t.Fatal(err)
	}
	kp, err := mls.NewKeyPackage(s, signer, mls.LeafOptions{Credential: mls.BasicCredential([]byte(name))})
	if err != nil {
		t.Fatal(err)
	}
	if err := kp.Package.Verify(s); err != nil {
		t.Fatal(err)
	}
	return &member{name: name, kp: kp}
}

type deliveryService struct {
	t      *testing.T
	s      *mls.CipherSuite
	signer *mls.SignaturePrivateKey
}

func (ds *deliveryService) propose(g *mls.Group, p mls.Proposal) *mls.Message {
	ds.t.Helper()
	m, err := mls.ExternalProposal(ds.s, ds.signer, 0, g.GroupID(), g.Epoch(), p)
	if err != nil {
		ds.t.Fatal(err)
	}
	return roundTrip(ds.t, m)
}

func roundTrip(t *testing.T, m *mls.Message) *mls.Message {
	t.Helper()
	var out mls.Message
	if err := mls.Unmarshal(mls.Marshal(m), &out); err != nil {
		t.Fatal(err)
	}
	return &out
}

func deliver(t *testing.T, members []*member, proposal *mls.Message) {
	t.Helper()
	for _, m := range members {
		ac, err := m.group.Verify(proposal)
		if err != nil {
			t.Fatalf("%s verify proposal: %v", m.name, err)
		}
		if _, err := m.group.CacheProposal(ac); err != nil {
			t.Fatalf("%s cache: %v", m.name, err)
		}
	}
}

func commit(t *testing.T, committer *member, others []*member, joiners ...*member) {
	t.Helper()
	res, err := committer.group.Commit(mls.CommitOptions{InlineTree: true})
	if err != nil {
		t.Fatalf("%s commit: %v", committer.name, err)
	}
	msg := roundTrip(t, res.Commit)
	for _, m := range others {
		ac, err := m.group.Verify(msg)
		if err != nil {
			t.Fatalf("%s verify commit: %v", m.name, err)
		}
		if m.group, err = m.group.HandleCommit(ac); err != nil {
			t.Fatalf("%s handle commit: %v", m.name, err)
		}
	}
	committer.group = res.Next
	if len(joiners) > 0 && res.Welcome == nil {
		t.Fatal("adds must produce a welcome")
	}
	for _, j := range joiners {
		welcome := roundTrip(t, &mls.Message{WireFormat: mls.WireWelcome, Welcome: res.Welcome}).Welcome
		if j.group, err = mls.Join(committer.group.Suite(), welcome, j.kp, mls.JoinOptions{}); err != nil {
			t.Fatalf("%s join: %v", j.name, err)
		}
	}
}

func agree(t *testing.T, epoch uint64, members ...*member) {
	t.Helper()
	ref := members[0].group
	for _, m := range members {
		if m.group.Epoch() != epoch {
			t.Fatalf("%s at epoch %d, want %d", m.name, m.group.Epoch(), epoch)
		}
		for _, other := range members {
			id := []byte(other.name)
			if !bytes.Equal(m.group.Export("Discord Secure Frames v0", id, 16), ref.Export("Discord Secure Frames v0", id, 16)) {
				t.Fatalf("%s disagrees on %s's secret at epoch %d", m.name, other.name, epoch)
			}
		}
		if !bytes.Equal(m.group.EpochAuthenticator(), ref.EpochAuthenticator()) {
			t.Fatalf("%s epoch authenticator differs", m.name)
		}
	}
}

func TestExternalSenderLifecycle(t *testing.T) {
	s := suite(t)
	dsKey, err := s.GenerateSignatureKey()
	if err != nil {
		t.Fatal(err)
	}
	ds := &deliveryService{t: t, s: s, signer: dsKey}
	ext := mls.Extensions{mls.ExternalSenders{{SignatureKey: dsKey.Public(), Credential: mls.BasicCredential([]byte("gateway"))}}.Extension()}

	alice, bob, carol, dave := newMember(t, s, "alice"), newMember(t, s, "bob"), newMember(t, s, "carol"), newMember(t, s, "dave")
	if alice.group, err = mls.NewGroup(s, []byte("channel"), alice.kp, ext); err != nil {
		t.Fatal(err)
	}

	deliver(t, []*member{alice}, ds.propose(alice.group, mls.Proposal{Type: mls.ProposalAdd, Add: &bob.kp.Package}))
	commit(t, alice, nil, bob)
	agree(t, 1, alice, bob)

	add := ds.propose(alice.group, mls.Proposal{Type: mls.ProposalAdd, Add: &carol.kp.Package})
	deliver(t, []*member{alice, bob}, add)
	commit(t, bob, []*member{alice}, carol)
	agree(t, 2, alice, bob, carol)

	remove := ds.propose(alice.group, mls.Proposal{Type: mls.ProposalRemove, Remove: alice.group.Self()})
	addDave := ds.propose(alice.group, mls.Proposal{Type: mls.ProposalAdd, Add: &dave.kp.Package})
	deliver(t, []*member{alice, bob, carol}, remove)
	deliver(t, []*member{alice, bob, carol}, addDave)
	commit(t, carol, []*member{bob}, dave)
	agree(t, 3, bob, carol, dave)
	if ac, err := alice.group.Verify(roundTrip(t, mustCommitOf(t, carol))); err == nil {
		if _, err := alice.group.HandleCommit(ac); err == nil {
			t.Fatal("removed member processed a commit for a later epoch")
		}
	}

	for i, committer := range []*member{bob, dave, carol} {
		others := []*member{}
		for _, m := range []*member{bob, carol, dave} {
			if m != committer {
				others = append(others, m)
			}
		}
		res, err := committer.group.Commit(mls.CommitOptions{ForcePath: true})
		if err != nil {
			t.Fatal(err)
		}
		for _, m := range others {
			ac, err := m.group.Verify(roundTrip(t, res.Commit))
			if err != nil {
				t.Fatalf("round %d %s: %v", i, m.name, err)
			}
			if m.group, err = m.group.HandleCommit(ac); err != nil {
				t.Fatalf("round %d %s: %v", i, m.name, err)
			}
		}
		committer.group = res.Next
		agree(t, uint64(4+i), bob, carol, dave)
	}
	if err := carol.group.Tree().VerifyParentHashes(s); err != nil {
		t.Fatal(err)
	}
}

func mustCommitOf(t *testing.T, m *member) *mls.Message {
	t.Helper()
	res, err := m.group.Commit(mls.CommitOptions{ForcePath: true})
	if err != nil {
		t.Fatal(err)
	}
	return res.Commit
}
