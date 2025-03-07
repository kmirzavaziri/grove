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
