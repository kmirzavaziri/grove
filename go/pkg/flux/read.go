package flux

import (
	"github.com/kmirzavaziri/grove/go/pkg/executor"
	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/greq"
	"github.com/kmirzavaziri/grove/go/pkg/grr"
)

type Read[T any] interface {
	PreExecutors(request greq.Request) []executor.Executor
	Data(request greq.Request, executorResults executor.Results) (T, error)
}

type readInline[T any] struct {
	preExecutors func(request greq.Request) []executor.Executor
	data         func(request greq.Request, executorResults executor.Results) (T, error)
}

func ReadInline[T any](
	preExecutors func(request greq.Request) []executor.Executor,
	data func(request greq.Request, executorResults executor.Results) (T, error),
) Read[T] {
	return &readInline[T]{
		preExecutors: preExecutors,
		data:         data,
	}
}

func (f *readInline[T]) PreExecutors(request greq.Request) []executor.Executor {
	if f.preExecutors == nil {
		return []executor.Executor{}
	}

	return f.preExecutors(request)
}

func (f *readInline[T]) Data(request greq.Request, executorResults executor.Results) (T, error) {
	if f.data == nil {
		return *new(T), nil
	}

	return f.data(request, executorResults)
}

func ReadValue[T any](f Read[T]) Read[*gex.Value] {
	return ReadInline(
		func(request greq.Request) []executor.Executor {
			if f == nil {
				return nil
			}

			return f.PreExecutors(request)
		},
		func(request greq.Request, executionResults executor.Results) (*gex.Value, error) {
			if f == nil {
				var zero T

				return gex.Marshal(zero), nil
			}

			r, err := f.Data(request, executionResults)

			return gex.Marshal(r), err
		},
	)
}

func ReadStatic[T any](value T) Read[T] {
	return ReadInline(
		func(_ greq.Request) []executor.Executor {
			return nil
		},
		func(_ greq.Request, _ executor.Results) (T, error) {
			return value, nil
		},
	)
}

func ReadStaticValue[T any](v T) Read[*gex.Value] {
	return ReadValue(ReadStatic(v))
}

type ReadAggregation[K comparable, V any] map[K]Read[V]

func (ff ReadAggregation[K, V]) Add(k K, f Read[V]) {
	ff[k] = f
}

func (ff ReadAggregation[K, V]) GetExecutors(request greq.Request) map[K][]executor.Executor {
	executorsMap := map[K][]executor.Executor{}

	for k, f := range ff {
		if f != nil {
			executorsMap[k] = f.PreExecutors(request)
		} else {
			executorsMap[k] = []executor.Executor{}
		}

	}

	return executorsMap

}

func (ff ReadAggregation[K, V]) Evaluate(request greq.Request, executionResultsAgg executor.ResultsAggregation[K]) (map[K]V, error) {
	result := map[K]V{}

	for k, f := range ff {
		if f == nil {
			result[k] = *new(V)
			continue
		}

		data, err := f.Data(request, executionResultsAgg[k])
		if err != nil {
			return nil, grr.Wrap(grr.ErrFailedToEvaluateFluxData, "failed to evaluate flux data for key %q: %w", k, err)
		}

		result[k] = data
	}

	return result, nil
}

func MustReadEvaluateStatic[T any](f Read[T]) T {
	request := greq.Request{}

	if len(f.PreExecutors(request)) != 0 {
		panic("read flux has pre executors but is expected to be static") // TODO gerr?
	}

	v, err := f.Data(request, nil)
	if err != nil {
		panic("read flux is returning error but is expected to be static") // TODO gerr?
	}

	return v
}
