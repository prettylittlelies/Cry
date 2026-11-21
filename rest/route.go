package rest

import (
	"fmt"
	"net/url"
	"strings"
)

type Param struct {
	Name  string
	Value string
	Raw   bool
}

type Route struct {
	Method   string
	Template string
	Params   []Param
}

var majorParams = []string{"channel_id", "guild_id", "webhook_id", "webhook_token"}

func NewRoute(method, template string, params ...Param) Route {
	return Route{Method: method, Template: template, Params: params}
}

type ParamValue interface {
	~string | ~uint64
}

func P[T ParamValue](name string, value T) Param {
	return Param{Name: name, Value: fmt.Sprint(value)}
}

func Raw(name, value string) Param {
	return Param{Name: name, Value: value, Raw: true}
}

func (r Route) Path() string {
	path := r.Template
	for _, p := range r.Params {
		v := p.Value
		if !p.Raw {
			v = url.PathEscape(v)
		}
		path = strings.ReplaceAll(path, "{"+p.Name+"}", v)
	}
	return path
}

func (r Route) Key() string {
	return r.Method + " " + r.Template
}

func (r Route) Major() string {
	if strings.HasPrefix(r.Template, "/users/@me") {
		return ""
	}
	var parts []string
	for _, name := range majorParams {
		for _, p := range r.Params {
			if p.Name == name {
				parts = append(parts, p.Value)
			}
		}
	}
	return strings.Join(parts, "+")
}
