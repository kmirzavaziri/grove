package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type LTabsArgs struct {
	Key   string
	Props flux.Read[LTabsProps]
}

type LTabsProps struct {
}

func LTabs(args LTabsArgs) *grove.Node {
	return &grove.Node{
		Type:  "grovex.LTabs",
		Key:   args.Key,
		Props: flux.ReadValue(args.Props),
	}
}
