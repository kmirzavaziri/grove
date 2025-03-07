package grove

import (
	"context"
	"github.com/kmirzavaziri/grove/go/pkg/executor"
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/greq"
	"github.com/kmirzavaziri/grove/go/pkg/grr"
	"github.com/kmirzavaziri/grove/go/pkg/guard"
	"google.golang.org/grpc/codes"
)

type RenderedNode struct {
	Type     string          `json:"type"`
	Key      string          `json:"key"`
	Role     string          `json:"role"`
	Props    *gex.Value      `json:"props"`
	Input    *RenderedInput  `json:"input"`
	Children []*RenderedNode `json:"children"`
}

type RenderedInput struct {
	Key   string      `json:"key"`
	Value *gex.Value  `json:"value"`
	Guard guard.Guard `json:"guard"`
}

type fluxKey struct {
	node     *Node
	forInput bool
}

func render(ctx context.Context, grove *Grove, path string, request greq.Request) (*RenderedNode, *grr.StdErr) {
	node, ok := grove.root.flatChildrenByPath["grove/"+path]
	if !ok {
		return nil, grr.NewStdErr(grr.Wrap(grr.ErrFailedToFindNode, "failed to find node %q", path), codes.NotFound)
	}

	var cancel context.CancelFunc

	if grove.config.PreExecutionTimeout != 0 {
		ctx, cancel = context.WithTimeout(ctx, grove.config.PreExecutionTimeout)
	} else {
		ctx, cancel = context.WithCancel(ctx)
	}

	defer cancel()

	nodeFluxesAgg := flux.ReadAggregation[fluxKey, *gex.Value]{}

	nodeFluxesAgg.Add(fluxKey{node: node}, node.Props)

	for _, c := range node.flatChildrenByPath {
		nodeFluxesAgg.Add(fluxKey{node: c}, c.Props)
	}

	inputFluxesAgg := flux.ReadAggregation[fluxKey, InputDef]{}

	for _, c := range node.flatChildrenByPath {
		if c.Input != nil {
			inputFluxesAgg.Add(fluxKey{node: c, forInput: true}, c.Input.Def)
		}
	}

	executorsAgg := executor.NewAggregation[fluxKey]()

	err := executorsAgg.Add(nodeFluxesAgg.GetExecutors(request))
	if err != nil {
		return nil, grr.NewStdErr(
			grr.Wrap(grr.ErrFailedToAddNodeFluxExecutors, "failed to add node flux executors: %w", err), codes.Internal,
		)
	}

	err = executorsAgg.Add(inputFluxesAgg.GetExecutors(request))
	if err != nil {
		return nil, grr.NewStdErr(
			grr.Wrap(grr.ErrFailedToAddInputFluxExecutors, "failed to add input flux executors: %w", err), codes.Internal,
		)
	}

	executionResultsAgg := executorsAgg.Execute(ctx)

	evaluatedNodeFluxes, err := nodeFluxesAgg.Evaluate(request, executionResultsAgg)
	if err != nil {
		return nil, grr.NewStdErr(
			grr.Wrap(grr.ErrFailedToEvaluateNodeFluxes, "failed to evaluate node fluxes: %w", err), codes.Internal,
		)
	}

	evaluatedInputFluxes, err := inputFluxesAgg.Evaluate(request, executionResultsAgg)
	if err != nil {
		return nil, grr.NewStdErr(
			grr.Wrap(grr.ErrFailedToEvaluateInputFluxes, "failed to evaluate input fluxes: %w", err), codes.Internal,
		)
	}

	return renderNode(node, request, evaluatedNodeFluxes, evaluatedInputFluxes)
}

func renderNode(
	node *Node, request greq.Request,
	evaluatedNodeFluxes map[fluxKey]*gex.Value, evaluatedInputFluxes map[fluxKey]InputDef,
) (*RenderedNode, *grr.StdErr) {
	var input *RenderedInput

	if node.Input != nil {
		inputFlux := evaluatedInputFluxes[fluxKey{node: node, forInput: true}]

		input = &RenderedInput{
			Key:   node.Input.Key,
			Value: inputFlux.Value,
			Guard: inputFlux.Guard,
		}
	}

	renderedNode := &RenderedNode{
		Type:     node.Type,
		Key:      node.Key,
		Role:     node.Role,
		Props:    evaluatedNodeFluxes[fluxKey{node: node}],
		Input:    input,
		Children: make([]*RenderedNode, 0, len(node.Children)),
	}

	for _, c := range node.Children {
		renderedChild, err := renderNode(c, request, evaluatedNodeFluxes, evaluatedInputFluxes)
		if err != nil {
			return nil, grr.NewStdErr(
				grr.Wrap(grr.ErrFailedToRenderChild, "failed to render child %q: %w", c.path, err), codes.Internal,
			)
		}

		if renderedChild == nil {
			continue
		}

		renderedNode.Children = append(renderedNode.Children, renderedChild)
	}

	return renderedNode, nil
}
