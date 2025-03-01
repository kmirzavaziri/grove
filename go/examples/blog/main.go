package main

import (
	"fmt"
	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/grovex"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"

	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

func main() {
	g, err := grove.New(
		[]*grove.Node{
			grovex.Page(
				gex.Static(grovex.PageFlux{Title: "Grove Example Home Page"}),
				grovex.PageData{
					Key: "home",
					Content: []*grove.Node{
						grovex.Card(
							gex.Static(grovex.CardFlux{Title: "Blog Card"}),
							grovex.CardData{
								Key: "post1",
								Body: grovex.Text(
									gex.Static(grovex.TextFlux{
										Text: "My First Blog Post",
										Size: grovex.TextSizeH1,
									}),
									grovex.TextData{Key: "title"},
								),
								Footer: grovex.Text(
									gex.Static(grovex.TextFlux{
										Text: "Read More",
										Size: grovex.TextSizePG,
									}),
									grovex.TextData{Key: "read_more"},
								),
							},
						),
					},
				},
			),
		},
		grove.Config{
			PreExecutionTimeout: time.Second,
		},
	)
	if err != nil {
		panic(fmt.Errorf("failed to initiate grove: %w", err))
	}

	http.Handle("/", g.HTTPHandler(grove.HTTPConfig{Logger: logrus.New()}))

	// TODO use log instead of println
	fmt.Println("starting server http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(fmt.Errorf("failed to listen and serve HTTP: %w", err))
	}
}
