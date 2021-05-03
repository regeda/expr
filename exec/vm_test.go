package exec_test

import (
	"testing"

	"github.com/regeda/expr/exec"
	"github.com/regeda/expr/internal/ast/value"
	"github.com/regeda/expr/internal/compiler"
	"github.com/regeda/expr/memory"
	"github.com/regeda/expr/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVM_Exec(t *testing.T) {
	comp := compiler.New()
	exec := exec.New(stdlib.Registry())

	t.Run("const boolean", func(t *testing.T) {
		for _, bb := range [...]bool{true, false} {
			bcode := comp.Compile(value.Nest(
				value.Exit(),
				value.Bool(bb),
			))

			addr, err := exec.Exec(bcode)
			require.NoError(t, err)
			require.NotNil(t, addr)

			require.Equal(t, memory.TypeBool, addr.Type())
			assert.Equal(t, bb, addr.Bool())
		}
	})

	t.Run("const int", func(t *testing.T) {
		bcode := comp.Compile(value.Nest(
			value.Exit(),
			value.Int(1),
		))

		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeInt64, addr.Type())
		assert.Equal(t, int64(1), addr.Int64())
	})

	t.Run("const array", func(t *testing.T) {
		bcode := comp.Compile(value.Nest(
			value.Exit(),
			value.Nest(
				value.Arr(),
				value.Int(1),
				value.Int(2),
				value.Int(3),
			),
		))

		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeVector, addr.Type())
		assert.Equal(t, 3, addr.VectorLen())

		for i, n := range [...]int64{1, 2, 3} {
			addrAt := addr.VectorAt(i)
			require.NotNil(t, addrAt)

			require.Equal(t, memory.TypeInt64, addrAt.Type())
			assert.Equal(t, n, addrAt.Int64())
		}
	})

	t.Run("array of array", func(t *testing.T) {
		bcode := comp.Compile(value.Nest(
			value.Exit(),
			value.Nest(
				value.Arr(),
				value.Nest(
					value.Arr(),
					value.Int(0xff),
				),
				value.Str("foo"),
			),
		))

		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeVector, addr.Type())
		assert.Equal(t, 2, addr.VectorLen())

		subarrayAtPos0 := addr.VectorAt(0)
		require.NotNil(t, subarrayAtPos0)
		require.Equal(t, memory.TypeVector, subarrayAtPos0.Type())
		assert.Equal(t, 1, subarrayAtPos0.VectorLen())

		intAtSubarrayAtPos0 := subarrayAtPos0.VectorAt(0)
		require.NotNil(t, intAtSubarrayAtPos0)
		require.Equal(t, memory.TypeInt64, intAtSubarrayAtPos0.Type())
		assert.Equal(t, int64(0xff), intAtSubarrayAtPos0.Int64())

		strAtPos1 := addr.VectorAt(1)
		require.NotNil(t, strAtPos1)
		require.Equal(t, memory.TypeBytes, strAtPos1.Type())
		assert.Equal(t, []byte("foo"), strAtPos1.Bytes())
	})

	t.Run("func join", func(t *testing.T) {
		bcode := comp.Compile(value.Nest(
			value.Exit(),
			value.Nest(
				value.Call("join"),
				value.Str("$"),
				value.Nest(
					value.Arr(),
					value.Str("a"),
					value.Str("b"),
					value.Str("c"),
				),
			),
		))
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBytes, addr.Type())
		assert.Equal(t, []byte("a$b$c"), addr.Bytes())
	})

	t.Run("func concat", func(t *testing.T) {
		bcode := comp.Compile(value.Nest(
			value.Exit(),
			value.Nest(
				value.Call("concat"),
				value.Str("foo"),
				value.Str("bar"),
				value.Str("baz"),
			),
		))
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBytes, addr.Type())
		assert.Equal(t, []byte("foobarbaz"), addr.Bytes())
	})

	t.Run("func equals int", func(t *testing.T) {
		bcode := comp.Compile(value.Nest(
			value.Exit(),
			value.Nest(
				value.Call("equals"),
				value.Int(1),
				value.Int(1),
			),
		))
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBool, addr.Type())
		assert.True(t, addr.Bool())
	})

	t.Run("func equals str", func(t *testing.T) {
		bcode := comp.Compile(value.Nest(
			value.Exit(),
			value.Nest(
				value.Call("equals"),
				value.Str("foo"),
				value.Str("foo"),
			),
		))
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBool, addr.Type())
		assert.True(t, addr.Bool())
	})

	t.Run("func equals concat", func(t *testing.T) {
		bcode := comp.Compile(value.Nest(
			value.Exit(),
			value.Nest(
				value.Call("equals"),
				value.Str("foobarbaz"),
				value.Nest(
					value.Call("concat"),
					value.Str("foo"),
					value.Str("bar"),
					value.Str("baz"),
				),
			),
		))
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBool, addr.Type())
		assert.True(t, addr.Bool())
	})

	t.Run("func contains", func(t *testing.T) {
		bcode := comp.Compile(value.Nest(
			value.Exit(),
			value.Nest(
				value.Call("contains"),
				value.Nest(
					value.Arr(),
					value.Str("a"),
					value.Str("b"),
					value.Str("c"),
				),
				value.Str("a"),
			),
		))
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBool, addr.Type())
		assert.True(t, addr.Bool())
	})

	t.Run("func intersects", func(t *testing.T) {
		bcode := comp.Compile(value.Nest(
			value.Exit(),
			value.Nest(
				value.Call("intersects"),
				value.Nest(
					value.Arr(),
					value.Int(1),
					value.Int(2),
				),
				value.Nest(
					value.Arr(),
					value.Int(2),
				),
			),
		))
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBool, addr.Type())
		assert.True(t, addr.Bool())
	})

	t.Run("delegator not exists", func(t *testing.T) {
		bcode := comp.Compile(value.Nest(
			value.Exit(),
			value.Call("NONAME"),
		))

		addr, err := exec.Exec(bcode)
		require.EqualError(t, err, "failed to exec frame at 0: delegator <noname> not exists")
		assert.Nil(t, addr)
	})
}
