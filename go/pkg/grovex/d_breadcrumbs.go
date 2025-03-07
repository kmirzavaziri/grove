package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type DBreadcrumbsArgs struct {
	Key   string
	Props flux.Read[DBreadcrumbsProps]
}

type DBreadcrumbsProps struct {
	Items []string
}

func DBreadcrumbs(args DBreadcrumbsArgs) *grove.Node {
	return &grove.Node{
		Type:  "grovex.DBreadcrumbs",
		Key:   args.Key,
		Props: flux.ReadValue(args.Props),
	}
}
