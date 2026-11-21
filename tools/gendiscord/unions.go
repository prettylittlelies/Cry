package main

import (
	"bytes"
	"fmt"
	"github.com/prettylittlelies/cry/tools/internal/naming"
	"strings"
)

type union struct {
	Enum      string
	Keyword   string
	Interface string
	Slice     string
	Single    string
	Unknown   string
	members   []unionMember
	enumName  string
}

type unionMember struct {
	structName string
	constName  string
}

var unions = []*union{{
	Enum:      "Component Type",
	Keyword:   "component",
	Interface: "Component",
	Slice:     "Components",
	Single:    "AnyComponent",
	Unknown:   "UnknownComponent",
}}

func bindUnions(enums []*namedEnum, r *resolver) {
	for _, u := range unions {
		for _, e := range enums {
			if e.enum.Title != u.Enum {
				continue
			}
			u.enumName = e.name
			for _, v := range e.enum.Values {
				if name, ok := r.structNamed(strings.ToLower(strings.ReplaceAll(v.Name, "_", " ")), e.enum.Doc); ok {
					u.members = append(u.members, unionMember{structName: name, constName: e.name + naming.GoName(v.Name)})
				}
			}
		}
		r.unions = append(r.unions, u)
	}
}

func (r *resolver) unionFor(name string) *union {
	for _, u := range r.unions {
		if u.enumName != "" && strings.HasSuffix(strings.ToLower(name), u.Keyword) {
			return u
		}
	}
	return nil
}

func emitUnions() []byte {
	var b bytes.Buffer
	for _, u := range unions {
		if u.enumName == "" {
			continue
		}
		et := u.enumName
		fmt.Fprintf(&b, "type %s interface {\n\t%s() %s\n}\n\n", u.Interface, et, et)
		fmt.Fprintf(&b, "type %s struct {\n\tType %s `json:\"type\"`\n\tRaw json.RawMessage `json:\"-\"`\n}\n\n", u.Unknown, et)
		fmt.Fprintf(&b, "func (u *%s) %s() %s { return u.Type }\n\n", u.Unknown, et, et)
		fmt.Fprintf(&b, "func (u *%s) MarshalJSON() ([]byte, error) { return u.Raw, nil }\n\n", u.Unknown)
		for _, m := range u.members {
			fmt.Fprintf(&b, "func (*%s) %s() %s { return %s }\n\n", m.structName, et, et, m.constName)
		}
		fmt.Fprintf(&b, "func decode%s(raw json.RawMessage) (%s, error) {\n", u.Interface, u.Interface)
		fmt.Fprintf(&b, "\tvar probe struct {\n\t\tType %s `json:\"type\"`\n\t}\n", et)
		b.WriteString("\tif err := json.Unmarshal(raw, &probe); err != nil {\n\t\treturn nil, err\n\t}\n")
		fmt.Fprintf(&b, "\tvar v %s\n\tswitch probe.Type {\n", u.Interface)
		for _, m := range u.members {
			fmt.Fprintf(&b, "\tcase %s:\n\t\tv = new(%s)\n", m.constName, m.structName)
		}
		fmt.Fprintf(&b, "\tdefault:\n\t\treturn &%s{Type: probe.Type, Raw: raw}, nil\n\t}\n", u.Unknown)
		b.WriteString("\treturn v, json.Unmarshal(raw, v)\n}\n\n")
		fmt.Fprintf(&b, "type %s []%s\n\n", u.Slice, u.Interface)
		fmt.Fprintf(&b, "func (c *%s) UnmarshalJSON(b []byte) error {\n", u.Slice)
		b.WriteString("\tvar raws []json.RawMessage\n\tif err := json.Unmarshal(b, &raws); err != nil {\n\t\treturn err\n\t}\n")
		fmt.Fprintf(&b, "\tout := make(%s, 0, len(raws))\n\tfor _, raw := range raws {\n", u.Slice)
		fmt.Fprintf(&b, "\t\tv, err := decode%s(raw)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tout = append(out, v)\n\t}\n\t*c = out\n\treturn nil\n}\n\n", u.Interface)
		fmt.Fprintf(&b, "type %s struct {\n\t%s\n}\n\n", u.Single, u.Interface)
		fmt.Fprintf(&b, "func (a *%s) UnmarshalJSON(b []byte) error {\n\tv, err := decode%s(b)\n\ta.%s = v\n\treturn err\n}\n\n", u.Single, u.Interface, u.Interface)
		fmt.Fprintf(&b, "func (a %s) MarshalJSON() ([]byte, error) { return json.Marshal(a.%s) }\n\n", u.Single, u.Interface)
	}
	return b.Bytes()
}
