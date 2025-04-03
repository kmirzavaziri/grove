package variants

import (
	"github.com/kmirzavaziri/grove/go/pkg/grr"
)

type DButtonVariant string

const (
	DButtonVariantItem      DButtonVariant = "item"
	DButtonVariantContained DButtonVariant = "contained"
	DButtonVariantOutlined  DButtonVariant = "outlined"
	DButtonVariantText      DButtonVariant = "text"
)

var dButtonVariantOptions = map[DButtonVariant]struct{}{
	DButtonVariantItem:      {},
	DButtonVariantContained: {},
	DButtonVariantOutlined:  {},
	DButtonVariantText:      {},
}

func (e *DButtonVariant) Visit() (*DButtonVariant, error) {
	if _, ok := dButtonVariantOptions[*e]; !ok {
		return nil, grr.Wrap(grr.ErrInvalidEnumValue, "invalid enum %s value %v", "DButtonVariant", *e)
	}

	return e, nil
}
