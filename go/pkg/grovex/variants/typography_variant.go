package variants

import (
	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/grr"
)

type TypographyVariant string

const (
	TypographyVariantInherit   TypographyVariant = "inherit"
	TypographyVariantBody1     TypographyVariant = "body1"
	TypographyVariantBody2     TypographyVariant = "body2"
	TypographyVariantButton    TypographyVariant = "button"
	TypographyVariantCaption   TypographyVariant = "caption"
	TypographyVariantH1        TypographyVariant = "h1"
	TypographyVariantH2        TypographyVariant = "h2"
	TypographyVariantH3        TypographyVariant = "h3"
	TypographyVariantH4        TypographyVariant = "h4"
	TypographyVariantH5        TypographyVariant = "h5"
	TypographyVariantH6        TypographyVariant = "h6"
	TypographyVariantOverline  TypographyVariant = "overline"
	TypographyVariantSubtitle1 TypographyVariant = "subtitle1"
	TypographyVariantSubtitle2 TypographyVariant = "subtitle2"
)

var typographyVariantOptions = map[TypographyVariant]struct{}{
	TypographyVariantInherit:   {},
	TypographyVariantBody1:     {},
	TypographyVariantBody2:     {},
	TypographyVariantButton:    {},
	TypographyVariantCaption:   {},
	TypographyVariantH1:        {},
	TypographyVariantH2:        {},
	TypographyVariantH3:        {},
	TypographyVariantH4:        {},
	TypographyVariantH5:        {},
	TypographyVariantH6:        {},
	TypographyVariantOverline:  {},
	TypographyVariantSubtitle1: {},
	TypographyVariantSubtitle2: {},
}

func (e *TypographyVariant) Visit() (*TypographyVariant, error) {
	if e == nil {
		return gex.P(TypographyVariantInherit), nil
	}

	if _, ok := typographyVariantOptions[*e]; !ok {
		return nil, grr.Wrap(grr.ErrInvalidEnumValue, "invalid enum %s value %v", "TypographyVariant", *e)
	}

	return e, nil
}
