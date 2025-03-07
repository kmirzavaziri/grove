package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type AFetchPayload struct {
	NodePath string
	request  *gex.Value
}

func AFetch(payload AFetchPayload) *grove.Action {
	return &grove.Action{
		Type:    "grovex.AFetch",
		Payload: gex.Marshal(payload),
	}
}
