package guards_test

import (
	"testing"

	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/guards"
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
		{t: "expect T - actual T", d: gex.Bool(true), v: guards.Equal{Value: true}, errMsg: ""},
		{t: "expect T - actual F", d: gex.Bool(false), v: guards.Equal{Value: true}, errMsg: "Guard failed"},
		{t: "expect F - actual T", d: gex.Bool(true), v: guards.Equal{Value: false}, errMsg: "Guard failed"},
		{t: "expect F - actual F", d: gex.Bool(false), v: guards.Equal{Value: false}, errMsg: ""},
	}.run(s.Suite)
}
