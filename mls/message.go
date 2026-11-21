package mls

import "fmt"

type WireFormat uint16

const (
	WirePublicMessage  WireFormat = 1
	WirePrivateMessage WireFormat = 2
	WireWelcome        WireFormat = 3
	WireGroupInfo      WireFormat = 4
	WireKeyPackage     WireFormat = 5
)

type SenderType uint8

const (
	SenderMember            SenderType = 1
	SenderExternal          SenderType = 2
	SenderNewMemberProposal SenderType = 3
	SenderNewMemberCommit   SenderType = 4
)

type Sender struct {
	Type  SenderType
	Index uint32
}

func (s *Sender) encode(w *Writer) {
	w.U8(uint8(s.Type))
	if s.Type == SenderMember || s.Type == SenderExternal {
		w.U32(s.Index)
	}
}

func (s *Sender) decode(r *Reader) {
	s.Type = SenderType(r.U8())
	switch s.Type {
	case SenderMember, SenderExternal:
		s.Index = r.U32()
	case SenderNewMemberProposal, SenderNewMemberCommit:
	default:
		r.Fail(fmt.Errorf("mls: invalid sender type %d", s.Type))
	}
}

type ContentType uint8

const (
	ContentApplication ContentType = 1
	ContentProposal    ContentType = 2
	ContentCommit      ContentType = 3
)

type FramedContent struct {
	GroupID           []byte
	Epoch             uint64
	Sender            Sender
	AuthenticatedData []byte
	Type              ContentType
	Application       []byte
	Proposal          *Proposal
	Commit            *Commit
}

func (c *FramedContent) encode(w *Writer) {
	w.Opaque(c.GroupID)
	w.U64(c.Epoch)
	c.Sender.encode(w)
	w.Opaque(c.AuthenticatedData)
	c.encodeBody(w)
}

func (c *FramedContent) encodeBody(w *Writer) {
	w.U8(uint8(c.Type))
	switch c.Type {
	case ContentApplication:
		w.Opaque(c.Application)
	case ContentProposal:
		c.Proposal.encode(w)
	case ContentCommit:
		c.Commit.encode(w)
	}
}

func (c *FramedContent) decode(r *Reader) {
	c.GroupID = r.Opaque()
	c.Epoch = r.U64()
	c.Sender.decode(r)
	c.AuthenticatedData = r.Opaque()
	c.decodeBody(r, ContentType(r.U8()))
}

func (c *FramedContent) decodeBody(r *Reader, t ContentType) {
	c.Type = t
	switch t {
	case ContentApplication:
		c.Application = r.Opaque()
	case ContentProposal:
		c.Proposal = new(Proposal)
		c.Proposal.decode(r)
	case ContentCommit:
		c.Commit = new(Commit)
		c.Commit.decode(r)
	default:
		r.Fail(fmt.Errorf("mls: invalid content type %d", t))
	}
}

type FramedContentAuthData struct {
	Signature       []byte
	ConfirmationTag []byte
}

func (a *FramedContentAuthData) encode(w *Writer, t ContentType) {
	w.Opaque(a.Signature)
	if t == ContentCommit {
		w.Opaque(a.ConfirmationTag)
	}
}

func (a *FramedContentAuthData) decode(r *Reader, t ContentType) {
	a.Signature = r.Opaque()
	if t == ContentCommit {
		a.ConfirmationTag = r.Opaque()
	}
}

type AuthenticatedContent struct {
	WireFormat WireFormat
	Content    FramedContent
	Auth       FramedContentAuthData
}

func (a *AuthenticatedContent) encode(w *Writer) {
	w.U16(uint16(a.WireFormat))
	a.Content.encode(w)
	a.Auth.encode(w, a.Content.Type)
}

func (a *AuthenticatedContent) tbs(ctx *GroupContext) []byte {
	var w Writer
	w.U16(uint16(MLS10))
	w.U16(uint16(a.WireFormat))
	a.Content.encode(&w)
	if t := a.Content.Sender.Type; t == SenderMember || t == SenderNewMemberCommit {
		ctx.encode(&w)
	}
	return w.buf
}

func (a *AuthenticatedContent) Sign(s *CipherSuite, key *SignaturePrivateKey, ctx *GroupContext) error {
	sig, err := s.SignWithLabel(key, "FramedContentTBS", a.tbs(ctx))
	a.Auth.Signature = sig
	return err
}

func (a *AuthenticatedContent) Verify(s *CipherSuite, public []byte, ctx *GroupContext) error {
	if !s.VerifyWithLabel(public, "FramedContentTBS", a.tbs(ctx), a.Auth.Signature) {
		return fmt.Errorf("%w: framed content", ErrInvalidSignature)
	}
	return nil
}

func (a *AuthenticatedContent) ProposalRef(s *CipherSuite) []byte {
	return s.RefHash("MLS 1.0 Proposal Reference", Marshal(a))
}

func (a *AuthenticatedContent) confirmedInput() []byte {
	var w Writer
	w.U16(uint16(a.WireFormat))
	a.Content.encode(&w)
	w.Opaque(a.Auth.Signature)
	return w.buf
}

type PublicMessage struct {
	Content       FramedContent
	Auth          FramedContentAuthData
	MembershipTag []byte
}

func (m *PublicMessage) encode(w *Writer) {
	m.Content.encode(w)
	m.Auth.encode(w, m.Content.Type)
	if m.Content.Sender.Type == SenderMember {
		w.Opaque(m.MembershipTag)
	}
}

func (m *PublicMessage) decode(r *Reader) {
	m.Content.decode(r)
	m.Auth.decode(r, m.Content.Type)
	if m.Content.Sender.Type == SenderMember {
		m.MembershipTag = r.Opaque()
	}
}

func (m *PublicMessage) Authenticated() AuthenticatedContent {
	return AuthenticatedContent{WireFormat: WirePublicMessage, Content: m.Content, Auth: m.Auth}
}

func membershipInput(a *AuthenticatedContent, ctx *GroupContext) []byte {
	var w Writer
	w.Raw(a.tbs(ctx))
	a.Auth.encode(&w, a.Content.Type)
	return w.buf
}

type PrivateMessage struct {
	GroupID             []byte
	Epoch               uint64
	Type                ContentType
	AuthenticatedData   []byte
	EncryptedSenderData []byte
	Ciphertext          []byte
}

func (m *PrivateMessage) encode(w *Writer) {
	w.Opaque(m.GroupID)
	w.U64(m.Epoch)
	w.U8(uint8(m.Type))
	w.Opaque(m.AuthenticatedData)
	w.Opaque(m.EncryptedSenderData)
	w.Opaque(m.Ciphertext)
}

func (m *PrivateMessage) decode(r *Reader) {
	m.GroupID = r.Opaque()
	m.Epoch = r.U64()
	m.Type = ContentType(r.U8())
	m.AuthenticatedData = r.Opaque()
	m.EncryptedSenderData = r.Opaque()
	m.Ciphertext = r.Opaque()
}

type Message struct {
	WireFormat WireFormat
	Public     *PublicMessage
	Private    *PrivateMessage
	Welcome    *Welcome
	GroupInfo  *GroupInfo
	KeyPackage *KeyPackage
}

func (m *Message) encode(w *Writer) {
	w.U16(uint16(MLS10))
	w.U16(uint16(m.WireFormat))
	switch m.WireFormat {
	case WirePublicMessage:
		m.Public.encode(w)
	case WirePrivateMessage:
		m.Private.encode(w)
	case WireWelcome:
		m.Welcome.encode(w)
	case WireGroupInfo:
		m.GroupInfo.encode(w)
	case WireKeyPackage:
		m.KeyPackage.encode(w)
	}
}

func (m *Message) decode(r *Reader) {
	if v := ProtocolVersion(r.U16()); v != MLS10 && r.Err() == nil {
		r.Fail(fmt.Errorf("mls: unsupported protocol version %d", v))
	}
	m.WireFormat = WireFormat(r.U16())
	switch m.WireFormat {
	case WirePublicMessage:
		m.Public = new(PublicMessage)
		m.Public.decode(r)
	case WirePrivateMessage:
		m.Private = new(PrivateMessage)
		m.Private.decode(r)
	case WireWelcome:
		m.Welcome = new(Welcome)
		m.Welcome.decode(r)
	case WireGroupInfo:
		m.GroupInfo = new(GroupInfo)
		m.GroupInfo.decode(r)
	case WireKeyPackage:
		m.KeyPackage = new(KeyPackage)
		m.KeyPackage.decode(r)
	default:
		r.Fail(fmt.Errorf("mls: invalid wire format %d", m.WireFormat))
	}
}

func (m *Message) GroupID() []byte {
	switch {
	case m.Public != nil:
		return m.Public.Content.GroupID
	case m.Private != nil:
		return m.Private.GroupID
	}
	return nil
}

func (m *Message) Epoch() uint64 {
	switch {
	case m.Public != nil:
		return m.Public.Content.Epoch
	case m.Private != nil:
		return m.Private.Epoch
	}
	return 0
}
