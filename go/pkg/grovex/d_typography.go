package grovex

import (
	"github.com/kmirzavaziri/grove/go/pkg/executor"
	"github.com/kmirzavaziri/grove/go/pkg/flux"
	"github.com/kmirzavaziri/grove/go/pkg/greq"
	"github.com/kmirzavaziri/grove/go/pkg/grove"
	"github.com/kmirzavaziri/grove/go/pkg/grovex/variants"
)

type DTypographyArgs struct {
	Key   string
	Props flux.Read[DTypographyProps]
}

type DTypographyProps struct {
	Text    string
	Align   *variants.Align
	Color   *variants.Color
	Variant *variants.DTypographyVariant
}

func DTypography(args DTypographyArgs) *grove.Node {
	return &grove.Node{
		Type: "grovex.DTypography",
		Key:  args.Key,
		Props: flux.ReadValue(flux.ReadInline(
			func(request greq.Request) []executor.Executor {
				return args.Props.PreExecutors(request)
			},
			func(request greq.Request, executorResults executor.Results) (DTypographyProps, error) {
				props, err := args.Props.Data(request, executorResults)
				if err != nil {
					return DTypographyProps{}, err
				}

				if props.Align, err = props.Align.Visit(); err != nil {
					return DTypographyProps{}, err
				}

				if props.Color, err = props.Color.Visit(); err != nil {
					return DTypographyProps{}, err
				}

				if props.Variant, err = props.Variant.Visit(); err != nil {
					return DTypographyProps{}, err
				}

				return props, nil
			},
		)),
	}
}
