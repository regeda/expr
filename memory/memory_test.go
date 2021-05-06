package memory_test

import (
	"testing"

	"github.com/regeda/expr/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var mem memory.Memory

func TestMemory_AllocBytes(t *testing.T) {
	t.Run("alloc zero", func(t *testing.T) {

		addr := mem.AllocBytes(0)
		assert.Equal(t, uint32(0), addr.Size())
	})
}

func TestMemory_AllocInt64(t *testing.T) {
	t.Run("alloc int64", func(t *testing.T) {
		addr := mem.AllocInt64(0xffffff)

		assert.Equal(t, memory.TypeInt64, addr.Type())
		assert.Equal(t, int64(0xffffff), addr.Int64())
	})
}

func TestMemory_AllocBytesAddr(t *testing.T) {
	addr := mem.AllocBytesAddr([]byte{1})
	require.NotNil(t, addr)

	assert.Equal(t, memory.TypeBytes, addr.Type())
	assert.Equal(t, []byte{1}, addr.Bytes())
}

func TestMemory_AllocVector(t *testing.T) {
	t.Run("alloc vector", func(t *testing.T) {
		addr := mem.AllocVector(3)

		assert.Equal(t, memory.TypeVector, addr.Type())
		assert.Equal(t, 3, addr.VectorLen())
	})

	t.Run("copy vector", func(t *testing.T) {
		elem := mem.AllocBytes(1)

		addr := mem.CopyVector(elem, elem, elem)
		require.NotNil(t, addr)

		assert.Equal(t, memory.TypeVector, addr.Type())
		assert.Equal(t, 3, addr.VectorLen())

		for i := 0; i < 3; i++ {
			assert.Equal(t, elem, addr.VectorAt(i))
		}
	})
}

func TestMemory_Alloc(t *testing.T) {
	t.Run("bytes", func(t *testing.T) {
		addr, err := mem.Alloc([]byte{1, 2, 3})
		require.NoError(t, err)
		require.NotNil(t, addr)

		assert.Equal(t, memory.TypeBytes, addr.Type())
		assert.Equal(t, []byte{1, 2, 3}, addr.Bytes())
	})

	t.Run("string", func(t *testing.T) {
		addr, err := mem.Alloc("foobar")
		require.NoError(t, err)
		require.NotNil(t, addr)

		assert.Equal(t, memory.TypeBytes, addr.Type())
		assert.Equal(t, []byte("foobar"), addr.Bytes())
	})

	t.Run("bool", func(t *testing.T) {
		addrT, err := mem.Alloc(true)
		require.NoError(t, err)
		assert.Equal(t, memory.True, addrT)

		addrF, err := mem.Alloc(false)
		require.NoError(t, err)
		assert.Equal(t, memory.False, addrF)
	})

	t.Run("int", func(t *testing.T) {
		addr, err := mem.Alloc(1)
		require.NoError(t, err)
		require.NotNil(t, addr)

		assert.Equal(t, memory.TypeInt64, addr.Type())
		assert.Equal(t, int64(1), addr.Int64())
	})

	t.Run("int64", func(t *testing.T) {
		addr, err := mem.Alloc(int64(1))
		require.NoError(t, err)
		require.NotNil(t, addr)

		assert.Equal(t, memory.TypeInt64, addr.Type())
		assert.Equal(t, int64(1), addr.Int64())
	})

	t.Run("untyped vector", func(t *testing.T) {
		addr, err := mem.Alloc([]interface{}{"foo", "bar"})
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeVector, addr.Type())
		require.Equal(t, 2, addr.VectorLen())

		require.Equal(t, memory.TypeBytes, addr.VectorAt(0).Type())
		assert.Equal(t, []byte("foo"), addr.VectorAt(0).Bytes())

		require.Equal(t, memory.TypeBytes, addr.VectorAt(1).Type())
		assert.Equal(t, []byte("bar"), addr.VectorAt(1).Bytes())
	})

	t.Run("unsupported type", func(t *testing.T) {
		addr, err := mem.Alloc(struct{}{})
		require.EqualError(t, err, "memory: unsupported type struct {}")
		assert.Equal(t, memory.Nil, addr)
	})
}
