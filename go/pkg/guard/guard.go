package guard

import (
	"fmt"

	luaPkg "github.com/kmirzavaziri/grove/go/internal/pkg/lua"
	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/grr"
	"github.com/yuin/gopher-lua"
)

type Guard string

type Config interface {
	Script() string
}

func New(c Config) Guard {
	return Guard(fmt.Sprintf("%s\n\nlocal config = %s\n\n%s",
		luaPkg.ErrorRender,
		luaPkg.Render(gex.Marshal(c).Lua()),
		c.Script(),
	))
}

func (g Guard) Guard(value *gex.Value) (string, error) {
	state := lua.NewState()
	defer state.Close()

	err := state.DoString(string(g))
	if err != nil {
		return "", grr.Wrap(grr.ErrInvalidLuaScript, "invalid lua script %q: %w", g, err)
	}

	err = state.CallByParam(lua.P{
		Fn:      state.GetGlobal("guard"),
		NRet:    1,
		Protect: true,
	}, value.Lua())
	if err != nil {
		return "", grr.Wrap(grr.ErrInvalidLuaScript, "invalid lua script %q: %w", g, err)
	}

	return state.Get(-1).String(), nil
}

func (g Guard) Lua() string {
	return string(g)
}
