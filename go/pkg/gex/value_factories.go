package gex

import (
	"reflect"
	"strings"

	"github.com/kmirzavaziri/grove/go/internal/pkg/cases"
)

func Marshal(value any) *Value {
	if m, ok := value.(Marshaller); ok {
		return m.Marshal()
	}

	val := reflect.ValueOf(value)

	switch val.Kind() {
	case reflect.Ptr:
		if val.IsNil() {
			return Nil()
		}

		return Marshal(val.Elem().Interface())
	case reflect.String:
		return Str(val.String())
	case reflect.Float64, reflect.Float32:
		return Num(val.Float())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return Num(float64(val.Int()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return Num(float64(val.Uint()))
	case reflect.Bool:
		return Bool(val.Bool())
	case reflect.Slice:
		var list l

		for i := 0; i < val.Len(); i++ {
			list = append(list, Marshal(val.Index(i).Interface()))
		}

		return List(list...)
	case reflect.Map:
		result := make(m)

		for _, key := range val.MapKeys() {
			strKey := key.String()
			result[strKey] = Marshal(val.MapIndex(key).Interface())
		}

		return Map(result)
	case reflect.Struct:
		result := make(m)

		t := reflect.TypeOf(value)

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldValue := val.Field(i)

			if !field.IsExported() {
				continue
			}

			key := strings.TrimSpace(field.Tag.Get("rip"))
			if key == "" {
				key = cases.ToSnakeCase(field.Name)
			}

			result[key] = Marshal(fieldValue.Interface())
		}

		return Map(result)
	default:
		return Nil()
	}
}

func Map(kv map[string]*Value) *Value {
	newKV := map[string]*Value{}

	for k, v := range kv {
		if !v.IsNil() {
			newKV[k] = v
		}
	}

	if len(newKV) == 0 {
		return Nil()
	}

	return &Value{
		t: TypeMap,
		m: newKV,
	}
}

func List(vv ...*Value) *Value {
	newVV := make([]*Value, 0, len(vv))

	for _, v := range vv {
		if !v.IsNil() {
			newVV = append(newVV, v)
		}
	}

	if len(newVV) == 0 {
		return Nil()
	}

	return &Value{
		t: TypeList,
		l: newVV,
	}
}

func Num(v float64) *Value {
	return &Value{
		t: TypeNumber,
		n: n(v),
	}
}

func Str(v string) *Value {
	return &Value{
		t: TypeString,
		s: s(v),
	}
}

func Bool(v bool) *Value {
	return &Value{
		t: TypeBool,
		b: b(v),
	}
}

func Nil() *Value {
	return &Value{
		t: TypeNil,
	}
}
