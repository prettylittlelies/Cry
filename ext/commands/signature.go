package commands

import (
	"reflect"
	"strconv"
	"strings"
)

func (c *Command) ShortDoc() string {
	if c.Brief != "" {
		return c.Brief
	}
	first, _, _ := strings.Cut(c.Help, "\n")
	return first
}

func (c *Command) Signature() string {
	if c.Usage != "" {
		return c.Usage
	}
	result := make([]string, 0, len(c.Params))
	for _, p := range c.Params {
		result = append(result, c.paramSignature(p))
	}
	return strings.Join(result, " ")
}

func (c *Command) paramSignature(p *Parameter) string {
	name := p.DisplayedName
	if name == "" {
		name = p.Name
	}
	if p.isAttachment() {
		return attachmentSignature(p, name)
	}
	if len(p.Literals) > 0 {
		quote := p.converter().Kind() == reflect.String
		values := make([]string, len(p.Literals))
		for i, l := range p.Literals {
			values[i] = l
			if quote {
				values[i] = strconv.Quote(l)
			}
		}
		name = strings.Join(values, "|")
	}
	if !p.Required() {
		suffix := ""
		if p.Greedy {
			suffix = "..."
		}
		if d, ok := p.DisplayDefault(); ok && d != "" {
			return "[" + name + "=" + d + "]" + suffix
		}
		return "[" + name + "]"
	}
	switch {
	case p.Kind == VarPositional && c.RequireVarPositional:
		return "<" + name + "...>"
	case p.Kind == VarPositional:
		return "[" + name + "...]"
	case p.Greedy:
		return "[" + name + "]..."
	case p.Optional():
		return "[" + name + "]"
	}
	return "<" + name + ">"
}
