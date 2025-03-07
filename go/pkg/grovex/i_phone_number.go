package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type IPhoneNumberArgs struct {
	Key   string
	Props flux.Read[IPhoneNumberProps]
	Input *grove.Input
}

type IPhoneNumberProps struct {
}

func IPhoneNumber(args IPhoneNumberArgs) *grove.Node {
	return &grove.Node{
		Type:  "grovex.IPhoneNumber",
		Key:   args.Key,
		Props: flux.ReadValue(args.Props),
		Input: args.Input,
	}
}
