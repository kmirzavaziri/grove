package lua

import _ "embed"

//go:embed error/render.lua
var ErrorRender string

//go:embed guards/type.lua
var GuardsType string

//go:embed guards/empty.lua
var GuardsEmpty string

//go:embed guards/not.lua
var GuardsNot string

//go:embed guards/and.lua
var GuardsAnd string

//go:embed guards/or.lua
var GuardsOr string

//go:embed guards/len.lua
var GuardsLen string

//go:embed guards/all.lua
var GuardsAll string

//go:embed guards/any.lua
var GuardsAny string

//go:embed guards/required_keys.lua
var GuardsRequiredKeys string

//go:embed guards/all_keys.lua
var GuardsAllKeys string

//go:embed guards/any_key.lua
var GuardsAnyKey string

//go:embed guards/nest_map.lua
var GuardsNestMap string

//go:embed guards/nest_list.lua
var GuardsNestList string

//go:embed guards/unique_items.lua
var GuardsUniqueItems string

//go:embed guards/range.lua
var GuardsRange string

//go:embed guards/integer.lua
var GuardsInteger string

//go:embed guards/match.lua
var GuardsMatch string

//go:embed guards/contains.lua
var GuardsContains string

//go:embed guards/has_prefix.lua
var GuardsHasPrefix string

//go:embed guards/has_suffix.lua
var GuardsHasSuffix string

//go:embed guards/enum.lua
var GuardsEnum string

//go:embed guards/email.lua
var GuardsEmail string

//go:embed guards/phone.lua
var GuardsPhone string

//go:embed guards/equal.lua
var GuardsEqual string
