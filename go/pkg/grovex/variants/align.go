package variants

import (
	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/grr"
)

type Align string

const (
	AlignInherit Align = "inherit"
	AlignCenter  Align = "center"
	AlignJustify Align = "justify"
	AlignStart   Align = "start"
	AlignEnd     Align = "end"
)

var alignOptions = map[Align]struct{}{
	AlignCenter:  {},
	AlignInherit: {},
	AlignJustify: {},
	AlignStart:   {},
	AlignEnd:     {},
}

func (e *Align) Visit() (*Align, error) {
	if e == nil {
		return gex.P(AlignInherit), nil
	}

	if _, ok := alignOptions[*e]; !ok {
		return nil, grr.Wrap(grr.ErrInvalidEnumValue, "invalid enum %s value %v", "Align", *e)
	}

	return e, nil
}
