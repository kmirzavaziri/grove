package gex

type Type uint

const (
	TypeNil Type = iota
	TypeMap
	TypeList
	TypeNumber
	TypeString
	TypeBool
)

var (
	Types = []Type{
		TypeNil,
		TypeMap,
		TypeList,
		TypeNumber,
		TypeString,
		TypeBool,
	}

	typeNames = map[Type]string{
		TypeNil:    "nil",
		TypeMap:    "map",
		TypeList:   "list",
		TypeNumber: "number",
		TypeString: "string",
		TypeBool:   "bool",
	}
)

func (t Type) String() string {
	return typeNames[t]
}
