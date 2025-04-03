package flux

import (
	"context"

	"github.com/kmirzavaziri/grove/go/pkg/executor"
	"github.com/kmirzavaziri/grove/go/pkg/greq"
)

type Act[T any] interface {
	PreExecutors(request greq.Request) []executor.Executor
	Data(ctx context.Context, request greq.Request, executorResults executor.Results) (T, error)
}

type actInline[T any] struct {
	preExecutors func(request greq.Request) []executor.Executor
	data         func(ctx context.Context, request greq.Request, executorResults executor.Results) (T, error)
}

func ActInline[T any](
	preExecutors func(request greq.Request) []executor.Executor,
	data func(ctx context.Context, request greq.Request, executorResults executor.Results) (T, error),
) Act[T] {
	return &actInline[T]{
		preExecutors: preExecutors,
		data:         data,
	}
}

func (f *actInline[T]) PreExecutors(request greq.Request) []executor.Executor {
	if f.preExecutors == nil {
		return []executor.Executor{}
	}

	return f.preExecutors(request)
}

func (f *actInline[T]) Data(ctx context.Context, request greq.Request, executorResults executor.Results) (T, error) {
	if f.data == nil {
		return *new(T), nil
	}

	return f.data(ctx, request, executorResults)
}
