package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type DDividerArgs struct {
	Key   string
	Props flux.Read[DDividerProps]
}

type DDividerProps struct {
}

func DDivider(args DDividerArgs) *grove.Node {
	return &grove.Node{
		Type:  "grovex.DDivider",
		Key:   args.Key,
		Props: flux.ReadValue(args.Props),
	}
}
