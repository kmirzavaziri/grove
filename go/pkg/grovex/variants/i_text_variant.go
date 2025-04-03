package variants

import "github.com/kmirzavaziri/grove/go/pkg/grr"

type ITextVariant string

const (
	ITextVariantText     ITextVariant = "text"
	ITextVariantTextarea ITextVariant = "textarea"
	ITextVariantPassword ITextVariant = "password"
	ITextVariantEmail    ITextVariant = "email"
	ITextVariantInt      ITextVariant = "int"
	ITextVariantFloat    ITextVariant = "float"
)

var iTextVariantOptions = map[ITextVariant]struct{}{
	ITextVariantText:     {},
	ITextVariantTextarea: {},
	ITextVariantPassword: {},
	ITextVariantEmail:    {},
	ITextVariantInt:      {},
	ITextVariantFloat:    {},
}

func (e *ITextVariant) Visit() (*ITextVariant, error) {
	if _, ok := iTextVariantOptions[*e]; !ok {
		return nil, grr.Wrap(grr.ErrInvalidEnumValue, "invalid enum %s value %v", "ButtonVariant", *e)
	}

	return e, nil
}
