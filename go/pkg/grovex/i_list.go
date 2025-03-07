package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type IListArgs struct {
	Key   string
	Props flux.Read[IListProps]
	Input *grove.Input
}

type IListProps struct {
}

func IList(args IListArgs) *grove.Node {
	return &grove.Node{
		Type:  "grovex.IList",
		Key:   args.Key,
		Props: flux.ReadValue(args.Props),
		Input: args.Input,
	}
}
