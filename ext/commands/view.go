package commands

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

var quotes = map[rune]rune{
	'"': '"', '\'': '\'', '‘': '’', '‚': '‛', '“': '”', '„': '‟', '⹂': '⹂',
	'「': '」', '『': '』', '〝': '〞', '﹁': '﹂', '﹃': '﹄', '＂': '＂', '｢': '｣',
	'«': '»', '‹': '›', '《': '》', '〈': '〉',
}

type View struct {
	buf string
	pos int
}

func NewView(s string) *View {
	return &View{buf: s}
}

func (v *View) EOF() bool {
	return v.pos >= len(v.buf)
}

func (v *View) SkipSpace() {
	for v.pos < len(v.buf) {
		r, n := utf8.DecodeRuneInString(v.buf[v.pos:])
		if !unicode.IsSpace(r) {
			return
		}
		v.pos += n
	}
}

func (v *View) Rest() string {
	v.SkipSpace()
	out := v.buf[v.pos:]
	v.pos = len(v.buf)
	return out
}

func (v *View) Word() (string, error) {
	v.SkipSpace()
	if v.EOF() {
		return "", errEndOfInput
	}
	open, n := utf8.DecodeRuneInString(v.buf[v.pos:])
	closing, quoted := quotes[open]
	if !quoted {
		start := v.pos
		for v.pos < len(v.buf) {
			r, size := utf8.DecodeRuneInString(v.buf[v.pos:])
			if unicode.IsSpace(r) {
				break
			}
			v.pos += size
		}
		return v.buf[start:v.pos], nil
	}
	v.pos += n
	var out strings.Builder
	for v.pos < len(v.buf) {
		r, size := utf8.DecodeRuneInString(v.buf[v.pos:])
		v.pos += size
		switch {
		case r == '\\' && v.pos < len(v.buf):
			next, nsize := utf8.DecodeRuneInString(v.buf[v.pos:])
			if next == closing || next == '\\' {
				out.WriteRune(next)
				v.pos += nsize
				continue
			}
			out.WriteRune(r)
		case r == closing:
			if v.pos < len(v.buf) {
				after, _ := utf8.DecodeRuneInString(v.buf[v.pos:])
				if !unicode.IsSpace(after) {
					return "", &UnexpectedQuoteError{Quote: closing}
				}
			}
			return out.String(), nil
		default:
			out.WriteRune(r)
		}
	}
	return "", &UnclosedQuoteError{Quote: closing}
}
