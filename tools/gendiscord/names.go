package main

import (
	"strings"
	"unicode"
)

var initialisms = map[string]string{
	"id": "ID", "ids": "IDs", "url": "URL", "urls": "URLs", "api": "API", "nsfw": "NSFW", "mfa": "MFA",
	"dm": "DM", "dms": "DMs", "ip": "IP", "ttl": "TTL", "uri": "URI", "uuid": "UUID", "rtc": "RTC",
	"sku": "SKU", "skus": "SKUs", "tts": "TTS", "afk": "AFK", "oauth2": "OAuth2", "ui": "UI",
	"json": "JSON", "html": "HTML", "http": "HTTP", "sms": "SMS", "vr": "VR", "pc": "PC",
	"cdn": "CDN", "gif": "GIF", "png": "PNG", "jpeg": "JPEG", "rpc": "RPC", "ssrc": "SSRC",
	"qos": "QoS", "etf": "ETF", "2fa": "2FA", "utc": "UTC", "os": "OS", "nfc": "NFC", "ws": "WS",
}

func words(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsDigit(r) })
}

func GoName(s string) string {
	var b strings.Builder
	for _, w := range words(s) {
		lw := strings.ToLower(w)
		if v, ok := initialisms[lw]; ok {
			b.WriteString(v)
			continue
		}
		b.WriteString(strings.ToUpper(lw[:1]))
		b.WriteString(lw[1:])
	}
	out := b.String()
	if out == "" {
		return "X"
	}
	if unicode.IsDigit(rune(out[0])) {
		out = "N" + out
	}
	return out
}

type namer struct {
	taken map[string]string
}

var docPriority = []string{"gateway-events", "user", "guild", "channel", "message", "presence", "relationships", "emoji", "sticker", "application", "user-settings"}

func priority(doc string) int {
	for i, d := range docPriority {
		if d == doc {
			return i
		}
	}
	return len(docPriority)
}

func (n *namer) assign(doc, title string, count map[string]int) string {
	base := GoName(title)
	if _, used := n.taken[base]; count[base] > 1 && used {
		base = GoName(doc) + base
	}
	for i := 2; ; i++ {
		if _, used := n.taken[base]; !used {
			n.taken[base] = title
			return base
		}
		base = GoName(title) + GoName(doc)
		if i > 2 {
			base += GoName(strings.Repeat("x", i))
		}
	}
}
