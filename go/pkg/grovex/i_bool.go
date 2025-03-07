package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type IBoolArgs struct {
	Key   string
	Props flux.Read[IBoolProps]
	Input *grove.Input
}

type IBoolProps struct {
}

func IBool(args IBoolArgs) *grove.Node {
	return &grove.Node{
		Type:  "grovex.IBool",
		Key:   args.Key,
		Props: flux.ReadValue(args.Props),
		Input: args.Input,
	}
}
