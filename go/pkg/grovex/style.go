package grovex

import "github.com/kmirzavaziri/grove/go/pkg/gex"

type TextSize string

func (s TextSize) Marshal() *gex.Value {
	return gex.Str(string(s))
}

const (
	TextSizePG TextSize = "PG"
	TextSizeH1 TextSize = "H1"
	TextSizeH2 TextSize = "H2"
	TextSizeH3 TextSize = "H3"
	TextSizeH4 TextSize = "H4"
	TextSizeH5 TextSize = "H5"
	TextSizeH6 TextSize = "H6"
)
