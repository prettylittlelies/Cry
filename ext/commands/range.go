package commands

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

func parseBounds(p *Parameter, tag reflect.StructTag) error {
	t := p.converter()
	for _, spec := range []struct {
		key string
		dst *any
	}{{"min", &p.Min}, {"max", &p.Max}} {
		raw, ok := tag.Lookup(spec.key)
		if !ok {
			continue
		}
		bound, err := parseBound(t, raw)
		if err != nil {
			return fmt.Errorf("commands: parameter %s: %w", p.Name, err)
		}
		*spec.dst = bound
	}
	return nil
}

func parseBound(t reflect.Type, raw string) (any, error) {
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		return strconv.ParseFloat(raw, 64)
	case reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.ParseInt(raw, 10, 64)
	}
	return nil, fmt.Errorf("expected int, float, or str as range type, received %s instead", t)
}

func toFloat(bound any) float64 {
	switch b := bound.(type) {
	case int64:
		return float64(b)
	case float64:
		return b
	}
	return 0
}

func formatBound(v any) string {
	switch b := v.(type) {
	case float64:
		s := strconv.FormatFloat(b, 'f', -1, 64)
		if !strings.ContainsAny(s, ".eEn") {
			s += ".0"
		}
		return s
	case float32:
		return formatBound(float64(b))
	}
	return fmt.Sprint(v)
}

func checkRange(p *Parameter, v reflect.Value) error {
	if p.Min == nil && p.Max == nil {
		return nil
	}
	var count float64
	switch v.Kind() {
	case reflect.String:
		count = float64(utf8.RuneCountInString(v.String()))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		count = float64(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		count = float64(v.Uint())
	case reflect.Float32, reflect.Float64:
		count = v.Float()
	default:
		return nil
	}
	if p.Min != nil && count < toFloat(p.Min) || p.Max != nil && count > toFloat(p.Max) {
		return errRange(v.Interface(), p.Min, p.Max)
	}
	return nil
}
