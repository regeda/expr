package stdlib

import (
	"github.com/regeda/expr/assert"
	"github.com/regeda/expr/delegate"
	"github.com/regeda/expr/memory"
)

func init() {
	Register("equals", delegate.DelegatorFunc(equals).
		Assert(assert.Every(
			assert.Len(2),
			assert.Any(
				assert.Type(memory.TypeInt64),
				assert.Type(memory.TypeBytes),
			),
		)))

	Register("contains", delegate.DelegatorFunc(contains).
		Assert(assert.Every(
			assert.Len(2),
			assert.TypeAt(0, memory.TypeVector),
			assert.Any(
				assert.Every(
					assert.VectorAt(0, assert.Type(memory.TypeInt64)),
					assert.TypeAt(1, memory.TypeInt64),
				),
				assert.Every(
					assert.VectorAt(0, assert.Type(memory.TypeBytes)),
					assert.TypeAt(1, memory.TypeBytes),
				),
			),
		)))

	Register("intersects", delegate.DelegatorFunc(intersects).
		Assert(assert.Every(
			assert.Len(2),
			assert.Type(memory.TypeVector),
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
		)))

}

func equals(mem *memory.Memory, argv []*memory.Addr) (*memory.Addr, error) {
	if argv[0].EqualBytes(argv[1]) {
		return memory.ConstTrue, nil
	}
	return memory.ConstFalse, nil
}

func contains(mem *memory.Memory, argv []*memory.Addr) (*memory.Addr, error) {
	for _, p := range argv[0].Vector() {
		if p.EqualBytes(argv[1]) {
			return memory.ConstTrue, nil
		}
	}
	return memory.ConstFalse, nil
}

func intersects(mem *memory.Memory, argv []*memory.Addr) (*memory.Addr, error) {
	for _, a := range argv[0].Vector() {
		for _, b := range argv[1].Vector() {
			if a.EqualBytes(b) {
				return memory.ConstTrue, nil
			}
		}
	}
	return memory.ConstFalse, nil
}
