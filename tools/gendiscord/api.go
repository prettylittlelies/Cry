package main

import (
	"bytes"
	"fmt"
	"go/token"
	"regexp"
	"strings"
)

var (
	pathParam  = regexp.MustCompile(`\{([^}]+)\}`)
	returnsObj = regexp.MustCompile(`(?i)returns\s+(?:a|an|the)?\s*(list of|array of)?\s*(?:partial\s+)?\[([^\]]+)\]`)
	returns204 = regexp.MustCompile(`(?i)returns a 204`)
)

type pathArg struct {
	wire, goName, typ string
	raw               bool
}

func lowerFirst(s string) string {
	upper := 0
	for upper < len(s) && s[upper] >= 'A' && s[upper] <= 'Z' {
		upper++
	}
	switch {
	case upper == len(s):
		return strings.ToLower(s)
	case upper > 1:
		return strings.ToLower(s[:upper-1]) + s[upper-1:]
	case upper == 1:
		return strings.ToLower(s[:1]) + s[1:]
	}
	return s
}

func argName(wire string) string {
	n := lowerFirst(GoName(wire))
	if token.IsKeyword(n) || n == "ctx" || n == "opts" || n == "query" || n == "params" {
		n += "Arg"
	}
	return n
}

func pathArgs(url string) ([]pathArg, string) {
	var args []pathArg
	seen := map[string]bool{}
	template := pathParam.ReplaceAllStringFunc(url, func(m string) string {
		name := strings.Trim(m, "{}")
		wire := strings.ReplaceAll(name, ".", "_")
		if !seen[wire] {
			seen[wire] = true
			a := pathArg{wire: wire, goName: argName(wire), typ: "string", raw: strings.Contains(wire, "filename")}
			if strings.HasSuffix(wire, "_id") || wire == "id" {
				a.typ = "snowflake.ID"
			}
			args = append(args, a)
		}
		return "{" + wire + "}"
	})
	return args, template
}

func qualify(t string) string {
	core := strings.TrimLeft(t, "[]*")
	prefix := t[:len(t)-len(core)]
	switch {
	case core == "" || strings.Contains(core, "."):
		return t
	case core == "string" || core == "int64" || core == "bool" || core == "float64" || core == "uint64":
		return t
	}
	return prefix + "discord." + core
}

type returnKind int

const (
	returnNone returnKind = iota
	returnRaw
	returnOne
	returnList
)

func endpointReturn(ep *Endpoint, r *resolver, fields map[string][]genField) (returnKind, string) {
	if ep.Response != nil {
		return returnOne, r.names[ep.Response]
	}
	prose := strings.Join(ep.Prose, " ")
	if m := returnsObj.FindStringSubmatch(prose); m != nil {
		if t := r.object(m[2], ep.Doc); fields[t] != nil {
			if m[1] != "" {
				return returnList, t
			}
			return returnOne, t
		}
	}
	if returns204.MatchString(prose) {
		return returnNone, ""
	}
	return returnRaw, ""
}

func emitAPI(spec *Spec, r *resolver, fields map[string][]genField) []byte {
	var body bytes.Buffer
	names := map[string]int{}
	for _, ep := range spec.Endpoints {
		names[GoName(ep.Name)]++
	}
	used := map[string]bool{}
	for _, ep := range spec.Endpoints {
		method := GoName(ep.Name)
		if names[method] > 1 || used[method] {
			method = GoName(ep.Doc) + method
		}
		if used[method] {
			continue
		}
		used[method] = true
		args, template := pathArgs(ep.URL)
		kind, ret := endpointReturn(ep, r, fields)
		var sig, params []string
		sig = append(sig, "ctx context.Context")
		for _, a := range args {
			sig = append(sig, a.goName+" "+a.typ)
			ctor := "rest.P"
			if a.raw {
				ctor = "rest.Raw"
			}
			params = append(params, fmt.Sprintf("%s(%q, %s)", ctor, a.wire, a.goName))
		}
		if ep.Query != nil {
			sig = append(sig, "query "+qualify(r.names[ep.Query]))
		}
		if ep.Params != nil {
			sig = append(sig, "params "+qualify(r.names[ep.Params]))
		}
		sig = append(sig, "opts ...Option")
		result := map[returnKind]string{
			returnNone: "error",
			returnRaw:  "(json.RawMessage, error)",
			returnOne:  "(*" + qualify(ret) + ", error)",
			returnList: "([]" + qualify(ret) + ", error)",
		}[kind]
		fmt.Fprintf(&body, "func (a *API) %s(%s) %s {\n", method, strings.Join(sig, ", "), result)
		route := fmt.Sprintf("rest.NewRoute(%q, %q", ep.Method, template)
		if len(params) > 0 {
			route += ", " + strings.Join(params, ", ")
		}
		route += ")"
		fmt.Fprintf(&body, "\treq := &rest.Request{Route: %s}\n", route)
		if ep.Query != nil {
			body.WriteString("\treq.Query = rest.Values(query)\n")
		}
		if ep.Params != nil {
			body.WriteString("\treq.JSON = params\n")
		}
		switch kind {
		case returnNone:
			body.WriteString("\treturn none(ctx, a, req, opts)\n")
		case returnRaw:
			body.WriteString("\treturn raw(ctx, a, req, opts)\n")
		case returnOne:
			fmt.Fprintf(&body, "\treturn call[%s](ctx, a, req, opts)\n", qualify(ret))
		case returnList:
			fmt.Fprintf(&body, "\treturn list[%s](ctx, a, req, opts)\n", qualify(ret))
		}
		body.WriteString("}\n\n")
	}
	src := body.String()
	var head bytes.Buffer
	head.WriteString("package api\n\nimport (\n\t\"context\"\n")
	if strings.Contains(src, "json.RawMessage") {
		head.WriteString("\t\"encoding/json\"\n")
	}
	head.WriteString("\n")
	if strings.Contains(src, "discord.") {
		head.WriteString("\t\"github.com/prettylittlelies/cry/discord\"\n")
	}
	head.WriteString("\t\"github.com/prettylittlelies/cry/rest\"\n")
	if strings.Contains(src, "snowflake.") {
		head.WriteString("\t\"github.com/prettylittlelies/cry/snowflake\"\n")
	}
	head.WriteString(")\n\n")
	head.WriteString(src)
	return gofmt(head.Bytes())
}
