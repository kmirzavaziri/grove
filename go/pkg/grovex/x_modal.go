package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type XModalArgs struct {
	Key   string
	Props flux.Read[XModalProps]
}

type XModalProps struct {
}

func XModal(args XModalArgs) *grove.Node {
	return &grove.Node{
		Type:  "grovex.XModal",
		Key:   args.Key,
		Props: flux.ReadValue(args.Props),
	}
}
