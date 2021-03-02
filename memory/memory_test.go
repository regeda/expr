package memory_test

import (
	"testing"

	"github.com/regeda/expr/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemory_Reset(t *testing.T) {
	var b memory.Memory

	assert.Equal(t, memory.HeapLimit, b.Heapfree())

	_, err := b.AllocBytes(0xff)
	require.NoError(t, err)

	assert.Equal(t, memory.HeapLimit-0xff, b.Heapfree())

	b.Reset()

	assert.Equal(t, memory.HeapLimit, b.Heapfree())
}

func TestMemory_AllocBytes(t *testing.T) {
	t.Run("alloc zero", func(t *testing.T) {
		var b memory.Memory

		addr, err := b.AllocBytes(0)
		require.NoError(t, err)
		assert.Equal(t, uint32(0), addr.Size())
	})

	t.Run("alloc bytes", func(t *testing.T) {
		var b memory.Memory

		addr, err := b.AllocBytes(memory.HeapLimit)
		require.NoError(t, err)
		assert.Equal(t, memory.HeapLimit, addr.Size())
	})

	t.Run("out of memory", func(t *testing.T) {
		var b memory.Memory

		_, err := b.AllocBytes(memory.HeapLimit)
		require.NoError(t, err)

		addr, err := b.AllocBytes(1)
		assert.EqualError(t, err, "memory: out of memory to alloc 1 bytes")
		assert.Nil(t, addr)
	})
}

func TestMemory_AllocInt64(t *testing.T) {
	t.Run("alloc int64", func(t *testing.T) {
		var b memory.Memory

		addr, err := b.AllocInt64(0xffffff)
		require.NoError(t, err)
		assert.Equal(t, memory.HeapLimit-8, b.Heapfree())

		assert.Equal(t, memory.TypeInt64, addr.Type())
		assert.Equal(t, int64(0xffffff), addr.Int64())
	})

	t.Run("out of memory", func(t *testing.T) {
		var b memory.Memory

		_, err := b.AllocBytes(memory.HeapLimit)
		require.NoError(t, err)

		addr, err := b.AllocInt64(0xffffff)
		require.EqualError(t, err, "memory: out of memory to alloc 8 bytes")
		assert.Nil(t, addr)
	})
}

func TestMemory_AllocBytesAddr(t *testing.T) {
	var b memory.Memory

	addr, err := b.AllocBytesAddr([]byte{1})
	require.NoError(t, err)
	require.NotNil(t, addr)

	assert.Equal(t, memory.HeapLimit, b.Heapfree())
	assert.Equal(t, memory.TypeBytes, addr.Type())
	assert.Equal(t, []byte{1}, addr.Bytes())
}

func TestMemory_AllocVector(t *testing.T) {
	t.Run("alloc vector", func(t *testing.T) {
		var b memory.Memory

		addr, err := b.AllocVector(3)
		require.NoError(t, err)

		assert.Equal(t, memory.TypeVector, addr.Type())
		assert.Equal(t, 3, addr.VectorLen())
	})

	t.Run("copy vector", func(t *testing.T) {
		var b memory.Memory

		elem, err := b.AllocBytes(1)
		require.NoError(t, err)

		addr, err := b.CopyVector(elem, elem, elem)
		require.NoError(t, err)
		require.NotNil(t, addr)

		assert.Equal(t, memory.TypeVector, addr.Type())
		assert.Equal(t, 3, addr.VectorLen())

		for i := 0; i < 3; i++ {
			assert.Equal(t, elem, addr.VectorAt(i))
		}
	})
}
