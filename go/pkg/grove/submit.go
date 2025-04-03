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

type Input struct {
	Key string
	Def flux.Read[InputDef]
}

type InputDef struct {
	Value *gex.Value
	Guard guard.Guard
}

func submit(ctx context.Context, grove *Grove, path string, request greq.Request) (*Action, *grr.StdErr) {
	node, ok := grove.root.flatChildrenByPath["grove/"+path]
	if !ok {
		return nil, grr.NewStdErr(grr.Wrap(grr.ErrFailedToFindNode, "failed to find node %q", path), codes.NotFound)
	}

	if node.Submit == nil {
		return nil, grr.NewStdErr(
			grr.Wrap(grr.ErrNodeIsNotSubmittable, "node %q is not submittable", path), codes.InvalidArgument,
		)
	}

	requestData, ok := request.Data.Map()
	if !ok {
		return nil, grr.NewStdErr(
			grr.Wrap(grr.ErrRequestDataIsNotAMap, "request data is not a map", path), codes.InvalidArgument,
		)
	}

	inputFluxesAgg := flux.ReadAggregation[string, InputDef]{}

	for key, input := range node.flatInputsByKey {
		inputFluxesAgg.Add(key, input.Def)
	}

	executorsAgg := executor.NewAggregation[string]()

	err := executorsAgg.Add(map[string][]executor.Executor{"": node.Submit.PreExecutors(request)})
	if err != nil {
		return nil, grr.NewStdErr(
			grr.Wrap(grr.ErrFailedToAddSubmitFluxExecutors, "failed to add submit flux executors: %w", err),
			codes.Internal,
		)
	}

	err = executorsAgg.Add(inputFluxesAgg.GetExecutors(request))
	if err != nil {
		return nil, grr.NewStdErr(
			grr.Wrap(grr.ErrFailedToAddInputFluxExecutors, "failed to add input flux executors: %w", err),
			codes.Internal,
		)
	}

	var (
		preExecutorsCtx    context.Context
		preExecutorsCancel context.CancelFunc
	)

	if grove.config.PreExecutionTimeout != 0 {
		preExecutorsCtx, preExecutorsCancel = context.WithTimeout(ctx, grove.config.PreExecutionTimeout)
	} else {
		preExecutorsCtx, preExecutorsCancel = context.WithCancel(ctx)
	}

	defer preExecutorsCancel()

	executionResultsAgg := executorsAgg.Execute(preExecutorsCtx)

	evaluatedInputFluxes, err := inputFluxesAgg.Evaluate(request, executionResultsAgg)
	if err != nil {
		return nil, grr.NewStdErr(
			grr.Wrap(grr.ErrFailedToEvaluateInputFluxes, "failed to evaluate input fluxes: %w", err), codes.Internal,
		)
	}

	renderedInputs := map[string]*RenderedInput{}

	for key, input := range node.flatInputsByKey {
		inputFlux := evaluatedInputFluxes[key]

		renderedInputs[key] = &RenderedInput{
			Key:   input.Key,
			Value: inputFlux.Value,
			Guard: inputFlux.Guard,
		}
	}

	errorMessages := map[string]string{}
	for key, input := range renderedInputs {
		value, ok := requestData[input.Key]
		if !ok {
			value = gex.Nil()
		}

		errorMessage, err := input.Guard.Guard(value)
		if err != nil {
			return nil, grr.NewStdErr(
				grr.Wrap(grr.ErrFailedToRunGuards, "failed to run guards: %w", err), codes.Internal,
			)
		}

		if errorMessage != "" {
			errorMessages[key] = errorMessage
		}
	}

	if len(errorMessages) != 0 {
		return ActionRenderErrorMessages(ActionRenderErrorMessagesPayload{ErrorMessages: errorMessages}), nil
	}

	var (
		submitCtx    context.Context
		submitCancel context.CancelFunc
	)

	if grove.config.SubmitTimeout != 0 {
		submitCtx, submitCancel = context.WithTimeout(ctx, grove.config.SubmitTimeout)
	} else {
		submitCtx, submitCancel = context.WithCancel(ctx)
	}

	defer submitCancel()

	action, err := node.Submit.Data(submitCtx, request, executionResultsAgg[""])
	if err != nil {
		return nil, grr.NewStdErr(
			grr.Wrap(grr.ErrFailedToSubmit, "failed to submit: %w", err), codes.Internal,
		)
	}

	return action, nil
}
