package variants

import (
	"github.com/kmirzavaziri/grove/go/pkg/gex"
)

type Size string

const (
	SizeDefault Size = ""
	SizeSmall   Size = "small"
	SizeMedium  Size = "medium"
	SizeLarge   Size = "large"
)

func (e *Size) Visit() (*Size, error) {
	if e == nil {
		return gex.P(SizeDefault), nil
	}

	return e, nil
}
