package mls

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/hkdf"
	"crypto/hmac"
	"crypto/hpke"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
)

type CipherSuiteID uint16

const P256AES128GCMSHA256P256 CipherSuiteID = 2

var ErrUnsupportedSuite = errors.New("mls: unsupported cipher suite")

const labelPrefix = "MLS 1.0 "

type CipherSuite struct {
	ID        CipherSuiteID
	hash      crypto.Hash
	kem       hpke.KEM
	kdf       hpke.KDF
	aead      hpke.AEAD
	curve     elliptic.Curve
	keySize   int
	nonceSize int
}

var suites = map[CipherSuiteID]*CipherSuite{
	P256AES128GCMSHA256P256: {
		ID:        P256AES128GCMSHA256P256,
		hash:      crypto.SHA256,
		kem:       hpke.DHKEM(ecdh.P256()),
		kdf:       hpke.HKDFSHA256(),
		aead:      hpke.AES128GCM(),
		curve:     elliptic.P256(),
		keySize:   16,
		nonceSize: 12,
	},
}

func Suite(id CipherSuiteID) (*CipherSuite, error) {
	if s, ok := suites[id]; ok {
		return s, nil
	}
	return nil, fmt.Errorf("%w: %d", ErrUnsupportedSuite, id)
}

func (s *CipherSuite) SecretSize() int { return s.hash.Size() }
func (s *CipherSuite) KeySize() int    { return s.keySize }
func (s *CipherSuite) NonceSize() int  { return s.nonceSize }

func (s *CipherSuite) Hash(data []byte) []byte {
	h := s.hash.New()
	h.Write(data)
	return h.Sum(nil)
}

func (s *CipherSuite) MAC(key, data []byte) []byte {
	m := hmac.New(s.hash.New, key)
	m.Write(data)
	return m.Sum(nil)
}

func (s *CipherSuite) Extract(salt, ikm []byte) []byte {
	if len(salt) == 0 {
		salt = make([]byte, s.SecretSize())
	}
	if len(ikm) == 0 {
		ikm = make([]byte, s.SecretSize())
	}
	out, err := hkdf.Extract(s.hash.New, ikm, salt)
	if err != nil {
		panic(err)
	}
	return out
}

func (s *CipherSuite) Expand(prk, info []byte, length int) []byte {
	out, err := hkdf.Expand(s.hash.New, prk, string(info), length)
	if err != nil {
		panic(err)
	}
	return out
}

func (s *CipherSuite) ExpandWithLabel(secret []byte, label string, context []byte, length int) []byte {
	var w Writer
	w.U16(uint16(length))
	w.Opaque([]byte(labelPrefix + label))
	w.Opaque(context)
	return s.Expand(secret, w.buf, length)
}

func (s *CipherSuite) DeriveSecret(secret []byte, label string) []byte {
	return s.ExpandWithLabel(secret, label, nil, s.SecretSize())
}

func (s *CipherSuite) DeriveTreeSecret(secret []byte, label string, generation uint32, length int) []byte {
	var w Writer
	w.U32(generation)
	return s.ExpandWithLabel(secret, label, w.buf, length)
}

func (s *CipherSuite) RefHash(label string, value []byte) []byte {
	var w Writer
	w.Opaque([]byte(label))
	w.Opaque(value)
	return s.Hash(w.buf)
}

func signContent(label string, content []byte) []byte {
	var w Writer
	w.Opaque([]byte(labelPrefix + label))
	w.Opaque(content)
	return w.buf
}

type SignaturePrivateKey struct {
	key *ecdsa.PrivateKey
}

func (s *CipherSuite) GenerateSignatureKey() (*SignaturePrivateKey, error) {
	k, err := ecdsa.GenerateKey(s.curve, rand.Reader)
	if err != nil {
		return nil, err
	}
	return &SignaturePrivateKey{k}, nil
}

func (s *CipherSuite) ParseSignaturePrivateKey(raw []byte) (*SignaturePrivateKey, error) {
	k, err := ecdsa.ParseRawPrivateKey(s.curve, raw)
	if err != nil {
		return nil, err
	}
	return &SignaturePrivateKey{k}, nil
}

func (k *SignaturePrivateKey) Public() []byte {
	b, err := k.key.PublicKey.Bytes()
	if err != nil {
		panic(err)
	}
	return b
}

func (k *SignaturePrivateKey) Bytes() []byte {
	b, err := k.key.Bytes()
	if err != nil {
		panic(err)
	}
	return b
}

func (s *CipherSuite) SignWithLabel(key *SignaturePrivateKey, label string, content []byte) ([]byte, error) {
	digest := sha256.Sum256(signContent(label, content))
	return ecdsa.SignASN1(rand.Reader, key.key, digest[:])
}

func (s *CipherSuite) VerifyWithLabel(public []byte, label string, content, signature []byte) bool {
	pub, err := ecdsa.ParseUncompressedPublicKey(s.curve, public)
	if err != nil {
		return false
	}
	digest := sha256.Sum256(signContent(label, content))
	return ecdsa.VerifyASN1(pub, digest[:], signature)
}

type HPKECiphertext struct {
	KEMOutput  []byte
	Ciphertext []byte
}

func (c *HPKECiphertext) encode(w *Writer) {
	w.Opaque(c.KEMOutput)
	w.Opaque(c.Ciphertext)
}

func (c *HPKECiphertext) decode(r *Reader) {
	c.KEMOutput = r.Opaque()
	c.Ciphertext = r.Opaque()
}

func encryptContext(label string, context []byte) []byte {
	var w Writer
	w.Opaque([]byte(labelPrefix + label))
	w.Opaque(context)
	return w.buf
}

func (s *CipherSuite) EncryptWithLabel(public []byte, label string, context, plaintext []byte) (HPKECiphertext, error) {
	pk, err := s.kem.NewPublicKey(public)
	if err != nil {
		return HPKECiphertext{}, err
	}
	enc, sender, err := hpke.NewSender(pk, s.kdf, s.aead, encryptContext(label, context))
	if err != nil {
		return HPKECiphertext{}, err
	}
	ct, err := sender.Seal(nil, plaintext)
	if err != nil {
		return HPKECiphertext{}, err
	}
	return HPKECiphertext{KEMOutput: enc, Ciphertext: ct}, nil
}

func (s *CipherSuite) DecryptWithLabel(private hpke.PrivateKey, label string, context []byte, ct HPKECiphertext) ([]byte, error) {
	recipient, err := hpke.NewRecipient(ct.KEMOutput, private, s.kdf, s.aead, encryptContext(label, context))
	if err != nil {
		return nil, err
	}
	return recipient.Open(nil, ct.Ciphertext)
}

func (s *CipherSuite) GenerateHPKEKey() (hpke.PrivateKey, error) {
	return s.kem.GenerateKey()
}

func (s *CipherSuite) ParseHPKEPrivateKey(raw []byte) (hpke.PrivateKey, error) {
	return s.kem.NewPrivateKey(raw)
}

func (s *CipherSuite) DeriveKeyPair(secret []byte) (hpke.PrivateKey, error) {
	return s.kem.DeriveKeyPair(secret)
}

func (s *CipherSuite) aeadCipher(key []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(block)
}

func (s *CipherSuite) Seal(key, nonce, aad, plaintext []byte) ([]byte, error) {
	a, err := s.aeadCipher(key)
	if err != nil {
		return nil, err
	}
	return a.Seal(nil, nonce, plaintext, aad), nil
}

func (s *CipherSuite) Open(key, nonce, aad, ciphertext []byte) ([]byte, error) {
	a, err := s.aeadCipher(key)
	if err != nil {
		return nil, err
	}
	return a.Open(nil, nonce, ciphertext, aad)
}

func (s *CipherSuite) RandomSecret() []byte {
	b := make([]byte, s.SecretSize())
	rand.Read(b)
	return b
}
