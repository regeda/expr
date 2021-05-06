package memory

import (
	"bytes"
	"encoding/binary"
)

var (
	True    = NewAddr(TypeBool, 1)
	False   = NewAddr(TypeBool, 0)
	NoBytes = NewAddr(TypeBytes)
	Nil     = NewAddr(TypeNil)
)

type Addr struct {
	typ Type
	dat []byte
	vec []Addr
}

func NewAddr(t Type, dat ...byte) Addr {
	return Addr{
		typ: t,
		dat: dat,
	}
}

func (a Addr) Size() uint32 { return uint32(len(a.dat)) }

func (a Addr) Vector() []Addr {
	return a.vec
}

func (a Addr) CopyBytes(src ...Addr) {
	var offset uint32
	for _, s := range src {
		copy(a.dat[offset:], s.dat)
		offset += s.Size()
	}
}

func (a Addr) CopyVector(v []Addr) {
	copy(a.vec, v)
}

func (a Addr) VectorAt(i int) Addr {
	return a.vec[i]
}

func (a Addr) SetVectorAt(i int, v Addr) {
	a.vec[i] = v
}

func (a Addr) VectorLen() int {
	return len(a.vec)
}

func (a Addr) Bytes() []byte {
	return a.dat
}

func (a Addr) Int64() int64 {
	return int64(binary.BigEndian.Uint64(a.dat))
}

func (a Addr) setInt64(n int64) {
	binary.BigEndian.PutUint64(a.dat, uint64(n))
}

func (a Addr) Bool() bool {
	return a.dat[0] == 1
}

func (a Addr) Type() Type {
	return a.typ
}

func (a Addr) EqualBytes(b Addr) bool {
	return bytes.Equal(a.dat, b.dat)
}
