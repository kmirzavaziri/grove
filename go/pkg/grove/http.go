package grove

import (
	"context"
	"encoding/json"
	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"net/http"
	"strings"
)

type HTTPConfig struct {
	Logger         Logger
	NodePathGetter gex.FastFlux[string]
}

type httpHandler struct {
	g *Grove
	c HTTPConfig
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO we're losing request metadata here

    // TODO only if not empty
	request := new(gex.Value)
// 	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
// 		h.c.Logger.Errorf("invalid request json payload: %s", err.Error())
// 		http.Error(w, "invalid request json payload", http.StatusBadRequest)
//
// 		return
// 	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodePath, _ := strings.CutPrefix(r.URL.Path, "/grove/")

	if h.c.NodePathGetter != nil {
		nodePath = h.c.NodePathGetter.Data(request)
	}

	rendered, err := h.g.Render(ctx, nodePath, request)
	if err != nil {
		h.c.Logger.Errorf("failed to render: %s", err.Error())
		http.Error(w, "failed to render page", http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(rendered)
	if err != nil {
		h.c.Logger.Errorf("failed to encode response: %s", err.Error())
		http.Error(w, "failed to encode response", http.StatusInternalServerError)

		return
	}
}
