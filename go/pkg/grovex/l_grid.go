package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type LGridArgs struct {
	Key   string
	Props flux.Read[LGridProps]
}

type LGridProps struct {
}

func LGrid(args LGridArgs) *grove.Node {
	return &grove.Node{
		Type:  "grovex.LGrid",
		Key:   args.Key,
		Props: flux.ReadValue(args.Props),
	}
}
