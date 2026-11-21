package converting

import (
	"fmt"
	"strconv"
	"strings"
)

type Colour uint32

func RGB(r, g, b uint8) Colour {
	return Colour(r)<<16 | Colour(g)<<8 | Colour(b)
}

func (c Colour) R() uint8 { return uint8(c >> 16) }

func (c Colour) G() uint8 { return uint8(c >> 8) }

func (c Colour) B() uint8 { return uint8(c) }

func (c Colour) String() string { return fmt.Sprintf("#%06x", uint32(c)) }

var namedColours = map[string]Colour{
	"default": 0, "teal": 0x1abc9c, "dark_teal": 0x11806a, "brand_green": 0x57f287, "green": 0x2ecc71,
	"dark_green": 0x1f8b4c, "blue": 0x3498db, "dark_blue": 0x206694, "purple": 0x9b59b6, "dark_purple": 0x71368a,
	"magenta": 0xe91e63, "dark_magenta": 0xad1457, "gold": 0xf1c40f, "dark_gold": 0xc27c0e, "orange": 0xe67e22,
	"dark_orange": 0xa84300, "brand_red": 0xed4245, "red": 0xe74c3c, "dark_red": 0x992d22,
	"lighter_grey": 0x95a5a6, "lighter_gray": 0x95a5a6, "dark_grey": 0x607d8b, "dark_gray": 0x607d8b,
	"light_grey": 0x979c9f, "light_gray": 0x979c9f, "darker_grey": 0x546e7a, "darker_gray": 0x546e7a,
	"og_blurple": 0x7289da, "blurple": 0x5865f2, "greyple": 0x99aab5, "dark_theme": 0x313338,
	"fuchsia": 0xeb459e, "yellow": 0xfee75c, "dark_embed": 0x2b2d31, "light_embed": 0xeeeff1, "pink": 0xeb459f,
}

func ParseColour(argument string) (Colour, bool) {
	if hex, ok := strings.CutPrefix(argument, "#"); ok {
		return parseHex(hex)
	}
	if hex, ok := strings.CutPrefix(argument, "0x"); ok {
		return parseHex(hex)
	}
	arg := strings.ToLower(argument)
	if body, ok := strings.CutPrefix(arg, "rgb("); ok {
		if body, ok := strings.CutSuffix(body, ")"); ok {
			return parseRGB(body)
		}
		return 0, false
	}
	if c, ok := namedColours[strings.ReplaceAll(arg, " ", "_")]; ok {
		return c, true
	}
	if len(arg) == 6 {
		return parseHex(arg)
	}
	return 0, false
}

func parseHex(hex string) (Colour, bool) {
	n, err := strconv.ParseUint(hex, 16, 32)
	return Colour(n), err == nil && n <= 0xffffff
}

func parseRGB(body string) (Colour, bool) {
	parts := strings.FieldsFunc(body, func(r rune) bool { return r == ',' || r == ' ' })
	if len(parts) != 3 {
		return 0, false
	}
	var rgb [3]uint8
	for i, p := range parts {
		v, ok := parseComponent(p)
		if !ok {
			return 0, false
		}
		rgb[i] = v
	}
	return RGB(rgb[0], rgb[1], rgb[2]), true
}

func parseComponent(part string) (uint8, bool) {
	if pct, ok := strings.CutSuffix(part, "%"); ok {
		f, err := strconv.ParseFloat(pct, 64)
		return uint8(f / 100 * 255), err == nil && f >= 0 && f <= 100
	}
	n, err := strconv.ParseUint(part, 10, 8)
	return uint8(n), err == nil
}
