package variants

import (
	"github.com/kmirzavaziri/grove/go/pkg/gex"
)

type Color string

const (
	ColorDefault   Color = ""
	ColorPrimary   Color = "primary"
	ColorSecondary Color = "secondary"
	ColorSuccess   Color = "success"
	ColorError     Color = "error"
	ColorInfo      Color = "info"
	ColorWarning   Color = "warning"
)

func (e *Color) Visit() (*Color, error) {
	if e == nil {
		return gex.P(ColorDefault), nil
	}

	return e, nil
}
