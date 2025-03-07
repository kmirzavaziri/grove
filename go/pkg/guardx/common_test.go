package guardx_test

import (
	"github.com/kmirzavaziri/grove/go/pkg/guardx"
	"testing"

	"github.com/kmirzavaziri/grove/go/pkg/gex"
	"github.com/kmirzavaziri/grove/go/pkg/guard"
	"github.com/stretchr/testify/suite"
)

type CommonSuite struct {
	suite.Suite
}

func TestCommonSuite(t *testing.T) {
	suite.Run(t, new(CommonSuite))
}

func (s *CommonSuite) TestType() {
	for _, tCase := range []testCase{
		{
			t: "nil",
			d: gex.Nil(),
		},

		{
			t: "map - empty map",
			d: gex.Map(map[string]*gex.Value{}),
		},
		{
			t: "map - empty key and nil value",
			d: gex.Map(map[string]*gex.Value{"": gex.Nil()}),
		},
		{
			t: "map - nil value",
			d: gex.Map(map[string]*gex.Value{"value": gex.Nil()}),
		},
		{
			t: "map - empty key and empty string value",
			d: gex.Map(map[string]*gex.Value{"": gex.Str("")}),
		},
		{
			t: "map - empty string value",
			d: gex.Map(map[string]*gex.Value{"value": gex.Str("")}),
		},
		{
			t: "map - empty key",
			d: gex.Map(map[string]*gex.Value{"": gex.Str("test")}),
		},
		{
			t: "map - one pair",
			d: gex.Map(map[string]*gex.Value{"value": gex.Str("test")}),
		},

		{
			t: "list - empty list",
			d: gex.List(),
		},
		{
			t: "list - nil item",
			d: gex.List(gex.Nil()),
		},
		{
			t: "list - empty string item",
			d: gex.List(gex.Str("")),
		},
		{
			t: "list - non-empty string item",
			d: gex.List(gex.Str("test")),
		},

		{
			t: "number - zero",
			d: gex.Num(0),
		},
		{
			t: "number - int",
			d: gex.Num(float64(42)),
		},
		{
			t: "number - float",
			d: gex.Num(42.42),
		},

		{
			t: "string - empty",
			d: gex.Str(""),
		},
		{
			t: "string - test",
			d: gex.Str("test"),
		},

		{
			t: "bool - false",
			d: gex.Bool(false),
		},
		{
			t: "bool - true",
			d: gex.Bool(true),
		},
	} {
		for _, typeName := range gex.Types {
			errMsg, err := guard.New(guardx.Type{TypeName: typeName}).Guard(tCase.d)
			s.Require().NoError(err)

			if tCase.d.Type() == gex.TypeNil || tCase.d.Type() == typeName {
				s.Emptyf(errMsg, tCase.t)
			} else {
				s.NotEmptyf(errMsg, tCase.t)
			}
		}
	}
}

func (s *CommonSuite) TestNotEmpty() {
	testCases{
		{t: "nil", d: gex.Nil(), v: guardx.Empty{}, errMsg: ""},

		{t: "string - empty", d: gex.Str(""), v: guardx.Empty{}, errMsg: ""},
		{t: "string - not empty", d: gex.Str("test"), v: guardx.Empty{}, errMsg: "Guard failed"},

		{t: "list - empty", d: gex.List(), v: guardx.Empty{}, errMsg: ""},
		{
			t: "list - not empty",
			d: gex.List(gex.Str("item")), v: guardx.Empty{}, errMsg: "Guard failed",
		},

		{
			t: "map - empty",
			d: gex.Map(map[string]*gex.Value{}), v: guardx.Empty{}, errMsg: "",
		},
		{
			t:      "map - not empty",
			d:      gex.Map(map[string]*gex.Value{"key": gex.Str("value")}),
			v:      guardx.Empty{},
			errMsg: "Guard failed",
		},

		{t: "number - zero", d: gex.Num(0), v: guardx.Empty{}, errMsg: "Guard failed"},
		{t: "number - negative", d: gex.Num(-42), v: guardx.Empty{}, errMsg: "Guard failed"},
		{t: "number - positive", d: gex.Num(42), v: guardx.Empty{}, errMsg: "Guard failed"},

		{t: "boolean - false", d: gex.Bool(false), v: guardx.Empty{}, errMsg: "Guard failed"},
		{t: "boolean - true", d: gex.Bool(true), v: guardx.Empty{}, errMsg: "Guard failed"},
	}.run(s.Suite)
}

func (s *CommonSuite) TestErrorMessage() {
	testCases{
		{
			t: "Type - Err",
			d: gex.Str("text"),
			v: guardx.Type{
				TypeName: gex.TypeNumber,
				Err:      "M - ${value} - ${config.type_name}",
			},
			errMsg: "M - text - number",
		},
		{
			t: "Empty - Err",
			d: gex.Str("text"),
			v: guardx.Empty{
				Err: "M - ${value}",
			},
			errMsg: "M - text",
		},
		{
			t: "Not - Err",
			d: gex.Str("text"),
			v: guardx.Not{
				Guard: guard.New(guardx.Type{TypeName: gex.TypeString}),
				Err:   "M - ${value}",
			},
			errMsg: "M - text",
		},
		{
			t: "And - Err (nested_err)",
			d: gex.Str("text"),
			v: guardx.And{
				Guards: []guard.Guard{
					guard.New(guardx.HasPrefix{Prefix: "t"}),
					guard.New(guardx.HasSuffix{
						Suffix: "f",
						Err:    "NM - ${value} - ${config.suffix}",
					}),
				},
				Err: "M - ${value} - ${nested_err}",
			},
			errMsg: "M - text - NM - text - f",
		},
		{
			t: "Or - Err",
			d: gex.Str("text"),
			v: guardx.Or{
				Guards: []guard.Guard{
					guard.New(guardx.HasPrefix{Prefix: "a"}),
					guard.New(guardx.HasSuffix{Suffix: "z"}),
				},
				Err: "M - ${value}",
			},
			errMsg: "M - text",
		},
		{
			t: "Len - ErrType",
			d: gex.Num(42),
			v: guardx.Len{
				Min:     gex.UintP(2),
				Max:     gex.UintP(4),
				ErrType: "M - ${value} - ${config.min} - ${config.max}",
			},
			errMsg: "M - 42 - 2 - 4",
		},
		{
			t: "Len - ErrMin",
			d: gex.Str("t"),
			v: guardx.Len{
				Min:    gex.UintP(2),
				Max:    gex.UintP(4),
				ErrMin: "M - ${value} - ${config.min} - ${config.max}",
			},
			errMsg: "M - t - 2 - 4",
		},
		{
			t: "Len - ErrMax",
			d: gex.Str("long text"),
			v: guardx.Len{
				Min:    gex.UintP(2),
				Max:    gex.UintP(4),
				ErrMax: "M - ${value} - ${config.min} - ${config.max}",
			},
			errMsg: "M - long text - 2 - 4",
		},
		{
			t: "All - ErrType",
			d: gex.Str("text"),
			v: guardx.All{
				Guard:   guard.New(guardx.HasPrefix{Prefix: "a"}),
				ErrType: "M - ${value}",
			},
			errMsg: "M - text",
		},
		{
			t: "All - Err (nested_err)",
			d: gex.List(gex.Str("text 1"), gex.Str("text 2")),
			v: guardx.All{
				Guard: guard.New(guardx.HasPrefix{
					Prefix: "a",
					Err:    "NM - ${value} - ${config.prefix}",
				}),
				Err: "M - ${value} - ${nested_err}",
			},
			errMsg: "M - text 1, text 2 - NM - text 1 - a",
		},
		{
			t: "Any - ErrType",
			d: gex.Str("text"),
			v: guardx.Any{
				Guard:   guard.New(guardx.HasPrefix{Prefix: "a"}),
				ErrType: "M - ${value}",
			},
			errMsg: "M - text",
		},
		{
			t: "Any - Err",
			d: gex.List(gex.Str("text 1"), gex.Str("text 2")),
			v: guardx.Any{
				Guard: guard.New(guardx.HasPrefix{Prefix: "a"}),
				Err:   "M - ${value}",
			},
			errMsg: "M - text 1, text 2",
		},
		{
			t: "RequiredKeys - ErrType",
			d: gex.Str("text"),
			v: guardx.RequiredKeys{
				Keys:    []string{"key1", "key2"},
				ErrType: "M - ${value} - ${config.keys}",
			},
			errMsg: "M - text - key1, key2",
		},
		{
			t: "RequiredKeys - Err (key)",
			d: gex.Map(map[string]*gex.Value{"key1": gex.Str("text")}),
			v: guardx.RequiredKeys{
				Keys: []string{"key1", "key2"},
				Err:  "M - ${value.key1} - ${config.keys} - ${key}",
			},
			errMsg: "M - text - key1, key2 - key2",
		},
		{
			t: "AllKeys - ErrType",
			d: gex.Str("text"),
			v: guardx.AllKeys{
				Pattern: "pattern",
				ErrType: "M - ${value} - ${config.pattern}",
			},
			errMsg: "M - text - pattern",
		},
		{
			t: "AllKeys - Err",
			d: gex.Map(map[string]*gex.Value{"key": gex.Str("text")}),
			v: guardx.AllKeys{
				Pattern: "pattern",
				Err:     "M - ${value.key} - ${config.pattern}",
			},
			errMsg: "M - text - pattern",
		},
		{
			t: "AnyKey - ErrType",
			d: gex.Str("text"),
			v: guardx.AnyKey{
				Pattern: "pattern",
				ErrType: "M - ${value} - ${config.pattern}",
			},
			errMsg: "M - text - pattern",
		},
		{
			t: "AnyKey - Err",
			d: gex.Map(map[string]*gex.Value{"key": gex.Str("text")}),
			v: guardx.AnyKey{
				Pattern: "pattern",
				Err:     "M - ${value.key} - ${config.pattern}",
			},
			errMsg: "M - text - pattern",
		},
		//{
		//	t: "NestMap - ErrType",
		//	d: ripd.String("text"),
		//	v: guards.NestMap{
		//		Key:        "",
		//		Guard: "",
		//		ErrType:    "M - ${value} - ${config.key}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "NestMap - Err (nested_err)",
		//	d: ripd.String("text"),
		//	v: guards.NestMap{
		//		Key:        "",
		//		Guard: "",
		//		Err:        "M - ${value} - ${config.key} - ${nested_err}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "NestList - ErrType",
		//	d: ripd.String("text"),
		//	v: guards.NestList{
		//		Index:      0,
		//		Guard: "",
		//		ErrType:    "M - ${value} - ${config.index}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "NestList - Err (nested_err)",
		//	d: ripd.String("text"),
		//	v: guards.NestList{
		//		Index:      0,
		//		Guard: "",
		//		Err:        "M - ${value} - ${config.index} - ${nested_err}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "UniqueItems - ErrType",
		//	d: ripd.String("text"),
		//	v: guards.UniqueItems{
		//		ErrType: "M - ${value}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "UniqueItems - Err (item)",
		//	d: ripd.String("text"),
		//	v: guards.UniqueItems{
		//		Err: "M - ${value} - ${item}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "Range - ErrType",
		//	d: ripd.String("text"),
		//	v: guards.Range{
		//		Min:     nil,
		//		Max:     nil,
		//		ErrType: "M - ${value} - ${config.min} - ${config.max}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "Range - ErrMin",
		//	d: ripd.String("text"),
		//	v: guards.Range{
		//		Min:    nil,
		//		Max:    nil,
		//		ErrMin: "M - ${value} - ${config.min} - ${config.max}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "Range - ErrMax",
		//	d: ripd.String("text"),
		//	v: guards.Range{
		//		Min:    nil,
		//		Max:    nil,
		//		ErrMax: "M - ${value} - ${config.min} - ${config.max}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "Integer - ErrType",
		//	d: ripd.String("text"),
		//	v: guards.Integer{
		//		ErrType: "M - ${value}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "Integer - Err",
		//	d: ripd.String("text"),
		//	v: guards.Integer{
		//		Err: "M - ${value}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "Match - ErrType",
		//	d: ripd.String("text"),
		//	v: guards.Match{
		//		Pattern: "",
		//		ErrType: "M - ${value} - ${config.pattern}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "Match - Err",
		//	d: ripd.String("text"),
		//	v: guards.Match{
		//		Pattern: "",
		//		Err:     "M - ${value} - ${config.pattern}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "Contains - ErrType",
		//	d: ripd.String("text"),
		//	v: guards.Contains{
		//		Substring: "",
		//		ErrType:   "M - ${value} - ${config.substring}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "Contains - Err",
		//	d: ripd.String("text"),
		//	v: guards.Contains{
		//		Substring: "",
		//		Err:       "M - ${value} - ${config.substring}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "HasPrefix - ErrType",
		//	d: ripd.String("text"),
		//	v: guards.HasPrefix{
		//		Prefix:  "",
		//		ErrType: "M - ${value} - ${config.prefix}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "HasPrefix - Err",
		//	d: ripd.String("text"),
		//	v: guards.HasPrefix{
		//		Prefix: "",
		//		Err:    "M - ${value} - ${config.prefix}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "HasSuffix - ErrType",
		//	d: ripd.String("text"),
		//	v: guards.HasSuffix{
		//		Suffix:  "",
		//		ErrType: "M - ${value} - ${config.suffix}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "HasSuffix - Err",
		//	d: ripd.String("text"),
		//	v: guards.HasSuffix{
		//		Suffix: "",
		//		Err:    "M - ${value} - ${config.suffix}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "Enum - ErrType",
		//	d: ripd.String("text"),
		//	v: guards.Enum{
		//		Values:  nil,
		//		ErrType: "M - ${value} - ${config.values}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "Enum - Err",
		//	d: ripd.String("text"),
		//	v: guards.Enum{
		//		Values: nil,
		//		Err:    "M - ${value} - ${config.values}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "Email - ErrType",
		//	d: ripd.String("text"),
		//	v: guards.Email{
		//		ErrType: "M - ${value}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "Email - Err",
		//	d: ripd.String("text"),
		//	v: guards.Email{
		//		Err: "M - ${value}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "Phone - ErrType",
		//	d: ripd.String("text"),
		//	v: guards.Phone{
		//		ErrType: "M - ${value}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "Phone - Err",
		//	d: ripd.String("text"),
		//	v: guards.Phone{
		//		Err: "M - ${value}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "Equal - ErrType",
		//	d: ripd.String("text"),
		//	v: guards.Equal{
		//		Value:   false,
		//		ErrType: "M - ${value} - ${config.value}",
		//	},
		//	errMessage: "",
		//},
		//{
		//	t: "Equal - Err",
		//	d: ripd.String("text"),
		//	v: guards.Equal{
		//		Value: false,
		//		Err:   "M - ${value} - ${config.value}",
		//	},
		//	errMessage: "",
		//},
	}.run(s.Suite)
}
