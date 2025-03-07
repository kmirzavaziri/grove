package guardx_test

import (
	"github.com/kmirzavaziri/grove/go/pkg/guardx"
	"testing"

	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/stretchr/testify/suite"
)

type BoolsSuite struct {
	suite.Suite
}

func TestBoolsSuite(t *testing.T) {
	suite.Run(t, new(BoolsSuite))
}

func (s *BoolsSuite) TestEqual() {
	testCases{
		{t: "expect T - actual T", d: gex.Bool(true), v: guardx.Equal{Value: true}, errMsg: ""},
		{t: "expect T - actual F", d: gex.Bool(false), v: guardx.Equal{Value: true}, errMsg: "Guard failed"},
		{t: "expect F - actual T", d: gex.Bool(true), v: guardx.Equal{Value: false}, errMsg: "Guard failed"},
		{t: "expect F - actual F", d: gex.Bool(false), v: guardx.Equal{Value: false}, errMsg: ""},
	}.run(s.Suite)
}
