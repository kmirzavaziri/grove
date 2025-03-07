package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type LAccordionArgs struct {
	Key   string
	Props flux.Read[LAccordionProps]
}

type LAccordionProps struct {
}

func LAccordion(args LAccordionArgs) *grove.Node {
	return &grove.Node{
		Type:  "grovex.LAccordion",
		Key:   args.Key,
		Props: flux.ReadValue(args.Props),
	}
}
