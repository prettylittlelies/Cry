package build

import (
	"fmt"
	"strconv"
)

type MobileTrack uint8

const (
	TrackStable MobileTrack = iota
	TrackBeta
	TrackAlpha
	TrackDevelopment
)

type MobileVersion struct {
	Major uint32
	Minor uint32
}

func ParseMobileVersion(s string) (MobileVersion, error) {
	v, err := ParseVersion(s)
	if err != nil || len(v) > 2 {
		return MobileVersion{}, &ParseError{Kind: "mobile version", Input: s, Err: err}
	}
	return MobileVersion{Major: v.Segment(0), Minor: v.Segment(1)}, nil
}

func (m MobileVersion) String() string {
	return fmt.Sprintf("%d.%d", m.Major, m.Minor)
}

func (m MobileVersion) AndroidCode(t MobileTrack) uint32 {
	return m.Major*1000 + uint32(t)*100 + m.Minor
}

func (m MobileVersion) AndroidUserAgent(t MobileTrack) string {
	return "Discord-Android/" + strconv.FormatUint(uint64(m.AndroidCode(t)), 10) + ";RNA"
}
