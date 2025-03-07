package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type LMasonryArgs struct {
	Key   string
	Props flux.Read[LMasonryProps]
}

type LMasonryProps struct {
}

func LMasonry(args LMasonryArgs) *grove.Node {
	return &grove.Node{
		Type:  "grovex.LMasonry",
		Key:   args.Key,
		Props: flux.ReadValue(args.Props),
	}
}
