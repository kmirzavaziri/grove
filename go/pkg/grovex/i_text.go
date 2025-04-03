package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
	"github.com/kmirzavaziri/grove/go/pkg/grovex/variants"
)

type ITextArgs struct {
	Key   string
	Props flux.Read[ITextProps]
	Input *grove.Input
}

type ITextProps struct {
	Label   string
	Variant *variants.ITextVariant

	Placeholder  string
	AutoComplete string
}

func IText(args ITextArgs) *grove.Node {
	if args.Input == nil {
		return grove.Error(args.Key, "input cannot be nil")
	}

	return &grove.Node{
		Type: "grovex.IText",
		Key:  args.Key,
		// TODO add per request validation by IText by adding middleware to Props so that we can verify user defined flux
		//  results, and raise error if the value is not in this list.
		Props: flux.ReadValue(args.Props),
		Input: args.Input,
	}
}
