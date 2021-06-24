package memory

import "fmt"

const (
	sizeInt8  = 1
	sizeInt16 = sizeInt8 << 1
	sizeInt32 = sizeInt16 << 1
	sizeInt64 = sizeInt32 << 1
)

type Memory struct {
	g grid
	h heap
}

type Opt func(*Memory)

func PreallocGrid(size uint32) Opt {
	return func(m *Memory) {
		m.g.grow(size)
	}
}

func PreallocHeap(size uint32) Opt {
	return func(m *Memory) {
		m.h.grow(size)
	}
}

func New(opts ...Opt) Memory {
	m := Memory{}

	for _, opt := range opts {
		opt(&m)
	}

	return m
}

func (b *Memory) Reset() {
	b.g.reset()
	b.h.reset()
}

func (b *Memory) alloc(t Type, size uint32) Addr {
	buf := b.h.alloc(size)
	return NewAddr(t, buf...)
}

func (b *Memory) Alloc(in interface{}) (Addr, error) {
	switch v := in.(type) {
	case []byte:
		return b.AllocBytesAddr(v), nil
	case string:
		return b.AllocBytesAddr([]byte(v)), nil
	case int:
		return b.AllocInt64(int64(v)), nil
	case int64:
		return b.AllocInt64(v), nil
	case bool:
		if v {
			return True, nil
		}
		return False, nil
	case []interface{}:
		vec := b.AllocVector(uint32(len(v)))
		for i, e := range v {
			addr, err := b.Alloc(e)
			if err != nil {
				return Nil, err
			}
			vec.SetVectorAt(i, addr)
		}
		return vec, nil
	default:
		return Nil, fmt.Errorf("memory: unsupported type %T", v)
	}
}

func (b *Memory) AllocBytesAddr(dat []byte) Addr {
	return NewAddr(TypeBytes, dat...)
}

func (b *Memory) AllocBytes(size uint32) Addr {
	if size == 0 {
		return NoBytes
	}
	return b.alloc(TypeBytes, size)
}

func (b *Memory) AllocInt64(v int64) Addr {
	addr := b.alloc(TypeInt64, sizeInt64)
	addr.SetInt64(v)
	return addr
}

func (b *Memory) AllocVector(size uint32) Addr {
	addr := NewAddr(TypeVector)
	addr.vec = b.g.alloc(size)
	return addr
}

func (b *Memory) CopyVector(elems ...Addr) Addr {
	addr := b.AllocVector(uint32(len(elems)))
	addr.CopyVector(elems)
	return addr
}
