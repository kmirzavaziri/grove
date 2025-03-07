package variants

import (
	"github.com/kmirzavaziri/grove/go/pkg/grr"
)

type ButtonVariant string

const (
	ButtonVariantItem      ButtonVariant = "item"
	ButtonVariantContained ButtonVariant = "contained"
	ButtonVariantOutlined  ButtonVariant = "outlined"
	ButtonVariantText      ButtonVariant = "text"
)

var buttonVariantOptions = map[ButtonVariant]struct{}{
	ButtonVariantItem:      {},
	ButtonVariantContained: {},
	ButtonVariantOutlined:  {},
	ButtonVariantText:      {},
}

func (e *ButtonVariant) Visit() (*ButtonVariant, error) {
	if _, ok := buttonVariantOptions[*e]; !ok {
		return nil, grr.Wrap(grr.ErrInvalidEnumValue, "invalid enum %s value %v", "ButtonVariant", *e)
	}

	return e, nil
}
