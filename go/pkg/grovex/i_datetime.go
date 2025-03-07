package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type IDatetimeArgs struct {
	Key   string
	Props flux.Read[IDatetimeProps]
	Input *grove.Input
}

type IDatetimeProps struct {
}

func IDatetime(args IDatetimeArgs) *grove.Node {
	return &grove.Node{
		Type:  "grovex.IDatetime",
		Key:   args.Key,
		Props: flux.ReadValue(args.Props),
		Input: args.Input,
	}
}
