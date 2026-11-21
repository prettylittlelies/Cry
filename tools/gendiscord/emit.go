package main

import (
	"bytes"
	"fmt"
	"go/format"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	shiftValue = regexp.MustCompile(`^(\d+)\s*<<\s*(\d+)$`)
	intValue   = regexp.MustCompile(`^-?\d+$`)
)

func enumKind(e *Enum) string {
	ints, shifts, strs := 0, 0, 0
	for _, v := range e.Values {
		switch {
		case shiftValue.MatchString(v.Value):
			if bit, _ := strconv.Atoi(shiftValue.FindStringSubmatch(v.Value)[2]); bit >= 64 {
				return ""
			}
			shifts++
		case intValue.MatchString(v.Value):
			ints++
		case v.Value != "" && !strings.ContainsAny(v.Value, " /"):
			strs++
		}
	}
	switch {
	case shifts > 0 && shifts+ints == len(e.Values):
		return "flags"
	case ints == len(e.Values):
		return "int"
	case strs == len(e.Values):
		return "string"
	}
	return ""
}

func emitEnums(enums []*namedEnum) []byte {
	var b bytes.Buffer
	b.WriteString("package discord\n\n")
	for _, ne := range enums {
		goType := map[string]string{"flags": "uint64", "int": "int64", "string": "string"}[ne.kind]
		fmt.Fprintf(&b, "type %s %s\n\nconst (\n", ne.name, goType)
		seen := map[string]bool{}
		for _, v := range ne.enum.Values {
			name := ne.name + GoName(v.Name)
			if seen[name] || GoName(v.Name) == "X" {
				continue
			}
			seen[name] = true
			switch ne.kind {
			case "flags":
				if m := shiftValue.FindStringSubmatch(v.Value); m != nil {
					fmt.Fprintf(&b, "\t%s %s = %s << %s\n", name, ne.name, m[1], m[2])
				} else {
					fmt.Fprintf(&b, "\t%s %s = %s\n", name, ne.name, v.Value)
				}
			case "int":
				fmt.Fprintf(&b, "\t%s %s = %s\n", name, ne.name, v.Value)
			case "string":
				fmt.Fprintf(&b, "\t%s %s = %s\n", name, ne.name, strconv.Quote(v.Value))
			}
		}
		b.WriteString(")\n\n")
	}
	return gofmt(b.Bytes())
}

type genField struct {
	name, typ, tag string
}

var extras = map[string][]genField{
	"GatewayGuild": {{name: "ID", typ: "snowflake.ID", tag: "id"}},
	"GuildMember":  {{name: "UserID", typ: "snowflake.ID", tag: "user_id,omitzero"}},
}

func structFields(spec *Spec, r *resolver, overrides map[string]string) map[string][]genField {
	out := map[string][]genField{}
	defer func() {
		for name, fs := range extras {
			out[name] = append(out[name], fs...)
		}
	}()
	for _, s := range spec.Structures {
		name := r.names[s]
		used := map[string]bool{}
		for _, f := range s.Fields {
			goField := GoName(f.Name)
			if used[goField] {
				continue
			}
			used[goField] = true
			t := r.field(f.Type, s.Doc, f.Description)
			if o, ok := overrides[name+"."+f.Name]; ok {
				t = o
			} else if o, ok := overrides["*."+f.Name]; ok && strings.TrimPrefix(t, "*") == "string" {
				t = strings.Replace(t, "string", o, 1)
			}
			tag := f.Name
			if f.Optional {
				tag += omit(t)
			}
			out[name] = append(out[name], genField{name: goField, typ: t, tag: tag})
		}
	}
	return out
}

func breakCycles(fields map[string][]genField) {
	reaches := func(from, to string) bool {
		seen := map[string]bool{}
		var walk func(string) bool
		walk = func(n string) bool {
			if n == to {
				return true
			}
			if seen[n] {
				return false
			}
			seen[n] = true
			for _, f := range fields[n] {
				if _, ok := fields[f.typ]; ok && walk(f.typ) {
					return true
				}
			}
			return false
		}
		return walk(from)
	}
	for changed := true; changed; {
		changed = false
		for owner, fs := range fields {
			for i, f := range fs {
				if _, ok := fields[f.typ]; ok && reaches(f.typ, owner) {
					fs[i].typ = "*" + f.typ
					fs[i].tag = strings.Replace(fs[i].tag, ",omitzero", ",omitempty", 1)
					changed = true
				}
			}
		}
	}
}

func emitStructs(spec *Spec, r *resolver, overrides map[string]string) ([]byte, map[string][]genField) {
	fields := structFields(spec, r, overrides)
	breakCycles(fields)
	names := make([]string, 0, len(fields))
	for n := range fields {
		names = append(names, n)
	}
	sort.Strings(names)
	var body bytes.Buffer
	for _, name := range names {
		fmt.Fprintf(&body, "type %s struct {\n", name)
		for _, f := range fields[name] {
			fmt.Fprintf(&body, "\t%s %s `json:%q`\n", f.name, f.typ, f.tag)
		}
		body.WriteString("}\n\n")
	}
	body.Write(eventsBody(spec, r, fields))
	var head bytes.Buffer
	head.WriteString("package discord\n\n")
	if len(r.imports) > 0 {
		head.WriteString("import (\n")
		keys := make([]string, 0, len(r.imports))
		for k := range r.imports {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Fprintf(&head, "\t%q\n", k)
		}
		head.WriteString(")\n\n")
	}
	head.Write(body.Bytes())
	return gofmt(head.Bytes()), fields
}

func omit(t string) string {
	if nilable(t) {
		return ",omitempty"
	}
	return ",omitzero"
}

func gofmt(src []byte) []byte {
	out, err := format.Source(src)
	if err != nil {
		panic(fmt.Sprintf("gendiscord: generated code does not format: %v", err))
	}
	return out
}
