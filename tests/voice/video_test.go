package voice_test

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"testing"

	"github.com/prettylittlelies/cry/voice"
	"github.com/prettylittlelies/cry/voice/dave"
)

func vp8Frames() (key, inter []byte) {
	key = append([]byte{0x10, 0x02, 0x00, 0x9d, 0x01, 0x2a, 0x80, 0x02, 0xe0, 0x01}, bytes.Repeat([]byte{0x55}, 200)...)
	inter = append([]byte{0x31}, bytes.Repeat([]byte{0x66}, 90)...)
	return key, inter
}

func TestVP8DAVERanges(t *testing.T) {
	key, inter := vp8Frames()
	if r := dave.VP8Ranges(key); len(r) != 1 || r[0].Size != 10 {
		t.Fatalf("keyframe ranges %v", r)
	}
	if r := dave.VP8Ranges(inter); len(r) != 1 || r[0].Size != 1 {
		t.Fatalf("interframe ranges %v", r)
	}
	rt := ratchet(t)
	var enc dave.Encryptor
	enc.SetRatchet(rt)
	var dec dave.Decryptor
	dec.SetRatchet(rt)
	for _, frame := range [][]byte{key, inter} {
		protected, err := enc.EncryptRanges(frame, dave.VP8Ranges(frame))
		if err != nil {
			t.Fatal(err)
		}
		clear := dave.VP8Ranges(frame)[0].Size
		if !bytes.Equal(protected[:clear], frame[:clear]) || bytes.Equal(protected[clear:len(frame)], frame[clear:]) {
			t.Fatal("unencrypted range must stay clear and the rest must be encrypted")
		}
		plain, err := dec.Decrypt(protected)
		if err != nil || !bytes.Equal(plain, frame) {
			t.Fatalf("decrypt: %v", err)
		}
	}
}

func ivfFile(frames ...[]byte) []byte {
	head := make([]byte, 32)
	copy(head, "DKIF")
	binary.LittleEndian.PutUint16(head[6:], 32)
	copy(head[8:], "VP80")
	binary.LittleEndian.PutUint16(head[12:], 640)
	binary.LittleEndian.PutUint16(head[14:], 480)
	binary.LittleEndian.PutUint32(head[16:], 30)
	binary.LittleEndian.PutUint32(head[20:], 1)
	for i, f := range frames {
		fh := binary.LittleEndian.AppendUint32(nil, uint32(len(f)))
		fh = binary.LittleEndian.AppendUint64(fh, uint64(i))
		head = append(append(head, fh...), f...)
	}
	return head
}

func TestIVFReader(t *testing.T) {
	key, inter := vp8Frames()
	ivf, err := voice.NewIVF(bytes.NewReader(ivfFile(key, inter)))
	if err != nil {
		t.Fatal(err)
	}
	h := ivf.Header
	if h.FourCC != "VP80" || h.Width != 640 || h.Height != 480 || h.FrameRate != 30 {
		t.Fatalf("header %+v", h)
	}
	for _, want := range [][]byte{key, inter} {
		got, err := ivf.Next()
		if err != nil || !bytes.Equal(got, want) {
			t.Fatalf("frame: %v", err)
		}
	}
	if _, err := ivf.Next(); !errors.Is(err, io.EOF) {
		t.Fatalf("end: %v", err)
	}
	if _, err := voice.NewIVF(bytes.NewReader(make([]byte, 32))); !errors.Is(err, voice.ErrNotIVF) {
		t.Fatal("non-ivf accepted")
	}
}
