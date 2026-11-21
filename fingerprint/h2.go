package fingerprint

import (
	"strconv"
	"strings"

	"golang.org/x/net/http2"
)

type Urgency uint8

const DefaultUrgency Urgency = 3

type Pseudo string

const (
	Method    Pseudo = ":method"
	Authority Pseudo = ":authority"
	Scheme    Pseudo = ":scheme"
	Path      Pseudo = ":path"
)

type H2 struct {
	Settings       []http2.Setting
	ConnectionFlow uint32
	PseudoOrder    []Pseudo
	Weights        []uint16
	Exclusive      bool
}

func Chromium152H2() H2 {
	return H2{
		Settings: []http2.Setting{
			{ID: http2.SettingHeaderTableSize, Val: 65536},
			{ID: http2.SettingEnablePush, Val: 0},
			{ID: http2.SettingInitialWindowSize, Val: 6291456},
			{ID: http2.SettingMaxHeaderListSize, Val: 262144},
		},
		ConnectionFlow: 15663105,
		PseudoOrder:    []Pseudo{Method, Authority, Scheme, Path},
		Weights:        []uint16{256, 220, 183, 147, 110, 74, 37, 1},
		Exclusive:      true,
	}
}

func (h H2) Setting(id http2.SettingID) (uint32, bool) {
	for _, s := range h.Settings {
		if s.ID == id {
			return s.Val, true
		}
	}
	return 0, false
}

func (h H2) Priority(u Urgency) http2.PriorityParam {
	w := h.Weights[min(int(u), len(h.Weights)-1)]
	return http2.PriorityParam{Exclusive: h.Exclusive, Weight: uint8(w - 1)}
}

func ParseUrgency(priority string) Urgency {
	for part := range strings.SplitSeq(priority, ",") {
		k, v, ok := strings.Cut(strings.TrimSpace(part), "=")
		if !ok || k != "u" {
			continue
		}
		if n, err := strconv.ParseUint(v, 10, 8); err == nil && n <= 7 {
			return Urgency(n)
		}
	}
	return DefaultUrgency
}
