package main

import (
	"fmt"
	"github.com/kmirzavaziri/grove/go/pkg/serve"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"

	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

func main() {
	g, err := grove.New(
		pages,
		grove.Config{
			PreExecutionTimeout: 500 * time.Millisecond,
			SubmitTimeout:       500 * time.Millisecond,
		},
	)
	if err != nil {
		panic(fmt.Errorf("failed to initiate grove: %w", err))
	}

	mux := http.NewServeMux()

	logger := logrus.New()

	mux.Handle("/", serve.NewHTTPHandler(g, serve.HTTPConfig{Logger: logger}))

	logger.Info("starting server http://localhost:8080")

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
