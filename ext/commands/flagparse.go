package commands

import (
	"reflect"
	"strings"
)

func (s *flagSpec) parse(argument string, ignoreExtra bool) (map[string][]string, error) {
	result := map[string][]string{}
	lastPosition := 0
	var lastFlag *Flag
	if s.positional != nil {
		var value string
		if loc := s.pattern.FindStringIndex(argument); loc != nil {
			value = strings.TrimSpace(argument[:loc[0]])
		} else {
			value = strings.TrimSpace(argument)
			lastPosition = len(argument)
		}
		if value != "" {
			result[s.key(s.positional.Name)] = []string{value}
		}
	}
	group := s.pattern.SubexpIndex("flag")
	for _, m := range s.pattern.FindAllStringSubmatchIndex(argument, -1) {
		begin, end := m[0], m[1]
		key := s.key(argument[m[2*group]:m[2*group+1]])
		if target, ok := s.aliases[key]; ok {
			key = target
		}
		if lastPosition > 0 && lastFlag != nil {
			var value string
			if begin-1 > lastPosition {
				value = strings.TrimLeft(argument[lastPosition:begin-1], " \t\n\r\v\f")
			}
			if value == "" {
				return nil, errMissingFlagArgument(lastFlag)
			}
			name := s.key(lastFlag.Name)
			result[name] = append(result[name], value)
		}
		lastPosition, lastFlag = end, s.byKey[key]
	}
	value := strings.TrimSpace(argument[lastPosition:])
	switch {
	case lastFlag != nil:
		if value == "" {
			return nil, errMissingFlagArgument(lastFlag)
		}
		name := s.key(lastFlag.Name)
		result[name] = append(result[name], value)
	case value != "" && !ignoreExtra:
		return nil, errTooManyArguments(s.typ.Name())
	}
	return result, nil
}

func (s *flagSpec) convert(ctx *Context, argument string) (reflect.Value, error) {
	ignoreExtra := true
	if ctx.Command != nil && ctx.CurrentParameter != nil && ctx.CurrentParameter.Kind == KeywordOnly {
		ignoreExtra = ctx.Command.IgnoreExtra
	}
	arguments, err := s.parse(argument, ignoreExtra)
	if err != nil {
		return reflect.Value{}, err
	}
	return s.build(ctx, arguments)
}

func (s *flagSpec) build(ctx *Context, arguments map[string][]string) (reflect.Value, error) {
	out := reflect.New(s.typ).Elem()
	for _, flag := range s.flags {
		values, ok := arguments[s.key(flag.Name)]
		if !ok {
			if flag.Required() {
				return out, errMissingRequiredFlag(flag)
			}
			if !flag.HasDefault {
				if nested, _ := flagSpecFor(flag.Type); nested != nil {
					v, err := nested.build(ctx, nil)
					if err != nil {
						return out, err
					}
					out.FieldByIndex(flag.field).Set(v)
				}
				continue
			}
			if flag.Default == "" && isCollection(flag.Type) {
				continue
			}
			values = []string{flag.Default}
		}
		if flag.MaxArgs > 0 && len(values) > flag.MaxArgs {
			if !flag.Override {
				return out, errTooManyFlags(flag, values)
			}
			values = values[len(values)-flag.MaxArgs:]
		}
		v, err := assignFlag(ctx, flag, values)
		if err != nil {
			return out, err
		}
		out.FieldByIndex(flag.field).Set(v)
	}
	return out, nil
}

func assignFlag(ctx *Context, flag *Flag, values []string) (reflect.Value, error) {
	t := flag.Type
	switch t.Kind() {
	case reflect.Slice:
		out := reflect.MakeSlice(t, 0, len(values))
		for _, value := range values {
			v, err := convertFlag(ctx, value, flag, t.Elem())
			if err != nil {
				return out, err
			}
			out = reflect.Append(out, v)
		}
		return out, nil
	case reflect.Map:
		out := reflect.MakeMapWithSize(t, len(values))
		pair := []reflect.Type{t.Key(), t.Elem()}
		for _, value := range values {
			kv, err := convertTuple(ctx, value, flag, pair)
			if err != nil {
				return out, err
			}
			out.SetMapIndex(kv[0], kv[1])
		}
		return out, nil
	}
	return convertFlag(ctx, values[0], flag, t)
}

func convertFlag(ctx *Context, argument string, flag *Flag, t reflect.Type) (reflect.Value, error) {
	switch t.Kind() {
	case reflect.Array:
		items, err := convertTuple(ctx, argument, flag, repeatType(t.Elem(), t.Len()))
		if err != nil {
			return reflect.Value{}, err
		}
		out := reflect.New(t).Elem()
		for i, item := range items {
			out.Index(i).Set(item)
		}
		return out, nil
	case reflect.Pointer:
		v, err := actualConversion(ctx, t.Elem(), argument, ctx.CurrentParameter)
		if err != nil {
			return v, err
		}
		ptr := reflect.New(t.Elem())
		ptr.Elem().Set(v)
		return ptr, nil
	}
	v, err := actualConversion(ctx, t, argument, ctx.CurrentParameter)
	if err != nil {
		return v, errBadFlagArgument(flag, argument, t, err)
	}
	return v, nil
}

func repeatType(t reflect.Type, n int) []reflect.Type {
	out := make([]reflect.Type, n)
	for i := range out {
		out[i] = t
	}
	return out
}

func convertTuple(ctx *Context, argument string, flag *Flag, types []reflect.Type) ([]reflect.Value, error) {
	view := NewStringView(argument)
	results := make([]reflect.Value, 0, len(types))
	for _, t := range types {
		view.SkipWS()
		if view.EOF() {
			break
		}
		word, _, err := view.GetQuotedWord()
		if err != nil {
			return nil, err
		}
		v, err := actualConversion(ctx, t, word, ctx.CurrentParameter)
		if err != nil {
			return nil, errBadFlagArgument(flag, word, t, err)
		}
		results = append(results, v)
	}
	if len(results) != len(types) {
		return nil, errMissingFlagArgument(flag)
	}
	return results, nil
}

func flagDefault(ctx *Context, t reflect.Type) (reflect.Value, bool, error) {
	spec, err := flagSpecFor(t)
	if err != nil || spec == nil || !spec.constructible() {
		return reflect.Value{}, false, err
	}
	v, err := spec.build(ctx, nil)
	return v, true, err
}

func isCollection(t reflect.Type) bool {
	return t.Kind() == reflect.Slice || t.Kind() == reflect.Map
}
