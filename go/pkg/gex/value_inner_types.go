package gex

type m map[string]*Value

func (v m) Marshal() *Value {
	return Map(v)
}

type l []*Value

func (v l) Marshal() *Value {
	return List(v...)
}

type n float64

func (v n) Marshal() *Value {
	return Num(float64(v))
}

type s string

func (v s) Marshal() *Value {
	return Str(string(v))
}

type b bool

func (v b) Marshal() *Value {
	return Bool(bool(v))
}
