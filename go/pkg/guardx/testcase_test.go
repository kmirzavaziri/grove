package guardx_test

import (
	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/guard"
	"github.com/stretchr/testify/suite"
)

type testCase struct {
	t      string
	d      *gex.Value
	v      guard.Config
	errMsg string
}

type testCases []testCase

func (cc testCases) run(s suite.Suite) {
	for _, c := range cc {
		errMsg, err := guard.New(c.v).Guard(c.d)
		s.Require().NoError(err)
		s.Equal(c.errMsg, errMsg, c.t)
	}
}
