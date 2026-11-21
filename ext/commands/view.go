package commands

import "unicode"

var quotes = map[rune]rune{
	'"': '"', '‘': '’', '‚': '‛', '“': '”', '„': '‟', '⹂': '⹂', '「': '」', '『': '』',
	'〝': '〞', '﹁': '﹂', '﹃': '﹄', '＂': '＂', '｢': '｣', '«': '»', '‹': '›', '《': '》', '〈': '〉',
}

var allQuotes = func() map[rune]bool {
	out := map[rune]bool{}
	for open, closing := range quotes {
		out[open], out[closing] = true, true
	}
	return out
}()

type StringView struct {
	buffer   []rune
	Index    int
	Previous int
}

func NewStringView(s string) *StringView {
	return &StringView{buffer: []rune(s)}
}

func (v *StringView) end() int {
	return len(v.buffer)
}

func (v *StringView) EOF() bool {
	return v.Index >= v.end()
}

func (v *StringView) Current() (rune, bool) {
	if v.EOF() {
		return 0, false
	}
	return v.buffer[v.Index], true
}

func (v *StringView) Undo() {
	v.Index = v.Previous
}

func (v *StringView) SkipWS() bool {
	pos := 0
	for v.Index+pos < v.end() && unicode.IsSpace(v.buffer[v.Index+pos]) {
		pos++
	}
	v.Previous = v.Index
	v.Index += pos
	return v.Previous != v.Index
}

func (v *StringView) SkipString(s string) bool {
	r := []rune(s)
	if v.Index+len(r) > v.end() || string(v.buffer[v.Index:v.Index+len(r)]) != s {
		return false
	}
	v.Previous = v.Index
	v.Index += len(r)
	return true
}

func (v *StringView) ReadRest() string {
	result := string(v.buffer[min(v.Index, v.end()):])
	v.Previous = v.Index
	v.Index = v.end()
	return result
}

func (v *StringView) Read(n int) string {
	stop := min(v.Index+n, v.end())
	result := string(v.buffer[min(v.Index, v.end()):stop])
	v.Previous = v.Index
	v.Index += n
	return result
}

func (v *StringView) get() (rune, bool) {
	var result rune
	ok := v.Index+1 < v.end()
	if ok {
		result = v.buffer[v.Index+1]
	}
	v.Previous = v.Index
	v.Index++
	return result, ok
}

func (v *StringView) GetWord() string {
	pos := 0
	for v.Index+pos < v.end() && !unicode.IsSpace(v.buffer[v.Index+pos]) {
		pos++
	}
	v.Previous = v.Index
	result := string(v.buffer[v.Index : v.Index+pos])
	v.Index += pos
	return result
}

func (v *StringView) GetQuotedWord() (string, bool, error) {
	current, ok := v.Current()
	if !ok {
		return "", false, nil
	}
	closing, quoted := quotes[current]
	var result []rune
	escaped := allQuotes
	if quoted {
		escaped = map[rune]bool{current: true, closing: true}
	} else {
		result = append(result, current)
	}
	for !v.EOF() {
		current, ok = v.get()
		if !ok {
			if quoted {
				return "", true, errExpectedClosingQuote(closing)
			}
			return string(result), true, nil
		}
		if current == '\\' {
			next, ok := v.get()
			if !ok {
				if quoted {
					return "", true, errExpectedClosingQuote(closing)
				}
				return string(result), true, nil
			}
			if escaped[next] {
				result = append(result, next)
			} else {
				v.Undo()
				result = append(result, current)
			}
			continue
		}
		if !quoted && allQuotes[current] {
			return "", true, errUnexpectedQuote(current)
		}
		if quoted && current == closing {
			next, ok := v.get()
			if ok && !unicode.IsSpace(next) {
				return "", true, errInvalidEndOfQuoted(next)
			}
			return string(result), true, nil
		}
		if unicode.IsSpace(current) && !quoted {
			return string(result), true, nil
		}
		result = append(result, current)
	}
	return string(result), true, nil
}
