package rest

import (
	"encoding"
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func Values(v any) url.Values {
	out := url.Values{}
	rv := reflect.Indirect(reflect.ValueOf(v))
	if !rv.IsValid() || rv.Kind() != reflect.Struct {
		return out
	}
	rt := rv.Type()
	for i := range rt.NumField() {
		f := rt.Field(i)
		name, opts, _ := strings.Cut(f.Tag.Get("json"), ",")
		if name == "" || name == "-" {
			continue
		}
		fv := rv.Field(i)
		if fv.IsZero() && (strings.Contains(opts, "omitempty") || strings.Contains(opts, "omitzero") || fv.Kind() == reflect.Pointer) {
			continue
		}
		for fv.Kind() == reflect.Pointer {
			fv = fv.Elem()
		}
		if fv.Kind() == reflect.Slice && fv.Type().Elem().Kind() != reflect.Uint8 {
			for j := range fv.Len() {
				out.Add(name, scalar(fv.Index(j)))
			}
			continue
		}
		out.Set(name, scalar(fv))
	}
	return out
}

func scalar(v reflect.Value) string {
	if m, ok := v.Interface().(encoding.TextMarshaler); ok {
		if b, err := m.MarshalText(); err == nil {
			return string(b)
		}
	}
	return fmt.Sprint(v.Interface())
}
