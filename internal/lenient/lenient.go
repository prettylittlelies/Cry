package lenient

import (
	"bytes"
	"encoding"
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"strings"
)

var (
	jsonUnmarshaler = reflect.TypeFor[json.Unmarshaler]()
	textUnmarshaler = reflect.TypeFor[encoding.TextUnmarshaler]()
)

func Unmarshal(data []byte, v any) error {
	err := json.Unmarshal(data, v)
	var typeErr *json.UnmarshalTypeError
	if !errors.As(err, &typeErr) {
		return err
	}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	var tree any
	if dec.Decode(&tree) != nil {
		return err
	}
	fixed, merr := json.Marshal(coerce(tree, reflect.TypeOf(v)))
	if merr != nil {
		return err
	}
	return json.Unmarshal(fixed, v)
}

func custom(t reflect.Type) bool {
	return t.Implements(jsonUnmarshaler) || reflect.PointerTo(t).Implements(jsonUnmarshaler) ||
		t.Implements(textUnmarshaler) || reflect.PointerTo(t).Implements(textUnmarshaler)
}

func coerce(v any, t reflect.Type) any {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if custom(t) || t.Kind() == reflect.Interface || v == nil {
		return v
	}
	numeric := func(k reflect.Kind) bool {
		return k >= reflect.Int && k <= reflect.Float64
	}
	switch val := v.(type) {
	case json.Number:
		switch {
		case t.Kind() == reflect.String:
			return val.String()
		case numeric(t.Kind()):
			return val
		}
	case string:
		switch {
		case t.Kind() == reflect.String:
			return val
		case numeric(t.Kind()):
			if _, err := strconv.ParseFloat(val, 64); err == nil {
				return json.Number(val)
			}
		}
	case bool:
		if t.Kind() == reflect.Bool {
			return val
		}
	case []any:
		if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
			for i := range val {
				val[i] = coerce(val[i], t.Elem())
			}
			return val
		}
	case map[string]any:
		switch t.Kind() {
		case reflect.Map:
			for k := range val {
				val[k] = coerce(val[k], t.Elem())
			}
			return val
		case reflect.Struct:
			fields := fieldTypes(t)
			for k := range val {
				if ft, ok := fields[strings.ToLower(k)]; ok {
					val[k] = coerce(val[k], ft)
				}
			}
			return val
		}
	}
	return nil
}

func fieldTypes(t reflect.Type) map[string]reflect.Type {
	out := map[string]reflect.Type{}
	for _, f := range reflect.VisibleFields(t) {
		if !f.IsExported() || f.Anonymous && f.Type.Kind() == reflect.Struct {
			continue
		}
		name, _, _ := strings.Cut(f.Tag.Get("json"), ",")
		if name == "-" {
			continue
		}
		if name == "" {
			name = f.Name
		}
		out[strings.ToLower(name)] = f.Type
	}
	return out
}
