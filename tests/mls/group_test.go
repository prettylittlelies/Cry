package mls_test

import (
	"testing"

	"github.com/prettylittlelies/cry/mls"
)

func message(t *testing.T, raw []byte) *mls.Message {
	t.Helper()
	var m mls.Message
	if err := mls.Unmarshal(raw, &m); err != nil {
		t.Fatal(err)
	}
	return &m
}

func keyPackage(t *testing.T, s *mls.CipherSuite, raw, initPriv, encPriv, sigPriv []byte) *mls.KeyPackageSecrets {
	t.Helper()
	m := message(t, raw)
	init, err := s.ParseHPKEPrivateKey(initPriv)
	if err != nil {
		t.Fatal(err)
	}
	out := &mls.KeyPackageSecrets{Package: *m.KeyPackage, Init: init}
	if encPriv != nil {
		if out.Encryption, err = s.ParseHPKEPrivateKey(encPriv); err != nil {
			t.Fatal(err)
		}
	}
	if sigPriv != nil {
		if out.Signer, err = s.ParseSignaturePrivateKey(sigPriv); err != nil {
			t.Fatal(err)
		}
	}
	return out
}

func TestPSKSecret(t *testing.T) {
	type vector struct {
		PSKs []struct {
			ID    hexBytes `json:"psk_id"`
			PSK   hexBytes
			Nonce hexBytes `json:"psk_nonce"`
		}
		Secret hexBytes `json:"psk_secret"`
	}
	s := suite(t)
	for _, v := range load[vector](t, "psk_secret.json") {
		ids := make([]mls.PreSharedKeyID, len(v.PSKs))
		values := make([][]byte, len(v.PSKs))
		for i, p := range v.PSKs {
			ids[i] = mls.PreSharedKeyID{Type: mls.PSKExternal, ID: p.ID, Nonce: p.Nonce}
			values[i] = p.PSK
		}
		equal(t, "psk secret", s.PSKSecret(ids, values), v.Secret)
	}
}

func TestWelcomeVector(t *testing.T) {
	type vector struct {
		InitPriv   hexBytes `json:"init_priv"`
		KeyPackage hexBytes `json:"key_package"`
		SignerPub  hexBytes `json:"signer_pub"`
		Welcome    hexBytes
	}
	s := suite(t)
	for _, v := range load[vector](t, "welcome.json") {
		kp := keyPackage(t, s, v.KeyPackage, v.InitPriv, nil, nil)
		secrets, info, psk, err := mls.OpenWelcome(s, message(t, v.Welcome).Welcome, kp, nil)
		if err != nil {
			t.Fatal(err)
		}
		if err := info.Verify(s, v.SignerPub); err != nil {
			t.Fatal(err)
		}
		ctx := info.GroupContext
		epoch := s.EpochSecrets(secrets.JoinerSecret, psk, &ctx)
		equal(t, "confirmation tag", s.MAC(epoch.Confirmation, ctx.ConfirmedTranscriptHash), info.ConfirmationTag)
	}
}

type passiveVector struct {
	PSKs []struct {
		ID  hexBytes `json:"psk_id"`
		PSK hexBytes
	} `json:"external_psks"`
	KeyPackage    hexBytes `json:"key_package"`
	SignaturePriv hexBytes `json:"signature_priv"`
	EncPriv       hexBytes `json:"encryption_priv"`
	InitPriv      hexBytes `json:"init_priv"`
	Welcome       hexBytes
	Tree          *hexBytes `json:"ratchet_tree"`
	InitialAuth   hexBytes  `json:"initial_epoch_authenticator"`
	Epochs        []struct {
		Proposals     []hexBytes
		Commit        hexBytes
		Authenticator hexBytes `json:"epoch_authenticator"`
	}
}

func runPassive(t *testing.T, file string) {
	s := suite(t)
	for vi, v := range load[passiveVector](t, file) {
		kp := keyPackage(t, s, v.KeyPackage, v.InitPriv, v.EncPriv, v.SignaturePriv)
		opts := mls.JoinOptions{PSKs: mls.PSKStore{}}
		for _, p := range v.PSKs {
			opts.PSKs[string(p.ID)] = p.PSK
		}
		if v.Tree != nil {
			opts.Tree = new(mls.RatchetTree)
			if err := mls.Unmarshal(*v.Tree, opts.Tree); err != nil {
				t.Fatal(err)
			}
		}
		g, err := mls.Join(s, message(t, v.Welcome).Welcome, kp, opts)
		if err != nil {
			t.Fatalf("%s %d: join: %v", file, vi, err)
		}
		equal(t, "initial authenticator", g.EpochAuthenticator(), v.InitialAuth)
		for ei, e := range v.Epochs {
			for _, raw := range e.Proposals {
				ac, err := g.Verify(message(t, raw))
				if err != nil {
					t.Fatalf("%s %d epoch %d: proposal: %v", file, vi, ei, err)
				}
				if _, err := g.CacheProposal(ac); err != nil {
					t.Fatalf("%s %d epoch %d: cache: %v", file, vi, ei, err)
				}
			}
			ac, err := g.Verify(message(t, e.Commit))
			if err != nil {
				t.Fatalf("%s %d epoch %d: commit: %v", file, vi, ei, err)
			}
			if g, err = g.HandleCommit(ac); err != nil {
				t.Fatalf("%s %d epoch %d: handle: %v", file, vi, ei, err)
			}
			equal(t, "epoch authenticator", g.EpochAuthenticator(), e.Authenticator)
		}
	}
}

func TestPassiveClientWelcome(t *testing.T) { runPassive(t, "passive-client-welcome.json") }

func TestPassiveClientCommits(t *testing.T) { runPassive(t, "passive-client-handling-commit.json") }
