package guards

import (
	"github.com/kmirzavaziri/grove/go/internal/pkg/lua"
)

type Equal struct {
	Value   bool
	ErrType string
	Err     string
}

func (v Equal) Script() string {
	return lua.GuardsEqual
}
