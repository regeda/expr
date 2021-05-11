package stdlib

import (
	"github.com/regeda/expr/assert"
	"github.com/regeda/expr/delegate"
	"github.com/regeda/expr/memory"
)

var Compare = delegate.Module{
	"equals": delegate.DelegatorFunc(equals).
		Assert(assert.Every(
			assert.Len(2),
			assert.Any(
				assert.TypeInt64,
				assert.TypeBytes,
				assert.Every(
					assert.TypeVector,
					assert.VectorAt(0, assert.Any(
						assert.TypeInt64,
						assert.TypeBytes,
					)),
					assert.VectorAt(1, assert.Any(
						assert.TypeInt64,
						assert.TypeBytes,
					)),
				),
			),
		)),
	"contains": delegate.DelegatorFunc(contains).
		Assert(assert.Every(
			assert.Len(2),
			assert.TypeAt(0, memory.TypeVector),
			assert.Any(
				assert.Every(
					assert.VectorAt(0, assert.TypeInt64),
					assert.TypeAt(1, memory.TypeInt64),
				),
				assert.Every(
					assert.VectorAt(0, assert.TypeBytes),
					assert.TypeAt(1, memory.TypeBytes),
				),
			),
		)),
	"intersects": delegate.DelegatorFunc(intersects).
		Assert(assert.Every(
			assert.Len(2),
			assert.TypeVector,
			assert.Any(
				assert.Every(
					assert.VectorAt(0, assert.Type(memory.TypeInt64)),
					assert.VectorAt(1, assert.Type(memory.TypeInt64)),
				),
				assert.Every(
					assert.VectorAt(0, assert.Type(memory.TypeBytes)),
					assert.VectorAt(1, assert.Type(memory.TypeBytes)),
				),
			),
		)),
}

func equals(mem *memory.Memory, argv []memory.Addr) (memory.Addr, error) {
	if !argv[0].EqualType(argv[1]) {
		return memory.False, nil
	}
	if argv[0].TypeOf(memory.TypeVector) {
		if argv[0].VectorLen() != argv[1].VectorLen() {
			return memory.False, nil
		}
		for i, a := range argv[0].Vector() {
			if !a.EqualBytes(argv[1].VectorAt(i)) {
				return memory.False, nil
			}
		}
		return memory.True, nil
	}
	if argv[0].EqualBytes(argv[1]) {
		return memory.True, nil
	}
	return memory.False, nil
}

func contains(mem *memory.Memory, argv []memory.Addr) (memory.Addr, error) {
	for _, p := range argv[0].Vector() {
		if p.EqualBytes(argv[1]) {
			return memory.True, nil
		}
	}
	return memory.False, nil
}

func intersects(mem *memory.Memory, argv []memory.Addr) (memory.Addr, error) {
	for _, a := range argv[0].Vector() {
		for _, b := range argv[1].Vector() {
			if a.EqualBytes(b) {
				return memory.True, nil
			}
		}
	}
	return memory.False, nil
}
