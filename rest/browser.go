package rest

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/prettylittlelies/cry/build"
	"github.com/prettylittlelies/cry/fingerprint"
	"github.com/prettylittlelies/cry/properties"
)

const (
	DefaultOrigin  = "https://discord.com"
	DefaultReferer = DefaultOrigin + "/channels/@me"
)

type Identity struct {
	Session *properties.Session
	Host    properties.Host
	Chrome  build.Version
	Mobile  bool
}

type site string

const (
	sameOrigin site = "same-origin"
	crossSite  site = "cross-site"
)

type browser struct {
	identity Identity
	secCHUA  string
	origin   string
	referer  string
	locale   string
	timezone string
	debug    string
}

func newBrowser(id Identity) *browser {
	return &browser{
		identity: id,
		secCHUA:  fingerprint.SecCHUA(fingerprint.Brands(id.Chrome.Major(), "")),
		origin:   DefaultOrigin,
		referer:  DefaultReferer,
		locale:   id.Host.Locale,
		timezone: id.Host.Timezone,
	}
}

func (b *browser) discord(h http.Header) {
	h.Set("X-Super-Properties", b.identity.Session.Encoded())
	h.Set("Accept-Language", b.identity.Host.AcceptLanguage())
	h.Set("X-Discord-Locale", b.locale)
	if b.timezone != "" {
		h.Set("X-Discord-Timezone", b.timezone)
	}
	if b.debug != "" {
		h.Set("X-Debug-Options", b.debug)
	}
}

func (b *browser) hints(h http.Header) {
	h.Set("Sec-CH-UA", b.secCHUA)
	h.Set("Sec-CH-UA-Mobile", "?"+strconv.Itoa(btoi(b.identity.Mobile)))
	h.Set("Sec-CH-UA-Platform", strconv.Quote(b.identity.Host.Platform.ClientHint()))
}

func (b *browser) fetch(h http.Header, s site, referer string) {
	h.Set("Sec-Fetch-Site", string(s))
	h.Set("Sec-Fetch-Mode", "cors")
	h.Set("Sec-Fetch-Dest", "empty")
	h.Set("Referer", referer)
	h.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	if h.Get("Accept-Language") == "" {
		h.Set("Accept-Language", b.locale)
	}
	h.Set("Priority", "u=1, i")
}

func (b *browser) headers(method string, h http.Header) {
	b.discord(h)
	b.hints(h)
	h.Set("User-Agent", b.identity.Session.UserAgent())
	h.Set("Accept", "*/*")
	if method != http.MethodGet && method != http.MethodHead {
		h.Set("Origin", b.origin)
	}
	b.fetch(h, sameOrigin, b.referer)
}

func (b *browser) crossSite(h http.Header) {
	b.hints(h)
	h.Set("User-Agent", b.identity.Session.UserAgent())
	h.Set("Accept", "*/*")
	h.Set("Origin", b.origin)
	b.fetch(h, crossSite, b.origin+"/")
}

func (b *browser) preflight(method string, unsafe []string) http.Header {
	h := http.Header{}
	h.Set("Accept", "*/*")
	h.Set("Access-Control-Request-Method", method)
	if len(unsafe) > 0 {
		h.Set("Access-Control-Request-Headers", strings.Join(unsafe, ","))
	}
	h.Set("Origin", b.origin)
	h.Set("User-Agent", b.identity.Session.UserAgent())
	b.fetch(h, crossSite, b.origin+"/")
	return h
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
