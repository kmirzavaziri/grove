package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type ISingleSelectArgs struct {
	Key   string
	Props flux.Read[ISingleSelectProps]
	Input *grove.Input
}

type ISingleSelectProps struct {
}

func ISingleSelect(args ISingleSelectArgs) *grove.Node {
	return &grove.Node{
		Type:  "grovex.ISingleSelect",
		Key:   args.Key,
		Props: flux.ReadValue(args.Props),
		Input: args.Input,
	}
}
