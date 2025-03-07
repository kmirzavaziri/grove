package executor

import (
	"context"
	"github.com/kmirzavaziri/grove/go/pkg/gex"
)

type Hash string

type Executor interface {
	Hash() (Hash, error)
	Execute(ctx context.Context) *gex.Value
}

type Results map[Hash]*gex.Value
