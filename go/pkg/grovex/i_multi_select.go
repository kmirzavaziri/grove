package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type IMultiSelectArgs struct {
	Key   string
	Props flux.Read[IMultiSelectProps]
	Input *grove.Input
}

type IMultiSelectProps struct {
}

func IMultiSelect(args IMultiSelectArgs) *grove.Node {
	return &grove.Node{
		Type:  "grovex.IMultiSelect",
		Key:   args.Key,
		Props: flux.ReadValue(args.Props),
		Input: args.Input,
	}
}
