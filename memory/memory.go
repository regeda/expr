package memory

import "fmt"

type Memory struct {
	g grid
	l links
	h heap
}

func (b *Memory) Heapfree() uint32 {
	return b.h.size() - b.h.p
}

func (b *Memory) Reset() {
	b.g.reset()
	b.l.reset()
	b.h.reset()
}

func (b *Memory) alloc(t Type, size uint32) (*Addr, error) {
	buf, err := b.h.alloc(size)
	if err != nil {
		return nil, err
	}
	return b.allocAddr(t, buf)
}

func (b *Memory) allocAddr(t Type, dat []byte) (*Addr, error) {
	return b.g.add(Addr{
		typ: t,
		dat: dat,
	})
}

func (b *Memory) Alloc(in interface{}) (*Addr, error) {
	switch v := in.(type) {
	case []byte:
		return b.AllocBytesAddr(v)
	case string:
		return b.AllocBytesAddr([]byte(v))
	case int:
		return b.AllocInt64(int64(v))
	case int64:
		return b.AllocInt64(v)
	case bool:
		if v {
			return ConstTrue, nil
		}
		return ConstFalse, nil
	case []interface{}:
		vec, err := b.AllocVector(uint32(len(v)))
		if err != nil {
			return nil, err
		}
		for i, e := range v {
			addr, err := b.Alloc(e)
			if err != nil {
				return nil, err
			}
			vec.SetVectorAt(i, addr)
		}
		return vec, nil
	default:
		return nil, fmt.Errorf("memory: unsupported type %T", v)
	}
}

func (b *Memory) AllocBytesAddr(dat []byte) (*Addr, error) {
	return b.allocAddr(TypeBytes, dat)
}

func (b *Memory) AllocBytes(size uint32) (*Addr, error) {
	if size == 0 {
		return ConstNoBytes, nil
	}
	return b.alloc(TypeBytes, size)
}

func (b *Memory) AllocInt64(v int64) (*Addr, error) {
	addr, err := b.alloc(TypeInt64, sizeInt64)
	if err != nil {
		return nil, err
	}
	addr.setInt64(v)
	return addr, nil
}

func (b *Memory) AllocVector(size uint32) (*Addr, error) {
	addr, err := b.allocAddr(TypeVector, nil)
	if err != nil {
		return nil, err
	}
	vec, err := b.l.alloc(size)
	if err != nil {
		return nil, err
	}
	addr.vec = vec
	return addr, nil
}

func (b *Memory) CopyVector(elems ...*Addr) (*Addr, error) {
	addr, err := b.AllocVector(uint32(len(elems)))
	if err != nil {
		return nil, err
	}
	addr.CopyVector(elems)
	return addr, nil
}
