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

	t.Run("links limit overflow", func(t *testing.T) {
		var b memory.Memory

		addr, err := b.AllocVector(memory.LinksLimit + 1)
		require.EqualError(t, err, "memory: links limit exceeded")
		require.Nil(t, addr)
	})

	t.Run("copy vector error", func(t *testing.T) {
		var b memory.Memory

		elems := make([]*memory.Addr, memory.LinksLimit+1)

		addr, err := b.CopyVector(elems...)
		require.EqualError(t, err, "memory: links limit exceeded")
		require.Nil(t, addr)
	})

	t.Run("grid limit overflow", func(t *testing.T) {
		var b memory.Memory

		bytes := []byte{1}

		var i uint32
		for i = 0; i < memory.GridLimit; i++ {
			addr, err := b.AllocBytesAddr(bytes)
			require.NoError(t, err)
			require.NotNil(t, addr)
		}

		addr, err := b.AllocBytesAddr(bytes)
		require.EqualError(t, err, "memory: grid limit exceeded")
		require.Nil(t, addr)
	})
}

func Test_Memory_Alloc(t *testing.T) {
	t.Run("bytes", func(t *testing.T) {
		var b memory.Memory

		addr, err := b.Alloc([]byte{1, 2, 3})
		require.NoError(t, err)
		require.NotNil(t, addr)

		assert.Equal(t, memory.TypeBytes, addr.Type())
		assert.Equal(t, []byte{1, 2, 3}, addr.Bytes())
	})

	t.Run("string", func(t *testing.T) {
		var b memory.Memory

		addr, err := b.Alloc("foobar")
		require.NoError(t, err)
		require.NotNil(t, addr)

		assert.Equal(t, memory.TypeBytes, addr.Type())
		assert.Equal(t, []byte("foobar"), addr.Bytes())
	})

	t.Run("bool", func(t *testing.T) {
		var b memory.Memory

		addrT, err := b.Alloc(true)
		require.NoError(t, err)
		assert.Equal(t, memory.ConstTrue, addrT)

		addrF, err := b.Alloc(false)
		require.NoError(t, err)
		assert.Equal(t, memory.ConstFalse, addrF)
	})

	t.Run("int", func(t *testing.T) {
		var b memory.Memory

		addr, err := b.Alloc(1)
		require.NoError(t, err)
		require.NotNil(t, addr)

		assert.Equal(t, memory.TypeInt64, addr.Type())
		assert.Equal(t, int64(1), addr.Int64())
	})

	t.Run("int64", func(t *testing.T) {
		var b memory.Memory

		addr, err := b.Alloc(int64(1))
		require.NoError(t, err)
		require.NotNil(t, addr)

		assert.Equal(t, memory.TypeInt64, addr.Type())
		assert.Equal(t, int64(1), addr.Int64())
	})

	t.Run("untyped vector", func(t *testing.T) {
		var b memory.Memory

		addr, err := b.Alloc([]interface{}{"foo", "bar"})
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
		var b memory.Memory

		addr, err := b.Alloc(struct{}{})
		require.EqualError(t, err, "memory: unsupported type struct {}")
		require.Nil(t, addr)
	})
}
