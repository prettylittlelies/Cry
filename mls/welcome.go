package mls

import "fmt"

type GroupContext struct {
	Version                 ProtocolVersion
	CipherSuite             CipherSuiteID
	GroupID                 []byte
	Epoch                   uint64
	TreeHash                []byte
	ConfirmedTranscriptHash []byte
	Extensions              Extensions
}

func (c *GroupContext) encode(w *Writer) {
	w.U16(uint16(c.Version))
	w.U16(uint16(c.CipherSuite))
	w.Opaque(c.GroupID)
	w.U64(c.Epoch)
	w.Opaque(c.TreeHash)
	w.Opaque(c.ConfirmedTranscriptHash)
	c.Extensions.encode(w)
}

func (c *GroupContext) decode(r *Reader) {
	c.Version = ProtocolVersion(r.U16())
	c.CipherSuite = CipherSuiteID(r.U16())
	c.GroupID = r.Opaque()
	c.Epoch = r.U64()
	c.TreeHash = r.Opaque()
	c.ConfirmedTranscriptHash = r.Opaque()
	c.Extensions.decode(r)
}

type GroupInfo struct {
	GroupContext    GroupContext
	Extensions      Extensions
	ConfirmationTag []byte
	Signer          LeafIndex
	Signature       []byte
}

func (g *GroupInfo) encodeContent(w *Writer) {
	g.GroupContext.encode(w)
	g.Extensions.encode(w)
	w.Opaque(g.ConfirmationTag)
	w.U32(uint32(g.Signer))
}

func (g *GroupInfo) encode(w *Writer) {
	g.encodeContent(w)
	w.Opaque(g.Signature)
}

func (g *GroupInfo) decode(r *Reader) {
	g.GroupContext.decode(r)
	g.Extensions.decode(r)
	g.ConfirmationTag = r.Opaque()
	g.Signer = LeafIndex(r.U32())
	g.Signature = r.Opaque()
}

func (g *GroupInfo) Sign(s *CipherSuite, key *SignaturePrivateKey) error {
	var w Writer
	g.encodeContent(&w)
	sig, err := s.SignWithLabel(key, "GroupInfoTBS", w.buf)
	g.Signature = sig
	return err
}

func (g *GroupInfo) Verify(s *CipherSuite, public []byte) error {
	var w Writer
	g.encodeContent(&w)
	if !s.VerifyWithLabel(public, "GroupInfoTBS", w.buf, g.Signature) {
		return fmt.Errorf("%w: group info", ErrInvalidSignature)
	}
	return nil
}

type GroupSecrets struct {
	JoinerSecret []byte
	PathSecret   []byte
	PSKs         []PreSharedKeyID
}

func (g *GroupSecrets) encode(w *Writer) {
	w.Opaque(g.JoinerSecret)
	w.Optional(g.PathSecret != nil, func(w *Writer) { w.Opaque(g.PathSecret) })
	writeList(w, g.PSKs)
}

func (g *GroupSecrets) decode(r *Reader) {
	g.JoinerSecret = r.Opaque()
	r.Optional(func(r *Reader) { g.PathSecret = r.Opaque() })
	g.PSKs = readList[PreSharedKeyID](r)
}

type EncryptedGroupSecrets struct {
	NewMember []byte
	Secrets   HPKECiphertext
}

func (e *EncryptedGroupSecrets) encode(w *Writer) {
	w.Opaque(e.NewMember)
	e.Secrets.encode(w)
}

func (e *EncryptedGroupSecrets) decode(r *Reader) {
	e.NewMember = r.Opaque()
	e.Secrets.decode(r)
}

type Welcome struct {
	CipherSuite        CipherSuiteID
	Secrets            []EncryptedGroupSecrets
	EncryptedGroupInfo []byte
}

func (wl *Welcome) encode(w *Writer) {
	w.U16(uint16(wl.CipherSuite))
	writeList(w, wl.Secrets)
	w.Opaque(wl.EncryptedGroupInfo)
}

func (wl *Welcome) decode(r *Reader) {
	wl.CipherSuite = CipherSuiteID(r.U16())
	wl.Secrets = readList[EncryptedGroupSecrets](r)
	wl.EncryptedGroupInfo = r.Opaque()
}
