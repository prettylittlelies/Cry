package commands

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type converterFunc func(ctx *Context, argument string) (reflect.Value, error)

var (
	convertersMu sync.RWMutex
	converters   = map[reflect.Type]converterFunc{}
)

func RegisterConverter[T any](fn func(ctx *Context, argument string) (T, error)) {
	t := reflect.TypeFor[T]()
	convertersMu.Lock()
	converters[t] = func(ctx *Context, argument string) (reflect.Value, error) {
		v, err := fn(ctx, argument)
		return reflect.ValueOf(&v).Elem(), err
	}
	convertersMu.Unlock()
}

func lookupConverter(t reflect.Type) (converterFunc, bool) {
	convertersMu.RLock()
	defer convertersMu.RUnlock()
	fn, ok := converters[t]
	return fn, ok
}

func ConvertBool(argument string) (bool, error) {
	switch strings.ToLower(argument) {
	case "yes", "y", "true", "t", "1", "enable", "on":
		return true, nil
	case "no", "n", "false", "f", "0", "disable", "off":
		return false, nil
	}
	return false, errBadBool(strings.ToLower(argument))
}

func builtin(t reflect.Type, argument string) (reflect.Value, error) {
	v := reflect.New(t).Elem()
	switch t.Kind() {
	case reflect.String:
		v.SetString(argument)
	case reflect.Bool:
		b, err := ConvertBool(argument)
		if err != nil {
			return v, err
		}
		v.SetBool(b)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(argument, 10, t.Bits())
		if err != nil {
			return v, err
		}
		v.SetInt(n)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n, err := strconv.ParseUint(argument, 10, t.Bits())
		if err != nil {
			return v, err
		}
		v.SetUint(n)
	case reflect.Float32, reflect.Float64:
		n, err := strconv.ParseFloat(argument, t.Bits())
		if err != nil {
			return v, err
		}
		v.SetFloat(n)
	default:
		return v, fmt.Errorf("commands: no converter registered for %s", t)
	}
	return v, nil
}

func actualConversion(ctx *Context, t reflect.Type, argument string, p *Parameter) (reflect.Value, error) {
	if spec, err := flagSpecFor(t); err != nil || spec != nil {
		if err != nil {
			return reflect.Value{}, err
		}
		return spec.convert(ctx, argument)
	}
	if fn, ok := lookupConverter(t); ok {
		v, err := fn(ctx, argument)
		if err != nil && !errors.Is(err, ErrCommand) {
			return v, &ConversionError{kinded: newKind(commandErr, "%v", err), Converter: t.String(), Original: err}
		}
		return v, err
	}
	v, err := builtin(t, argument)
	if err != nil && !errors.Is(err, ErrCommand) {
		return v, &BadArgument{newKind(badArg, "Converting to %q failed for parameter %q.", t.String(), p.Name)}
	}
	return v, err
}

func runConverters(ctx *Context, p *Parameter, argument string) (reflect.Value, error) {
	t := p.converter()
	v, err := actualConversion(ctx, t, argument, p)
	if err != nil {
		if p.Optional() && p.Kind != VarPositional && errors.Is(err, ErrCommand) {
			ctx.View.Undo()
			return reflect.Zero(reflect.PointerTo(t)), nil
		}
		return v, err
	}
	if len(p.Literals) > 0 && !slices.Contains(p.Literals, fmt.Sprint(v.Interface())) {
		quoted := make([]string, len(p.Literals))
		for i, l := range p.Literals {
			quoted[i] = "'" + l + "'"
		}
		return v, &BadLiteralArgument{
			kinded:   newKind(userInput, "Could not convert %q into the literal %s.", displayed(p), humanJoin(quoted, "or")),
			Param:    p,
			Literals: p.Literals,
			Argument: argument,
		}
	}
	if err := checkRange(p, v); err != nil {
		return v, err
	}
	if p.Optional() {
		ptr := reflect.New(t)
		ptr.Elem().Set(v)
		return ptr, nil
	}
	return v, nil
}
