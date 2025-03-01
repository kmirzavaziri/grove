package gex

type FastFlux[T any] interface {
	Data(requestData *Value) T
}

type Flux[T any] interface {
	PreExecutors(requestData *Value) []Executor
	Data(requestData *Value, executionResults ExecutionResults) (T, error)
}

type fluxV[T any] struct {
	f Flux[T]
}

func FluxV[T any](f Flux[T]) Flux[*Value] {
	return &fluxV[T]{f}
}

func (f *fluxV[T]) PreExecutors(requestData *Value) []Executor {
	if f.f == nil {
		return nil
	}

	return f.f.PreExecutors(requestData)
}

func (f *fluxV[T]) Data(requestData *Value, executionResults ExecutionResults) (*Value, error) {
	if f.f == nil {
		var zero T

		return Marshal(zero), nil
	}

	r, err := f.f.Data(requestData, executionResults)

	return Marshal(r), err
}

func Static[T any](value T) Flux[T] {
	return &static[T]{value: value}
}

type static[T any] struct {
	value T
}

func (s *static[T]) PreExecutors(_ *Value) []Executor {
	return nil
}

func (s *static[T]) Data(_ *Value, _ ExecutionResults) (T, error) {
	return s.value, nil
}

func StaticV[T any](value T) Flux[*Value] {
	return &fluxV[T]{Static(value)}
}
