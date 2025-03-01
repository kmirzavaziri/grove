package lua

import (
	"fmt"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

func Render(value lua.LValue) string {
	switch value.Type() {
	case lua.LTNil:
		return "nil"
	case lua.LTBool:
		if lua.LVAsBool(value) {
			return "true"
		}

		return "false"
	case lua.LTNumber:
		return fmt.Sprintf("%v", value)
	case lua.LTString:
		return fmt.Sprintf("%q", value)
	case lua.LTTable:
		var builder strings.Builder

		builder.WriteString("{")

		table := value.(*lua.LTable)

		table.ForEach(func(key, value lua.LValue) {
			builder.WriteString("[")
			builder.WriteString(Render(key))
			builder.WriteString("]")
			builder.WriteString(" = ")
			builder.WriteString(Render(value))
			builder.WriteString(", ")
		})

		builder.WriteString("}")

		return builder.String()
	default:
		return "<unsupported type>"
	}
}
