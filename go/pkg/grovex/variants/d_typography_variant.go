package variants

import (
	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/grr"
)

type DTypographyVariant string

const (
	DTypographyVariantInherit   DTypographyVariant = "inherit"
	DTypographyVariantBody1     DTypographyVariant = "body1"
	DTypographyVariantBody2     DTypographyVariant = "body2"
	DTypographyVariantButton    DTypographyVariant = "button"
	DTypographyVariantCaption   DTypographyVariant = "caption"
	DTypographyVariantH1        DTypographyVariant = "h1"
	DTypographyVariantH2        DTypographyVariant = "h2"
	DTypographyVariantH3        DTypographyVariant = "h3"
	DTypographyVariantH4        DTypographyVariant = "h4"
	DTypographyVariantH5        DTypographyVariant = "h5"
	DTypographyVariantH6        DTypographyVariant = "h6"
	DTypographyVariantOverline  DTypographyVariant = "overline"
	DTypographyVariantSubtitle1 DTypographyVariant = "subtitle1"
	DTypographyVariantSubtitle2 DTypographyVariant = "subtitle2"
)

var dTypographyVariantOptions = map[DTypographyVariant]struct{}{
	DTypographyVariantInherit:   {},
	DTypographyVariantBody1:     {},
	DTypographyVariantBody2:     {},
	DTypographyVariantButton:    {},
	DTypographyVariantCaption:   {},
	DTypographyVariantH1:        {},
	DTypographyVariantH2:        {},
	DTypographyVariantH3:        {},
	DTypographyVariantH4:        {},
	DTypographyVariantH5:        {},
	DTypographyVariantH6:        {},
	DTypographyVariantOverline:  {},
	DTypographyVariantSubtitle1: {},
	DTypographyVariantSubtitle2: {},
}

func (e *DTypographyVariant) Visit() (*DTypographyVariant, error) {
	if e == nil {
		return gex.P(DTypographyVariantInherit), nil
	}

	if _, ok := dTypographyVariantOptions[*e]; !ok {
		return nil, grr.Wrap(grr.ErrInvalidEnumValue, "invalid enum %s value %v", "DTypographyVariant", *e)
	}

	return e, nil
}
