package serve

import (
	"encoding/json"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/greq"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
	"net/http"
	"net/url"
	"strings"
)

type HTTPMetaData struct {
	Method  string
	URL     *url.URL
	Header  http.Header
	Trailer http.Header
}

type HTTPConfig struct {
	Logger         Logger
	NodePathGetter flux.Pure[string]
}

type HTTPHandler struct {
	g *grove.Grove
	c HTTPConfig
}

func NewHTTPHandler(g *grove.Grove, config HTTPConfig) http.Handler {
	if config.NodePathGetter == nil {
		config.NodePathGetter = flux.PureInline(func(request greq.Request) string {
			// TODO unmarshal request.Metadata to its type instead of chaining
			path, err := request.Metadata.Chain().Get("url").Get("path").String()
			if err != nil {
				return ""
			}

			// TODO can we make it more robust?
			nodePath, ok := strings.CutPrefix(path, "/grove/fetch/")
			if !ok {
				nodePath, _ = strings.CutPrefix(path, "/grove/submit/")
			}

			return nodePath
		})
	}

	return &HTTPHandler{g: g, c: config}
}

func (h *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.error(w, "invalid method", fmt.Errorf("method should be POST"), http.StatusMethodNotAllowed)
		return
	}

	switch true {
	case strings.HasPrefix(r.URL.Path, "/grove/fetch/"):
		h.render(w, r)
	case strings.HasPrefix(r.URL.Path, "/grove/submit/"):
		h.submit(w, r)
	default:
		h.error(w, "not found", fmt.Errorf("should start with /grove/fetch or /grove/submit"), http.StatusNotFound)
	}
}

func (h *HTTPHandler) render(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request, err := h.parseRequest(r)
	if err != nil {
		h.error(w, "invalid request json", err, http.StatusBadRequest)
	}

	nodePath := h.getNodePath(r, request)

	res, stdErr := h.g.Render(r.Context(), nodePath, request)
	if stdErr != nil {
		h.error(w, "failed to render", stdErr, runtime.HTTPStatusFromCode(stdErr.Code()))

		return
	}

	h.respond(w, res)
}

func (h *HTTPHandler) submit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request, err := h.parseRequest(r)
	if err != nil {
		h.error(w, "invalid request json", err, http.StatusBadRequest)
	}

	nodePath := h.getNodePath(r, request)

	res, stdErr := h.g.Submit(r.Context(), nodePath, request)
	if stdErr != nil {
		h.error(w, "failed to submit", stdErr, runtime.HTTPStatusFromCode(stdErr.Code()))

		return
	}

	h.respond(w, res)
}

func (h *HTTPHandler) respond(w http.ResponseWriter, res any) {
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		h.c.Logger.Errorf("failed to encode response: %s", err)
	}
}

func (h *HTTPHandler) getNodePath(r *http.Request, request greq.Request) string {
	return h.c.NodePathGetter.Data(request)
}

func (h *HTTPHandler) parseRequest(r *http.Request) (greq.Request, error) {
	requestData := gex.Nil()
	if r.ContentLength > 0 {
		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			return greq.Request{}, err
		}
	}
	return greq.Request{
		Metadata: gex.Marshal(HTTPMetaData{
			Method:  r.Method,
			URL:     r.URL,
			Header:  r.Header,
			Trailer: r.Trailer,
		}),
		Data: requestData,
	}, nil
}

func (h *HTTPHandler) error(w http.ResponseWriter, errorMessage string, err error, status int) {
	h.c.Logger.Errorf("%s: %s", errorMessage, err)
	w.WriteHeader(status)
	writeErr := json.NewEncoder(w).Encode(map[string]string{"error": errorMessage, "details": err.Error()})
	if writeErr != nil {
		h.c.Logger.Errorf("failed to encode response: %s", err)
	}
}
