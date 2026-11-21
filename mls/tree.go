package mls

import (
	"bytes"
	"errors"
	"fmt"
	"slices"
)

var ErrInvalidTree = errors.New("mls: invalid ratchet tree")

type NodeType uint8

const (
	NodeLeaf   NodeType = 1
	NodeParent NodeType = 2
)

type Node struct {
	Leaf   *LeafNode
	Parent *ParentNode
}

func (n *Node) blank() bool { return n == nil || n.Leaf == nil && n.Parent == nil }

func (n *Node) EncryptionKey() []byte {
	switch {
	case n.Leaf != nil:
		return n.Leaf.EncryptionKey
	case n.Parent != nil:
		return n.Parent.EncryptionKey
	}
	return nil
}

func (n *Node) parentHash() []byte {
	switch {
	case n.Leaf != nil:
		return n.Leaf.ParentHash
	case n.Parent != nil:
		return n.Parent.ParentHash
	}
	return nil
}

func (n *Node) encode(w *Writer) {
	if n.Leaf != nil {
		w.U8(uint8(NodeLeaf))
		n.Leaf.encode(w)
		return
	}
	w.U8(uint8(NodeParent))
	n.Parent.encode(w)
}

func (n *Node) decode(r *Reader) {
	switch t := NodeType(r.U8()); t {
	case NodeLeaf:
		n.Leaf = new(LeafNode)
		n.Leaf.decode(r)
	case NodeParent:
		n.Parent = new(ParentNode)
		n.Parent.decode(r)
	default:
		r.Fail(fmt.Errorf("mls: invalid node type %d", t))
	}
}

type RatchetTree struct {
	nodes []*Node
}

func (t *RatchetTree) encode(w *Writer) {
	end := len(t.nodes)
	for end > 0 && t.nodes[end-1].blank() {
		end--
	}
	w.Vector(func(w *Writer) {
		for _, n := range t.nodes[:end] {
			w.Optional(!n.blank(), func(w *Writer) { n.encode(w) })
		}
	})
}

func (t *RatchetTree) decode(r *Reader) {
	t.nodes = nil
	r.Vector(func(r *Reader) {
		var n *Node
		r.Optional(func(r *Reader) {
			n = new(Node)
			n.decode(r)
		})
		t.nodes = append(t.nodes, n)
	})
	if r.Err() != nil {
		return
	}
	if len(t.nodes)%2 == 0 || t.nodes[len(t.nodes)-1].blank() {
		r.Fail(fmt.Errorf("%w: malformed node list", ErrInvalidTree))
		return
	}
	leaves := uint32(1)
	for NodeWidth(leaves) < uint32(len(t.nodes)) {
		leaves *= 2
	}
	for i, n := range t.nodes {
		if n.blank() {
			continue
		}
		if NodeIndex(i).IsLeaf() != (n.Leaf != nil) {
			r.Fail(fmt.Errorf("%w: node %d has wrong type", ErrInvalidTree, i))
			return
		}
	}
	t.nodes = append(t.nodes, make([]*Node, int(NodeWidth(leaves))-len(t.nodes))...)
}

func NewRatchetTree(first LeafNode) *RatchetTree {
	return &RatchetTree{nodes: []*Node{{Leaf: &first}}}
}

func (t *RatchetTree) Clone() *RatchetTree {
	var out RatchetTree
	if err := Unmarshal(Marshal(t), &out); err != nil {
		panic(err)
	}
	if len(out.nodes) < len(t.nodes) {
		out.nodes = append(out.nodes, make([]*Node, len(t.nodes)-len(out.nodes))...)
	}
	return &out
}

func (t *RatchetTree) Leaves() uint32 { return uint32(len(t.nodes)+1) / 2 }

func (t *RatchetTree) node(i NodeIndex) *Node {
	if int(i) >= len(t.nodes) {
		return nil
	}
	return t.nodes[i]
}

func (t *RatchetTree) Leaf(l LeafIndex) *LeafNode {
	if n := t.node(l.Node()); n != nil {
		return n.Leaf
	}
	return nil
}

func (t *RatchetTree) Parent(i NodeIndex) *ParentNode {
	if n := t.node(i); n != nil {
		return n.Parent
	}
	return nil
}

func (t *RatchetTree) Blank(i NodeIndex) bool { return t.node(i).blank() }

func (t *RatchetTree) Members(yield func(LeafIndex, *LeafNode) bool) {
	for l := range LeafIndex(t.Leaves()) {
		if leaf := t.Leaf(l); leaf != nil && !yield(l, leaf) {
			return
		}
	}
}

func (t *RatchetTree) setLeaf(l LeafIndex, leaf *LeafNode) {
	t.nodes[l.Node()] = &Node{Leaf: leaf}
}

func (t *RatchetTree) blankPath(l LeafIndex) {
	t.nodes[l.Node()] = nil
	for _, p := range DirectPath(l.Node(), t.Leaves()) {
		t.nodes[p] = nil
	}
}

func (t *RatchetTree) AddLeaf(leaf LeafNode) LeafIndex {
	l := LeafIndex(0)
	for ; l < LeafIndex(t.Leaves()); l++ {
		if t.Blank(l.Node()) {
			break
		}
	}
	if l == LeafIndex(t.Leaves()) {
		t.nodes = append(t.nodes, make([]*Node, len(t.nodes)+1)...)
	}
	t.setLeaf(l, &leaf)
	for _, p := range DirectPath(l.Node(), t.Leaves()) {
		if parent := t.Parent(p); parent != nil {
			parent.UnmergedLeaves = append(parent.UnmergedLeaves, l)
		}
	}
	return l
}

func (t *RatchetTree) UpdateLeaf(l LeafIndex, leaf LeafNode) {
	t.blankPath(l)
	t.setLeaf(l, &leaf)
}

func (t *RatchetTree) RemoveLeaf(l LeafIndex) {
	t.blankPath(l)
	for t.Leaves() > 1 {
		half := len(t.nodes) / 2
		if !slices.ContainsFunc(t.nodes[half:], func(n *Node) bool { return !n.blank() }) {
			t.nodes = t.nodes[:half]
			continue
		}
		break
	}
}

func (t *RatchetTree) Resolution(i NodeIndex) []NodeIndex {
	n := t.node(i)
	switch {
	case !n.blank():
		out := []NodeIndex{i}
		if n.Parent != nil {
			for _, l := range n.Parent.UnmergedLeaves {
				out = append(out, l.Node())
			}
		}
		return out
	case i.IsLeaf():
		return nil
	}
	return append(t.Resolution(i.Left()), t.Resolution(i.Right())...)
}

func (t *RatchetTree) FilteredDirectPath(l LeafIndex) []NodeIndex {
	var out []NodeIndex
	child := l.Node()
	for _, p := range DirectPath(child, t.Leaves()) {
		if len(t.Resolution(child.Sibling())) > 0 {
			out = append(out, p)
		}
		child = p
	}
	return out
}

func (t *RatchetTree) Find(match func(*LeafNode) bool) (LeafIndex, bool) {
	for l, leaf := range t.Members {
		if match(leaf) {
			return l, true
		}
	}
	return 0, false
}

func (t *RatchetTree) TreeHash(s *CipherSuite) []byte {
	return t.hash(s, Root(t.Leaves()))
}

func (t *RatchetTree) NodeHash(s *CipherSuite, i NodeIndex) []byte {
	return t.hash(s, i)
}

func (t *RatchetTree) hash(s *CipherSuite, i NodeIndex) []byte {
	var w Writer
	n := t.node(i)
	if i.IsLeaf() {
		w.U8(uint8(NodeLeaf))
		w.U32(uint32(i.Leaf()))
		w.Optional(!n.blank(), func(w *Writer) { n.Leaf.encode(w) })
	} else {
		w.U8(uint8(NodeParent))
		w.Optional(!n.blank(), func(w *Writer) { n.Parent.encode(w) })
		w.Opaque(t.hash(s, i.Left()))
		w.Opaque(t.hash(s, i.Right()))
	}
	return s.Hash(w.buf)
}

func (t *RatchetTree) ParentHash(s *CipherSuite, parent NodeIndex, sibling NodeIndex) []byte {
	p := t.Parent(parent)
	pruned := t.Clone()
	for _, l := range p.UnmergedLeaves {
		pruned.nodes[l.Node()] = nil
		for _, a := range DirectPath(l.Node(), pruned.Leaves()) {
			if ap := pruned.Parent(a); ap != nil {
				ap.UnmergedLeaves = slices.DeleteFunc(ap.UnmergedLeaves, func(x LeafIndex) bool { return x == l })
			}
		}
	}
	var w Writer
	w.Opaque(p.EncryptionKey)
	w.Opaque(p.ParentHash)
	w.Opaque(pruned.hash(s, sibling))
	return s.Hash(w.buf)
}

func (t *RatchetTree) VerifyParentHashes(s *CipherSuite) error {
	for i := range NodeIndex(len(t.nodes)) {
		p := t.Parent(i)
		if i.IsLeaf() || p == nil {
			continue
		}
		valid := false
		for _, side := range [2][2]NodeIndex{{i.Left(), i.Right()}, {i.Right(), i.Left()}} {
			want := t.ParentHash(s, i, side[1])
			for _, d := range t.Resolution(side[0]) {
				if d.IsLeaf() && slices.Contains(p.UnmergedLeaves, d.Leaf()) {
					continue
				}
				if bytes.Equal(t.node(d).parentHash(), want) {
					valid = true
				}
			}
		}
		if !valid {
			return fmt.Errorf("%w: parent hash of node %d", ErrInvalidTree, i)
		}
	}
	return nil
}

func (t *RatchetTree) VerifyLeaves(s *CipherSuite, groupID []byte) error {
	seenSig := map[string]bool{}
	seenEnc := map[string]bool{}
	for l, leaf := range t.Members {
		if err := leaf.Verify(s, groupID, l); err != nil {
			return fmt.Errorf("leaf %d: %w", l, err)
		}
		if seenSig[string(leaf.SignatureKey)] || seenEnc[string(leaf.EncryptionKey)] {
			return fmt.Errorf("%w: duplicate key at leaf %d", ErrInvalidTree, l)
		}
		seenSig[string(leaf.SignatureKey)], seenEnc[string(leaf.EncryptionKey)] = true, true
	}
	for i, n := range t.nodes {
		if n == nil || n.Parent == nil {
			continue
		}
		for _, u := range n.Parent.UnmergedLeaves {
			if t.Leaf(u) == nil || !slices.Contains(DirectPath(u.Node(), t.Leaves()), NodeIndex(i)) {
				return fmt.Errorf("%w: unmerged leaf %d at node %d", ErrInvalidTree, u, i)
			}
		}
	}
	return nil
}
