package guards

import (
	"github.com/kmirzavaziri/grove/go/internal/pkg/lua"
	"github.com/kmirzavaziri/grove/go/pkg/gex"
)

var luaType = map[gex.Type]string{
	gex.TypeNil:    "nil",
	gex.TypeMap:    "map",
	gex.TypeList:   "list",
	gex.TypeNumber: "number",
	gex.TypeString: "string",
	gex.TypeBool:   "boolean",
}

type Type struct {
	TypeName gex.Type
	Err      string
}

func (v Type) Marshal() *gex.Value {
	return gex.Marshal(struct {
		TypeName string
		Err      string
	}{
		TypeName: luaType[v.TypeName],
		Err:      v.Err,
	})
}

func (v Type) Script() string {
	return lua.GuardsType
}

type Empty struct {
	Err string
}

func (v Empty) Script() string {
	return lua.GuardsEmpty
}

type Not struct {
	Guard Guard
	Err   string
}

func (v Not) Script() string {
	return lua.GuardsNot
}

type And struct {
	Guards []Guard
	Err    string
}

func (v And) Script() string {
	return lua.GuardsAnd
}

type Or struct {
	Guards []Guard
	Err    string
}

func (v Or) Script() string {
	return lua.GuardsOr
}
