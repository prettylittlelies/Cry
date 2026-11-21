package mls

import (
	"bytes"
	"crypto/hpke"
	"errors"
	"fmt"
	"maps"
	"slices"
)

var ErrPathMismatch = errors.New("mls: update path inconsistent with tree")

type TreeSecrets struct {
	Leaf LeafIndex
	keys map[NodeIndex]hpke.PrivateKey
}

func NewTreeSecrets(leaf LeafIndex, encryption hpke.PrivateKey) *TreeSecrets {
	return &TreeSecrets{Leaf: leaf, keys: map[NodeIndex]hpke.PrivateKey{leaf.Node(): encryption}}
}

func (ts *TreeSecrets) Set(n NodeIndex, key hpke.PrivateKey) { ts.keys[n] = key }

func (ts *TreeSecrets) Key(t *RatchetTree, n NodeIndex) (hpke.PrivateKey, bool) {
	key, ok := ts.keys[n]
	if !ok || t.Blank(n) || !bytes.Equal(key.PublicKey().Bytes(), t.node(n).EncryptionKey()) {
		return nil, false
	}
	return key, true
}

func (ts *TreeSecrets) Clone() *TreeSecrets {
	out := &TreeSecrets{Leaf: ts.Leaf, keys: make(map[NodeIndex]hpke.PrivateKey, len(ts.keys))}
	maps.Copy(out.keys, ts.keys)
	return out
}

func (ts *TreeSecrets) prune(t *RatchetTree) {
	for n := range ts.keys {
		if _, ok := ts.Key(t, n); !ok {
			delete(ts.keys, n)
		}
	}
}

func pathChild(parent, leaf NodeIndex) NodeIndex {
	if leaf < parent {
		return parent.Left()
	}
	return parent.Right()
}

func (t *RatchetTree) mergePath(s *CipherSuite, sender LeafIndex, keys [][]byte) ([]NodeIndex, []byte, error) {
	filtered := t.FilteredDirectPath(sender)
	if len(filtered) != len(keys) {
		return nil, nil, fmt.Errorf("%w: %d path nodes for %d filtered nodes", ErrPathMismatch, len(keys), len(filtered))
	}
	for _, p := range DirectPath(sender.Node(), t.Leaves()) {
		t.nodes[p] = nil
	}
	for i, p := range filtered {
		t.nodes[p] = &Node{Parent: &ParentNode{EncryptionKey: keys[i]}}
	}
	leaf := sender.Node()
	for i := range slices.Backward(filtered) {
		var hash []byte
		if i+1 < len(filtered) {
			above := filtered[i+1]
			hash = t.ParentHash(s, above, pathChild(above, leaf).Sibling())
		}
		t.Parent(filtered[i]).ParentHash = hash
	}
	if len(filtered) == 0 {
		return filtered, nil, nil
	}
	return filtered, t.ParentHash(s, filtered[0], pathChild(filtered[0], leaf).Sibling()), nil
}

type PathResult struct {
	Path         UpdatePath
	CommitSecret []byte
	PathSecrets  map[NodeIndex][]byte
	Secrets      *TreeSecrets
}

type GroupContextFunc func(treeHash []byte) *GroupContext

func (t *RatchetTree) Encap(s *CipherSuite, own *TreeSecrets, leaf LeafNode, signer *SignaturePrivateKey, groupID []byte, context GroupContextFunc, exclude []LeafIndex) (*PathResult, error) {
	sender := own.Leaf
	filtered := t.FilteredDirectPath(sender)
	secrets := make([][]byte, len(filtered))
	privs := make([]hpke.PrivateKey, len(filtered))
	keys := make([][]byte, len(filtered))
	secret := s.RandomSecret()
	for i := range filtered {
		if i > 0 {
			secret = s.DeriveSecret(secret, "path")
		}
		secrets[i] = secret
		priv, err := s.DeriveKeyPair(s.DeriveSecret(secret, "node"))
		if err != nil {
			return nil, err
		}
		privs[i], keys[i] = priv, priv.PublicKey().Bytes()
	}
	encryption, err := s.GenerateHPKEKey()
	if err != nil {
		return nil, err
	}
	_, parentHash, err := t.mergePath(s, sender, keys)
	if err != nil {
		return nil, err
	}
	leaf.EncryptionKey = encryption.PublicKey().Bytes()
	leaf.Source = SourceCommit
	leaf.ParentHash = parentHash
	leaf.Lifetime = Lifetime{}
	if err := leaf.Sign(s, signer, groupID, sender); err != nil {
		return nil, err
	}
	t.setLeaf(sender, &leaf)
	ctx := Marshal(context(t.TreeHash(s)))

	result := &PathResult{Path: UpdatePath{LeafNode: leaf}, PathSecrets: map[NodeIndex][]byte{}, Secrets: NewTreeSecrets(sender, encryption)}
	for i, p := range filtered {
		node := UpdatePathNode{EncryptionKey: keys[i]}
		for _, r := range t.Resolution(pathChild(p, sender.Node()).Sibling()) {
			if r.IsLeaf() && slices.Contains(exclude, r.Leaf()) {
				continue
			}
			ct, err := s.EncryptWithLabel(t.node(r).EncryptionKey(), "UpdatePathNode", ctx, secrets[i])
			if err != nil {
				return nil, err
			}
			node.EncryptedPathSecret = append(node.EncryptedPathSecret, ct)
		}
		result.Path.Nodes = append(result.Path.Nodes, node)
		result.PathSecrets[p] = secrets[i]
		result.Secrets.Set(p, privs[i])
	}
	if len(secrets) > 0 {
		result.CommitSecret = s.DeriveSecret(secrets[len(secrets)-1], "path")
	} else {
		result.CommitSecret = make([]byte, s.SecretSize())
	}
	return result, nil
}

func (t *RatchetTree) Decap(s *CipherSuite, own *TreeSecrets, sender LeafIndex, path *UpdatePath, groupID []byte, context GroupContextFunc, exclude []LeafIndex) ([]byte, error) {
	if sender == own.Leaf {
		return nil, fmt.Errorf("%w: cannot decap own path", ErrPathMismatch)
	}
	if path.LeafNode.Source != SourceCommit {
		return nil, fmt.Errorf("%w: leaf source %d", ErrPathMismatch, path.LeafNode.Source)
	}
	if err := path.LeafNode.Verify(s, groupID, sender); err != nil {
		return nil, err
	}
	keys := make([][]byte, len(path.Nodes))
	for i, n := range path.Nodes {
		keys[i] = n.EncryptionKey
	}
	original := t.Clone()
	filtered, parentHash, err := t.mergePath(s, sender, keys)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(parentHash, path.LeafNode.ParentHash) {
		return nil, fmt.Errorf("%w: leaf parent hash", ErrPathMismatch)
	}
	leaf := path.LeafNode
	t.setLeaf(sender, &leaf)
	ctx := Marshal(context(t.TreeHash(s)))

	ownNode := own.Leaf.Node()
	at := slices.IndexFunc(filtered, func(p NodeIndex) bool {
		return CommonAncestor(p, ownNode) == p
	})
	if at < 0 {
		return nil, fmt.Errorf("%w: receiver not below sender path", ErrPathMismatch)
	}
	resolution := slices.DeleteFunc(original.Resolution(pathChild(filtered[at], sender.Node()).Sibling()), func(r NodeIndex) bool {
		return r.IsLeaf() && slices.Contains(exclude, r.Leaf())
	})
	if len(resolution) != len(path.Nodes[at].EncryptedPathSecret) {
		return nil, fmt.Errorf("%w: %d ciphertexts for resolution of %d", ErrPathMismatch, len(path.Nodes[at].EncryptedPathSecret), len(resolution))
	}
	var secret []byte
	for i, r := range resolution {
		key, ok := own.Key(original, r)
		if !ok {
			continue
		}
		if secret, err = s.DecryptWithLabel(key, "UpdatePathNode", ctx, path.Nodes[at].EncryptedPathSecret[i]); err != nil {
			return nil, err
		}
		break
	}
	if secret == nil {
		return nil, fmt.Errorf("%w: no private key in resolution", ErrPathMismatch)
	}
	return t.absorbPath(s, own, filtered[at:], secret)
}

func (t *RatchetTree) absorbPath(s *CipherSuite, own *TreeSecrets, nodes []NodeIndex, secret []byte) ([]byte, error) {
	for i, p := range nodes {
		if i > 0 {
			secret = s.DeriveSecret(secret, "path")
		}
		priv, err := s.DeriveKeyPair(s.DeriveSecret(secret, "node"))
		if err != nil {
			return nil, err
		}
		if !bytes.Equal(priv.PublicKey().Bytes(), t.node(p).EncryptionKey()) {
			return nil, fmt.Errorf("%w: derived key mismatch at node %d", ErrPathMismatch, p)
		}
		own.Set(p, priv)
	}
	own.prune(t)
	return s.DeriveSecret(secret, "path"), nil
}

func (t *RatchetTree) ImportPathSecret(s *CipherSuite, own *TreeSecrets, from NodeIndex, secret []byte) ([]byte, error) {
	path := DirectPath(own.Leaf.Node(), t.Leaves())
	at := slices.Index(path, from)
	if at < 0 {
		return nil, fmt.Errorf("%w: node %d not on direct path", ErrPathMismatch, from)
	}
	var nodes []NodeIndex
	for _, p := range path[at:] {
		if !t.Blank(p) {
			nodes = append(nodes, p)
		}
	}
	return t.absorbPath(s, own, nodes, secret)
}
