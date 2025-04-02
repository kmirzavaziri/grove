package main

import (
	"github.com/kmirzavaziri/grove/go/examples/blog/internal/pkg/components"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

var pages = []*grove.Node{
	components.Page(components.PageArgs{
		Key:   "home",
		Title: "Home",
	}),
	components.Page(components.PageArgs{
		Key:   "users",
		Title: "Users",
	}),
	components.Page(components.PageArgs{
		Key:   "permissions",
		Title: "Permissions",
	}),
	components.Page(components.PageArgs{
		Key:   "events",
		Title: "Events",
	}),
}
