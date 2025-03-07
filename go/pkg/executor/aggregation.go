package executor

import (
	"context"
	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/grr"
	"reflect"
	"sync"
)

type Aggregation[K comparable] struct {
	executors    map[Hash]Executor
	resultHashes map[K][]Hash
}

func NewAggregation[K comparable]() Aggregation[K] {
	return Aggregation[K]{
		executors:    map[Hash]Executor{},
		resultHashes: map[K][]Hash{},
	}
}

type ResultsAggregation[K comparable] map[K]Results

func (ee Aggregation[K]) Add(executorsMap map[K][]Executor) error {
	for key, executors := range executorsMap {
		ee.resultHashes[key] = make([]Hash, len(executors))

		for i, executor := range executors {
			hash, err := executor.Hash()
			if err != nil {
				return grr.Wrap(
					grr.ErrFailedToGetExecutorHash,
					"failed to get executor %s hash: %w", reflect.TypeOf(executor).String(), err,
				)
			}

			ee.executors[hash] = executor

			ee.resultHashes[key][i] = hash
		}
	}

	return nil
}

func (ee Aggregation[K]) Execute(ctx context.Context) ResultsAggregation[K] {
	executionResultsAgg := ResultsAggregation[K]{}

	for key := range ee.resultHashes {
		executionResultsAgg[key] = Results{}
	}

	if len(ee.executors) == 0 {
		return executionResultsAgg
	}

	var wg sync.WaitGroup

	resultCh := make(chan struct {
		hash  Hash
		value *gex.Value
	}, len(ee.executors))

	for hash, executor := range ee.executors {
		wg.Add(1)

		go func(h Hash, e Executor) {
			defer wg.Done()
			resultCh <- struct {
				hash  Hash
				value *gex.Value
			}{h, e.Execute(ctx)}
		}(hash, executor)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	executionResults := Results{}

	for res := range resultCh {
		executionResults[res.hash] = res.value
	}

	for key := range ee.resultHashes {
		for _, hash := range ee.resultHashes[key] {
			executionResultsAgg[key][hash] = executionResults[hash]
		}
	}

	return executionResultsAgg
}
