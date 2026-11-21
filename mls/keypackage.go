package mls

import (
	"crypto/hpke"
	"fmt"
)

type KeyPackage struct {
	Version     ProtocolVersion
	CipherSuite CipherSuiteID
	InitKey     []byte
	LeafNode    LeafNode
	Extensions  Extensions
	Signature   []byte
}

func (k *KeyPackage) encodeContent(w *Writer) {
	w.U16(uint16(k.Version))
	w.U16(uint16(k.CipherSuite))
	w.Opaque(k.InitKey)
	k.LeafNode.encode(w)
	k.Extensions.encode(w)
}

func (k *KeyPackage) encode(w *Writer) {
	k.encodeContent(w)
	w.Opaque(k.Signature)
}

func (k *KeyPackage) decode(r *Reader) {
	k.Version = ProtocolVersion(r.U16())
	k.CipherSuite = CipherSuiteID(r.U16())
	k.InitKey = r.Opaque()
	k.LeafNode.decode(r)
	k.Extensions.decode(r)
	k.Signature = r.Opaque()
}

func (k *KeyPackage) Ref(s *CipherSuite) []byte {
	return s.RefHash("MLS 1.0 KeyPackage Reference", Marshal(k))
}

func (k *KeyPackage) Verify(s *CipherSuite) error {
	if k.Version != MLS10 || k.CipherSuite != s.ID {
		return fmt.Errorf("mls: key package version %d suite %d", k.Version, k.CipherSuite)
	}
	if k.LeafNode.Source != SourceKeyPackage {
		return fmt.Errorf("mls: key package leaf source %d", k.LeafNode.Source)
	}
	if err := k.LeafNode.Verify(s, nil, 0); err != nil {
		return err
	}
	var w Writer
	k.encodeContent(&w)
	if !s.VerifyWithLabel(k.LeafNode.SignatureKey, "KeyPackageTBS", w.buf, k.Signature) {
		return fmt.Errorf("%w: key package", ErrInvalidSignature)
	}
	return nil
}

type KeyPackageSecrets struct {
	Package    KeyPackage
	Init       hpke.PrivateKey
	Encryption hpke.PrivateKey
	Signer     *SignaturePrivateKey
}

type LeafOptions struct {
	Credential   Credential
	Capabilities Capabilities
	Extensions   Extensions
}

func (o LeafOptions) capabilities(s *CipherSuite) Capabilities {
	c := o.Capabilities
	if len(c.Versions) == 0 {
		c.Versions = []ProtocolVersion{MLS10}
	}
	if len(c.CipherSuites) == 0 {
		c.CipherSuites = []CipherSuiteID{s.ID}
	}
	if len(c.Credentials) == 0 {
		c.Credentials = []CredentialType{o.Credential.Type}
	}
	return c
}

func NewKeyPackage(s *CipherSuite, signer *SignaturePrivateKey, opts LeafOptions) (*KeyPackageSecrets, error) {
	encryption, err := s.GenerateHPKEKey()
	if err != nil {
		return nil, err
	}
	init, err := s.GenerateHPKEKey()
	if err != nil {
		return nil, err
	}
	kp := KeyPackage{
		Version:     MLS10,
		CipherSuite: s.ID,
		InitKey:     init.PublicKey().Bytes(),
		LeafNode: LeafNode{
			EncryptionKey: encryption.PublicKey().Bytes(),
			SignatureKey:  signer.Public(),
			Credential:    opts.Credential,
			Capabilities:  opts.capabilities(s),
			Source:        SourceKeyPackage,
			Lifetime:      InfiniteLifetime(),
			Extensions:    opts.Extensions,
		},
	}
	if err := kp.LeafNode.Sign(s, signer, nil, 0); err != nil {
		return nil, err
	}
	var w Writer
	kp.encodeContent(&w)
	if kp.Signature, err = s.SignWithLabel(signer, "KeyPackageTBS", w.buf); err != nil {
		return nil, err
	}
	return &KeyPackageSecrets{Package: kp, Init: init, Encryption: encryption, Signer: signer}, nil
}
