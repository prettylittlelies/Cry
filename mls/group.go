package mls

import (
	"crypto/hpke"
	"errors"
	"slices"
)

var (
	ErrWrongGroup         = errors.New("mls: message for another group")
	ErrWrongEpoch         = errors.New("mls: message for another epoch")
	ErrUnknownSender      = errors.New("mls: unknown sender")
	ErrUnknownProposal    = errors.New("mls: commit references unknown proposal")
	ErrInvalidProposal    = errors.New("mls: invalid proposal")
	ErrInvalidCommit      = errors.New("mls: invalid commit")
	ErrInvalidWelcome     = errors.New("mls: invalid welcome")
	ErrMembershipTag      = errors.New("mls: invalid membership tag")
	ErrConfirmationTag    = errors.New("mls: invalid confirmation tag")
	ErrUnsupportedFeature = errors.New("mls: unsupported feature")
)

type CachedProposal struct {
	Ref      []byte
	Proposal Proposal
	Sender   Sender
}

type Group struct {
	suite      *CipherSuite
	context    GroupContext
	tree       *RatchetTree
	own        *TreeSecrets
	signer     *SignaturePrivateKey
	secrets    EpochSecrets
	interim    []byte
	proposals  []CachedProposal
	psks       PSKStore
	resumption map[uint64][]byte
}

func NewGroup(s *CipherSuite, groupID []byte, kp *KeyPackageSecrets, extensions Extensions) (*Group, error) {
	tree := NewRatchetTree(kp.Package.LeafNode.Clone())
	g := &Group{
		suite:  s,
		tree:   tree,
		own:    NewTreeSecrets(0, kp.Encryption),
		signer: kp.Signer,
		context: GroupContext{
			Version:     MLS10,
			CipherSuite: s.ID,
			GroupID:     groupID,
			Epoch:       0,
			TreeHash:    tree.TreeHash(s),
			Extensions:  extensions,
		},
	}
	g.secrets = s.EpochSecrets(s.RandomSecret(), nil, &g.context)
	g.interim = s.InterimTranscriptHash(nil, s.MAC(g.secrets.Confirmation, nil))
	g.rememberResumption()
	return g, nil
}

func (g *Group) Suite() *CipherSuite          { return g.suite }
func (g *Group) Context() GroupContext        { return g.context }
func (g *Group) GroupID() []byte              { return g.context.GroupID }
func (g *Group) Epoch() uint64                { return g.context.Epoch }
func (g *Group) Tree() *RatchetTree           { return g.tree }
func (g *Group) Self() LeafIndex              { return g.own.Leaf }
func (g *Group) Signer() *SignaturePrivateKey { return g.signer }
func (g *Group) Proposals() []CachedProposal  { return slices.Clone(g.proposals) }

func (g *Group) EpochAuthenticator() []byte { return g.secrets.Authenticator }

func (g *Group) Export(label string, context []byte, length int) []byte {
	return g.suite.Export(g.secrets.Exporter, label, context, length)
}

func (g *Group) ExternalSenders() (ExternalSenders, error) {
	raw, ok := g.context.Extensions.Find(ExtensionExternalSenders)
	if !ok {
		return nil, nil
	}
	var out ExternalSenders
	err := Unmarshal(raw, &out)
	return out, err
}

func (g *Group) Clone() *Group {
	out := *g
	out.tree = g.tree.Clone()
	out.own = g.own.Clone()
	out.proposals = slices.Clone(g.proposals)
	out.context.Extensions = slices.Clone(g.context.Extensions)
	return &out
}

func (g *Group) AddPSK(id, secret []byte) {
	if g.psks == nil {
		g.psks = PSKStore{}
	}
	g.psks[string(id)] = secret
}

func (g *Group) EncryptionKey(n NodeIndex) (hpke.PrivateKey, bool) {
	return g.own.Key(g.tree, n)
}
