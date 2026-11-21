package mls

import (
	"errors"
	"fmt"
)

type LeafNodeSource uint8

const (
	SourceKeyPackage LeafNodeSource = 1
	SourceUpdate     LeafNodeSource = 2
	SourceCommit     LeafNodeSource = 3
)

var ErrInvalidSignature = errors.New("mls: invalid signature")

type LeafNode struct {
	EncryptionKey []byte
	SignatureKey  []byte
	Credential    Credential
	Capabilities  Capabilities
	Source        LeafNodeSource
	Lifetime      Lifetime
	ParentHash    []byte
	Extensions    Extensions
	Signature     []byte
}

func (l *LeafNode) encodeContent(w *Writer) {
	w.Opaque(l.EncryptionKey)
	w.Opaque(l.SignatureKey)
	l.Credential.encode(w)
	l.Capabilities.encode(w)
	w.U8(uint8(l.Source))
	switch l.Source {
	case SourceKeyPackage:
		l.Lifetime.encode(w)
	case SourceCommit:
		w.Opaque(l.ParentHash)
	}
	l.Extensions.encode(w)
}

func (l *LeafNode) encode(w *Writer) {
	l.encodeContent(w)
	w.Opaque(l.Signature)
}

func (l *LeafNode) decode(r *Reader) {
	l.EncryptionKey = r.Opaque()
	l.SignatureKey = r.Opaque()
	l.Credential.decode(r)
	l.Capabilities.decode(r)
	l.Source = LeafNodeSource(r.U8())
	switch l.Source {
	case SourceKeyPackage:
		l.Lifetime.decode(r)
	case SourceUpdate:
	case SourceCommit:
		l.ParentHash = r.Opaque()
	default:
		r.Fail(fmt.Errorf("mls: invalid leaf node source %d", l.Source))
	}
	l.Extensions.decode(r)
	l.Signature = r.Opaque()
}

func (l *LeafNode) tbs(groupID []byte, leaf LeafIndex) []byte {
	var w Writer
	l.encodeContent(&w)
	if l.Source == SourceUpdate || l.Source == SourceCommit {
		w.Opaque(groupID)
		w.U32(uint32(leaf))
	}
	return w.buf
}

func (l *LeafNode) Sign(s *CipherSuite, key *SignaturePrivateKey, groupID []byte, leaf LeafIndex) error {
	sig, err := s.SignWithLabel(key, "LeafNodeTBS", l.tbs(groupID, leaf))
	l.Signature = sig
	return err
}

func (l *LeafNode) Verify(s *CipherSuite, groupID []byte, leaf LeafIndex) error {
	if !s.VerifyWithLabel(l.SignatureKey, "LeafNodeTBS", l.tbs(groupID, leaf), l.Signature) {
		return fmt.Errorf("%w: leaf node", ErrInvalidSignature)
	}
	return nil
}

func (l *LeafNode) Clone() LeafNode {
	var out LeafNode
	if err := Unmarshal(Marshal(l), &out); err != nil {
		panic(err)
	}
	return out
}

type ParentNode struct {
	EncryptionKey  []byte
	ParentHash     []byte
	UnmergedLeaves []LeafIndex
}

func (p *ParentNode) encode(w *Writer) {
	w.Opaque(p.EncryptionKey)
	w.Opaque(p.ParentHash)
	w.Vector(func(w *Writer) {
		for _, l := range p.UnmergedLeaves {
			w.U32(uint32(l))
		}
	})
}

func (p *ParentNode) decode(r *Reader) {
	p.EncryptionKey = r.Opaque()
	p.ParentHash = r.Opaque()
	r.Vector(func(r *Reader) { p.UnmergedLeaves = append(p.UnmergedLeaves, LeafIndex(r.U32())) })
}
