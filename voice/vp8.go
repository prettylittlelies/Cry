package voice

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"sync"
)

const (
	vp8PayloadType    = 103
	vp8RTXPayloadType = 104
	videoClock        = 90000
	maxVideoPayload   = 1100
	extPlayoutDelay   = 6
	extRID            = 11
	primaryRID        = "100"
)

var ErrNotIVF = errors.New("voice: not an ivf stream")

var videoExtensions = func() []byte {
	elems := []byte{extPlayoutDelay<<4 | 2, 0, 0, 0, extRID<<4 | byte(len(primaryRID)-1)}
	elems = append(elems, primaryRID...)
	for len(elems)%4 != 0 {
		elems = append(elems, 0)
	}
	return elems
}()

type videoSender struct {
	mu       sync.Mutex
	ssrc     uint32
	rtx      uint32
	sequence uint16
	stamp    uint32
	stats    senderStats
}

type videoPacket struct {
	header  []byte
	payload []byte
}

func (s *videoSender) packetize(frame []byte, step uint32) []videoPacket {
	s.mu.Lock()
	defer s.mu.Unlock()
	var out []videoPacket
	for off := 0; off == 0 || off < len(frame); {
		end := min(off+maxVideoPayload, len(frame))
		h := make([]byte, rtpHeaderLen, rtpHeaderLen+4)
		h[0], h[1] = rtpVersion|0x10, vp8PayloadType
		if end == len(frame) {
			h[1] |= 0x80
		}
		binary.BigEndian.PutUint16(h[2:], s.sequence)
		binary.BigEndian.PutUint32(h[4:], s.stamp)
		binary.BigEndian.PutUint32(h[8:], s.ssrc)
		h = binary.BigEndian.AppendUint16(h, 0xBEDE)
		h = binary.BigEndian.AppendUint16(h, uint16(len(videoExtensions)/4))
		descriptor := byte(0)
		if off == 0 {
			descriptor = 0x10
		}
		payload := append(append(append([]byte(nil), videoExtensions...), descriptor), frame[off:end]...)
		out = append(out, videoPacket{header: h, payload: payload})
		s.sequence++
		s.stats.packets++
		s.stats.octets += uint32(len(payload))
		off = end
		if end == len(frame) {
			break
		}
	}
	s.stats.timestamp = s.stamp
	s.stamp += step
	return out
}

func (s *videoSender) snapshot() (uint32, senderStats) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.ssrc, s.stats
}

type IVFHeader struct {
	FourCC    string
	Width     int
	Height    int
	Rate      uint32
	Scale     uint32
	FrameRate float64
}

type IVF struct {
	r      *bufio.Reader
	Header IVFHeader
}

func NewIVF(r io.Reader) (*IVF, error) {
	br := bufio.NewReader(r)
	head := make([]byte, 32)
	if _, err := io.ReadFull(br, head); err != nil {
		return nil, err
	}
	if !bytes.Equal(head[:4], []byte("DKIF")) {
		return nil, ErrNotIVF
	}
	size := int(binary.LittleEndian.Uint16(head[6:]))
	if size > 32 {
		if _, err := br.Discard(size - 32); err != nil {
			return nil, err
		}
	}
	h := IVFHeader{
		FourCC: string(head[8:12]),
		Width:  int(binary.LittleEndian.Uint16(head[12:])),
		Height: int(binary.LittleEndian.Uint16(head[14:])),
		Rate:   binary.LittleEndian.Uint32(head[16:]),
		Scale:  binary.LittleEndian.Uint32(head[20:]),
	}
	if h.Scale > 0 {
		h.FrameRate = float64(h.Rate) / float64(h.Scale)
	}
	return &IVF{r: br, Header: h}, nil
}

func (v *IVF) Next() ([]byte, error) {
	head := make([]byte, 12)
	if _, err := io.ReadFull(v.r, head); err != nil {
		return nil, err
	}
	frame := make([]byte, binary.LittleEndian.Uint32(head))
	if _, err := io.ReadFull(v.r, frame); err != nil {
		return nil, err
	}
	return frame, nil
}
