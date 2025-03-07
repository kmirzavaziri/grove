package guardx_test

import (
	"github.com/kmirzavaziri/grove/go/pkg/guardx"
	"testing"

	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/stretchr/testify/suite"
)

type StringsSuite struct {
	suite.Suite
}

func TestStringsSuite(t *testing.T) {
	suite.Run(t, new(StringsSuite))
}

func (s *StringsSuite) TestMatch() {
	testCases{
		{t: "matching value", d: gex.Str("test"), v: guardx.Match{Pattern: "^test$"}, errMsg: ""},
		{t: "non-matching value", d: gex.Str("fail"), v: guardx.Match{Pattern: "^test$"}, errMsg: "Guard failed"},
	}.run(s.Suite)
}

func (s *StringsSuite) TestContains() {
	testCases{
		{t: "contains", d: gex.Str("hello world"), v: guardx.Contains{Substring: "world"}, errMsg: ""},
		{t: "does not contain", d: gex.Str("hello"), v: guardx.Contains{Substring: "world"}, errMsg: "Guard failed"},
	}.run(s.Suite)
}

func (s *StringsSuite) TestHasPrefix() {
	testCases{
		{t: "prefix", d: gex.Str("hello"), v: guardx.HasPrefix{Prefix: "he"}, errMsg: ""},
		{t: "no prefix", d: gex.Str("hello"), v: guardx.HasPrefix{Prefix: "x"}, errMsg: "Guard failed"},
	}.run(s.Suite)
}

func (s *StringsSuite) TestHasSuffix() {
	testCases{
		{t: "suffix", d: gex.Str("hello"), v: guardx.HasSuffix{Suffix: "lo"}, errMsg: ""},
		{t: "no suffix", d: gex.Str("hello"), v: guardx.HasSuffix{Suffix: "x"}, errMsg: "Guard failed"},
	}.run(s.Suite)
}

func (s *StringsSuite) TestEnum() {
	testCases{
		{t: "valid enum v1", d: gex.Str("v1"), v: guardx.Enum{Values: []string{"v1", "v2"}}, errMsg: ""},
		{t: "valid enum v2", d: gex.Str("v2"), v: guardx.Enum{Values: []string{"v1", "v2"}}, errMsg: ""},
		{t: "invalid enum", d: gex.Str("v3"), v: guardx.Enum{Values: []string{"v1", "v2"}}, errMsg: "Guard failed"},
	}.run(s.Suite)
}

func (s *StringsSuite) TestEmail() {
	testCases{
		{t: "valid email", d: gex.Str("test@example.com"), v: guardx.Email{}, errMsg: ""},
		{t: "invalid email", d: gex.Str("test@"), v: guardx.Email{}, errMsg: "Guard failed"},
	}.run(s.Suite)
}

func (s *StringsSuite) TestPhone() {
	testCases{
		{t: "valid phone", d: gex.Str("971541234567"), v: guardx.Phone{}, errMsg: ""},
		{t: "invalid phone", d: gex.Str("123"), v: guardx.Phone{}, errMsg: "Guard failed"},
	}.run(s.Suite)
}
