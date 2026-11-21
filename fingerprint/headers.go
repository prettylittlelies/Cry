package fingerprint

import (
	"net/http"
	"slices"
	"strings"

	"golang.org/x/net/http2/hpack"
)

type HeaderOrder struct {
	Leading  []string
	Custom   []string
	Trailing []string
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

func ChromiumRendererOrder() HeaderOrder {
	return HeaderOrder{
		Leading: []string{"sec-ch-ua-platform"},
		Custom: []string{
			"authorization",
			"x-captcha-key",
			"x-captcha-rqtoken",
			"x-captcha-session-id",
			"x-super-properties",
			"x-context-properties",
			"x-discord-locale",
			"x-discord-timezone",
			"x-debug-options",
			"x-audit-log-reason",
			"content-type",
		},
		Trailing: []string{
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
		},
	}
}

func (o HeaderOrder) Fields(h http.Header) []hpack.HeaderField {
	lower := make(map[string][]string, len(h))
	for k, v := range h {
		name := strings.ToLower(k)
		if _, skip := hopByHop[name]; !skip {
			lower[name] = append(lower[name], v...)
		}
	}
	fields := make([]hpack.HeaderField, 0, len(lower))
	take := func(name string) {
		for _, v := range lower[name] {
			fields = append(fields, hpack.HeaderField{Name: name, Value: v})
		}
		delete(lower, name)
	}
	for _, n := range o.Leading {
		take(n)
	}
	for _, n := range o.Custom {
		take(n)
	}
	trailing := make(map[string]struct{}, len(o.Trailing))
	for _, n := range o.Trailing {
		trailing[n] = struct{}{}
	}
	var unknown []string
	for n := range lower {
		if _, ok := trailing[n]; !ok {
			unknown = append(unknown, n)
		}
	}
	slices.Sort(unknown)
	for _, n := range unknown {
		take(n)
	}
	for _, n := range o.Trailing {
		take(n)
	}
	return fields
}
