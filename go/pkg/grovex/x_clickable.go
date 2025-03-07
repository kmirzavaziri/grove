package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type XClickableArgs struct {
	Key      string
	Props    flux.Read[XClickableProps]
	Children []*grove.Node
}

type XClickableProps struct {
	Action *grove.Action
}

func XClickable(args XClickableArgs) *grove.Node {
	return &grove.Node{
		Type:     "grovex.XClickable",
		Key:      args.Key,
		Props:    flux.ReadValue(args.Props),
		Children: args.Children,
	}
}
