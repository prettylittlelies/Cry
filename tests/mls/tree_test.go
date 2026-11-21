package mls_test

import (
	"slices"
	"testing"

	"github.com/prettylittlelies/cry/mls"
)

func TestTreeValidation(t *testing.T) {
	type vector struct {
		Tree        hexBytes
		GroupID     hexBytes   `json:"group_id"`
		TreeHashes  []hexBytes `json:"tree_hashes"`
		Resolutions [][]uint32
	}
	s := suite(t)
	for i, v := range load[vector](t, "tree-validation.json") {
		var tree mls.RatchetTree
		if err := mls.Unmarshal(v.Tree, &tree); err != nil {
			t.Fatalf("%d: %v", i, err)
		}
		equal(t, "re-encode", mls.Marshal(&tree), v.Tree)
		for n, want := range v.TreeHashes {
			equal(t, "node hash", tree.NodeHash(s, mls.NodeIndex(n)), want)
		}
		for n, want := range v.Resolutions {
			got := tree.Resolution(mls.NodeIndex(n))
			if !slices.Equal(got, toNodes(want)) {
				t.Errorf("%d: resolution(%d) = %v, want %v", i, n, got, want)
			}
		}
		if err := tree.VerifyParentHashes(s); err != nil {
			t.Errorf("%d: %v", i, err)
		}
		if err := tree.VerifyLeaves(s, v.GroupID); err != nil {
			t.Errorf("%d: %v", i, err)
		}
	}
}

func toNodes(in []uint32) []mls.NodeIndex {
	out := make([]mls.NodeIndex, len(in))
	for i, v := range in {
		out[i] = mls.NodeIndex(v)
	}
	return out
}

func TestTreeOperations(t *testing.T) {
	type vector struct {
		TreeBefore     hexBytes `json:"tree_before"`
		TreeAfter      hexBytes `json:"tree_after"`
		Proposal       hexBytes
		ProposalSender uint32   `json:"proposal_sender"`
		TreeHashBefore hexBytes `json:"tree_hash_before"`
		TreeHashAfter  hexBytes `json:"tree_hash_after"`
	}
	s := suite(t)
	for i, v := range load[vector](t, "tree-operations.json") {
		var tree mls.RatchetTree
		if err := mls.Unmarshal(v.TreeBefore, &tree); err != nil {
			t.Fatalf("%d: %v", i, err)
		}
		equal(t, "hash before", tree.TreeHash(s), v.TreeHashBefore)
		var p mls.Proposal
		if err := mls.Unmarshal(v.Proposal, &p); err != nil {
			t.Fatalf("%d: %v", i, err)
		}
		switch p.Type {
		case mls.ProposalAdd:
			tree.AddLeaf(p.Add.LeafNode)
		case mls.ProposalUpdate:
			tree.UpdateLeaf(mls.LeafIndex(v.ProposalSender), *p.Update)
		case mls.ProposalRemove:
			tree.RemoveLeaf(p.Remove)
		}
		equal(t, "tree after", mls.Marshal(&tree), v.TreeAfter)
		equal(t, "hash after", tree.TreeHash(s), v.TreeHashAfter)
	}
}
