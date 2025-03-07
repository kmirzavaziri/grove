package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type LBoxArgs struct {
	Key      string
	Role     string
	Props    flux.Read[LBoxProps]
	Children []*grove.Node
}

type LBoxProps struct{}

func LBox(args LBoxArgs) *grove.Node {
	return &grove.Node{
		Type:     "grovex.LBox",
		Key:      args.Key,
		Role:     args.Role,
		Props:    flux.ReadValue(args.Props),
		Children: args.Children,
	}
}
