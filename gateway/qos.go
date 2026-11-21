package gateway

import (
	"encoding/base64"
	"slices"
	"strings"

	"google.golang.org/protobuf/encoding/protowire"
)

const QoSVersion = 31

type QoSReason string

const (
	Foregrounded QoSReason = "foregrounded"
	RTCConnected QoSReason = "rtc_connected"
)

type QoS struct {
	Active  bool        `json:"active"`
	Ver     int         `json:"ver"`
	Reasons []QoSReason `json:"reasons"`
}

func NewQoS(reasons ...QoSReason) QoS {
	if reasons == nil {
		reasons = []QoSReason{}
	}
	return QoS{Active: len(reasons) > 0, Ver: QoSVersion, Reasons: reasons}
}

type qosState struct {
	current  *QoS
	upcoming *QoS
}

func (s *qosState) change(q QoS) (wake bool) {
	wasInactive := s.current == nil || !s.current.Active
	if s.current == nil {
		c := q
		c.Reasons = slices.Clone(q.Reasons)
		s.current = &c
	}
	if q.Active {
		s.current.Active = true
		merged := slices.Concat(s.current.Reasons, q.Reasons)
		slices.Sort(merged)
		s.current.Reasons = slices.Compact(merged)
		wake = wasInactive
	}
	up := q
	s.upcoming = &up
	return wake
}

func (s *qosState) consume() *QoS {
	out := s.current
	if s.upcoming != nil {
		s.current = s.upcoming
	}
	s.upcoming = nil
	return out
}

func (s *qosState) active() bool {
	switch {
	case s.upcoming != nil:
		return s.upcoming.Active
	case s.current != nil:
		return s.current.Active
	}
	return false
}

func decodeLenient(s string) []byte {
	s = strings.NewReplacer("-", "+", "_", "/").Replace(strings.TrimRight(s, "="))
	b, err := base64.RawStdEncoding.DecodeString(s)
	if err != nil {
		return nil
	}
	return b
}

func qosToken(active bool, derived string) string {
	var client []byte
	if active {
		client = protowire.AppendTag(client, 1, protowire.VarintType)
		client = protowire.AppendVarint(client, 1)
	}
	var token []byte
	token = protowire.AppendTag(token, 1, protowire.BytesType)
	token = protowire.AppendBytes(token, client)
	if raw := decodeLenient(derived); len(raw) > 0 {
		token = protowire.AppendTag(token, 2, protowire.BytesType)
		token = protowire.AppendBytes(token, raw)
	}
	return base64.StdEncoding.EncodeToString(token)
}
