package guards_test

import (
	"testing"

	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/guards"
	"github.com/stretchr/testify/suite"
)

type CollectionsSuite struct {
	suite.Suite
}

func TestCollectionsSuite(t *testing.T) {
	suite.Run(t, new(CollectionsSuite))
}

func (s *CollectionsSuite) TestLen() {
	testCases{
		{
			t:      "valid string",
			d:      gex.Str("test"),
			v:      guards.Len{Min: gex.UintP(1), Max: gex.UintP(10)},
			errMsg: "",
		},
		{
			t:      "valid list",
			d:      gex.List(gex.Str("item1"), gex.Str("item2")),
			v:      guards.Len{Min: gex.UintP(1), Max: gex.UintP(3)},
			errMsg: "",
		},
		{
			t:      "too short string",
			d:      gex.Str(""),
			v:      guards.Len{Min: gex.UintP(1), Max: gex.UintP(10)},
			errMsg: "Guard failed",
		},
		{
			t:      "too long string",
			d:      gex.Str("this is too long"),
			v:      guards.Len{Min: gex.UintP(1), Max: gex.UintP(10)},
			errMsg: "Guard failed",
		},
		{
			t:      "too short list",
			d:      gex.List(),
			v:      guards.Len{Min: gex.UintP(1), Max: gex.UintP(2)},
			errMsg: "Guard failed",
		},
		{
			t:      "too long list",
			d:      gex.List(gex.Str("1"), gex.Str("2"), gex.Str("3")),
			v:      guards.Len{Min: gex.UintP(1), Max: gex.UintP(2)},
			errMsg: "Guard failed",
		},
		{
			t:      "optional max",
			d:      gex.Str("123"),
			v:      guards.Len{Min: gex.UintP(1)},
			errMsg: "",
		},
	}.run(s.Suite)
}

func (s *CollectionsSuite) TestAll() {
	testCases{
		{
			t: "all valid values",
			d: gex.List(gex.Num(1), gex.Num(2)),
			v: guards.All{
				Guard: guards.New(guards.Range{Min: gex.FloatP(1), Max: gex.FloatP(3)}),
			},
			errMsg: "",
		},
		{
			t: "one invalid value",
			d: gex.List(gex.Num(1), gex.Num(4)),
			v: guards.All{
				Guard: guards.New(guards.Range{Min: gex.FloatP(1), Max: gex.FloatP(3)}),
			},
			errMsg: "Guard failed",
		},
	}.run(s.Suite)
}

func (s *CollectionsSuite) TestAny() {
	testCases{
		{
			t: "all invalid values",
			d: gex.List(gex.Num(0), gex.Num(4)),
			v: guards.Any{
				Guard: guards.New(guards.Range{Min: gex.FloatP(1), Max: gex.FloatP(3)}),
			},
			errMsg: "Guard failed",
		},
		{
			t: "one valid value",
			d: gex.List(gex.Num(1), gex.Num(4)),
			v: guards.Any{
				Guard: guards.New(guards.Range{Min: gex.FloatP(1), Max: gex.FloatP(3)}),
			},
			errMsg: "",
		},
	}.run(s.Suite)
}

func (s *CollectionsSuite) TestRequiredKeys() {
	testCases{
		{
			t: "all keys present",
			d: gex.Map(map[string]*gex.Value{"key1": gex.Str("value")}),
			v: guards.RequiredKeys{
				Keys: []string{"key1"},
			},
			errMsg: "",
		},
		{
			t: "missing key",
			d: gex.Map(map[string]*gex.Value{"key1": gex.Str("value")}),
			v: guards.RequiredKeys{
				Keys: []string{"key2"},
			},
			errMsg: "Guard failed",
		},
	}.run(s.Suite)
}

func (s *CollectionsSuite) TestAllKeys() {
	testCases{
		{
			t: "all valid keys",
			d: gex.Map(map[string]*gex.Value{"key1": gex.Str("value"), "key2": gex.Str("value")}),
			v: guards.AllKeys{
				Pattern: "key[0-9]",
			},
			errMsg: "",
		},
		{
			t: "one invalid key",
			d: gex.Map(map[string]*gex.Value{"key1": gex.Str("value"), "key": gex.Str("value")}),
			v: guards.AllKeys{
				Pattern: "key[0-9]",
			},
			errMsg: "Guard failed",
		},
	}.run(s.Suite)
}

func (s *CollectionsSuite) TestAnyKey() {
	testCases{
		{
			t: "all invalid keys",
			d: gex.Map(map[string]*gex.Value{"key": gex.Str("value"), "key_": gex.Str("value")}),
			v: guards.AnyKey{
				Pattern: "key[0-9]",
			},
			errMsg: "Guard failed",
		},
		{
			t: "one valid key",
			d: gex.Map(map[string]*gex.Value{"key": gex.Str("value"), "key1": gex.Str("value")}),
			v: guards.AnyKey{
				Pattern: "key[0-9]",
			},
			errMsg: "",
		},
	}.run(s.Suite)
}

func (s *CollectionsSuite) TestNestMap() {
	testCases{
		{
			t: "valid nested value",
			d: gex.Map(map[string]*gex.Value{"key": gex.Num(1)}),
			v: guards.NestMap{
				Key:   "key",
				Guard: guards.New(guards.Range{Min: gex.FloatP(0), Max: gex.FloatP(2)}),
			},
			errMsg: "",
		},
		{
			t: "invalid nested value",
			d: gex.Map(map[string]*gex.Value{"key": gex.Num(3)}),
			v: guards.NestMap{
				Key:   "key",
				Guard: guards.New(guards.Range{Min: gex.FloatP(0), Max: gex.FloatP(2)}),
			},
			errMsg: "Guard failed",
		},
	}.run(s.Suite)
}

func (s *CollectionsSuite) TestNestList() {
	testCases{
		{
			t: "valid nested value",
			d: gex.List(gex.Num(1), gex.Num(2), gex.Num(3)),
			v: guards.NestList{
				Index: 0,
				Guard: guards.New(guards.Range{Min: gex.FloatP(0), Max: gex.FloatP(2)}),
			},
			errMsg: "",
		},
		{
			t: "invalid nested value",
			d: gex.List(gex.Num(1), gex.Num(2), gex.Num(3)),
			v: guards.NestList{
				Index: 2,
				Guard: guards.New(guards.Range{Min: gex.FloatP(0), Max: gex.FloatP(2)}),
			},
			errMsg: "Guard failed",
		},
	}.run(s.Suite)
}

func (s *CollectionsSuite) TestUniqueItems() {
	testCases{
		{
			t:      "unique items",
			d:      gex.List(gex.Str("item1"), gex.Str("item2")),
			v:      guards.UniqueItems{},
			errMsg: "",
		},
		{
			t:      "duplicate items",
			d:      gex.List(gex.Str("item1"), gex.Str("item1")),
			v:      guards.UniqueItems{},
			errMsg: "Guard failed",
		},
	}.run(s.Suite)
}
