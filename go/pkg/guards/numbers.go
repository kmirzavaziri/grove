package guards

import (
	"github.com/kmirzavaziri/grove/go/internal/pkg/lua"
)

type Range struct {
	Min     *float64
	Max     *float64
	ErrType string
	ErrMin  string
	ErrMax  string
}

func (v Range) Script() string {
	return lua.GuardsRange
}

type Integer struct {
	ErrType string
	Err     string
}

func (v Integer) Script() string {
	return lua.GuardsInteger
}
