package commands

import (
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

type FlagConverter struct{}

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
	field       []int
}

func (f *Flag) Required() bool {
	return !f.HasDefault && f.Type.Kind() != reflect.Pointer && !constructible(f.Type)
}

type flagSpec struct {
	typ             reflect.Type
	flags           []*Flag
	byKey           map[string]*Flag
	aliases         map[string]string
	positional      *Flag
	caseInsensitive bool
	pattern         *regexp.Regexp
}

type flagSpecResult struct {
	spec *flagSpec
	err  error
}

var (
	flagMarker = reflect.TypeFor[FlagConverter]()
	flagSpecs  sync.Map
)

func IsFlagConverter(t reflect.Type) bool {
	if t.Kind() != reflect.Struct {
		return false
	}
	return slices.ContainsFunc(reflect.VisibleFields(t), func(f reflect.StructField) bool {
		return f.Anonymous && f.Type == flagMarker
	})
}

func FlagsOf(t reflect.Type) ([]*Flag, error) {
	spec, err := flagSpecFor(t)
	if err != nil {
		return nil, err
	}
	return slices.Clone(spec.flags), nil
}

func constructible(t reflect.Type) bool {
	spec, err := flagSpecFor(t)
	return err == nil && spec != nil && spec.constructible()
}

func (s *flagSpec) constructible() bool {
	return !slices.ContainsFunc(s.flags, (*Flag).Required)
}

func (s *flagSpec) key(name string) string {
	if s.caseInsensitive {
		return strings.ToLower(name)
	}
	return name
}

func flagSpecFor(t reflect.Type) (*flagSpec, error) {
	if !IsFlagConverter(t) {
		return nil, nil
	}
	if cached, ok := flagSpecs.Load(t); ok {
		r := cached.(flagSpecResult)
		return r.spec, r.err
	}
	spec, err := buildFlagSpec(t)
	cached, _ := flagSpecs.LoadOrStore(t, flagSpecResult{spec, err})
	r := cached.(flagSpecResult)
	return r.spec, r.err
}

func buildFlagSpec(t reflect.Type) (*flagSpec, error) {
	spec := &flagSpec{typ: t, byKey: map[string]*Flag{}, aliases: map[string]string{}}
	delimiter, prefix := ":", ""
	for i := range t.NumField() {
		f := t.Field(i)
		if f.Anonymous && f.Type == flagMarker {
			if d, ok := f.Tag.Lookup("delimiter"); ok {
				delimiter = d
			}
			prefix = f.Tag.Get("prefix")
			spec.caseInsensitive = f.Tag.Get("case_insensitive") == "true"
			continue
		}
		if !f.IsExported() {
			continue
		}
		flag, err := newFlag(f)
		if err != nil {
			return nil, err
		}
		if err := spec.add(flag); err != nil {
			return nil, err
		}
	}
	forbidden := delimiter + prefix
	for _, f := range spec.flags {
		for _, name := range append([]string{f.Name}, f.Aliases...) {
			if err := validateFlagName(name, forbidden); err != nil {
				return nil, err
			}
		}
	}
	spec.pattern = compileFlagPattern(spec, prefix, delimiter)
	return spec, nil
}

func newFlag(f reflect.StructField) (*Flag, error) {
	flag := &Flag{
		Name:        f.Tag.Get("flag"),
		Attribute:   f.Name,
		Type:        f.Type,
		Description: f.Tag.Get("description"),
		Positional:  f.Tag.Get("positional") == "true",
		field:       f.Index,
	}
	if flag.Name == "" {
		flag.Name = snake(f.Name)
	}
	if a := f.Tag.Get("aliases"); a != "" {
		flag.Aliases = strings.Split(a, "|")
	}
	flag.Default, flag.HasDefault = f.Tag.Lookup("default")
	flag.MaxArgs, flag.Override = 1, false
	switch f.Type.Kind() {
	case reflect.Slice:
		flag.MaxArgs = -1
	case reflect.Map:
		flag.MaxArgs, flag.Override, flag.CastToDict = -1, true, true
	}
	if v, ok := f.Tag.Lookup("max_args"); ok {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("commands: flag %q has invalid max_args %q", flag.Name, v)
		}
		flag.MaxArgs = n
	}
	if v, ok := f.Tag.Lookup("override"); ok {
		flag.Override = v == "true"
	}
	return flag, nil
}

func (s *flagSpec) add(flag *Flag) error {
	if flag.Positional {
		if s.positional != nil {
			return fmt.Errorf("%q positional flag conflicts with %q flag.", flag.Name, s.positional.Name)
		}
		s.positional = flag
	}
	key := s.key(flag.Name)
	if _, dup := s.byKey[key]; dup || s.aliases[key] != "" {
		return fmt.Errorf("%q flag conflicts with previous flag or alias.", flag.Name)
	}
	s.byKey[key] = flag
	for _, alias := range flag.Aliases {
		a := s.key(alias)
		if _, dup := s.byKey[a]; dup || s.aliases[a] != "" {
			return fmt.Errorf("%q flag alias %q conflicts with previous flag or alias.", flag.Name, alias)
		}
		s.aliases[a] = key
	}
	s.flags = append(s.flags, flag)
	return nil
}

func validateFlagName(name, forbidden string) error {
	if name == "" {
		return fmt.Errorf("flag names should not be empty")
	}
	for _, ch := range name {
		switch {
		case unicode.IsSpace(ch):
			return fmt.Errorf("flag name %q cannot have spaces", name)
		case ch == '\\':
			return fmt.Errorf("flag name %q cannot have backslashes", name)
		case strings.ContainsRune(forbidden, ch):
			return fmt.Errorf("flag name %q cannot have any of %q within them", name, forbidden)
		}
	}
	return nil
}

func compileFlagPattern(s *flagSpec, prefix, delimiter string) *regexp.Regexp {
	keys := make([]string, 0, len(s.byKey)+len(s.aliases))
	for k := range s.byKey {
		keys = append(keys, regexp.QuoteMeta(k))
	}
	for a := range s.aliases {
		keys = append(keys, regexp.QuoteMeta(a))
	}
	slices.SortStableFunc(keys, func(a, b string) int { return len(b) - len(a) })
	mode := ""
	if s.caseInsensitive {
		mode = "(?i)"
	}
	return regexp.MustCompile(mode + "((" + regexp.QuoteMeta(prefix) + ")(?P<flag>" + strings.Join(keys, "|") + ")" + regexp.QuoteMeta(delimiter) + ")")
}
