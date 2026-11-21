package rest

import (
	"net/http"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	corsDefaultMaxAge = 5 * time.Second
	corsMaxAge        = 2 * time.Hour
)

var corsSafelistedContentTypes = map[string]struct{}{
	"application/x-www-form-urlencoded": {},
	"multipart/form-data":               {},
	"text/plain":                        {},
}

type corsEntry struct {
	methods   map[string]struct{}
	headers   map[string]struct{}
	anyMethod bool
	anyHeader bool
	expiresAt time.Time
}

type corsCache struct {
	mu      sync.Mutex
	entries map[string]corsEntry
	now     func() time.Time
}

func newCORSCache(now func() time.Time) *corsCache {
	return &corsCache{entries: make(map[string]corsEntry), now: now}
}

func simpleMethod(m string) bool {
	return m == http.MethodGet || m == http.MethodHead || m == http.MethodPost
}

func safelisted(name, value string) bool {
	switch name {
	case "accept", "accept-language", "content-language":
		return true
	case "content-type":
		mt, _, _ := strings.Cut(strings.ToLower(value), ";")
		_, ok := corsSafelistedContentTypes[strings.TrimSpace(mt)]
		return ok
	}
	return false
}

func unsafeHeaders(h http.Header) []string {
	var out []string
	for k, v := range h {
		name := strings.ToLower(k)
		if !safelisted(name, strings.Join(v, ",")) {
			out = append(out, name)
		}
	}
	slices.Sort(out)
	return out
}

func needsPreflight(method string, unsafe []string) bool {
	return !simpleMethod(method) || len(unsafe) > 0
}

func (c *corsCache) covers(key, method string, unsafe []string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	e, ok := c.entries[key]
	if !ok || !c.now().Before(e.expiresAt) {
		delete(c.entries, key)
		return false
	}
	if _, ok := e.methods[method]; !ok && !e.anyMethod && !simpleMethod(method) {
		return false
	}
	for _, h := range unsafe {
		if _, ok := e.headers[h]; !ok && !e.anyHeader {
			return false
		}
	}
	return true
}

func tokens(v string) map[string]struct{} {
	out := make(map[string]struct{})
	for t := range strings.SplitSeq(v, ",") {
		if t = strings.ToLower(strings.TrimSpace(t)); t != "" {
			out[t] = struct{}{}
		}
	}
	return out
}

func (c *corsCache) store(key string, h http.Header) {
	methods := tokens(strings.ToUpper(h.Get("Access-Control-Allow-Methods")))
	upper := make(map[string]struct{}, len(methods))
	for m := range methods {
		upper[strings.ToUpper(m)] = struct{}{}
	}
	headers := tokens(h.Get("Access-Control-Allow-Headers"))
	_, anyMethod := upper["*"]
	_, anyHeader := headers["*"]
	age := corsDefaultMaxAge
	if s, err := strconv.Atoi(h.Get("Access-Control-Max-Age")); err == nil && s >= 0 {
		age = min(time.Duration(s)*time.Second, corsMaxAge)
	}
	c.mu.Lock()
	c.entries[key] = corsEntry{methods: upper, headers: headers, anyMethod: anyMethod, anyHeader: anyHeader, expiresAt: c.now().Add(age)}
	c.mu.Unlock()
}
