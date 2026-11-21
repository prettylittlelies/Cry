package mls

import "fmt"

type ProposalType uint16

const (
	ProposalAdd                    ProposalType = 1
	ProposalUpdate                 ProposalType = 2
	ProposalRemove                 ProposalType = 3
	ProposalPSK                    ProposalType = 4
	ProposalReInit                 ProposalType = 5
	ProposalExternalInit           ProposalType = 6
	ProposalGroupContextExtensions ProposalType = 7
)

type PSKType uint8

const (
	PSKExternal   PSKType = 1
	PSKResumption PSKType = 2
)

type PreSharedKeyID struct {
	Type    PSKType
	ID      []byte
	Usage   uint8
	GroupID []byte
	Epoch   uint64
	Nonce   []byte
}

func (p *PreSharedKeyID) encode(w *Writer) {
	w.U8(uint8(p.Type))
	switch p.Type {
	case PSKExternal:
		w.Opaque(p.ID)
	case PSKResumption:
		w.U8(p.Usage)
		w.Opaque(p.GroupID)
		w.U64(p.Epoch)
	}
	w.Opaque(p.Nonce)
}

func (p *PreSharedKeyID) decode(r *Reader) {
	p.Type = PSKType(r.U8())
	switch p.Type {
	case PSKExternal:
		p.ID = r.Opaque()
	case PSKResumption:
		p.Usage = r.U8()
		p.GroupID = r.Opaque()
		p.Epoch = r.U64()
	default:
		r.Fail(fmt.Errorf("mls: invalid psk type %d", p.Type))
	}
	p.Nonce = r.Opaque()
}

type ReInit struct {
	GroupID     []byte
	Version     ProtocolVersion
	CipherSuite CipherSuiteID
	Extensions  Extensions
}

type Proposal struct {
	Type       ProposalType
	Add        *KeyPackage
	Update     *LeafNode
	Remove     LeafIndex
	PSK        *PreSharedKeyID
	ReInit     *ReInit
	KEMOutput  []byte
	Extensions Extensions
}

func (p *Proposal) encode(w *Writer) {
	w.U16(uint16(p.Type))
	switch p.Type {
	case ProposalAdd:
		p.Add.encode(w)
	case ProposalUpdate:
		p.Update.encode(w)
	case ProposalRemove:
		w.U32(uint32(p.Remove))
	case ProposalPSK:
		p.PSK.encode(w)
	case ProposalReInit:
		w.Opaque(p.ReInit.GroupID)
		w.U16(uint16(p.ReInit.Version))
		w.U16(uint16(p.ReInit.CipherSuite))
		p.ReInit.Extensions.encode(w)
	case ProposalExternalInit:
		w.Opaque(p.KEMOutput)
	case ProposalGroupContextExtensions:
		p.Extensions.encode(w)
	}
}

func (p *Proposal) decode(r *Reader) {
	p.Type = ProposalType(r.U16())
	switch p.Type {
	case ProposalAdd:
		p.Add = new(KeyPackage)
		p.Add.decode(r)
	case ProposalUpdate:
		p.Update = new(LeafNode)
		p.Update.decode(r)
	case ProposalRemove:
		p.Remove = LeafIndex(r.U32())
	case ProposalPSK:
		p.PSK = new(PreSharedKeyID)
		p.PSK.decode(r)
	case ProposalReInit:
		p.ReInit = &ReInit{GroupID: r.Opaque(), Version: ProtocolVersion(r.U16()), CipherSuite: CipherSuiteID(r.U16())}
		p.ReInit.Extensions.decode(r)
	case ProposalExternalInit:
		p.KEMOutput = r.Opaque()
	case ProposalGroupContextExtensions:
		p.Extensions.decode(r)
	default:
		r.Fail(fmt.Errorf("mls: unknown proposal type %d", p.Type))
	}
}

type ProposalOrRefType uint8

const (
	ProposalValue     ProposalOrRefType = 1
	ProposalReference ProposalOrRefType = 2
)

type ProposalOrRef struct {
	Type     ProposalOrRefType
	Proposal *Proposal
	Ref      []byte
}

func (p *ProposalOrRef) encode(w *Writer) {
	w.U8(uint8(p.Type))
	if p.Type == ProposalValue {
		p.Proposal.encode(w)
		return
	}
	w.Opaque(p.Ref)
}

func (p *ProposalOrRef) decode(r *Reader) {
	p.Type = ProposalOrRefType(r.U8())
	switch p.Type {
	case ProposalValue:
		p.Proposal = new(Proposal)
		p.Proposal.decode(r)
	case ProposalReference:
		p.Ref = r.Opaque()
	default:
		r.Fail(fmt.Errorf("mls: invalid proposal or ref %d", p.Type))
	}
}

type UpdatePathNode struct {
	EncryptionKey       []byte
	EncryptedPathSecret []HPKECiphertext
}

func (n *UpdatePathNode) encode(w *Writer) {
	w.Opaque(n.EncryptionKey)
	writeList(w, n.EncryptedPathSecret)
}

func (n *UpdatePathNode) decode(r *Reader) {
	n.EncryptionKey = r.Opaque()
	n.EncryptedPathSecret = readList[HPKECiphertext](r)
}

type UpdatePath struct {
	LeafNode LeafNode
	Nodes    []UpdatePathNode
}

func (p *UpdatePath) encode(w *Writer) {
	p.LeafNode.encode(w)
	writeList(w, p.Nodes)
}

func (p *UpdatePath) decode(r *Reader) {
	p.LeafNode.decode(r)
	p.Nodes = readList[UpdatePathNode](r)
}

type Commit struct {
	Proposals []ProposalOrRef
	Path      *UpdatePath
}

func (c *Commit) encode(w *Writer) {
	writeList(w, c.Proposals)
	w.Optional(c.Path != nil, func(w *Writer) { c.Path.encode(w) })
}

func (c *Commit) decode(r *Reader) {
	c.Proposals = readList[ProposalOrRef](r)
	r.Optional(func(r *Reader) {
		c.Path = new(UpdatePath)
		c.Path.decode(r)
	})
}
