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
	b.WriteString("package discord\n\nimport \"strconv\"\n\n")
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
		if ne.kind != "string" {
			writeStringer(&b, ne, seen)
		}
	}
	return gofmt(b.Bytes())
}

func writeStringer(b *bytes.Buffer, ne *namedEnum, seen map[string]bool) {
	fmt.Fprintf(b, "func (v %s) String() string {\n\tswitch v {\n", ne.name)
	done := map[uint64]bool{}
	for _, v := range ne.enum.Values {
		name := ne.name + GoName(v.Name)
		key, ok := numericValue(v.Value)
		if !ok || !seen[name] || done[key] || GoName(v.Name) == "X" {
			continue
		}
		done[key] = true
		fmt.Fprintf(b, "\tcase %s:\n\t\treturn %q\n", name, v.Name)
	}
	if ne.kind == "flags" {
		b.WriteString("\t}\n\treturn strconv.FormatUint(uint64(v), 10)\n}\n\n")
		return
	}
	b.WriteString("\t}\n\treturn strconv.FormatInt(int64(v), 10)\n}\n\n")
}

type genField struct {
	name, typ, tag string
}

var extras = map[string][]genField{
	"GatewayGuild": {{name: "ID", typ: "snowflake.ID", tag: "id"}},
	"GuildMember":  {{name: "UserID", typ: "snowflake.ID", tag: "user_id,omitzero"}},
	"PartialGuild": {{name: "OwnerID", typ: "snowflake.ID", tag: "owner_id,omitzero"}},
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
	body.Write(emitUnions())
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

func numericValue(v string) (uint64, bool) {
	if m := shiftValue.FindStringSubmatch(v); m != nil {
		base, _ := strconv.ParseUint(m[1], 10, 64)
		bit, _ := strconv.Atoi(m[2])
		return base << bit, true
	}
	n, err := strconv.ParseInt(v, 10, 64)
	return uint64(n), err == nil
}
