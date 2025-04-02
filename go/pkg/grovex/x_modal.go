package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type XModalArgs struct {
	Key      string
	Props    flux.Read[XModalProps]
	Children []*grove.Node
}

type XModalProps struct {
	Title       string
	Description string
	Open        bool
}

func XModal(args XModalArgs) *grove.Node {
	return &grove.Node{
		Type:     "grovex.XModal",
		Key:      args.Key,
		Props:    flux.ReadValue(args.Props),
		Children: args.Children,
	}
}
