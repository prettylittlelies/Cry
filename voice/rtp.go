package voice

import (
	"encoding/binary"
	"sync"
)

const (
	rtpHeaderLen     = 12
	rtpVersion       = 0x80
	opusPayloadType  = 120
	samplesPerFrame  = 960
	rtcpSenderReport = 200
	rtcpPayloadMax   = 206
)

type rtpSender struct {
	mu        sync.Mutex
	ssrc      uint32
	sequence  uint16
	timestamp uint32
}

func (s *rtpSender) next() []byte {
	s.mu.Lock()
	defer s.mu.Unlock()
	h := make([]byte, rtpHeaderLen)
	h[0], h[1] = rtpVersion, opusPayloadType
	binary.BigEndian.PutUint16(h[2:], s.sequence)
	binary.BigEndian.PutUint32(h[4:], s.timestamp)
	binary.BigEndian.PutUint32(h[8:], s.ssrc)
	s.sequence++
	s.timestamp += samplesPerFrame
	return h
}

type rtpPacket struct {
	payloadType uint8
	sequence    uint16
	timestamp   uint32
	ssrc        uint32
}

func parseRTP(p []byte) (rtpPacket, bool) {
	if len(p) < rtpHeaderLen || p[0]>>6 != 2 {
		return rtpPacket{}, false
	}
	pt := p[1] & 0x7f
	if p[1] >= rtcpSenderReport && p[1] <= rtcpPayloadMax {
		return rtpPacket{}, false
	}
	return rtpPacket{
		payloadType: pt,
		sequence:    binary.BigEndian.Uint16(p[2:]),
		timestamp:   binary.BigEndian.Uint32(p[4:]),
		ssrc:        binary.BigEndian.Uint32(p[8:]),
	}, true
}
