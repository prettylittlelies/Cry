package commands

import "reflect"

type Flag struct {
	Name        string
	Aliases     []string
	Attribute   string
	Type        reflect.Type
	Default     string
	HasDefault  bool
	MaxArgs     int
	Override    bool
	Description string
	Positional  bool
	CastToDict  bool
}

func (f *Flag) Required() bool {
	return !f.HasDefault
}
