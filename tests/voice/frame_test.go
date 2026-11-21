package voice_test

import (
	"bytes"
	"testing"

	"github.com/prettylittlelies/cry/mls"
	"github.com/prettylittlelies/cry/voice/dave"
)

func ratchet(t *testing.T) *dave.KeyRatchet {
	t.Helper()
	s, err := mls.Suite(mls.P256AES128GCMSHA256P256)
	if err != nil {
		t.Fatal(err)
	}
	return dave.NewKeyRatchet(s, bytes.Repeat([]byte{7}, 16))
}

func TestFrameProtection(t *testing.T) {
	r := ratchet(t)
	var enc dave.Encryptor
	enc.SetRatchet(r)
	var dec dave.Decryptor
	dec.SetRatchet(r)
	frame := []byte("opus frame payload")
	protected, err := enc.Encrypt(frame)
	if err != nil {
		t.Fatal(err)
	}
	n := len(protected)
	if protected[n-1] != 0xFA || protected[n-2] != 0xFA || int(protected[n-3]) != 8+1+1+2 || n != len(frame)+12 {
		t.Fatalf("footer % x", protected[len(frame):])
	}
	if got, err := dec.Decrypt(protected); err != nil || !bytes.Equal(got, frame) {
		t.Fatalf("decrypt %q %v", got, err)
	}
	if _, err := dec.Decrypt(protected); err == nil {
		t.Fatal("replayed nonce must be rejected")
	}
	second, _ := enc.Encrypt(frame)
	if bytes.Equal(second[:len(frame)], protected[:len(frame)]) {
		t.Fatal("nonce must advance")
	}
	tampered := bytes.Clone(second)
	tampered[0] ^= 1
	if _, err := dec.Decrypt(tampered); err == nil {
		t.Fatal("tampered frame must fail")
	}
	if got, err := dec.Decrypt([]byte{0xF8, 0xFF, 0xFE}); err != nil || len(got) != 3 {
		t.Fatal("silence passes through")
	}
	if _, err := dec.Decrypt([]byte("plain opus")); err != dave.ErrNotProtocolFrame {
		t.Fatalf("plain frame: %v", err)
	}
}

func TestRatchetGenerations(t *testing.T) {
	r := ratchet(t)
	k0, _ := r.Key(0)
	k5, _ := r.Key(5)
	again, _ := r.Key(0)
	if bytes.Equal(k0, k5) || !bytes.Equal(k0, again) || len(k0) != 16 {
		t.Fatal("ratchet keys")
	}
	r.Erase(0)
	if _, err := r.Key(0); err != dave.ErrErasedGeneration {
		t.Fatalf("erased: %v", err)
	}
}
