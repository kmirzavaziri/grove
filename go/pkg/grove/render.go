package grove

import (
	"context"
	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/grr"
)

type RenderedNode struct {
	Type     string          `json:"type"`
	Key      string          `json:"key"`
	Role     string          `json:"role"`
	Data     *gex.Value      `json:"data"`
	Children []*RenderedNode `json:"children"`
}

func render(ctx context.Context, grove *Grove, path string, request *gex.Value) (*RenderedNode, error) {
	node, ok := grove.lookupByPath["grove/"+path]
	if !ok {
		return nil, grr.Wrap(grr.ErrFailedToFindNode, "failed to find node %q", path)
	}

	ctx, cancel := context.WithTimeout(ctx, grove.config.PreExecutionTimeout)
	defer cancel()

	executors := gex.Executors{}
	nodeToExecutorHashes := map[*walkedNode]map[string]struct{}{}

	for _, children := range node.flatChildren {
		nodeToExecutorHashes[children] = map[string]struct{}{}

		if children.data == nil {
			continue
		}

		nodeExecutors := children.data.PreExecutors(request)

		hashes, err := executors.Add(nodeExecutors)
		if err != nil {
			return nil, grr.Wrap(
				grr.ErrFailedToAddFluxExecutors, "failed to add flux executors for node %q: %w", children.path, err,
			)
		}

		nodeToExecutorHashes[children] = hashes
	}

	executionResults := executors.Execute(ctx)

	return renderNode(node, request, executionResults, nodeToExecutorHashes)
}

func renderNode(
	node *walkedNode, request *gex.Value, executionResults gex.ExecutionResults,
	nodeToExecutorHashes map[*walkedNode]map[string]struct{},
) (*RenderedNode, error) {
	if node.data == nil {
		return nil, nil // nolint: nilnil
	}

	data, err := node.data.Data(request, executionResults.Limit(nodeToExecutorHashes[node]))
	if err != nil {
		return nil, grr.Wrap(
			grr.ErrFailedToEvaluateFluxData, "failed to evaluate flux data for node %q: %w", node.path, err,
		)
	}

	renderedNode := &RenderedNode{
		Type:     node.type_,
		Key:      node.key,
		Role:     node.role,
		Data:     data,
		Children: make([]*RenderedNode, 0, len(node.children)),
	}

	for _, c := range node.children {
		renderedChild, err := renderNode(c, request, executionResults, nodeToExecutorHashes)
		if err != nil {
			return nil, grr.Wrap(grr.ErrFailedToRenderChild, "failed to render child %q: %w", c.path, err)
		}

		if renderedChild == nil || renderedChild.Data.IsNil() {
			continue
		}

		renderedNode.Children = append(renderedNode.Children, renderedChild)
	}

	return renderedNode, nil
}
