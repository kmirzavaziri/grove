package components

import (
	"context"

	"github.com/kmirzavaziri/grove/go/pkg/executor"
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/greq"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
	"github.com/kmirzavaziri/grove/go/pkg/grovex"
	"github.com/kmirzavaziri/grove/go/pkg/grovex/variants"
)

type PageArgs struct {
	Key         string
	Title       string
	Main        []*grove.Node
	Breadcrumbs []string
}

func Page(args PageArgs) *grove.Node {
	breadcrumbItems := make([]grovex.DBreadcrumbsItem, 0, len(args.Breadcrumbs)+1)

	breadcrumbItems = append(breadcrumbItems, grovex.DBreadcrumbsItem{
		Title: "Dashboard",
		Action: grovex.ARender(grovex.ARenderPayload{
			NodePath:      []string{"home"},
			UpdateHistory: true,
		}),
	})

	for _, bc := range args.Breadcrumbs {
		breadcrumbItems = append(breadcrumbItems, grovex.DBreadcrumbsItem{
			Title: bc,
		})
	}

	navbarStart := []*grove.Node{
		grovex.DBreadcrumbs(grovex.DBreadcrumbsArgs{
			Key: "breadcrumbs",
			Props: flux.ReadStatic(grovex.DBreadcrumbsProps{
				Items: breadcrumbItems,
			}),
		}),
	}

	navbarEnd := []*grove.Node{
		grovex.DTypography(grovex.DTypographyArgs{
			Key: "date",
			Props: flux.ReadStatic(grovex.DTypographyProps{
				Text:    "March 30, 2025",
				Variant: gex.P(variants.DTypographyVariantBody2),
			}),
		}),
	}

	return grovex.LPage(grovex.LPageArgs{
		Key:            args.Key,
		Props:          flux.ReadStatic(grovex.LPageProps{Title: args.Title}),
		SidebarStart:   sidebarStart(args.Key),
		SidebarEnd:     sidebarEnd(),
		NavbarStart:    navbarStart,
		NavbarEnd:      navbarEnd,
		SidebarStartXS: sidebarStart(args.Key),
		SidebarEndXS:   sidebarEnd(),
		NavbarStartXS:  navbarStart,
		NavbarEndXS:    navbarEnd,
		Main:           args.Main,
	})
}

func sidebarStart(pageKey string) []*grove.Node {
	return []*grove.Node{
		// grovex.XClickable(grovex.XClickableArgs{
		// 	Key: "avatar_xc",
		// 	Props: flux.ReadStatic(grovex.XClickableProps{
		// 		Action: grovex.ARender(grovex.ARenderPayload{
		// 			NodePath:      []string{"home"},
		// 			UpdateHistory: true,
		// 		}),
		// 	}),
		// 	Children: []*grove.Node{
		// 		grovex.DProfile(grovex.DProfileArgs{
		// 			Key: "avatar",
		// 			Props: flux.ReadStatic(grovex.DProfileProps{
		// 				Image:    "/static/images/avatar/7.jpg",
		// 				Title:    "Kamyar Mirzavaziri",
		// 				Subtitle: "u/kmirzavaziri",
		// 			}),
		// 		}),
		// 	},
		// }),
		grovex.XClickable(grovex.XClickableArgs{
			Key: "avatar",
			Props: flux.ReadStatic(grovex.XClickableProps{
				Action: grovex.ARender(grovex.ARenderPayload{
					NodePath: []string{pageKey, "sidebar_start", "auth"},
					Patch:    true,
					Node: &grove.RenderedNode{
						Props: gex.Marshal(grovex.XModalProps{Open: true}),
					},
				}),
			}),
			Children: []*grove.Node{
				grovex.DProfile(grovex.DProfileArgs{
					Key: "profile",
					Props: flux.ReadStatic(grovex.DProfileProps{
						Title:    "Guest",
						Subtitle: "guest",
					}),
				}),
			},
		}),
		grovex.MForm(
			grovex.XModal(grovex.XModalArgs{
				Key:   "auth",
				Props: flux.ReadStatic(grovex.XModalProps{}),
				Children: []*grove.Node{
					grovex.DTypography(grovex.DTypographyArgs{
						Key: "title",
						Props: flux.ReadStatic(grovex.DTypographyProps{
							Text:    "Authenticate",
							Variant: gex.P(variants.DTypographyVariantH6),
						}),
					}),
					grovex.IText(grovex.ITextArgs{
						Key: "username",
						Input: &grove.Input{
							Key: "username",
							Def: flux.ReadStatic(grove.InputDef{}),
						},
						Props: flux.ReadStatic(grovex.ITextProps{
							Variant:      gex.P(variants.ITextVariantText),
							Label:        "Username",
							AutoComplete: "username",
						}),
					}),
					grovex.IText(grovex.ITextArgs{
						Key: "password",
						Input: &grove.Input{
							Key: "password",
							Def: flux.ReadStatic(grove.InputDef{}),
						},
						Props: flux.ReadStatic(grovex.ITextProps{
							Variant:      gex.P(variants.ITextVariantPassword),
							Label:        "Password",
							AutoComplete: "password",
						}),
					}),
					grovex.XClickable(grovex.XClickableArgs{
						Key: "login",
						Props: flux.ReadStatic(grovex.XClickableProps{
							Action: grovex.ASubmit(grovex.ASubmitPayload{
								NodePath: []string{pageKey, "sidebar_start", "auth"},
							}),
						}),
						Children: []*grove.Node{
							grovex.DButton(grovex.DButtonArgs{
								Key: "login",
								Props: flux.ReadStatic(grovex.DButtonProps{
									Text:      "Login",
									Variant:   gex.P(variants.DButtonVariantOutlined),
									FullWidth: true,
								}),
							}),
						},
					}),
				},
			}),
			flux.ActInline(
				nil,
				func(ctx context.Context, request greq.Request, executorResults executor.Results) (*grove.Action, error) {
					return nil, nil
				},
			),
		),
		grovex.DDivider(grovex.DDividerArgs{Key: "avatar_divider"}),
		menu(pageKey),
	}
}

func sidebarEnd() []*grove.Node {
	return []*grove.Node{
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
						Variant:   gex.P(variants.DButtonVariantOutlined),
						FullWidth: true,
						StartIcon: gex.P(variants.IconLogoutRounded),
					}),
				}),
			},
		}),
	}
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
			Action: grovex.ARender(grovex.ARenderPayload{
				NodePath:      []string{key},
				UpdateHistory: true,
			}),
		}),
		Children: []*grove.Node{
			grovex.DButton(grovex.DButtonArgs{
				Key: "button",
				Props: flux.ReadStatic(grovex.DButtonProps{
					Text:      text,
					Variant:   gex.P(variants.DButtonVariantItem),
					FullWidth: true,
					Selected:  selected,
					StartIcon: gex.P(icon),
				}),
			}),
		},
	})
}
