package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type INumberArgs struct {
	Key   string
	Props flux.Read[INumberProps]
	Input *grove.Input
}

type INumberProps struct {
}

func INumber(args INumberArgs) *grove.Node {
	return &grove.Node{
		Type:  "grovex.INumber",
		Key:   args.Key,
		Props: flux.ReadValue(args.Props),
		Input: args.Input,
	}
}
