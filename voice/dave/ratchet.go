package dave

import (
	"errors"
	"sync"

	"github.com/prettylittlelies/cry/mls"
)

var ErrErasedGeneration = errors.New("dave: key generation erased")

const maxGenerationGap = 250

type KeyRatchet struct {
	mu     sync.Mutex
	suite  *mls.CipherSuite
	secret []byte
	next   uint32
	keys   map[uint32][]byte
}

func NewKeyRatchet(suite *mls.CipherSuite, base []byte) *KeyRatchet {
	return &KeyRatchet{suite: suite, secret: base, keys: map[uint32][]byte{}}
}

func (r *KeyRatchet) Key(generation uint32) ([]byte, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if k, ok := r.keys[generation]; ok {
		return k, nil
	}
	if generation < r.next {
		return nil, ErrErasedGeneration
	}
	if generation-r.next > maxGenerationGap {
		return nil, errors.New("dave: key generation gap too large")
	}
	for r.next <= generation {
		g := r.next
		r.keys[g] = r.suite.DeriveTreeSecret(r.secret, "key", g, keySize)
		r.secret = r.suite.DeriveTreeSecret(r.secret, "secret", g, r.suite.SecretSize())
		r.next++
	}
	return r.keys[generation], nil
}

func (r *KeyRatchet) Erase(generation uint32) {
	r.mu.Lock()
	delete(r.keys, generation)
	r.mu.Unlock()
}

func wrappedGeneration(oldest, generation uint32) uint32 {
	const wrap = 1 << 8
	remainder := oldest % wrap
	factor := oldest / wrap
	if generation < remainder {
		factor++
	}
	return factor*wrap + generation
}
