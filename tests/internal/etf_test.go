package internal_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/prettylittlelies/cry/internal/etf"
)

type orderedDict struct {
	A     int   `json:"a"`
	Two   int   `json:"2"`
	Three []int `json:"3"`
}

func TestErlpackEncoderVectors(t *testing.T) {
	cases := []struct {
		name string
		in   any
		want string
	}{
		{"string with null byte", "hello\x00 world", "\x83m\x00\x00\x00\x0chello\x00 world"},
		{"string", "hello world", "\x83m\x00\x00\x00\x0bhello world"},
		{"dictionary", orderedDict{A: 1, Two: 2, Three: []int{1, 2, 3}}, "\x83t\x00\x00\x00\x03a\x02a\x02a\x03l\x00\x00\x00\x03a\x01a\x02a\x03jm\x00\x00\x00\x01aa\x01"},
		{"false", false, "\x83s\x05false"},
		{"true", true, "\x83s\x04true"},
		{"nil", nil, "\x83s\x03nil"},
		{"float", 2.5, "\x83F\x40\x04\x00\x00\x00\x00\x00\x00"},
		{"large float", 51512123841234.31423412341435123412341342, "\x83F\x42\xc7\x6c\xcc\xeb\xed\x69\x28"},
		{"int32", 1024, "\x83b\x00\x00\x04\x00"},
		{"int32 min", -2147483648, "\x83b\x80\x00\x00\x00"},
		{"int32 max", 2147483647, "\x83b\x7f\xff\xff\xff"},
		{"uint32", uint32(4294967295), "\x83n\x04\x00\xff\xff\xff\xff"},
		{"beyond uint32 is double", int64(1795228334734), "\x83F\x42\x7a\x1f\xbf\x27\xa8\xe0\x00"},
		{"list", []any{1, "two", 3.1, "four", []string{"five"}}, "\x83l\x00\x00\x00\x05a\x01m\x00\x00\x00\x03twoF\x40\x08\xcc\xcc\xcc\xcc\xcc\xcdm\x00\x00\x00\x04fourl\x00\x00\x00\x01m\x00\x00\x00\x04fivejj"},
		{"empty list", []int{}, "\x83j"},
	}
	for _, c := range cases {
		got, err := etf.Marshal(c.in)
		if err != nil {
			t.Fatalf("%s: %v", c.name, err)
		}
		if !bytes.Equal(got, []byte(c.want)) {
			t.Errorf("%s:\n got  %q\n want %q", c.name, got, c.want)
		}
	}
	for i := range 256 {
		got, _ := etf.Marshal(i)
		if !bytes.Equal(got, []byte{0x83, 'a', byte(i)}) {
			t.Fatalf("small int %d: %q", i, got)
		}
	}
}

func TestErlpackDecoderVectors(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"string ext is byte list", "\x83k\x00\x03abc", `[97,98,99]`},
		{"binary", "\x83m\x00\x00\x00\x0chello\x00 world", `"hello\u0000 world"`},
		{"dictionary", "\x83t\x00\x00\x00\x03a\x02a\x02a\x03l\x00\x00\x00\x03a\x01a\x02a\x03jm\x00\x00\x00\x01aa\x01", `{"2":2,"3":[1,2,3],"a":1}`},
		{"false", "\x83s\x05false", `false`},
		{"true", "\x83s\x04true", `true`},
		{"nil token", "\x83j", `[]`},
		{"nil atom", "\x83s\x03nil", `null`},
		{"null atom", "\x83s\x04null", `null`},
		{"float ext", "\x83c2.50000000000000000000e+00\x00\x00\x00\x00\x00", `2.5`},
		{"new float", "\x83F\x40\x04\x00\x00\x00\x00\x00\x00", `2.5`},
		{"int32 min", "\x83b\x80\x00\x00\x00", `-2147483648`},
		{"small big negative", "\x83n\x04\x01\x01\x02\x03\x04", `-67305985`},
		{"small big", "\x83n\x08\x00\x01\x02\x03\x04\x05\x06\x07\x08", `"578437695752307201"`},
		{"large big", "\x83o\x00\x00\x00\x08\x01\x01\x02\x03\x04\x05\x06\x07\x08", `"-578437695752307201"`},
		{"large big small", "\x83o\x00\x00\x00\x04\x00\x01\x02\x03\x04", `67305985`},
		{"atom", "\x83d\x00\x0Dguild members", `"guild members"`},
		{"small tuple", "\x83h\x03m\x00\x00\x00\x06vanisha\x01a\x04", `["vanish",1,4]`},
		{"large tuple", "\x83i\x00\x00\x00\x03m\x00\x00\x00\x06vanisha\x01a\x04", `["vanish",1,4]`},
		{"compressed", "\x83P\x00\x00\x00\x24\x78\x9C\xCB\x61\x60\x60\x60\x4A\x64\xCA\x66\x90\xCC\x2C\x51\x2F\x56\x48\x4F\x2D\x29\xC9\xCC\x4B\x57\xC8\xC8\x2F\x51\xC8\xCC\x53\xC8\x48\x2D\x4A\xD5\xCB\x02\x00\xA8\xA8\x0A\x9D", `[2,[105,116,39,115,32,103,101,116,116,105,110,103,32,104,111,116,32,105,110,32,104,101,114,101,46]]`},
	}
	for _, c := range cases {
		got, err := etf.ToJSON([]byte(c.in))
		if err != nil {
			t.Fatalf("%s: %v", c.name, err)
		}
		if string(got) != c.want {
			t.Errorf("%s:\n got  %s\n want %s", c.name, got, c.want)
		}
	}
}

func TestErlpackMalformed(t *testing.T) {
	for _, in := range []string{
		"\x83k\x00",
		"\x83t\x00\x00\x00\x03a\x02a\x02a\x03",
		"\x83s\x09true",
		"\x83b\x00\x00\x04",
		"\x83m\x00\x00\x00\x0chel",
		"\x83\xff",
		"a\x01",
	} {
		if _, err := etf.ToJSON([]byte(in)); err == nil {
			t.Errorf("%q: expected error", in)
		}
	}
}

func TestRoundTripKeepsFieldOrder(t *testing.T) {
	type identify struct {
		Token        string `json:"token"`
		Capabilities int    `json:"capabilities"`
		Properties   struct {
			OS            string `json:"os"`
			IsFastConnect bool   `json:"is_fast_connect"`
		} `json:"properties"`
		ClientState struct {
			GuildVersions map[string]int `json:"guild_versions"`
		} `json:"client_state"`
	}
	var in identify
	in.Token, in.Capabilities, in.Properties.OS, in.Properties.IsFastConnect = "t", 1767421, "Mac OS X", true
	in.ClientState.GuildVersions = map[string]int{}
	packed, err := etf.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	raw, err := etf.ToJSON(packed)
	if err != nil {
		t.Fatal(err)
	}
	want, _ := json.Marshal(in)
	if !bytes.Equal(raw, want) {
		t.Fatalf("got  %s\nwant %s", raw, want)
	}
}

func TestDecoderStreamsTerms(t *testing.T) {
	a, _ := etf.Marshal(map[string]int{"op": 10})
	b, _ := etf.Marshal(map[string]int{"op": 11})
	dec := etf.NewDecoder(bytes.NewReader(append(a, b...)))
	for _, want := range []string{`{"op":10}`, `{"op":11}`} {
		got, err := dec.Next()
		if err != nil || string(got) != want {
			t.Fatalf("got %s err=%v", got, err)
		}
	}
}
