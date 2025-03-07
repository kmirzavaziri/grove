package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type DProfileArgs struct {
	Key   string
	Props flux.Read[DProfileProps]
}

type DProfileProps struct {
	Image    string
	Title    string
	Subtitle string
}

func DProfile(args DProfileArgs) *grove.Node {
	return &grove.Node{
		Type:  "grovex.DProfile",
		Key:   args.Key,
		Props: flux.ReadValue(args.Props),
	}
}
