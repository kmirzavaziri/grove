package guardx_test

import (
	"github.com/kmirzavaziri/grove/go/pkg/guardx"
	"testing"

	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/stretchr/testify/suite"
)

type NumbersSuite struct {
	suite.Suite
}

func TestNumbersSuite(t *testing.T) {
	suite.Run(t, new(NumbersSuite))
}

func (s *NumbersSuite) TestRange() {
	testCases{
		{t: "within", d: gex.Num(5), v: guardx.Range{Min: gex.FloatP(0), Max: gex.FloatP(10)}, errMsg: ""},
		{t: "incl - l", d: gex.Num(0), v: guardx.Range{Min: gex.FloatP(0), Max: gex.FloatP(10)}, errMsg: ""},
		{t: "incl - r", d: gex.Num(10), v: guardx.Range{Min: gex.FloatP(0), Max: gex.FloatP(10)}, errMsg: ""},
		{t: "below", d: gex.Num(-1), v: guardx.Range{Min: gex.FloatP(0), Max: gex.FloatP(10)}, errMsg: "Guard failed"},
		{t: "above", d: gex.Num(11), v: guardx.Range{Min: gex.FloatP(0), Max: gex.FloatP(10)}, errMsg: "Guard failed"},
		{t: "open l", d: gex.Num(-1000), v: guardx.Range{Max: gex.FloatP(10)}, errMsg: ""},
		{t: "open r", d: gex.Num(1000), v: guardx.Range{Min: gex.FloatP(0)}, errMsg: ""},
		{t: "open l - above", d: gex.Num(11), v: guardx.Range{Max: gex.FloatP(10)}, errMsg: "Guard failed"},
		{t: "open r - below", d: gex.Num(-1), v: guardx.Range{Min: gex.FloatP(0)}, errMsg: "Guard failed"},
	}.run(s.Suite)
}

func (s *NumbersSuite) TestInteger() {
	testCases{
		{t: "integer", d: gex.Num(5), v: guardx.Integer{}, errMsg: ""},
		{t: "float", d: gex.Num(5.5), v: guardx.Integer{}, errMsg: "Guard failed"},
	}.run(s.Suite)
}
