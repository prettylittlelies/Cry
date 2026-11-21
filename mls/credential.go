package mls

import (
	"bytes"
	"fmt"
	"slices"
)

type ProtocolVersion uint16

const MLS10 ProtocolVersion = 1

type CredentialType uint16

const (
	CredentialBasic CredentialType = 1
	CredentialX509  CredentialType = 2
)

type ExtensionType uint16

const (
	ExtensionApplicationID        ExtensionType = 1
	ExtensionRatchetTree          ExtensionType = 2
	ExtensionRequiredCapabilities ExtensionType = 3
	ExtensionExternalPub          ExtensionType = 4
	ExtensionExternalSenders      ExtensionType = 5
)

type Credential struct {
	Type         CredentialType
	Identity     []byte
	Certificates [][]byte
}

func BasicCredential(identity []byte) Credential {
	return Credential{Type: CredentialBasic, Identity: identity}
}

func (c *Credential) Equal(o Credential) bool {
	return bytes.Equal(Marshal(c), Marshal(&o))
}

func (c *Credential) encode(w *Writer) {
	w.U16(uint16(c.Type))
	switch c.Type {
	case CredentialBasic:
		w.Opaque(c.Identity)
	case CredentialX509:
		w.Vector(func(w *Writer) {
			for _, cert := range c.Certificates {
				w.Opaque(cert)
			}
		})
	}
}

func (c *Credential) decode(r *Reader) {
	c.Type = CredentialType(r.U16())
	switch c.Type {
	case CredentialBasic:
		c.Identity = r.Opaque()
	case CredentialX509:
		r.Vector(func(r *Reader) { c.Certificates = append(c.Certificates, r.Opaque()) })
	default:
		r.Fail(fmt.Errorf("mls: unknown credential type %d", c.Type))
	}
}

type Capabilities struct {
	Versions     []ProtocolVersion
	CipherSuites []CipherSuiteID
	Extensions   []ExtensionType
	Proposals    []ProposalType
	Credentials  []CredentialType
}

func writeU16s[T ~uint16](w *Writer, items []T) {
	w.Vector(func(w *Writer) {
		for _, it := range items {
			w.U16(uint16(it))
		}
	})
}

func readU16s[T ~uint16](r *Reader) []T {
	var out []T
	r.Vector(func(r *Reader) { out = append(out, T(r.U16())) })
	return out
}

func (c *Capabilities) encode(w *Writer) {
	writeU16s(w, c.Versions)
	writeU16s(w, c.CipherSuites)
	writeU16s(w, c.Extensions)
	writeU16s(w, c.Proposals)
	writeU16s(w, c.Credentials)
}

func (c *Capabilities) decode(r *Reader) {
	c.Versions = readU16s[ProtocolVersion](r)
	c.CipherSuites = readU16s[CipherSuiteID](r)
	c.Extensions = readU16s[ExtensionType](r)
	c.Proposals = readU16s[ProposalType](r)
	c.Credentials = readU16s[CredentialType](r)
}

func (c *Capabilities) SupportsExtension(t ExtensionType) bool {
	return t <= ExtensionExternalSenders || slices.Contains(c.Extensions, t)
}

func (c *Capabilities) SupportsProposal(t ProposalType) bool {
	return t <= ProposalGroupContextExtensions || slices.Contains(c.Proposals, t)
}

type Lifetime struct {
	NotBefore uint64
	NotAfter  uint64
}

func InfiniteLifetime() Lifetime { return Lifetime{NotBefore: 0, NotAfter: ^uint64(0)} }

func (l *Lifetime) encode(w *Writer) {
	w.U64(l.NotBefore)
	w.U64(l.NotAfter)
}

func (l *Lifetime) decode(r *Reader) {
	l.NotBefore = r.U64()
	l.NotAfter = r.U64()
}

type Extension struct {
	Type ExtensionType
	Data []byte
}

func (e *Extension) encode(w *Writer) {
	w.U16(uint16(e.Type))
	w.Opaque(e.Data)
}

func (e *Extension) decode(r *Reader) {
	e.Type = ExtensionType(r.U16())
	e.Data = r.Opaque()
}

type Extensions []Extension

func (es Extensions) encode(w *Writer) { writeList(w, es) }

func (es *Extensions) decode(r *Reader) {
	*es = readList[Extension](r)
	seen := map[ExtensionType]bool{}
	for _, e := range *es {
		if seen[e.Type] {
			r.Fail(fmt.Errorf("mls: duplicate extension %d", e.Type))
		}
		seen[e.Type] = true
	}
}

func (es Extensions) Find(t ExtensionType) ([]byte, bool) {
	for _, e := range es {
		if e.Type == t {
			return e.Data, true
		}
	}
	return nil, false
}

func (es Extensions) With(t ExtensionType, data []byte) Extensions {
	out := slices.DeleteFunc(slices.Clone(es), func(e Extension) bool { return e.Type == t })
	return append(out, Extension{Type: t, Data: data})
}

type ExternalSender struct {
	SignatureKey []byte
	Credential   Credential
}

func (e *ExternalSender) encode(w *Writer) {
	w.Opaque(e.SignatureKey)
	e.Credential.encode(w)
}

func (e *ExternalSender) decode(r *Reader) {
	e.SignatureKey = r.Opaque()
	e.Credential.decode(r)
}

type ExternalSenders []ExternalSender

func (es ExternalSenders) encode(w *Writer) { writeList(w, es) }

func (es *ExternalSenders) decode(r *Reader) { *es = readList[ExternalSender](r) }

func (es ExternalSenders) Extension() Extension {
	return Extension{Type: ExtensionExternalSenders, Data: Marshal(es)}
}
