package rest

import (
	"net/http"
	"strconv"

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

func (b *browser) headers(method string, h http.Header) {
	h.Set("X-Super-Properties", b.identity.Session.Encoded())
	h.Set("Accept-Language", b.identity.Host.AcceptLanguage())
	h.Set("X-Discord-Locale", b.locale)
	if b.timezone != "" {
		h.Set("X-Discord-Timezone", b.timezone)
	}
	if b.debug != "" {
		h.Set("X-Debug-Options", b.debug)
	}
	h.Set("Sec-CH-UA", b.secCHUA)
	h.Set("Sec-CH-UA-Mobile", "?"+strconv.Itoa(btoi(b.identity.Mobile)))
	h.Set("Sec-CH-UA-Platform", strconv.Quote(b.identity.Host.Platform.ClientHint()))
	h.Set("User-Agent", b.identity.Session.UserAgent())
	h.Set("Accept", "*/*")
	if method != http.MethodGet && method != http.MethodHead {
		h.Set("Origin", b.origin)
	}
	h.Set("Sec-Fetch-Site", "same-origin")
	h.Set("Sec-Fetch-Mode", "cors")
	h.Set("Sec-Fetch-Dest", "empty")
	h.Set("Referer", b.referer)
	h.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	h.Set("Priority", "u=1, i")
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
