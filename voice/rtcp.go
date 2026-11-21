package voice

import (
	"encoding/binary"
	"time"
)

const (
	rtcpVersion     = 0x80
	rtcpHeaderClear = 8
	ntpEpochOffset  = 2208988800
)

type senderStats struct {
	packets   uint32
	octets    uint32
	timestamp uint32
}

func ntpTime(t time.Time) uint64 {
	secs := uint64(t.Unix()) + ntpEpochOffset
	frac := uint64(t.Nanosecond()) << 32 / uint64(time.Second)
	return secs<<32 | frac
}

func senderReport(ssrc uint32, now time.Time, s senderStats) (header, body []byte) {
	header = make([]byte, rtcpHeaderClear)
	header[0], header[1] = rtcpVersion, rtcpSenderReport
	binary.BigEndian.PutUint16(header[2:], 6)
	binary.BigEndian.PutUint32(header[4:], ssrc)
	body = binary.BigEndian.AppendUint64(nil, ntpTime(now))
	body = binary.BigEndian.AppendUint32(body, s.timestamp)
	body = binary.BigEndian.AppendUint32(body, s.packets)
	body = binary.BigEndian.AppendUint32(body, s.octets)
	return header, body
}
