package guardx

import (
	"github.com/kmirzavaziri/grove/go/internal/pkg/lua"
)

type Match struct {
	Pattern string
	ErrType string
	Err     string
}

func (v Match) Script() string {
	return lua.GuardsMatch
}

type Contains struct {
	Substring string
	ErrType   string
	Err       string
}

func (v Contains) Script() string {
	return lua.GuardsContains
}

type HasPrefix struct {
	Prefix  string
	ErrType string
	Err     string
}

func (v HasPrefix) Script() string {
	return lua.GuardsHasPrefix
}

type HasSuffix struct {
	Suffix  string
	ErrType string
	Err     string
}

func (v HasSuffix) Script() string {
	return lua.GuardsHasSuffix
}

type Enum struct {
	Values  []string
	ErrType string
	Err     string
}

func (v Enum) Script() string {
	return lua.GuardsEnum
}

type Email struct {
	ErrType string
	Err     string
}

func (v Email) Script() string {
	return lua.GuardsEmail
}

type Phone struct {
	ErrType string
	Err     string
}

func (v Phone) Script() string {
	return lua.GuardsPhone
}
