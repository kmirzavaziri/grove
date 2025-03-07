package guardx

import (
	"github.com/kmirzavaziri/grove/go/internal/pkg/lua"
	"github.com/kmirzavaziri/grove/go/pkg/guard"
)

type Len struct {
	Min     *uint64
	Max     *uint64
	ErrType string
	ErrMin  string
	ErrMax  string
}

func (v Len) Script() string {
	return lua.GuardsLen
}

type All struct {
	Guard   guard.Guard
	ErrType string
	Err     string
}

func (v All) Script() string {
	return lua.GuardsAll
}

type Any struct {
	Guard   guard.Guard
	ErrType string
	Err     string
}

func (v Any) Script() string {
	return lua.GuardsAny
}

type RequiredKeys struct {
	Keys    []string
	ErrType string
	Err     string
}

func (v RequiredKeys) Script() string {
	return lua.GuardsRequiredKeys
}

type AllKeys struct {
	Pattern string
	ErrType string
	Err     string
}

func (v AllKeys) Script() string {
	return lua.GuardsAllKeys
}

type AnyKey struct {
	Pattern string
	ErrType string
	Err     string
}

func (v AnyKey) Script() string {
	return lua.GuardsAnyKey
}

type NestMap struct {
	Key     string
	Guard   guard.Guard
	ErrType string
	Err     string
}

func (v NestMap) Script() string {
	return lua.GuardsNestMap
}

type NestList struct {
	Index   uint64
	Guard   guard.Guard
	ErrType string
	Err     string
}

func (v NestList) Script() string {
	return lua.GuardsNestList
}

type UniqueItems struct {
	ErrType string
	Err     string
}

func (v UniqueItems) Script() string {
	return lua.GuardsUniqueItems
}
