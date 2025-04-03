package main

import (
	"github.com/kmirzavaziri/grove/go/examples/dashboard/internal/pkg/components"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

var pages = []*grove.Node{
	components.Page(components.PageArgs{
		Key:   "home",
		Title: "Home",
	}),
	components.Page(components.PageArgs{
		Key:         "users",
		Title:       "Users",
		Breadcrumbs: []string{"Users"},
	}),
	components.Page(components.PageArgs{
		Key:         "permissions",
		Title:       "Permissions",
		Breadcrumbs: []string{"Permissions"},
	}),
	components.Page(components.PageArgs{
		Key:         "events",
		Title:       "Events",
		Breadcrumbs: []string{"Events"},
	}),
}
