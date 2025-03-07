package components

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
	"github.com/kmirzavaziri/grove/go/pkg/grovex"
	"github.com/kmirzavaziri/grove/go/pkg/grovex/variants"
)

type PageArgs struct {
	Key   string
	Title string
	Main  []*grove.Node
}

func Page(args PageArgs) *grove.Node {
	navbarStart := []*grove.Node{
		grovex.DBreadcrumbs(grovex.DBreadcrumbsArgs{
			Key: "breadcrumbs",
			Props: flux.ReadStatic(grovex.DBreadcrumbsProps{
				Items: []string{"Dashboard", args.Title},
			}),
		}),
	}

	navbarEnd := []*grove.Node{
		grovex.DTypography(grovex.DTypographyArgs{
			Key: "date",
			Props: flux.ReadStatic(grovex.DTypographyProps{
				Text:    "March 30, 2025",
				Variant: gex.P(variants.TypographyVariantBody2),
			}),
		}),
	}

	sidebarStart := []*grove.Node{
		grovex.XClickable(grovex.XClickableArgs{
			Key: "avatar_xc",
			Props: flux.ReadStatic(grovex.XClickableProps{
				Action: grovex.AFetch(grovex.AFetchPayload{
					NodePath: "home",
				}),
			}),
			Children: []*grove.Node{
				grovex.DProfile(grovex.DProfileArgs{
					Key: "avatar",
					Props: flux.ReadStatic(grovex.DProfileProps{
						Image:    "/static/images/avatar/7.jpg",
						Title:    "Kamyar Mirzavaziri",
						Subtitle: "u/kmirzavaziri",
					}),
				}),
			},
		}),
		grovex.DDivider(grovex.DDividerArgs{Key: "avatar_divider"}),
		menu(args.Key),
	}

	sidebarEnd := []*grove.Node{
		grovex.XClickable(grovex.XClickableArgs{
			Key: "logout",
			Props: flux.ReadStatic(grovex.XClickableProps{
				// TODO
				Action: nil,
			}),
			Children: []*grove.Node{
				grovex.DButton(grovex.DButtonArgs{
					Key: "logout",
					Props: flux.ReadStatic(grovex.DButtonProps{
						Text:      "Logout",
						Variant:   gex.P(variants.ButtonVariantOutlined),
						FullWidth: true,
						StartIcon: gex.P(variants.IconLogoutRounded),
					}),
				}),
			},
		}),
	}

	return grovex.LPage(grovex.LPageArgs{
		Key:            args.Key,
		Props:          flux.ReadStatic(grovex.LPageProps{Title: args.Title}),
		SidebarStart:   sidebarStart,
		SidebarEnd:     sidebarEnd,
		NavbarStart:    navbarStart,
		NavbarEnd:      navbarEnd,
		SidebarStartXS: sidebarStart,
		SidebarEndXS:   sidebarEnd,
		NavbarStartXS:  navbarStart,
		NavbarEndXS:    navbarEnd,
		Main:           args.Main,
	})
}

func menu(pageKey string) *grove.Node {
	return grovex.LBox(grovex.LBoxArgs{
		Key: "menu",
		Children: []*grove.Node{
			menuItem("users", variants.IconPeople, "Users", pageKey == "users"),
			menuItem("permissions", variants.IconKey, "Permissions", pageKey == "permissions"),
			menuItem("events", variants.IconEvent, "Events", pageKey == "events"),
		},
	})
}

func menuItem(key string, icon variants.Icon, text string, selected bool) *grove.Node {
	return grovex.XClickable(grovex.XClickableArgs{
		Key: key,
		Props: flux.ReadStatic(grovex.XClickableProps{
			Action: grovex.AFetch(grovex.AFetchPayload{
				NodePath: key,
			}),
		}),
		Children: []*grove.Node{
			grovex.DButton(grovex.DButtonArgs{
				Key: "button",
				Props: flux.ReadStatic(grovex.DButtonProps{
					Text:      text,
					Variant:   gex.P(variants.ButtonVariantItem),
					FullWidth: true,
					Selected:  selected,
					StartIcon: gex.P(icon),
				}),
			}),
		},
	})
}
