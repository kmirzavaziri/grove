package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type PageFlux struct {
	Title string
}

type PageData struct {
	Key     string
	Content []*grove.Node
}

func Page(flux gex.Flux[PageFlux], data PageData) *grove.Node {
	return &grove.Node{
		Type:     "grovex.Page",
		Key:      data.Key,
		Data:     gex.FluxV(flux),
		Children: data.Content,
	}
}
