package grove

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/grr"
)

type Config struct {
	PreExecutionTimeout time.Duration
}

type Grove struct {
	root   *walkedNode
	config Config

	lookupByPath map[string]*walkedNode
}

func New(trees []*Node, config Config) (*Grove, error) {
	root, err := newWalkedNode(&Node{
		Type:     "grove",
		Key:      "grove",
		Role:     "grove",
		Children: trees,
	}, nil)
	if err != nil {
		return nil, grr.Wrap(grr.ErrFailedToWalkNode, "failed to walk over grove: %w", err)
	}

	lookupByPath := map[string]*walkedNode{strings.Join(root.path, "/"): root}

	for _, c := range root.flatChildren {
		lookupByPath[strings.Join(c.path, "/")] = c
	}

	return &Grove{root: root, config: config, lookupByPath: lookupByPath}, nil
}

func (g *Grove) Render(ctx context.Context, path string, request *gex.Value) (*RenderedNode, error) {
	return render(ctx, g, path, request)
}

func (g *Grove) HTTPHandler(config HTTPConfig) http.Handler {
	return &httpHandler{g: g, c: config}
}
