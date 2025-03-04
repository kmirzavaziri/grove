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

	mux := http.NewServeMux()

	mux.Handle("/", g.HTTPHandler(grove.HTTPConfig{Logger: logrus.New()}))

	// TODO use log instead of println
	fmt.Println("starting server http://localhost:8080")

	if err := http.ListenAndServe(":8080", corsMiddleware(mux)); err != nil {
		panic(fmt.Errorf("failed to listen and serve HTTP: %w", err))
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
