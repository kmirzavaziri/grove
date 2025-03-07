package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type DImageArgs struct {
	Key   string
	Props flux.Read[DImageProps]
}

type DImageProps struct {
}

func DImage(args DImageArgs) *grove.Node {
	return &grove.Node{
		Type:  "grovex.DImage",
		Key:   args.Key,
		Props: flux.ReadValue(args.Props),
	}
}
