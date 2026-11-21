package mls

import "math/bits"

type LeafIndex uint32

type NodeIndex uint32

func (l LeafIndex) Node() NodeIndex { return NodeIndex(2 * l) }

func (n NodeIndex) IsLeaf() bool { return n%2 == 0 }

func (n NodeIndex) Leaf() LeafIndex { return LeafIndex(n / 2) }

func (n NodeIndex) Level() int { return bits.TrailingZeros32(^uint32(n)) }

func NodeWidth(leaves uint32) uint32 {
	if leaves == 0 {
		return 0
	}
	return 2*(leaves-1) + 1
}

func Root(leaves uint32) NodeIndex {
	w := NodeWidth(leaves)
	return NodeIndex(1<<(bits.Len32(w)-1) - 1)
}

func (n NodeIndex) Left() NodeIndex {
	k := n.Level()
	if k == 0 {
		return n
	}
	return n ^ (1 << (k - 1))
}

func (n NodeIndex) Right() NodeIndex {
	k := n.Level()
	if k == 0 {
		return n
	}
	return n ^ (3 << (k - 1))
}

func (n NodeIndex) Parent() NodeIndex {
	k := n.Level()
	b := uint32(n>>(k+1)) & 1
	return NodeIndex((uint32(n) | 1<<k) ^ b<<(k+1))
}

func (n NodeIndex) Sibling() NodeIndex {
	p := n.Parent()
	if n < p {
		return p.Right()
	}
	return p.Left()
}

func DirectPath(n NodeIndex, leaves uint32) []NodeIndex {
	root := Root(leaves)
	var out []NodeIndex
	for n != root {
		n = n.Parent()
		out = append(out, n)
	}
	return out
}

func Copath(n NodeIndex, leaves uint32) []NodeIndex {
	root := Root(leaves)
	if n == root {
		return nil
	}
	path := append([]NodeIndex{n}, DirectPath(n, leaves)...)
	path = path[:len(path)-1]
	out := make([]NodeIndex, len(path))
	for i, p := range path {
		out[i] = p.Sibling()
	}
	return out
}

func CommonAncestor(x, y NodeIndex) NodeIndex {
	lx, ly := x.Level()+1, y.Level()+1
	if lx <= ly && x>>ly == y>>ly {
		return y
	}
	if ly <= lx && x>>lx == y>>lx {
		return x
	}
	k := 0
	for x != y {
		x, y = x>>1, y>>1
		k++
	}
	return x<<k + (1<<(k-1) - 1)
}
