package stdlib

import (
	"errors"

	"github.com/regeda/expr/delegate"
	"github.com/regeda/expr/memory"
)

var Compare = delegate.Module{
	"equals":     delegate.DelegatorFunc(equals),
	"contains":   delegate.DelegatorFunc(contains),
	"intersects": delegate.DelegatorFunc(intersects),
}

var (
	errEqualsExpectedTwoArgs = errors.New("equals: expected 2 args")
)

func equals(mem *memory.Memory, argv []memory.Addr) (memory.Addr, error) {
	if len(argv) != 2 {
		return memory.Nil, errEqualsExpectedTwoArgs
	}
	if argv[0].Equal(argv[1]) {
		return memory.True, nil
	}
	return memory.False, nil
}

var (
	errContainsExpectedTwoArgs   = errors.New("contains: expected 2 args")
	errContainsExpectedArrayAt0  = errors.New("contains: expected array at 0")
	errContainsExpectedScalarAt1 = errors.New("contains: expected scalar at 1")
)

func contains(mem *memory.Memory, argv []memory.Addr) (memory.Addr, error) {
	if len(argv) != 2 {
		return memory.Nil, errContainsExpectedTwoArgs
	}
	if !argv[0].TypeOf(memory.TypeVector) {
		return memory.Nil, errContainsExpectedArrayAt0
	}
	if argv[1].TypeOf(memory.TypeVector) {
		return memory.Nil, errContainsExpectedScalarAt1
	}
	for _, p := range argv[0].Vector() {
		if p.Equal(argv[1]) {
			return memory.True, nil
		}
	}
	return memory.False, nil
}

var (
	errIntersectsExpectedTwoArgs = errors.New("intersects: expected 2 args")
	errIntersectsExpectedArrays  = errors.New("intersects: expected arrays")
)

func intersects(mem *memory.Memory, argv []memory.Addr) (memory.Addr, error) {
	if len(argv) != 2 {
		return memory.Nil, errIntersectsExpectedTwoArgs
	}
	if !argv[0].EqualType(argv[1]) || !argv[0].TypeOf(memory.TypeVector) {
		return memory.Nil, errIntersectsExpectedArrays
	}
	for _, a := range argv[0].Vector() {
		for _, b := range argv[1].Vector() {
			if a.Equal(b) {
				return memory.True, nil
			}
		}
	}
	return memory.False, nil
}
