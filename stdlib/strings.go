package stdlib

import (
	"github.com/regeda/expr/assert"
	"github.com/regeda/expr/delegate"
	"github.com/regeda/expr/memory"
)

func init() {
	Register("concat", delegate.DelegatorFunc(concat).Assert(
		assert.Type(memory.TypeBytes),
	))

	Register("join", delegate.DelegatorFunc(join).Assert(assert.Every(
		assert.Len(2),
		assert.TypeAt(0, memory.TypeBytes),
		assert.TypeAt(1, memory.TypeVector),
		assert.VectorAt(1, assert.Type(memory.TypeBytes)),
	)))
}

func concat(mem *memory.Memory, argv []memory.Addr) (memory.Addr, error) {
	var size uint32
	for _, a := range argv {
		size += a.Size()
	}
	addr := mem.AllocBytes(size)
	addr.CopyBytes(argv...)
	return addr, nil
}

func join(mem *memory.Memory, argv []memory.Addr) (memory.Addr, error) {
	sep, srcs := argv[0], argv[1].Vector()
	if len(srcs) == 0 {
		return memory.NoBytes, nil
	}
	size := sep.Size() * uint32(len(srcs)-1)
	for _, src := range srcs {
		size += src.Size()
	}
	addr := mem.AllocBytes(size)
	buf := addr.Bytes()
	var offset uint32
	for i, src := range srcs {
		if i > 0 {
			copy(buf[offset:], sep.Bytes())
			offset += sep.Size()
		}
		copy(buf[offset:], src.Bytes())
		offset += src.Size()
	}
	return addr, nil
}
