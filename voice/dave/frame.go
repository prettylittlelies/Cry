package dave

import (
	"bytes"
	"errors"
	"sync"
	"time"
)

const (
	markerByte       = 0xFA
	footerFixed      = truncatedTag + 1 + 2
	previousKeyGrace = 10 * time.Second
)

var (
	ErrNotProtocolFrame = errors.New("dave: not a protocol frame")
	ErrNoKey            = errors.New("dave: no key for sender")
)

var silence = []byte{0xF8, 0xFF, 0xFE}

func IsSilence(frame []byte) bool { return bytes.Equal(frame, silence) }

type Encryptor struct {
	mu         sync.Mutex
	ratchet    *KeyRatchet
	nonce      uint32
	generation uint32
	cipher     *frameCipher
}

func (e *Encryptor) SetRatchet(r *KeyRatchet) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if r == e.ratchet {
		return
	}
	e.ratchet, e.nonce, e.generation, e.cipher = r, 0, 0, nil
}

func (e *Encryptor) Active() bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.ratchet != nil
}

func (e *Encryptor) Encrypt(frame []byte) ([]byte, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.ratchet == nil {
		return nil, ErrNoKey
	}
	e.nonce++
	generation := wrappedGeneration(e.generation, e.nonce>>generationShift)
	if generation != e.generation || e.cipher == nil {
		key, err := e.ratchet.Key(generation)
		if err != nil {
			return nil, err
		}
		if e.cipher, err = newFrameCipher(key); err != nil {
			return nil, err
		}
		e.generation = generation
	}
	ciphertext, tag := e.cipher.seal(e.nonce, frame, nil)
	out := make([]byte, 0, len(ciphertext)+footerFixed+5)
	out = append(out, ciphertext...)
	out = append(out, tag...)
	out = appendULEB128(out, uint64(e.nonce))
	out = append(out, byte(len(out)-len(ciphertext)+3), markerByte, markerByte)
	return out, nil
}

type unencryptedRange struct{ offset, size uint64 }

type footer struct {
	tag    []byte
	nonce  uint32
	ranges []unencryptedRange
	body   []byte
}

func parseFrame(frame []byte) (*footer, error) {
	n := len(frame)
	if n < footerFixed+1 || frame[n-1] != markerByte || frame[n-2] != markerByte {
		return nil, ErrNotProtocolFrame
	}
	size := int(frame[n-3])
	if size < footerFixed+1 || size > n {
		return nil, ErrNotProtocolFrame
	}
	f := &footer{body: frame[:n-size]}
	supp := frame[n-size : n-3]
	f.tag, supp = supp[:truncatedTag], supp[truncatedTag:]
	nonce, used, err := readULEB128(supp)
	if err != nil || nonce > 0xFFFFFFFF {
		return nil, ErrNotProtocolFrame
	}
	f.nonce, supp = uint32(nonce), supp[used:]
	var end uint64
	for len(supp) > 0 {
		off, a, err := readULEB128(supp)
		if err != nil {
			return nil, ErrNotProtocolFrame
		}
		size, b, err := readULEB128(supp[a:])
		if err != nil || off < end || off+size > uint64(len(f.body)) {
			return nil, ErrNotProtocolFrame
		}
		f.ranges = append(f.ranges, unencryptedRange{off, size})
		end, supp = off+size, supp[a+b:]
	}
	return f, nil
}

func (f *footer) split() (aad, ciphertext []byte) {
	var pos uint64
	for _, r := range f.ranges {
		ciphertext = append(ciphertext, f.body[pos:r.offset]...)
		aad = append(aad, f.body[r.offset:r.offset+r.size]...)
		pos = r.offset + r.size
	}
	return aad, append(ciphertext, f.body[pos:]...)
}

func (f *footer) join(plaintext []byte) []byte {
	out := make([]byte, 0, len(f.body))
	var pos uint64
	for _, r := range f.ranges {
		n := r.offset - pos
		out = append(out, plaintext[:n]...)
		plaintext = plaintext[n:]
		out = append(out, f.body[r.offset:r.offset+r.size]...)
		pos = r.offset + r.size
	}
	return append(out, plaintext...)
}

type epochRatchet struct {
	ratchet *KeyRatchet
	expires time.Time
}

type Decryptor struct {
	mu       sync.Mutex
	ratchets []epochRatchet
	seen     map[*KeyRatchet]map[uint32]bool
}

func (d *Decryptor) SetRatchet(r *KeyRatchet) {
	d.mu.Lock()
	defer d.mu.Unlock()
	now := time.Now()
	kept := d.ratchets[:0]
	for _, er := range d.ratchets {
		if er.expires.IsZero() {
			er.expires = now.Add(previousKeyGrace)
		}
		if now.Before(er.expires) {
			kept = append(kept, er)
		}
	}
	d.ratchets = append(kept, epochRatchet{ratchet: r})
	if d.seen == nil {
		d.seen = map[*KeyRatchet]map[uint32]bool{}
	}
	d.seen[r] = map[uint32]bool{}
}

func (d *Decryptor) Decrypt(frame []byte) ([]byte, error) {
	if IsSilence(frame) {
		return frame, nil
	}
	f, err := parseFrame(frame)
	if err != nil {
		return nil, err
	}
	aad, ciphertext := f.split()
	d.mu.Lock()
	defer d.mu.Unlock()
	now := time.Now()
	for i := len(d.ratchets) - 1; i >= 0; i-- {
		er := d.ratchets[i]
		if !er.expires.IsZero() && now.After(er.expires) {
			continue
		}
		if d.seen[er.ratchet][f.nonce] {
			continue
		}
		key, err := er.ratchet.Key(wrappedGeneration(0, f.nonce>>generationShift))
		if err != nil {
			continue
		}
		c, err := newFrameCipher(key)
		if err != nil {
			return nil, err
		}
		plaintext, err := c.open(f.nonce, ciphertext, f.tag, aad)
		if err != nil {
			continue
		}
		d.seen[er.ratchet][f.nonce] = true
		return f.join(plaintext), nil
	}
	return nil, ErrAuthentication
}
