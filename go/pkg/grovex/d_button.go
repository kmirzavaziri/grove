package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
	"github.com/kmirzavaziri/grove/go/pkg/grovex/variants"
)

type DButtonArgs struct {
	Key   string
	Props flux.Read[DButtonProps]
}

type DButtonProps struct {
	Text string

	Color     *variants.Color
	Variant   *variants.ButtonVariant
	FullWidth bool
	Disabled  bool
	Selected  bool

	StartIcon *variants.Icon
	EndIcon   *variants.Icon
	Size      *variants.Size
}

func DButton(args DButtonArgs) *grove.Node {
	return &grove.Node{
		Type: "grovex.DButton",
		Key:  args.Key,
		// TODO checks and defaults, selected is only compatible with item, full width is not compatible with item
		Props: flux.ReadValue(args.Props),
	}
}
