package main

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var (
	innerPayload = regexp.MustCompile(`(?is)inner payload(.{0,60}?)\[([^\]]+)\]`)
	extraLink    = regexp.MustCompile(`\(#([a-z0-9-]*extra-fields)\)`)
)

func wireName(event string) string {
	return strings.ToUpper(strings.Join(words(event), "_"))
}

func tagName(tag string) string {
	name, _, _ := strings.Cut(tag, ",")
	return name
}

func merge(base []genField, extra ...[]genField) []genField {
	out := append([]genField(nil), base...)
	for _, fs := range extra {
		for _, f := range fs {
			replaced := false
			for i := range out {
				if tagName(out[i].tag) == tagName(f.tag) {
					out[i], replaced = f, true
				}
			}
			if !replaced {
				out = append(out, f)
			}
		}
	}
	return out
}

func eventsBody(spec *Spec, r *resolver, fields map[string][]genField) []byte {
	byAnchor := map[string]*Structure{}
	for _, s := range spec.Structures {
		byAnchor[s.Anchor] = s
	}
	type entry struct {
		wire, goName string
	}
	var entries []entry
	var body bytes.Buffer
	seen := map[string]bool{}
	for _, ev := range spec.Events {
		prose := strings.Join(ev.Prose, " ")
		var base []genField
		var extras [][]genField
		for _, s := range ev.Structs {
			switch {
			case strings.HasSuffix(s.Title, "Extra Fields"):
				extras = append(extras, fields[r.names[s]])
			case base == nil && strings.EqualFold(s.Title, ev.Name):
				base = fields[r.names[s]]
			}
		}
		sliceOf := ""
		if base == nil {
			if m := innerPayload.FindStringSubmatch(prose); m != nil {
				if t := r.object(m[2], "gateway-events"); fields[t] != nil {
					if lead := strings.ToLower(m[1]); strings.Contains(lead, "list of") || strings.Contains(lead, "array of") {
						sliceOf = t
					} else {
						base = fields[t]
					}
				}
			}
		}
		if sliceOf != "" {
			name := GoName(ev.Name) + "Event"
			if !seen[name] {
				seen[name] = true
				entries = append(entries, entry{wire: wireName(ev.Name), goName: name})
				fmt.Fprintf(&body, "type %s []%s\n\n", name, sliceOf)
			}
			continue
		}
		for _, m := range extraLink.FindAllStringSubmatch(prose, -1) {
			if s, ok := byAnchor[m[1]]; ok {
				extras = append(extras, fields[r.names[s]])
			}
		}
		if base == nil && len(extras) == 0 {
			continue
		}
		name := GoName(ev.Name) + "Event"
		if seen[name] {
			continue
		}
		seen[name] = true
		entries = append(entries, entry{wire: wireName(ev.Name), goName: name})
		fmt.Fprintf(&body, "type %s struct {\n", name)
		for _, f := range merge(base, extras...) {
			fmt.Fprintf(&body, "\t%s %s `json:%q`\n", f.name, f.typ, f.tag)
		}
		body.WriteString("}\n\n")
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].wire < entries[j].wire })
	body.WriteString("var Events = map[string]func() any{\n")
	for _, e := range entries {
		fmt.Fprintf(&body, "\t%q: func() any { return new(%s) },\n", e.wire, e.goName)
	}
	body.WriteString("}\n")
	return body.Bytes()
}
