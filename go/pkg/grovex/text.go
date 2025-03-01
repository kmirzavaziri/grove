package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type TextFlux struct {
	Text string
	Size TextSize
}

type TextData struct {
	Key string
}

func Text(flux gex.Flux[TextFlux], data TextData) *grove.Node {
	return &grove.Node{
		Type: "grovex.Text",
		Key:  data.Key,
		Data: gex.FluxV(flux),
	}
}
