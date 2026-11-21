package fingerprint

import (
	"net/http"
	"slices"
	"strings"

	"golang.org/x/net/http2/hpack"
)

const Wildcard = "*"

type HeaderPolicy struct {
	Insertion []string
	Leading   []string
	Hints     []string
	Trailing  []string
}

var hopByHop = map[string]struct{}{
	"connection":        {},
	"host":              {},
	"keep-alive":        {},
	"proxy-connection":  {},
	"te":                {},
	"transfer-encoding": {},
	"upgrade":           {},
}

var (
	discordInsertion = []string{
		"content-type",
		Wildcard,
		"x-audit-log-reason",
		"x-context-properties",
		"x-failed-requests",
		"authorization",
		"x-super-properties",
		"x-fingerprint",
		"x-installation-id",
		"accept-language",
		"x-discord-locale",
		"x-discord-timezone",
		"x-debug-options",
		"x-routing-key",
		"x-client-trace-id",
	}
	chromiumHints    = []string{"sec-ch-ua", "sec-ch-ua-mobile", "sec-ch-ua-platform", "user-agent"}
	chromiumTrailing = []string{
		"accept",
		"origin",
		"sec-fetch-site",
		"sec-fetch-mode",
		"sec-fetch-dest",
		"sec-fetch-storage-access",
		"referer",
		"accept-encoding",
		"accept-language",
		"cookie",
		"priority",
	}
)

func DiscordWebPolicy() HeaderPolicy {
	return HeaderPolicy{
		Insertion: discordInsertion,
		Leading:   []string{"content-length"},
		Hints:     chromiumHints,
		Trailing:  chromiumTrailing,
	}
}

func DiscordDesktopPolicy() HeaderPolicy {
	p := DiscordWebPolicy()
	p.Trailing = slices.DeleteFunc(slices.Clone(chromiumTrailing), func(s string) bool { return s == "accept-language" })
	return p
}

func (p HeaderPolicy) Arrange(inserted []string) []string {
	managed := make(map[string]struct{}, len(p.Leading)+len(p.Hints)+len(p.Trailing))
	for _, group := range [][]string{p.Leading, p.Hints, p.Trailing} {
		for _, n := range group {
			managed[n] = struct{}{}
		}
	}
	present := make(map[string]struct{}, len(inserted))
	var page []string
	for _, n := range inserted {
		present[n] = struct{}{}
		if _, ok := managed[n]; !ok {
			page = append(page, n)
		}
	}
	keep := func(group []string) []string {
		return slices.DeleteFunc(slices.Clone(group), func(n string) bool {
			_, ok := present[n]
			return !ok
		})
	}
	return slices.Concat(keep(p.Leading), BlinkOrder(BlinkOrder(page), keep(p.Hints)), keep(p.Trailing))
}

func (p HeaderPolicy) insertion(names map[string]struct{}) []string {
	known := make(map[string]struct{}, len(p.Insertion))
	for _, n := range p.Insertion {
		known[n] = struct{}{}
	}
	var unknown []string
	for n := range names {
		if _, ok := known[n]; !ok {
			unknown = append(unknown, n)
		}
	}
	slices.Sort(unknown)
	out := make([]string, 0, len(names))
	for _, n := range p.Insertion {
		if n == Wildcard {
			out = append(out, unknown...)
			unknown = nil
		} else if _, ok := names[n]; ok {
			out = append(out, n)
		}
	}
	return append(out, unknown...)
}

func (p HeaderPolicy) Fields(h http.Header) []hpack.HeaderField {
	values := make(map[string][]string, len(h))
	names := make(map[string]struct{}, len(h))
	for k, v := range h {
		name := strings.ToLower(k)
		if _, skip := hopByHop[name]; !skip {
			values[name] = append(values[name], v...)
			names[name] = struct{}{}
		}
	}
	fields := make([]hpack.HeaderField, 0, len(values))
	for _, n := range p.Arrange(p.insertion(names)) {
		for _, v := range values[n] {
			fields = append(fields, hpack.HeaderField{Name: n, Value: v})
		}
	}
	return fields
}
