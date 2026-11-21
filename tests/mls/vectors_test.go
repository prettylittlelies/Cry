package mls_test

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/prettylittlelies/cry/mls"
)

type hexBytes []byte

func (h *hexBytes) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	out, err := hex.DecodeString(s)
	*h = out
	return err
}

func load[T any](t *testing.T, name string) []T {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join("..", "fixtures", "mls", name))
	if err != nil {
		t.Fatal(err)
	}
	var out []T
	if err := json.Unmarshal(raw, &out); err != nil {
		t.Fatal(err)
	}
	return out
}

func suite(t *testing.T) *mls.CipherSuite {
	t.Helper()
	s, err := mls.Suite(mls.P256AES128GCMSHA256P256)
	if err != nil {
		t.Fatal(err)
	}
	return s
}

func equal(t *testing.T, what string, got, want []byte) {
	t.Helper()
	if !bytes.Equal(got, want) {
		t.Errorf("%s:\n got  %x\n want %x", what, got, want)
	}
}

func TestCryptoBasics(t *testing.T) {
	type vector struct {
		RefHash struct {
			Label string
			Value hexBytes
			Out   hexBytes
		} `json:"ref_hash"`
		ExpandWithLabel struct {
			Secret  hexBytes
			Label   string
			Context hexBytes
			Length  int
			Out     hexBytes
		} `json:"expand_with_label"`
		DeriveSecret struct {
			Secret hexBytes
			Label  string
			Out    hexBytes
		} `json:"derive_secret"`
		DeriveTreeSecret struct {
			Secret     hexBytes
			Label      string
			Generation uint32
			Length     int
			Out        hexBytes
		} `json:"derive_tree_secret"`
		SignWithLabel struct {
			Priv, Pub, Content, Signature hexBytes
			Label                         string
		} `json:"sign_with_label"`
		EncryptWithLabel struct {
			Priv, Pub, Context, Plaintext hexBytes
			KEMOutput                     hexBytes `json:"kem_output"`
			Ciphertext                    hexBytes
			Label                         string
		} `json:"encrypt_with_label"`
	}
	s := suite(t)
	for _, v := range load[vector](t, "crypto-basics.json") {
		equal(t, "ref_hash", s.RefHash(v.RefHash.Label, v.RefHash.Value), v.RefHash.Out)
		e := v.ExpandWithLabel
		equal(t, "expand_with_label", s.ExpandWithLabel(e.Secret, e.Label, e.Context, e.Length), e.Out)
		equal(t, "derive_secret", s.DeriveSecret(v.DeriveSecret.Secret, v.DeriveSecret.Label), v.DeriveSecret.Out)
		d := v.DeriveTreeSecret
		equal(t, "derive_tree_secret", s.DeriveTreeSecret(d.Secret, d.Label, d.Generation, d.Length), d.Out)

		sg := v.SignWithLabel
		if !s.VerifyWithLabel(sg.Pub, sg.Label, sg.Content, sg.Signature) {
			t.Error("vector signature does not verify")
		}
		key, err := s.ParseSignaturePrivateKey(sg.Priv)
		if err != nil {
			t.Fatal(err)
		}
		equal(t, "signature public key", key.Public(), sg.Pub)
		sig, err := s.SignWithLabel(key, sg.Label, sg.Content)
		if err != nil || !s.VerifyWithLabel(sg.Pub, sg.Label, sg.Content, sig) {
			t.Errorf("own signature: %v", err)
		}

		en := v.EncryptWithLabel
		priv, err := s.ParseHPKEPrivateKey(en.Priv)
		if err != nil {
			t.Fatal(err)
		}
		equal(t, "hpke public key", priv.PublicKey().Bytes(), en.Pub)
		pt, err := s.DecryptWithLabel(priv, en.Label, en.Context, mls.HPKECiphertext{KEMOutput: en.KEMOutput, Ciphertext: en.Ciphertext})
		if err != nil {
			t.Fatal(err)
		}
		equal(t, "decrypt_with_label", pt, en.Plaintext)
		ct, err := s.EncryptWithLabel(en.Pub, en.Label, en.Context, en.Plaintext)
		if err != nil {
			t.Fatal(err)
		}
		pt, err = s.DecryptWithLabel(priv, en.Label, en.Context, ct)
		if err != nil {
			t.Fatal(err)
		}
		equal(t, "round trip", pt, en.Plaintext)
	}
}

func TestVarintHeaders(t *testing.T) {
	type vector struct {
		Header hexBytes `json:"vlbytes_header"`
		Length int
	}
	for _, v := range load[vector](t, "deserialization.json") {
		var w mls.Writer
		w.Varint(v.Length)
		equal(t, "encode", w.Bytes(), v.Header)
		r := mls.NewReader(v.Header)
		if got := r.Varint(); got != v.Length || r.Err() != nil {
			t.Errorf("decode %x: %d %v", v.Header, got, r.Err())
		}
	}
	for _, bad := range [][]byte{{0x40, 0x01}, {0x80, 0x00, 0x00, 0x3f}, {0xc0}} {
		if r := mls.NewReader(bad); r.Varint() >= 0 && r.Err() == nil {
			t.Errorf("%x must be rejected", bad)
		}
	}
}

func TestTreeMath(t *testing.T) {
	type vector struct {
		Leaves  uint32 `json:"n_leaves"`
		Nodes   uint32 `json:"n_nodes"`
		Root    uint32
		Left    []*uint32
		Right   []*uint32
		Parent  []*uint32
		Sibling []*uint32
	}
	check := func(name string, i int, want *uint32, got mls.NodeIndex, defined bool) {
		t.Helper()
		if want == nil {
			if defined {
				t.Errorf("%s(%d) = %d, want none", name, i, got)
			}
			return
		}
		if !defined || uint32(got) != *want {
			t.Errorf("%s(%d) = %d, want %d", name, i, got, *want)
		}
	}
	for _, v := range load[vector](t, "tree-math.json") {
		if mls.NodeWidth(v.Leaves) != v.Nodes || uint32(mls.Root(v.Leaves)) != v.Root {
			t.Fatalf("width/root for %d leaves", v.Leaves)
		}
		root := mls.Root(v.Leaves)
		for i := range v.Nodes {
			n := mls.NodeIndex(i)
			check("left", int(i), v.Left[i], n.Left(), !n.IsLeaf())
			check("right", int(i), v.Right[i], n.Right(), !n.IsLeaf())
			check("parent", int(i), v.Parent[i], n.Parent(), n != root)
			check("sibling", int(i), v.Sibling[i], n.Sibling(), n != root)
		}
	}
	if got := mls.CommonAncestor(0, 6); got != 3 {
		t.Errorf("common ancestor 0,6 = %d", got)
	}
	if got := mls.CommonAncestor(1, 2); got != 1 {
		t.Errorf("common ancestor 1,2 = %d", got)
	}
}
