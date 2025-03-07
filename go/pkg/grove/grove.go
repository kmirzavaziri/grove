package grove

import (
	"context"
	"github.com/kmirzavaziri/grove/go/pkg/greq"
	"time"

	"github.com/kmirzavaziri/grove/go/pkg/grr"
)

type Config struct {
	PreExecutionTimeout time.Duration
	SubmitTimeout       time.Duration
}

type Grove struct {
	root   *Node
	config Config
}

func New(trees []*Node, config Config) (*Grove, error) {
	root := &Node{
		Type:     "grove",
		Key:      "grove",
		Role:     "grove",
		Children: trees,
	}

	err := root.validateAndPopulate()
	if err != nil {
		return nil, grr.Wrap(grr.ErrFailedToValidateAndPopulateGrove, "failed to validate and populate grove: %w", err)
	}

	return &Grove{root: root, config: config}, nil
}

func (g *Grove) Render(ctx context.Context, path string, request greq.Request) (*RenderedNode, *grr.StdErr) {
	return render(ctx, g, path, request)
}

func (g *Grove) Submit(ctx context.Context, path string, request greq.Request) (*Action, *grr.StdErr) {
	return submit(ctx, g, path, request)
}
