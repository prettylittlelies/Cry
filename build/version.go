package build

import (
	"encoding/json"
	"slices"
	"strconv"
	"strings"
)

type Version []uint32

func ParseVersion(s string) (Version, error) {
	parts := strings.Split(strings.TrimSpace(s), ".")
	v := make(Version, len(parts))
	for i, p := range parts {
		n, err := strconv.ParseUint(p, 10, 32)
		if err != nil {
			return nil, &ParseError{Kind: "version", Input: s, Err: err}
		}
		v[i] = uint32(n)
	}
	return v, nil
}

func (v Version) Segment(i int) uint32 {
	if i < len(v) {
		return v[i]
	}
	return 0
}

func (v Version) Major() uint32 {
	return v.Segment(0)
}

func (v Version) Compare(o Version) int {
	return slices.Compare(v, o)
}

func (v Version) String() string {
	parts := make([]string, len(v))
	for i, n := range v {
		parts[i] = strconv.FormatUint(uint64(n), 10)
	}
	return strings.Join(parts, ".")
}

func (v Version) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v *Version) UnmarshalJSON(b []byte) error {
	if len(b) > 0 && b[0] == '[' {
		var segs []uint32
		if err := json.Unmarshal(b, &segs); err != nil {
			return &ParseError{Kind: "version", Input: string(b), Err: err}
		}
		*v = segs
		return nil
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return &ParseError{Kind: "version", Input: string(b), Err: err}
	}
	return v.UnmarshalText([]byte(s))
}

func (v *Version) UnmarshalText(b []byte) error {
	p, err := ParseVersion(string(b))
	if err != nil {
		return err
	}
	*v = p
	return nil
}
