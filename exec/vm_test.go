package exec_test

import (
	"bytes"
	"testing"

	"github.com/regeda/expr/ast"
	"github.com/regeda/expr/compiler"
	"github.com/regeda/expr/delegate"
	"github.com/regeda/expr/exec"
	"github.com/regeda/expr/memory"
	"github.com/regeda/expr/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVM_Exec(t *testing.T) {
	tracing := bytes.NewBuffer(nil)

	registry := delegate.NewRegistry(delegate.RegistryWithTracing(tracing))
	registry.Import(stdlib.Compare, stdlib.Strings)

	var comp compiler.Compiler

	exec := exec.New(exec.WithRegistry(registry))

	t.Run("const boolean", func(t *testing.T) {
		for _, bb := range [...]bool{true, false} {
			tracing.Reset()

			bcode := comp.Compile(
				ast.Exit().Nest(
					ast.Bool(bb),
				))

			addr, err := exec.Exec(bcode)
			require.NoError(t, err)
			require.NotNil(t, addr)

			require.Equal(t, memory.TypeBool, addr.Type())
			assert.Equal(t, bb, addr.Bool())

			assert.Empty(t, tracing.String())
		}
	})

	t.Run("const int", func(t *testing.T) {
		tracing.Reset()

		bcode := comp.Compile(
			ast.Exit().Nest(
				ast.Int(1),
			))

		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeInt64, addr.Type())
		assert.Equal(t, int64(1), addr.Int64())

		assert.Empty(t, tracing.String())
	})

	t.Run("const array", func(t *testing.T) {
		tracing.Reset()

		bcode := comp.Compile(
			ast.Exit().Nest(
				ast.Arr().Nest(
					ast.Int(1),
					ast.Int(2),
					ast.Int(3),
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

		assert.Empty(t, tracing.String())
	})

	t.Run("array of array", func(t *testing.T) {
		tracing.Reset()

		bcode := comp.Compile(
			ast.Exit().Nest(
				ast.Arr().Nest(
					ast.Arr().Nest(
						ast.Int(0xff),
					),
					ast.Str("foo"),
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

		assert.Empty(t, tracing.String())
	})

	t.Run("func join", func(t *testing.T) {
		tracing.Reset()

		bcode := comp.Compile(
			ast.Exit().Nest(
				ast.Call("join").Nest(
					ast.Str("$"),
					ast.Arr().Nest(
						ast.Str("a"),
						ast.Str("b"),
						ast.Str("c"),
					),
				),
			))
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBytes, addr.Type())
		assert.Equal(t, []byte("a$b$c"), addr.Bytes())

		assert.Equal(t, "join(bytes=\"$\", vector=[bytes=\"a\", bytes=\"b\", bytes=\"c\"]) -> bytes=\"a$b$c\"\n", tracing.String())
	})

	t.Run("func concat", func(t *testing.T) {
		tracing.Reset()

		bcode := comp.Compile(
			ast.Exit().Nest(
				ast.Call("concat").Nest(
					ast.Str("foo"),
					ast.Str("bar"),
					ast.Str("baz"),
				),
			))
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBytes, addr.Type())
		assert.Equal(t, []byte("foobarbaz"), addr.Bytes())

		assert.Equal(t, "concat(bytes=\"foo\", bytes=\"bar\", bytes=\"baz\") -> bytes=\"foobarbaz\"\n", tracing.String())
	})

	t.Run("func equals int", func(t *testing.T) {
		tracing.Reset()

		bcode := comp.Compile(
			ast.Exit().Nest(
				ast.Call("equals").Nest(
					ast.Int(1),
					ast.Int(1),
				),
			))
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBool, addr.Type())
		assert.True(t, addr.Bool())

		assert.Equal(t, "equals(int64=1, int64=1) -> bool=1\n", tracing.String())
	})

	t.Run("func equals str", func(t *testing.T) {
		tracing.Reset()

		bcode := comp.Compile(
			ast.Exit().Nest(
				ast.Call("equals").Nest(
					ast.Str("foo"),
					ast.Str("foo"),
				),
			))
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBool, addr.Type())
		assert.True(t, addr.Bool())

		assert.Equal(t, "equals(bytes=\"foo\", bytes=\"foo\") -> bool=1\n", tracing.String())
	})

	t.Run("func equals vector", func(t *testing.T) {
		tracing.Reset()

		bcode := comp.Compile(
			ast.Exit().Nest(
				ast.Call("equals").Nest(
					ast.Arr().Nest(
						ast.Str("foo"),
						ast.Str("bar"),
					),
					ast.Arr().Nest(
						ast.Str("foo"),
						ast.Str("bar"),
					),
				),
			))

		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBool, addr.Type())
		assert.True(t, addr.Bool())

		assert.Equal(t, "equals(vector=[bytes=\"foo\", bytes=\"bar\"], vector=[bytes=\"foo\", bytes=\"bar\"]) -> bool=1\n", tracing.String())
	})

	t.Run("func equals concat", func(t *testing.T) {
		tracing.Reset()

		bcode := comp.Compile(
			ast.Exit().Nest(
				ast.Call("equals").Nest(
					ast.Str("foobarbaz"),
					ast.Call("concat").Nest(
						ast.Str("foo"),
						ast.Str("bar"),
						ast.Str("baz"),
					),
				),
			))
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBool, addr.Type())
		assert.True(t, addr.Bool())

		assert.Equal(t, "concat(bytes=\"foo\", bytes=\"bar\", bytes=\"baz\") -> bytes=\"foobarbaz\"\nequals(bytes=\"foobarbaz\", bytes=\"foobarbaz\") -> bool=1\n", tracing.String())
	})

	t.Run("func contains", func(t *testing.T) {
		tracing.Reset()

		bcode := comp.Compile(
			ast.Exit().Nest(
				ast.Call("contains").Nest(
					ast.Arr().Nest(
						ast.Str("a"),
						ast.Str("b"),
						ast.Str("c"),
					),
					ast.Str("a"),
				),
			))
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBool, addr.Type())
		assert.True(t, addr.Bool())

		assert.Equal(t, "contains(vector=[bytes=\"a\", bytes=\"b\", bytes=\"c\"], bytes=\"a\") -> bool=1\n", tracing.String())
	})

	t.Run("func intersects", func(t *testing.T) {
		tracing.Reset()

		bcode := comp.Compile(
			ast.Exit().Nest(
				ast.Call("intersects").Nest(
					ast.Arr().Nest(
						ast.Int(1),
						ast.Int(2),
					),
					ast.Arr().Nest(
						ast.Int(2),
					),
				),
			))
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBool, addr.Type())
		assert.True(t, addr.Bool())

		assert.Equal(t, "intersects(vector=[int64=1, int64=2], vector=[int64=2]) -> bool=1\n", tracing.String())
	})

	t.Run("delegator not exists", func(t *testing.T) {
		tracing.Reset()

		bcode := comp.Compile(
			ast.Exit().Nest(
				ast.Call("NONAME"),
			))

		addr, err := exec.Exec(bcode)
		require.EqualError(t, err, "failed to exec frame at 0: delegator <noname> not exists")
		assert.Equal(t, memory.Nil, addr)
	})
}
