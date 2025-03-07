package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type LFragmentArgs struct {
	Key      string
	Role     string
	Props    flux.Read[LFragmentProps]
	Children []*grove.Node
}

type LFragmentProps struct{}

func LFragment(args LFragmentArgs) *grove.Node {
	return &grove.Node{
		Type:     "grovex.LFragment",
		Key:      args.Key,
		Role:     args.Role,
		Props:    flux.ReadValue(args.Props),
		Children: args.Children,
	}
}
