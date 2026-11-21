package mls

import "crypto/hpke"

type EpochSecrets struct {
	SenderData    []byte
	Encryption    []byte
	Exporter      []byte
	External      []byte
	Confirmation  []byte
	Membership    []byte
	Resumption    []byte
	Authenticator []byte
	Init          []byte
}

func (s *CipherSuite) JoinerSecret(initSecret, commitSecret []byte, ctx *GroupContext) []byte {
	return s.ExpandWithLabel(s.Extract(initSecret, commitSecret), "joiner", Marshal(ctx), s.SecretSize())
}

func (s *CipherSuite) WelcomeSecret(joiner, psk []byte) []byte {
	return s.DeriveSecret(s.Extract(joiner, psk), "welcome")
}

func (s *CipherSuite) EpochSecrets(joiner, psk []byte, ctx *GroupContext) EpochSecrets {
	epoch := s.ExpandWithLabel(s.Extract(joiner, psk), "epoch", Marshal(ctx), s.SecretSize())
	return EpochSecrets{
		SenderData:    s.DeriveSecret(epoch, "sender data"),
		Encryption:    s.DeriveSecret(epoch, "encryption"),
		Exporter:      s.DeriveSecret(epoch, "exporter"),
		External:      s.DeriveSecret(epoch, "external"),
		Confirmation:  s.DeriveSecret(epoch, "confirm"),
		Membership:    s.DeriveSecret(epoch, "membership"),
		Resumption:    s.DeriveSecret(epoch, "resumption"),
		Authenticator: s.DeriveSecret(epoch, "authentication"),
		Init:          s.DeriveSecret(epoch, "init"),
	}
}

func (s *CipherSuite) Export(exporter []byte, label string, context []byte, length int) []byte {
	return s.ExpandWithLabel(s.DeriveSecret(exporter, label), "exported", s.Hash(context), length)
}

func (s *CipherSuite) ExternalKey(external []byte) (hpke.PrivateKey, error) {
	return s.DeriveKeyPair(external)
}

func (s *CipherSuite) welcomeKeyNonce(welcome []byte) (key, nonce []byte) {
	return s.ExpandWithLabel(welcome, "key", nil, s.KeySize()), s.ExpandWithLabel(welcome, "nonce", nil, s.NonceSize())
}

func (s *CipherSuite) ConfirmedTranscriptHash(interim []byte, content *AuthenticatedContent) []byte {
	return s.Hash(append(append([]byte(nil), interim...), content.confirmedInput()...))
}

func (s *CipherSuite) InterimTranscriptHash(confirmed, confirmationTag []byte) []byte {
	var w Writer
	w.Raw(confirmed)
	w.Opaque(confirmationTag)
	return s.Hash(w.buf)
}

func (a *AuthenticatedContent) decode(r *Reader) {
	a.WireFormat = WireFormat(r.U16())
	a.Content.decode(r)
	a.Auth.decode(r, a.Content.Type)
}
