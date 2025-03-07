package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type LPageArgs struct {
	Key   string
	Props flux.Read[LPageProps]

	SidebarStart []*grove.Node
	SidebarEnd   []*grove.Node
	NavbarStart  []*grove.Node
	NavbarEnd    []*grove.Node

	SidebarStartXS []*grove.Node
	SidebarEndXS   []*grove.Node
	NavbarStartXS  []*grove.Node
	NavbarEndXS    []*grove.Node

	Main []*grove.Node
}

type LPageProps struct {
	Title string
}

func LPage(args LPageArgs) *grove.Node {
	return &grove.Node{
		Type:  "grovex.LPage",
		Key:   args.Key,
		Props: flux.ReadValue(args.Props),
		Children: []*grove.Node{
			LFragment(LFragmentArgs{Role: "main", Children: args.Main}),
			LFragment(LFragmentArgs{Role: "sidebar_start", Children: args.SidebarStart}),
			LFragment(LFragmentArgs{Role: "sidebar_end", Children: args.SidebarEnd}),
			LFragment(LFragmentArgs{Role: "sidebar_start_xs", Children: args.SidebarStartXS}),
			LFragment(LFragmentArgs{Role: "sidebar_end_xs", Children: args.SidebarEndXS}),
			LFragment(LFragmentArgs{Role: "navbar_start", Children: args.NavbarStart}),
			LFragment(LFragmentArgs{Role: "navbar_end", Children: args.NavbarEnd}),
			LFragment(LFragmentArgs{Role: "navbar_start_xs", Children: args.NavbarStartXS}),
			LFragment(LFragmentArgs{Role: "navbar_end_xs", Children: args.NavbarEndXS}),
		},
	}
}
