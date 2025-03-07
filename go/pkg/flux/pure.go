package flux

import (
	"github.com/kmirzavaziri/grove/go/pkg/greq"
)

type Pure[T any] interface {
	Data(request greq.Request) T
}

type pureInline[T any] struct {
	data func(request greq.Request) T
}

func PureInline[T any](data func(request greq.Request) T) Pure[T] {
	return &pureInline[T]{
		data: data,
	}
}

func (f *pureInline[T]) Data(request greq.Request) T {
	return f.data(request)
}
