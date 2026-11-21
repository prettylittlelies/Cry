package mls_test

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/prettylittlelies/cry/mls"
)

type treekemVector struct {
	GroupID   hexBytes `json:"group_id"`
	Epoch     uint64
	Confirmed hexBytes `json:"confirmed_transcript_hash"`
	Tree      hexBytes `json:"ratchet_tree"`
	Leaves    []struct {
		Index      uint32
		Encryption hexBytes `json:"encryption_priv"`
		Signature  hexBytes `json:"signature_priv"`
		Paths      []struct {
			Node   uint32
			Secret hexBytes `json:"path_secret"`
		} `json:"path_secrets"`
	} `json:"leaves_private"`
	Updates []struct {
		Sender       uint32
		Path         hexBytes   `json:"update_path"`
		PathSecrets  []hexBytes `json:"path_secrets"`
		CommitSecret hexBytes   `json:"commit_secret"`
		TreeHash     hexBytes   `json:"tree_hash_after"`
	} `json:"update_paths"`
}

func (v *treekemVector) context(treeHash []byte) *mls.GroupContext {
	return &mls.GroupContext{Version: mls.MLS10, CipherSuite: mls.P256AES128GCMSHA256P256, GroupID: v.GroupID, Epoch: v.Epoch, TreeHash: treeHash, ConfirmedTranscriptHash: v.Confirmed}
}

func (v *treekemVector) secrets(t *testing.T, s *mls.CipherSuite, i int) *mls.TreeSecrets {
	t.Helper()
	l := v.Leaves[i]
	enc, err := s.ParseHPKEPrivateKey(l.Encryption)
	if err != nil {
		t.Fatal(err)
	}
	ts := mls.NewTreeSecrets(mls.LeafIndex(l.Index), enc)
	for _, p := range l.Paths {
		key, err := s.DeriveKeyPair(s.DeriveSecret(p.Secret, "node"))
		if err != nil {
			t.Fatal(err)
		}
		ts.Set(mls.NodeIndex(p.Node), key)
	}
	return ts
}

func (v *treekemVector) tree(t *testing.T) *mls.RatchetTree {
	t.Helper()
	var tree mls.RatchetTree
	if err := mls.Unmarshal(v.Tree, &tree); err != nil {
		t.Fatal(err)
	}
	return &tree
}

func TestTreeKEMDecap(t *testing.T) {
	s := suite(t)
	for vi, v := range load[treekemVector](t, "treekem.json") {
		for ui, u := range v.Updates {
			var path mls.UpdatePath
			if err := mls.Unmarshal(u.Path, &path); err != nil {
				t.Fatal(err)
			}
			for li, l := range v.Leaves {
				if l.Index == u.Sender {
					continue
				}
				tree := v.tree(t)
				commit, err := tree.Decap(s, v.secrets(t, s, li), mls.LeafIndex(u.Sender), &path, v.GroupID, v.context, nil)
				if err != nil {
					t.Fatalf("%d/%d leaf %d: %v", vi, ui, l.Index, err)
				}
				equal(t, "commit secret "+strconv.Itoa(vi), commit, u.CommitSecret)
				equal(t, "tree hash", tree.TreeHash(s), u.TreeHash)
			}
		}
	}
}

func TestTreeKEMEncapRoundTrip(t *testing.T) {
	s := suite(t)
	for vi, v := range load[treekemVector](t, "treekem.json") {
		for si, sender := range v.Leaves {
			tree := v.tree(t)
			signer, err := s.ParseSignaturePrivateKey(sender.Signature)
			if err != nil {
				t.Fatal(err)
			}
			leaf := tree.Leaf(mls.LeafIndex(sender.Index)).Clone()
			res, err := tree.Encap(s, v.secrets(t, s, si), leaf, signer, v.GroupID, v.context, nil)
			if err != nil {
				t.Fatalf("%d sender %d: %v", vi, sender.Index, err)
			}
			for ri, r := range v.Leaves {
				if ri == si {
					continue
				}
				other := v.tree(t)
				commit, err := other.Decap(s, v.secrets(t, s, ri), mls.LeafIndex(sender.Index), &res.Path, v.GroupID, v.context, nil)
				if err != nil {
					t.Fatalf("%d sender %d receiver %d: %v", vi, sender.Index, r.Index, err)
				}
				if !bytes.Equal(commit, res.CommitSecret) || !bytes.Equal(other.TreeHash(s), tree.TreeHash(s)) {
					t.Fatalf("%d sender %d receiver %d diverged", vi, sender.Index, r.Index)
				}
			}
			if err := tree.VerifyParentHashes(s); err != nil {
				t.Fatalf("%d sender %d: %v", vi, sender.Index, err)
			}
		}
	}
}
