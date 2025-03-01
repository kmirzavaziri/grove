package gex

import (
	"encoding/json"

	lua "github.com/yuin/gopher-lua"
)

type Value struct {
	t Type

	m m
	l l
	n n
	s s
	b b
}

func (v *Value) Type() Type {
	if v == nil {
		return Nil().t
	}

	return v.t
}

func (v *Value) Map() (map[string]*Value, bool) {
	if v == nil {
		return m{}, false
	}

	if v.t != TypeMap {
		return m{}, false
	}

	return v.m, true
}

func (v *Value) List() ([]*Value, bool) {
	if v == nil {
		return l{}, false
	}

	if v.t != TypeList {
		return l{}, false
	}

	return v.l, true
}

func (v *Value) Number() (float64, bool) {
	if v == nil {
		return 0, false
	}

	if v.t != TypeNumber {
		return 0, false
	}

	return float64(v.n), true
}

func (v *Value) String() (string, bool) {
	if v == nil {
		return "", false
	}

	if v.t != TypeString {
		return "", false
	}

	return string(v.s), true
}

func (v *Value) Bool() (bool, bool) {
	if v == nil {
		return false, false
	}

	if v.t != TypeBool {
		return false, false
	}

	return bool(v.b), true
}

func (v *Value) IsNil() bool {
	if v == nil {
		return true
	}

	return v.t == TypeNil
}

func (v *Value) Chain() Chain {
	if v == nil {
		return Chain{v: Nil()}
	}

	return Chain{v: v}
}

// Marshalling

func (v *Value) Marshal() *Value {
	return v
}

func (v *Value) Unmarshal() any {
	if v == nil {
		return nil
	}

	switch v.t {
	case TypeNil:
		return nil
	case TypeMap:
		m := make(map[string]any, len(v.m))
		for k, val := range v.m {
			m[k] = val.Unmarshal()
		}

		return m
	case TypeList:
		l := make([]any, len(v.l))
		for i, val := range v.l {
			l[i] = val.Unmarshal()
		}

		return l
	case TypeNumber:
		return float64(v.n)
	case TypeString:
		return string(v.s)
	case TypeBool:
		return bool(v.b)
	default:
		return nil
	}
}

// Lua

func (v *Value) Lua() lua.LValue {
	if v == nil {
		return lua.LNil
	}

	switch v.t {
	case TypeMap:
		lTable := &lua.LTable{}

		for k, val := range v.m {
			lTable.RawSetString(k, val.Lua())
		}

		return lTable
	case TypeList:
		lTable := &lua.LTable{}

		for i, val := range v.l {
			lTable.RawSetInt(i+1, val.Lua())
		}

		return lTable
	case TypeNumber:
		return lua.LNumber(v.n)
	case TypeString:
		return lua.LString(v.s)
	case TypeBool:
		return lua.LBool(v.b)
	default:
		return lua.LNil
	}
}

// JSON

func (v *Value) UnmarshalJSON(b []byte) error {
	var raw any

	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	*v = *Marshal(raw)

	return nil
}

func (v *Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Unmarshal())
}
