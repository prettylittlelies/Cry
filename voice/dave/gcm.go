package dave

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/subtle"
	"encoding/binary"
	"errors"
)

const (
	keySize         = 16
	nonceSize       = 12
	truncatedNonce  = 4
	truncatedTag    = 8
	generationShift = 24
)

var ErrAuthentication = errors.New("dave: frame authentication failed")

type frameCipher struct {
	block cipher.Block
	aead  cipher.AEAD
}

func newFrameCipher(key []byte) (*frameCipher, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return &frameCipher{block: block, aead: aead}, nil
}

func expandNonce(n uint32) []byte {
	full := make([]byte, nonceSize)
	binary.LittleEndian.PutUint32(full[nonceSize-truncatedNonce:], n)
	return full
}

func (c *frameCipher) seal(nonce uint32, plaintext, aad []byte) (ciphertext, tag []byte) {
	out := c.aead.Seal(nil, expandNonce(nonce), plaintext, aad)
	return out[:len(plaintext)], out[len(plaintext) : len(plaintext)+truncatedTag]
}

func (c *frameCipher) open(nonce uint32, ciphertext, tag, aad []byte) ([]byte, error) {
	full := expandNonce(nonce)
	counter := make([]byte, aes.BlockSize)
	copy(counter, full)
	binary.BigEndian.PutUint32(counter[nonceSize:], 2)
	plaintext := make([]byte, len(ciphertext))
	cipher.NewCTR(c.block, counter).XORKeyStream(plaintext, ciphertext)
	check := c.aead.Seal(nil, full, plaintext, aad)
	if subtle.ConstantTimeCompare(check[len(plaintext):len(plaintext)+truncatedTag], tag) != 1 {
		return nil, ErrAuthentication
	}
	return plaintext, nil
}
