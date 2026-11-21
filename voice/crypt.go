package voice

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"errors"
	"fmt"
	"slices"
	"sync/atomic"

	"golang.org/x/crypto/chacha20poly1305"
)

const (
	ModeAES256GCM         = "aead_aes256_gcm_rtpsize"
	ModeXChaCha20Poly1305 = "aead_xchacha20_poly1305_rtpsize"
	nonceSuffix           = 4
)

var (
	ErrNoMode       = errors.New("voice: no supported transport encryption mode")
	ErrShortPacket  = errors.New("voice: packet too short")
	ErrTransportTag = errors.New("voice: transport decryption failed")
)

func chooseMode(offered []string) (string, error) {
	for _, m := range []string{ModeAES256GCM, ModeXChaCha20Poly1305} {
		if slices.Contains(offered, m) {
			return m, nil
		}
	}
	return "", fmt.Errorf("%w: %v", ErrNoMode, offered)
}

type transport struct {
	aead    cipher.AEAD
	counter atomic.Uint32
}

func newTransport(mode string, key []byte) (*transport, error) {
	var aead cipher.AEAD
	var err error
	switch mode {
	case ModeAES256GCM:
		var block cipher.Block
		if block, err = aes.NewCipher(key); err == nil {
			aead, err = cipher.NewGCM(block)
		}
	case ModeXChaCha20Poly1305:
		aead, err = chacha20poly1305.NewX(key)
	default:
		return nil, fmt.Errorf("%w: %s", ErrNoMode, mode)
	}
	if err != nil {
		return nil, err
	}
	return &transport{aead: aead}, nil
}

func (t *transport) nonce(suffix []byte) []byte {
	n := make([]byte, t.aead.NonceSize())
	copy(n, suffix)
	return n
}

func (t *transport) seal(header, payload []byte) []byte {
	suffix := binary.BigEndian.AppendUint32(nil, t.counter.Add(1))
	out := append(slices.Clip(header), t.aead.Seal(nil, t.nonce(suffix), payload, header)...)
	return append(out, suffix...)
}

func rtpHeaderSize(p []byte) (clear, extension int, err error) {
	if len(p) < rtpHeaderLen {
		return 0, 0, ErrShortPacket
	}
	clear = rtpHeaderLen + 4*int(p[0]&0x0f)
	if p[0]&0x10 != 0 {
		clear += 4
		if len(p) < clear {
			return 0, 0, ErrShortPacket
		}
		extension = 4 * int(binary.BigEndian.Uint16(p[clear-2:clear]))
	}
	if len(p) < clear {
		return 0, 0, ErrShortPacket
	}
	return clear, extension, nil
}

func (t *transport) open(p []byte) (header, payload []byte, err error) {
	clear, extension, err := rtpHeaderSize(p)
	if err != nil {
		return nil, nil, err
	}
	if len(p) < clear+nonceSuffix+t.aead.Overhead() {
		return nil, nil, ErrShortPacket
	}
	body, suffix := p[clear:len(p)-nonceSuffix], p[len(p)-nonceSuffix:]
	plain, err := t.aead.Open(nil, t.nonce(suffix), body, p[:clear])
	if err != nil {
		return nil, nil, ErrTransportTag
	}
	if extension > len(plain) {
		return nil, nil, ErrShortPacket
	}
	return p[:clear], plain[extension:], nil
}
