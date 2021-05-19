package stdlib

import (
	"errors"

	"github.com/regeda/expr/delegate"
	"github.com/regeda/expr/memory"
)

var Strings = delegate.Module{
	"concat": delegate.DelegatorFunc(concat),
	"join":   delegate.DelegatorFunc(join),
}

var (
	errConcatExpectedBytes = errors.New("concat: expected bytes")
)

func concat(mem *memory.Memory, argv []memory.Addr) (memory.Addr, error) {
	var size uint32
	for _, a := range argv {
		if !a.TypeOf(memory.TypeBytes) {
			return memory.Nil, errConcatExpectedBytes
		}
		size += a.Size()
	}
	addr := mem.AllocBytes(size)
	addr.CopyBytes(argv...)
	return addr, nil
}

var (
	errJoinExpectedTwoArgs         = errors.New("join: expected 2 args")
	errJoinExpectedBytesAt0        = errors.New("join: expected bytes at 0")
	errJoinExpectedArrayOfBytesAt1 = errors.New("join: expected array of bytes at 1")
)

func join(mem *memory.Memory, argv []memory.Addr) (memory.Addr, error) {
	if len(argv) != 2 {
		return memory.Nil, errJoinExpectedTwoArgs
	}
	if !argv[0].TypeOf(memory.TypeBytes) {
		return memory.Nil, errJoinExpectedBytesAt0
	}
	if !argv[1].TypeOf(memory.TypeVector) {
		return memory.Nil, errJoinExpectedArrayOfBytesAt1
	}
	sep, srcs := argv[0], argv[1].Vector()
	if len(srcs) == 0 {
		return memory.NoBytes, nil
	}
	size := sep.Size() * uint32(len(srcs)-1)
	for _, src := range srcs {
		if !src.TypeOf(memory.TypeBytes) {
			return memory.Nil, errJoinExpectedArrayOfBytesAt1
		}
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
