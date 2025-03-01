package gex

import (
	"context"
	"reflect"
	"sync"

	"github.com/kmirzavaziri/grove/go/pkg/grr"
)

type Executor interface {
	Hash() (string, error)
	Execute(ctx context.Context) Value
}

type Executors map[string]Executor

func (ee Executors) Add(executors []Executor) (map[string]struct{}, error) {
	hashes := map[string]struct{}{}

	for _, executor := range executors {
		hash, err := executor.Hash()
		if err != nil {
			return nil, grr.Wrap(
				grr.ErrFailedToGetExecutorHash,
				"failed to get executor %s hash: %w", reflect.TypeOf(executor).String(), err,
			)
		}

		ee[hash] = executor

		hashes[hash] = struct{}{}
	}

	return hashes, nil
}

func (ee Executors) Execute(ctx context.Context) ExecutionResults {
	if len(ee) == 0 {
		return ExecutionResults{}
	}

	results := make(ExecutionResults, len(ee))

	var wg sync.WaitGroup

	resultCh := make(chan struct {
		hash  string
		value Value
	}, len(ee))

	for hash, executor := range ee {
		wg.Add(1)

		go func(h string, e Executor) {
			defer wg.Done()
			resultCh <- struct {
				hash  string
				value Value
			}{h, e.Execute(ctx)}
		}(hash, executor)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for res := range resultCh {
		results[res.hash] = res.value
	}

	return results
}

type ExecutionResults map[string]Value

func (r ExecutionResults) Limit(hashes map[string]struct{}) ExecutionResults {
	limited := make(ExecutionResults, len(hashes))

	for hash := range hashes {
		if value, exists := r[hash]; exists {
			limited[hash] = value
		}
	}

	return limited
}
