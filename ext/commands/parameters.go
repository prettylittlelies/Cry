package commands

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

type ParamKind int

const (
	Positional ParamKind = iota
	KeywordOnly
	VarPositional
)

type Parameter struct {
	Name             string
	DisplayedName    string
	Description      string
	DisplayedDefault string
	Kind             ParamKind
	Type             reflect.Type
	Default          string
	HasDefault       bool
	Greedy           bool
	Literals         []string
	field            []int
}

func (p *Parameter) Optional() bool {
	return p.Type.Kind() == reflect.Pointer
}

func (p *Parameter) Required() bool {
	return !p.HasDefault
}

func (p *Parameter) converter() reflect.Type {
	t := p.Type
	if p.Kind == VarPositional || p.Greedy {
		t = t.Elem()
	}
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t
}

func snake(s string) string {
	var b strings.Builder
	runes := []rune(s)
	for i, r := range runes {
		if unicode.IsUpper(r) {
			if i > 0 && (unicode.IsLower(runes[i-1]) || i+1 < len(runes) && unicode.IsLower(runes[i+1])) {
				b.WriteByte('_')
			}
			r = unicode.ToLower(r)
		}
		b.WriteRune(r)
	}
	return b.String()
}

func signature(t reflect.Type) ([]*Parameter, error) {
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("commands: parameters must be a struct, got %s", t)
	}
	var params []*Parameter
	for i := range t.NumField() {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}
		p := &Parameter{
			Name:             f.Tag.Get("name"),
			DisplayedName:    f.Tag.Get("displayed_name"),
			Description:      f.Tag.Get("description"),
			DisplayedDefault: f.Tag.Get("displayed_default"),
			Type:             f.Type,
			field:            f.Index,
		}
		if p.Name == "" {
			p.Name = snake(f.Name)
		}
		p.Default, p.HasDefault = f.Tag.Lookup("default")
		if lit := f.Tag.Get("literal"); lit != "" {
			p.Literals = strings.Split(lit, "|")
		}
		switch f.Tag.Get("commands") {
		case "rest":
			p.Kind = KeywordOnly
		case "varargs":
			p.Kind = VarPositional
		case "greedy":
			p.Greedy = true
		}
		if (p.Kind == VarPositional || p.Greedy) && f.Type.Kind() != reflect.Slice {
			return nil, fmt.Errorf("commands: parameter %s must be a slice to be %s", p.Name, f.Tag.Get("commands"))
		}
		params = append(params, p)
		if p.Kind == KeywordOnly {
			break
		}
	}
	return params, nil
}
