package dave

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"slices"
	"sync"

	"github.com/prettylittlelies/cry/mls"
	"github.com/prettylittlelies/cry/snowflake"
)

const (
	ProtocolVersion   = 1
	exporterLabel     = "Discord Secure Frames v0"
	maxOutboundCommit = 4
)

var (
	ErrNoGroup         = errors.New("dave: no mls group")
	ErrUnexpectedUser  = errors.New("dave: unexpected user in proposal")
	ErrDuplicateUser   = errors.New("dave: duplicate user credential in group")
	ErrExternalSender  = errors.New("dave: group external sender mismatch")
	ErrForbiddenChange = errors.New("dave: committer changed its identity")
)

type Recognizer func(user snowflake.ID) bool

type outboundCommit struct {
	encoded []byte
	next    *mls.Group
}

type Session struct {
	mu       sync.Mutex
	suite    *mls.CipherSuite
	self     snowflake.ID
	signer   *mls.SignaturePrivateKey
	groupID  []byte
	external *mls.ExternalSender
	join     *mls.KeyPackageSecrets
	pending  *mls.Group
	current  *mls.Group
	queued   *mls.Group
	outbound []outboundCommit
	version  int
}

func NewSession(self snowflake.ID, signer *mls.SignaturePrivateKey) (*Session, error) {
	suite, err := mls.Suite(mls.P256AES128GCMSHA256P256)
	if err != nil {
		return nil, err
	}
	if signer == nil {
		if signer, err = suite.GenerateSignatureKey(); err != nil {
			return nil, err
		}
	}
	return &Session{suite: suite, self: self, signer: signer}, nil
}

func identity(user snowflake.ID) []byte {
	return binary.BigEndian.AppendUint64(nil, uint64(user))
}

func UserOf(c mls.Credential) (snowflake.ID, bool) {
	if c.Type != mls.CredentialBasic || len(c.Identity) != 8 {
		return 0, false
	}
	return snowflake.ID(binary.BigEndian.Uint64(c.Identity)), true
}

func (s *Session) Suite() *mls.CipherSuite { return s.suite }

func (s *Session) SignaturePublicKey() []byte { return s.signer.Public() }

func (s *Session) Reset(version int, group snowflake.ID) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.version, s.groupID = version, identity(group)
	s.current, s.pending, s.queued, s.outbound, s.join = nil, nil, nil, nil, nil
}

func (s *Session) leafOptions() mls.LeafOptions {
	return mls.LeafOptions{
		Credential:   mls.BasicCredential(identity(s.self)),
		Capabilities: mls.Capabilities{Credentials: []mls.CredentialType{mls.CredentialBasic}},
	}
}

func (s *Session) SetExternalSender(e *mls.ExternalSender) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.external = e
	return s.createPending()
}

func (s *Session) createPending() error {
	s.pending, s.queued, s.outbound = nil, nil, nil
	if s.groupID == nil || s.external == nil || s.current != nil {
		return nil
	}
	kp, err := s.keyPackage()
	if err != nil {
		return err
	}
	ext := mls.Extensions{mls.ExternalSenders{*s.external}.Extension()}
	s.pending, err = mls.NewGroup(s.suite, s.groupID, kp, ext)
	return err
}

func (s *Session) keyPackage() (*mls.KeyPackageSecrets, error) {
	return mls.NewKeyPackage(s.suite, s.signer, s.leafOptions())
}

func (s *Session) KeyPackage() ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	kp, err := s.keyPackage()
	if err != nil {
		return nil, err
	}
	s.join = kp
	if s.pending == nil && s.current == nil {
		if err := s.createPending(); err != nil {
			return nil, err
		}
	}
	return EncodeKeyPackage(&kp.Package), nil
}

func (s *Session) base() *mls.Group {
	if s.current != nil {
		return s.current
	}
	return s.pending
}

func (s *Session) ProcessProposals(p *Proposals, recognized Recognizer) ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	base := s.base()
	if base == nil {
		return nil, nil
	}
	next := base.Clone()
	if s.queued != nil {
		next = s.queued.Clone()
	}
	switch p.Operation {
	case ProposalsRevoke:
		for _, ref := range p.Refs {
			if !next.RevokeProposal(ref) {
				return nil, fmt.Errorf("%w: unknown revocation", mls.ErrUnknownProposal)
			}
		}
	case ProposalsAppend:
		for _, m := range p.Messages {
			ac, err := next.Verify(m)
			if err != nil {
				return nil, err
			}
			if err := validateProposal(ac, recognized); err != nil {
				return nil, err
			}
			if _, err := next.CacheProposal(ac); err != nil {
				return nil, err
			}
		}
	}
	s.queued = next
	if len(next.Proposals()) == 0 {
		return nil, nil
	}
	res, err := next.Commit(mls.CommitOptions{InlineTree: true})
	if err != nil {
		return nil, err
	}
	encoded := mls.Marshal(res.Commit)
	s.outbound = append(s.outbound, outboundCommit{encoded: encoded, next: res.Next})
	if len(s.outbound) > maxOutboundCommit {
		s.outbound = s.outbound[1:]
	}
	return EncodeCommitWelcome(res.Commit, res.Welcome), nil
}

func validateProposal(ac *mls.AuthenticatedContent, recognized Recognizer) error {
	if ac.Content.Sender.Type != mls.SenderExternal {
		return fmt.Errorf("%w: proposal not from external sender", mls.ErrInvalidProposal)
	}
	p := ac.Content.Proposal
	switch p.Type {
	case mls.ProposalRemove:
		return nil
	case mls.ProposalAdd:
		user, ok := UserOf(p.Add.LeafNode.Credential)
		if !ok || recognized != nil && !recognized(user) {
			return fmt.Errorf("%w: %d", ErrUnexpectedUser, user)
		}
		return nil
	}
	return fmt.Errorf("%w: type %d", mls.ErrInvalidProposal, p.Type)
}

func (s *Session) ProcessCommit(raw []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, o := range s.outbound {
		if bytes.Equal(o.encoded, raw) {
			return s.install(o.next)
		}
	}
	if s.current == nil || s.queued == nil {
		return ErrNoGroup
	}
	var m mls.Message
	if err := mls.Unmarshal(raw, &m); err != nil {
		return err
	}
	ac, err := s.queued.Verify(&m)
	if err != nil {
		return err
	}
	committer := mls.LeafIndex(ac.Content.Sender.Index)
	before := s.queued.Tree().Leaf(committer)
	next, err := s.queued.HandleCommit(ac)
	if err != nil {
		return err
	}
	if after := next.Tree().Leaf(committer); before == nil || after == nil || !before.Credential.Equal(after.Credential) || !bytes.Equal(before.SignatureKey, after.SignatureKey) {
		return ErrForbiddenChange
	}
	return s.install(next)
}

func (s *Session) ProcessWelcome(raw []byte, recognized Recognizer) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.join == nil {
		return ErrNoGroup
	}
	var w mls.Welcome
	if err := mls.Unmarshal(raw, &w); err != nil {
		return err
	}
	g, err := mls.Join(s.suite, &w, s.join, mls.JoinOptions{})
	if err != nil {
		return err
	}
	if !bytes.Equal(g.GroupID(), s.groupID) {
		return mls.ErrWrongGroup
	}
	senders, err := g.ExternalSenders()
	if err != nil {
		return err
	}
	if s.external == nil || len(senders) != 1 || !bytes.Equal(senders[0].SignatureKey, s.external.SignatureKey) || !senders[0].Credential.Equal(s.external.Credential) {
		return ErrExternalSender
	}
	for _, leaf := range g.Tree().Members {
		if user, ok := UserOf(leaf.Credential); !ok || user != s.self && recognized != nil && !recognized(user) {
			return fmt.Errorf("%w: %d", ErrUnexpectedUser, user)
		}
	}
	return s.install(g)
}

func (s *Session) install(g *mls.Group) error {
	seen := map[snowflake.ID]bool{}
	for _, leaf := range g.Tree().Members {
		user, ok := UserOf(leaf.Credential)
		if !ok || seen[user] {
			return fmt.Errorf("%w: %d", ErrDuplicateUser, user)
		}
		seen[user] = true
	}
	s.current, s.pending, s.queued, s.outbound, s.join = g, nil, nil, nil, nil
	return nil
}

func (s *Session) Established() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.current != nil
}

func (s *Session) Epoch() uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.current == nil {
		return 0
	}
	return s.current.Epoch()
}

func (s *Session) EpochAuthenticator() []byte {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.current == nil {
		return nil
	}
	return slices.Clone(s.current.EpochAuthenticator())
}

func (s *Session) Members() []snowflake.ID {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.current == nil {
		return nil
	}
	var out []snowflake.ID
	for _, leaf := range s.current.Tree().Members {
		if user, ok := UserOf(leaf.Credential); ok {
			out = append(out, user)
		}
	}
	return out
}

func (s *Session) Ratchet(user snowflake.ID) (*KeyRatchet, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.current == nil {
		return nil, ErrNoGroup
	}
	context := binary.LittleEndian.AppendUint64(nil, uint64(user))
	return NewKeyRatchet(s.suite, s.current.Export(exporterLabel, context, 16)), nil
}

func (s *Session) Group() *mls.Group {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.current == nil {
		return nil
	}
	return s.current.Clone()
}
