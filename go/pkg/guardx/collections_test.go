package guardx_test

import (
	"github.com/kmirzavaziri/grove/go/pkg/guardx"
	"testing"

	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/guard"
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
			v:      guardx.Len{Min: gex.UintP(1), Max: gex.UintP(10)},
			errMsg: "",
		},
		{
			t:      "valid list",
			d:      gex.List(gex.Str("item1"), gex.Str("item2")),
			v:      guardx.Len{Min: gex.UintP(1), Max: gex.UintP(3)},
			errMsg: "",
		},
		{
			t:      "too short string",
			d:      gex.Str(""),
			v:      guardx.Len{Min: gex.UintP(1), Max: gex.UintP(10)},
			errMsg: "Guard failed",
		},
		{
			t:      "too long string",
			d:      gex.Str("this is too long"),
			v:      guardx.Len{Min: gex.UintP(1), Max: gex.UintP(10)},
			errMsg: "Guard failed",
		},
		{
			t:      "too short list",
			d:      gex.List(),
			v:      guardx.Len{Min: gex.UintP(1), Max: gex.UintP(2)},
			errMsg: "Guard failed",
		},
		{
			t:      "too long list",
			d:      gex.List(gex.Str("1"), gex.Str("2"), gex.Str("3")),
			v:      guardx.Len{Min: gex.UintP(1), Max: gex.UintP(2)},
			errMsg: "Guard failed",
		},
		{
			t:      "optional max",
			d:      gex.Str("123"),
			v:      guardx.Len{Min: gex.UintP(1)},
			errMsg: "",
		},
	}.run(s.Suite)
}

func (s *CollectionsSuite) TestAll() {
	testCases{
		{
			t: "all valid values",
			d: gex.List(gex.Num(1), gex.Num(2)),
			v: guardx.All{
				Guard: guard.New(guardx.Range{Min: gex.FloatP(1), Max: gex.FloatP(3)}),
			},
			errMsg: "",
		},
		{
			t: "one invalid value",
			d: gex.List(gex.Num(1), gex.Num(4)),
			v: guardx.All{
				Guard: guard.New(guardx.Range{Min: gex.FloatP(1), Max: gex.FloatP(3)}),
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
			v: guardx.Any{
				Guard: guard.New(guardx.Range{Min: gex.FloatP(1), Max: gex.FloatP(3)}),
			},
			errMsg: "Guard failed",
		},
		{
			t: "one valid value",
			d: gex.List(gex.Num(1), gex.Num(4)),
			v: guardx.Any{
				Guard: guard.New(guardx.Range{Min: gex.FloatP(1), Max: gex.FloatP(3)}),
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
			v: guardx.RequiredKeys{
				Keys: []string{"key1"},
			},
			errMsg: "",
		},
		{
			t: "missing key",
			d: gex.Map(map[string]*gex.Value{"key1": gex.Str("value")}),
			v: guardx.RequiredKeys{
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
			v: guardx.AllKeys{
				Pattern: "key[0-9]",
			},
			errMsg: "",
		},
		{
			t: "one invalid key",
			d: gex.Map(map[string]*gex.Value{"key1": gex.Str("value"), "key": gex.Str("value")}),
			v: guardx.AllKeys{
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
			v: guardx.AnyKey{
				Pattern: "key[0-9]",
			},
			errMsg: "Guard failed",
		},
		{
			t: "one valid key",
			d: gex.Map(map[string]*gex.Value{"key": gex.Str("value"), "key1": gex.Str("value")}),
			v: guardx.AnyKey{
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
			v: guardx.NestMap{
				Key:   "key",
				Guard: guard.New(guardx.Range{Min: gex.FloatP(0), Max: gex.FloatP(2)}),
			},
			errMsg: "",
		},
		{
			t: "invalid nested value",
			d: gex.Map(map[string]*gex.Value{"key": gex.Num(3)}),
			v: guardx.NestMap{
				Key:   "key",
				Guard: guard.New(guardx.Range{Min: gex.FloatP(0), Max: gex.FloatP(2)}),
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
			v: guardx.NestList{
				Index: 0,
				Guard: guard.New(guardx.Range{Min: gex.FloatP(0), Max: gex.FloatP(2)}),
			},
			errMsg: "",
		},
		{
			t: "invalid nested value",
			d: gex.List(gex.Num(1), gex.Num(2), gex.Num(3)),
			v: guardx.NestList{
				Index: 2,
				Guard: guard.New(guardx.Range{Min: gex.FloatP(0), Max: gex.FloatP(2)}),
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
			v:      guardx.UniqueItems{},
			errMsg: "",
		},
		{
			t:      "duplicate items",
			d:      gex.List(gex.Str("item1"), gex.Str("item1")),
			v:      guardx.UniqueItems{},
			errMsg: "Guard failed",
		},
	}.run(s.Suite)
}
