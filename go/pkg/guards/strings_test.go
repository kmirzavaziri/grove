package guards_test

import (
	"testing"

	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/guards"
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
		{t: "matching value", d: gex.Str("test"), v: guards.Match{Pattern: "^test$"}, errMsg: ""},
		{t: "non-matching value", d: gex.Str("fail"), v: guards.Match{Pattern: "^test$"}, errMsg: "Guard failed"},
	}.run(s.Suite)
}

func (s *StringsSuite) TestContains() {
	testCases{
		{t: "contains", d: gex.Str("hello world"), v: guards.Contains{Substring: "world"}, errMsg: ""},
		{t: "does not contain", d: gex.Str("hello"), v: guards.Contains{Substring: "world"}, errMsg: "Guard failed"},
	}.run(s.Suite)
}

func (s *StringsSuite) TestHasPrefix() {
	testCases{
		{t: "prefix", d: gex.Str("hello"), v: guards.HasPrefix{Prefix: "he"}, errMsg: ""},
		{t: "no prefix", d: gex.Str("hello"), v: guards.HasPrefix{Prefix: "x"}, errMsg: "Guard failed"},
	}.run(s.Suite)
}

func (s *StringsSuite) TestHasSuffix() {
	testCases{
		{t: "suffix", d: gex.Str("hello"), v: guards.HasSuffix{Suffix: "lo"}, errMsg: ""},
		{t: "no suffix", d: gex.Str("hello"), v: guards.HasSuffix{Suffix: "x"}, errMsg: "Guard failed"},
	}.run(s.Suite)
}

func (s *StringsSuite) TestEnum() {
	testCases{
		{t: "valid enum v1", d: gex.Str("v1"), v: guards.Enum{Values: []string{"v1", "v2"}}, errMsg: ""},
		{t: "valid enum v2", d: gex.Str("v2"), v: guards.Enum{Values: []string{"v1", "v2"}}, errMsg: ""},
		{t: "invalid enum", d: gex.Str("v3"), v: guards.Enum{Values: []string{"v1", "v2"}}, errMsg: "Guard failed"},
	}.run(s.Suite)
}

func (s *StringsSuite) TestEmail() {
	testCases{
		{t: "valid email", d: gex.Str("test@example.com"), v: guards.Email{}, errMsg: ""},
		{t: "invalid email", d: gex.Str("test@"), v: guards.Email{}, errMsg: "Guard failed"},
	}.run(s.Suite)
}

func (s *StringsSuite) TestPhone() {
	testCases{
		{t: "valid phone", d: gex.Str("971541234567"), v: guards.Phone{}, errMsg: ""},
		{t: "invalid phone", d: gex.Str("123"), v: guards.Phone{}, errMsg: "Guard failed"},
	}.run(s.Suite)
}
