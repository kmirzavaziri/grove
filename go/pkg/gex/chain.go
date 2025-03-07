package gex

import (
	"github.com/kmirzavaziri/grove/go/pkg/grr"
)

type Chain struct {
	v   *Value
	err error
}

func (c Chain) Get(key string) Chain {
	if c.err != nil {
		return Chain{err: c.err}
	}

	mv, ok := c.v.Map()
	if !ok {
		return Chain{err: grr.Wrap(grr.ErrInvalidValueType, "expected map to get key %q, got %s", key, c.v.Type())}
	}

	v, ok := mv[key]
	if !ok {
		return Chain{err: grr.Wrap(grr.ErrKeyNotFound, "key %q not found in map", key)}
	}

	return Chain{v: v}
}

func (c Chain) At(index int) Chain {
	if c.err != nil {
		return Chain{err: c.err}
	}

	listValue, ok := c.v.List()
	if !ok {
		return Chain{err: grr.Wrap(grr.ErrInvalidValueType, "expected list to get at %d, got %s", index, c.v.Type())}
	}

	if index < 0 || index > len(listValue) {
		return Chain{err: grr.Wrap(grr.ErrIndexOutOfRange, "index %d out of range in list", index)}
	}

	return Chain{v: listValue[index]}
}

func (c Chain) Number() (float64, error) {
	if c.err != nil {
		return 0, c.err
	}

	v, ok := c.v.Number()
	if !ok {
		return 0, grr.Wrap(grr.ErrInvalidValueType, "expected number, got %s", c.v.Type())
	}

	return v, nil
}

func (c Chain) String() (string, error) {
	if c.err != nil {
		return "", c.err
	}

	v, ok := c.v.String()
	if !ok {
		return "", grr.Wrap(grr.ErrInvalidValueType, "expected string, got %s", c.v.Type())
	}

	return v, nil
}

func (c Chain) Bool() (bool, error) {
	if c.err != nil {
		return false, c.err
	}

	v, ok := c.v.Bool()
	if !ok {
		return false, grr.Wrap(grr.ErrInvalidValueType, "expected bool, got %s", c.v.Type())
	}

	return v, nil
}

func (c Chain) TryNumber() float64 {
	if c.err != nil {
		return 0
	}

	v, ok := c.v.Number()
	if !ok {
		return 0
	}

	return v
}

func (c Chain) TryString() string {
	if c.err != nil {
		return ""
	}

	v, ok := c.v.String()
	if !ok {
		return ""
	}

	return v
}

func (c Chain) TryBool() bool {
	if c.err != nil {
		return false
	}

	v, ok := c.v.Bool()
	if !ok {
		return false
	}

	return v
}
