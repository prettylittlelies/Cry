package discord

import (
	"regexp"
	"strings"
)

const (
	markdownCommon = `^>(?:>>)?\s|\[.+\]\(.+\)|^#{1,3}|^\s*-`
	markdownURL    = `(?P<url><[^: >]+:/[^ >]+>|(?:https?|steam)://[^\s<]+[^<.,:;"'\]\s])`
	markdownStock  = `(?P<markdown>[_\\~|*` + "`" + `]|` + markdownCommon + `)`
	markdownChars  = "*`_~|"
)

var (
	markdownStockPattern = regexp.MustCompile(`(?m)` + markdownStock)
	markdownLinkPattern  = regexp.MustCompile(`(?m)(?:` + markdownURL + `|` + markdownStock + `)`)
	mentionEscapePattern = regexp.MustCompile(`@(everyone|here|[!&]?[0-9]{17,20})`)
	markdownAnchored     = []*regexp.Regexp{
		regexp.MustCompile(`\A>(?:>>)?\s`),
		regexp.MustCompile(`\A\[.+\]\(.+\)`),
		regexp.MustCompile(`\A#{1,3}`),
		regexp.MustCompile(`\A\s*-`),
	}
)

func markdownPattern(ignoreLinks bool) *regexp.Regexp {
	if ignoreLinks {
		return markdownLinkPattern
	}
	return markdownStockPattern
}

func replaceMarkdown(text string, ignoreLinks bool, replace func(url, markdown string) string) string {
	re := markdownPattern(ignoreLinks)
	urlGroup, mdGroup := re.SubexpIndex("url"), re.SubexpIndex("markdown")
	group := func(m []int, g int) string {
		if g < 0 || m[2*g] < 0 {
			return ""
		}
		return text[m[2*g]:m[2*g+1]]
	}
	var b strings.Builder
	last := 0
	for _, m := range re.FindAllStringSubmatchIndex(text, -1) {
		b.WriteString(text[last:m[0]])
		b.WriteString(replace(group(m, urlGroup), group(m, mdGroup)))
		last = m[1]
	}
	b.WriteString(text[last:])
	return b.String()
}

func RemoveMarkdown(text string, ignoreLinks bool) string {
	return replaceMarkdown(text, ignoreLinks, func(url, _ string) string { return url })
}

func EscapeMarkdown(text string, asNeeded, ignoreLinks bool) string {
	if asNeeded {
		return escapeNeededMarkdown(strings.ReplaceAll(text, `\`, `\\`))
	}
	return replaceMarkdown(text, ignoreLinks, func(url, markdown string) string {
		if url != "" {
			return url
		}
		return `\` + markdown
	})
}

func escapeNeededMarkdown(text string) string {
	closers := map[byte]int{}
	for j := 1; j < len(text); j++ {
		if c := text[j]; strings.IndexByte(markdownChars, c) >= 0 && text[j-1] != c {
			closers[c] = j
		}
	}
	var b strings.Builder
	for i := 0; i < len(text); {
		if n := markdownMatch(text, i, closers); n > 0 {
			b.WriteByte('\\')
			b.WriteString(text[i : i+n])
			i += n
			continue
		}
		b.WriteByte(text[i])
		i++
	}
	return b.String()
}

func markdownMatch(text string, i int, closers map[byte]int) int {
	if c := text[i]; strings.IndexByte(markdownChars, c) >= 0 {
		if j, ok := closers[c]; ok && j > i {
			return 1
		}
	}
	lineStart := i == 0 || text[i-1] == '\n'
	for k, re := range markdownAnchored {
		if k != 1 && !lineStart {
			continue
		}
		if loc := re.FindStringIndex(text[i:]); loc != nil && loc[1] > 0 {
			return loc[1]
		}
	}
	return 0
}

func EscapeMentions(text string) string {
	return mentionEscapePattern.ReplaceAllString(text, "@​$1")
}
