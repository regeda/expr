package memory_test

import (
	"testing"

	"github.com/regeda/expr/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Addr_Size(t *testing.T) {
	addr := memory.NewAddr(memory.TypeBytes, 1, 2, 3)
	assert.Equal(t, uint32(3), addr.Size())
}

func Test_Addr_Bool(t *testing.T) {
	addrTrue := memory.NewAddr(memory.TypeBool, 1)
	assert.True(t, addrTrue.Bool())

	addrFalse := memory.NewAddr(memory.TypeBool, 0)
	assert.False(t, addrFalse.Bool())
}

func Test_Addr_NilVector(t *testing.T) {
	addr := memory.NewAddr(memory.TypeVector)

	require.Nil(t, addr.Vector())
}

func Test_Addr_CopyBytes(t *testing.T) {
	addr1 := memory.NewAddr(memory.TypeBytes, 1, 2, 3)
	addr2 := memory.NewAddr(memory.TypeBytes, 4, 5, 6)

	addrX := memory.NewAddr(memory.TypeBytes, make([]byte, 6)...)
	addrX.CopyBytes(addr1, addr2)

	assert.Equal(t, []byte{1, 2, 3, 4, 5, 6}, addrX.Bytes())
}

func Test_Addr_EqualBytes(t *testing.T) {
	assert.True(t, memory.ConstTrue.EqualBytes(memory.ConstTrue))
	assert.False(t, memory.ConstTrue.EqualBytes(memory.ConstFalse))
}
