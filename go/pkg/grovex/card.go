package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type CardFlux struct {
	Title string
}

type CardData struct {
	Key    string
	Body   *grove.Node
	Footer *grove.Node
}

func Card(flux gex.Flux[CardFlux], data CardData) *grove.Node {
	if data.Body == nil {
		return grove.Error(data.Key, "card body cannot be nil")
	}

	if data.Footer == nil {
		return grove.Error(data.Key, "card footer cannot be nil")
	}

	return &grove.Node{
		Type: "grovex.Card",
		Key:  data.Key,
		Data: gex.FluxV(flux),
		Children: []*grove.Node{
			{
				Type:     data.Body.Type,
				Key:      data.Body.Key,
				Role:     "body",
				Data:     data.Body.Data,
				Children: data.Body.Children,
			},
			{
				Type:     data.Footer.Type,
				Key:      data.Footer.Key,
				Role:     "footer",
				Data:     data.Footer.Data,
				Children: data.Footer.Children,
			},
		},
	}
}
