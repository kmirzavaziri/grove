package cases

import (
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	return strings.ToLower(
		strings.ReplaceAll(
			matchAllCap.ReplaceAllString(
				matchFirstCap.ReplaceAllString(str, "${1}_${2}"),
				"${1}_${2}",
			),
			" ", "_",
		),
	)
}
