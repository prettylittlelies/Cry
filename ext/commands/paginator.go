package commands

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/width"
)

type Paginator struct {
	Prefix  string
	Suffix  string
	MaxSize int
	Linesep string

	current []string
	count   int
	pages   []string
	started bool
}

func NewPaginator(prefix, suffix string) Paginator {
	p := Paginator{Prefix: prefix, Suffix: suffix, MaxSize: 2000, Linesep: "\n"}
	p.Clear()
	return p
}

func (p *Paginator) Clear() {
	p.pages = nil
	p.reset()
	p.started = true
}

func (p *Paginator) reset() {
	p.current, p.count = nil, 0
	if p.Prefix != "" {
		p.current = []string{p.Prefix}
		p.count = runes(p.Prefix) + runes(p.Linesep)
	}
}

func (p *Paginator) ensure() {
	if !p.started {
		p.Clear()
	}
}

func (p *Paginator) AddLine(line string) error {
	return p.add(line, false)
}

func (p *Paginator) AddParagraph(line string) error {
	return p.add(line, true)
}

func (p *Paginator) add(line string, empty bool) error {
	p.ensure()
	sep := runes(p.Linesep)
	limit := p.MaxSize - runes(p.Prefix) - runes(p.Suffix) - 2*sep
	if runes(line) > limit {
		return fmt.Errorf("Line exceeds maximum page size %d", limit)
	}
	if p.count+runes(line)+sep > p.MaxSize-runes(p.Suffix) {
		p.ClosePage()
	}
	p.count += runes(line) + sep
	p.current = append(p.current, line)
	if empty {
		p.current = append(p.current, "")
		p.count += sep
	}
	return nil
}

func (p *Paginator) ClosePage() {
	p.ensure()
	if p.Suffix != "" {
		p.current = append(p.current, p.Suffix)
	}
	p.pages = append(p.pages, strings.Join(p.current, p.Linesep))
	p.reset()
}

func (p *Paginator) Len() int {
	p.ensure()
	total := p.count
	for _, page := range p.pages {
		total += runes(page)
	}
	return total
}

func (p *Paginator) Pages() []string {
	p.ensure()
	floor := 0
	if p.Prefix != "" {
		floor = 1
	}
	if len(p.current) > floor {
		p.ClosePage()
	}
	return p.pages
}

func runes(s string) int {
	return utf8.RuneCountInString(s)
}

func StringWidth(s string) int {
	n := 0
	for _, r := range s {
		switch width.LookupRune(r).Kind() {
		case width.EastAsianWide, width.EastAsianFullwidth, width.EastAsianAmbiguous:
			n += 2
		default:
			n++
		}
	}
	return n
}
