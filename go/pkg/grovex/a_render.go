package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type ARenderPayload struct {
	NodePath      []string
	Request       *gex.Value
	Node          *grove.RenderedNode
	UpdateHistory bool
	Patch         bool
}

func ARender(payload ARenderPayload) *grove.Action {
	return &grove.Action{
		Type:    "grovex.ARender",
		Payload: gex.Marshal(payload),
	}
}
