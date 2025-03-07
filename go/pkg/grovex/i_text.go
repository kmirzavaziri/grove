package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
)

type ITextType string

const (
	ITextTypeText     ITextType = "TEXT"
	ITextTypeTextarea ITextType = "TEXTAREA"
	ITextTypePassword ITextType = "PASSWORD"
	ITextTypeEmail    ITextType = "EMAIL"
	ITextTypeInt      ITextType = "INT"
	ITextTypeFloat    ITextType = "FLOAT"
)

// TODO add per request validation by IText by adding middleware to Props so that we can verify user defined flux
//  results, and raise error if the value is not in this list.

var ITextTypeOptions = map[ITextType]struct{}{
	ITextTypeText:     {},
	ITextTypeTextarea: {},
	ITextTypePassword: {},
	ITextTypeEmail:    {},
	ITextTypeInt:      {},
	ITextTypeFloat:    {},
}

type ITextArgs struct {
	Key   string
	Props flux.Read[ITextProps]
	Input *grove.Input
}

type ITextProps struct {
	Label       string
	Hint        string
	Placeholder string
	Type        ITextType
	Prefix      string
	Suffix      string
}

func IText(args ITextArgs) *grove.Node {
	if args.Input == nil {
		return grove.Error(args.Key, "input cannot be nil")
	}

	return &grove.Node{
		Type:  "grovex.IText",
		Key:   args.Key,
		Props: flux.ReadValue(args.Props),
		Input: args.Input,
	}
}
