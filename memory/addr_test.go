package memory_test

import (
	"bytes"
	"testing"

	"github.com/regeda/expr/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddr_Size(t *testing.T) {
	addr := memory.NewAddr(memory.TypeBytes, 1, 2, 3)
	assert.Equal(t, uint32(3), addr.Size())
}

func TestAddr_Bool(t *testing.T) {
	addrTrue := memory.NewAddr(memory.TypeBool, 1)
	assert.True(t, addrTrue.Bool())

	addrFalse := memory.NewAddr(memory.TypeBool, 0)
	assert.False(t, addrFalse.Bool())
}

func TestAddr_NilVector(t *testing.T) {
	addr := memory.NewAddr(memory.TypeVector)

	require.Nil(t, addr.Vector())
}

func TestAddr_CopyBytes(t *testing.T) {
	addr1 := memory.NewAddr(memory.TypeBytes, 1, 2, 3)
	addr2 := memory.NewAddr(memory.TypeBytes, 4, 5, 6)

	addrX := memory.NewAddr(memory.TypeBytes, make([]byte, 6)...)
	addrX.CopyBytes(addr1, addr2)

	assert.Equal(t, []byte{1, 2, 3, 4, 5, 6}, addrX.Bytes())
}

func TestAddr_TypeOf(t *testing.T) {
	assert.True(t, memory.True.TypeOf(memory.TypeBool))
	assert.False(t, memory.True.TypeOf(memory.TypeNil))
}

func TestAddr_EqualType(t *testing.T) {
	assert.True(t, memory.True.EqualType(memory.False))
}

func TestAddr_EqualBytes(t *testing.T) {
	assert.True(t, memory.True.EqualBytes(memory.True))
	assert.False(t, memory.True.EqualBytes(memory.False))
}

func TestAddr_Print(t *testing.T) {
	mem := memory.New()

	cases := []struct {
		name     string
		v        interface{}
		expected string
	}{
		{"string", "hello world", `bytes="hello world"`},
		{"quoted string", `hi "dude"`, `bytes="hi \"dude\""`},
		{"bool true", true, "bool=1"},
		{"bool false", false, "bool=0"},
		{"int", 123, "int64=123"},
		{"vector", []interface{}{1, true, "foo"}, `vector=[int64=1, bool=1, bytes="foo"]`},
		{"vector of vector", []interface{}{[]interface{}{true}, 1}, `vector=[vector=[bool=1], int64=1]`},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			addr, err := mem.Alloc(c.v)
			require.NoError(t, err)

			addr.Print(buf)

			assert.Equal(t, c.expected, buf.String())
		})
	}
}
