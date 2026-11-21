package fingerprint

import (
	"net/http"
	"slices"
	"strings"

	"golang.org/x/net/http2/hpack"
)

const Wildcard = "*"

type HeaderOrder []string

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
	discordRequest = []string{
		"content-type",
		Wildcard,
		"x-audit-log-reason",
		"x-context-properties",
		"x-failed-requests",
	}
	discordPrepared = []string{
		"authorization",
		"x-super-properties",
		"x-fingerprint",
		"x-installation-id",
	}
	discordLocale = []string{
		"x-discord-locale",
		"x-discord-timezone",
		"x-debug-options",
		"x-routing-key",
		"x-client-trace-id",
	}
	chromiumTrailing = []string{
		"content-length",
		"user-agent",
		"sec-ch-ua",
		"sec-ch-ua-mobile",
		"accept",
		"origin",
		"sec-fetch-site",
		"sec-fetch-mode",
		"sec-fetch-dest",
		"referer",
		"accept-encoding",
		"accept-language",
		"cookie",
		"priority",
	}
)

func DiscordWebOrder() HeaderOrder {
	return slices.Concat([]string{"sec-ch-ua-platform"}, discordRequest, discordPrepared, discordLocale, chromiumTrailing)
}

func DiscordDesktopOrder() HeaderOrder {
	trailing := slices.DeleteFunc(slices.Clone(chromiumTrailing), func(s string) bool { return s == "accept-language" })
	return slices.Concat([]string{"sec-ch-ua-platform"}, discordRequest, discordPrepared, []string{"accept-language"}, discordLocale, trailing)
}

func (o HeaderOrder) Fields(h http.Header) []hpack.HeaderField {
	values := make(map[string][]string, len(h))
	for k, v := range h {
		name := strings.ToLower(k)
		if _, skip := hopByHop[name]; !skip {
			values[name] = append(values[name], v...)
		}
	}
	known := make(map[string]struct{}, len(o))
	for _, n := range o {
		known[n] = struct{}{}
	}
	var unknown []string
	for n := range values {
		if _, ok := known[n]; !ok {
			unknown = append(unknown, n)
		}
	}
	slices.Sort(unknown)
	fields := make([]hpack.HeaderField, 0, len(values))
	emit := func(name string) {
		for _, v := range values[name] {
			fields = append(fields, hpack.HeaderField{Name: name, Value: v})
		}
	}
	for _, n := range o {
		if n != Wildcard {
			emit(n)
			continue
		}
		for _, u := range unknown {
			emit(u)
		}
		unknown = nil
	}
	for _, u := range unknown {
		emit(u)
	}
	return fields
}
