package main

import (
	"regexp"
	"strings"
)

type resolver struct {
	structs map[string][]*Structure
	names   map[*Structure]string
	enums   map[string]*namedEnum
	imports map[string]bool
}

type namedEnum struct {
	enum *Enum
	name string
	kind string
}

var (
	anchorLink   = regexp.MustCompile(`\]\(([^)#]*)#([^)]+)\)`)
	objectSuffix = regexp.MustCompile(`(?i)\s+(object|objcet|objects)$`)
)

func newResolver(spec *Spec, names map[*Structure]string, enums []*namedEnum) *resolver {
	r := &resolver{structs: map[string][]*Structure{}, names: names, enums: map[string]*namedEnum{}, imports: map[string]bool{}}
	for _, s := range spec.Structures {
		k := strings.ToLower(s.Title)
		r.structs[k] = append(r.structs[k], s)
	}
	for _, e := range enums {
		r.enums[e.enum.Doc+"#"+e.enum.Anchor] = e
	}
	return r
}

func (r *resolver) structNamed(title, doc string) (string, bool) {
	cands := r.structs[strings.ToLower(title)]
	for _, s := range cands {
		if s.Doc == doc {
			return r.names[s], true
		}
	}
	if len(cands) > 0 {
		return r.names[cands[0]], true
	}
	return "", false
}

var wrappers = map[string]string{
	"stringvalue": "string",
	"boolvalue":   "bool",
	"int32value":  "int64",
	"int64value":  "int64",
	"uint32value": "int64",
	"uint64value": "uint64",
	"floatvalue":  "float64",
	"doublevalue": "float64",
	"bytesvalue":  "[]byte",
}

var qualifiers = []string{"partial ", "private ", "full ", "public ", "extended "}

func candidates(name string) []string {
	out := []string{name}
	for changed := true; changed; {
		changed = false
		for _, q := range qualifiers {
			if rest, ok := strings.CutPrefix(name, q); ok {
				name, changed = rest, true
				out = append(out, name)
			}
		}
	}
	return out
}

func (r *resolver) object(name, doc string) string {
	for _, candidate := range candidates(strings.ToLower(strings.TrimSpace(name))) {
		for _, form := range []string{candidate, candidate + "s", strings.TrimSuffix(candidate, "s")} {
			if t, ok := r.structNamed(form, doc); ok {
				return t
			}
		}
	}
	r.imports["encoding/json"] = true
	return "json.RawMessage"
}

func (r *resolver) linkedEnum(desc, doc string) *namedEnum {
	for _, m := range anchorLink.FindAllStringSubmatch(desc, -1) {
		target := doc
		if m[1] != "" {
			parts := strings.Split(strings.Trim(m[1], "/"), "/")
			target = parts[len(parts)-1]
		}
		if e, ok := r.enums[target+"#"+m[2]]; ok {
			return e
		}
	}
	return nil
}

func splitTop(s string) []string {
	var parts []string
	depth, start := 0, 0
	for i, c := range s {
		switch c {
		case '[':
			depth++
		case ']':
			depth--
		case ',':
			if depth == 0 {
				parts = append(parts, strings.TrimSpace(s[start:i]))
				start = i + 1
			}
		}
	}
	return append(parts, strings.TrimSpace(s[start:]))
}

func (r *resolver) base(t, doc string) string {
	t = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(t), "?"))
	lt := strings.ToLower(t)
	switch {
	case strings.HasPrefix(lt, "array[") && strings.HasSuffix(lt, "]"):
		elems := splitTop(t[6 : len(t)-1])
		first := r.base(elems[0], doc)
		for _, e := range elems[1:] {
			if r.base(e, doc) != first {
				r.imports["encoding/json"] = true
				return "[]json.RawMessage"
			}
		}
		return "[]" + r.field(elems[0], doc, "")
	case strings.HasPrefix(lt, "map[") && strings.HasSuffix(lt, "]"):
		kv := splitTop(t[4 : len(t)-1])
		if len(kv) == 2 {
			key := r.base(kv[0], doc)
			if key != "string" && key != "snowflake.ID" && key != "int64" {
				key = "string"
			}
			return "map[" + key + "]" + r.field(kv[1], doc, "")
		}
	case lt == "snowflake":
		r.imports["github.com/prettylittlelies/cry/snowflake"] = true
		return "snowflake.ID"
	case lt == "string" || lt == "iso8601 date" || lt == "iso4217 currency code" || lt == "locale":
		return "string"
	case lt == "integer" || lt == "unix timestamp" || lt == "int":
		return "int64"
	case lt == "boolean" || lt == "bool":
		return "bool"
	case lt == "float" || lt == "number" || lt == "double":
		return "float64"
	case lt == "iso8601 timestamp":
		r.imports["time"] = true
		return "time.Time"
	case lt == "image data":
		return "ImageData"
	case wrappers[lt] != "":
		return wrappers[lt]
	case objectSuffix.MatchString(lt):
		return r.object(objectSuffix.ReplaceAllString(t, ""), doc)
	}
	if name, ok := r.structNamed(lt, doc); ok {
		return name
	}
	r.imports["encoding/json"] = true
	return "json.RawMessage"
}

func nilable(t string) bool {
	return strings.HasPrefix(t, "[]") || strings.HasPrefix(t, "map[") || t == "json.RawMessage" || strings.HasPrefix(t, "*")
}

func (r *resolver) field(t, doc, desc string) string {
	nullable := strings.HasPrefix(strings.TrimSpace(t), "?")
	g := r.base(t, doc)
	if e := r.linkedEnum(desc, doc); e != nil && (g == "int64" && e.kind != "string" || g == "string" && e.kind == "string") {
		g = e.name
	}
	if nullable && !nilable(g) {
		return "*" + g
	}
	return g
}
