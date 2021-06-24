package exec_test

import (
	"bytes"
	"testing"

	"github.com/regeda/expr/compiler"
	"github.com/regeda/expr/delegate"
	"github.com/regeda/expr/exec"
	"github.com/regeda/expr/lexer"
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

			typ := lexer.TypFalse
			if bb {
				typ = lexer.TypTrue
			}

			bcode := comp.Compile([]lexer.Node{
				{Typ: typ},
			})
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

		bcode := comp.Compile([]lexer.Node{
			{Typ: lexer.TypInt, DatN: 1},
		})
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeInt64, addr.Type())
		assert.Equal(t, int64(1), addr.Int64())

		assert.Empty(t, tracing.String())
	})

	t.Run("const array", func(t *testing.T) {
		tracing.Reset()

		bcode := comp.Compile([]lexer.Node{
			{Typ: lexer.TypInt, DatN: 1},
			{Typ: lexer.TypInt, DatN: 2},
			{Typ: lexer.TypInt, DatN: 3},
			{Typ: lexer.TypVector, Cap: 3},
		})
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

		bcode := comp.Compile([]lexer.Node{
			{Typ: lexer.TypInt, DatN: 0xff},
			{Typ: lexer.TypVector, Cap: 1},
			{Typ: lexer.TypStr, DatS: "foo"},
			{Typ: lexer.TypVector, Cap: 2},
		})
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

		bcode := comp.Compile([]lexer.Node{
			{Typ: lexer.TypStr, DatS: "$"},
			{Typ: lexer.TypStr, DatS: "a"},
			{Typ: lexer.TypStr, DatS: "b"},
			{Typ: lexer.TypStr, DatS: "c"},
			{Typ: lexer.TypVector, Cap: 3},
			{Typ: lexer.TypInvoke, DatS: "join", Cap: 2},
		})
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBytes, addr.Type())
		assert.Equal(t, []byte("a$b$c"), addr.Bytes())

		assert.Equal(t, "join(bytes=\"$\", vector=[bytes=\"a\", bytes=\"b\", bytes=\"c\"]) -> bytes=\"a$b$c\"\n", tracing.String())
	})

	t.Run("func concat", func(t *testing.T) {
		tracing.Reset()

		bcode := comp.Compile([]lexer.Node{
			{Typ: lexer.TypStr, DatS: "foo"},
			{Typ: lexer.TypStr, DatS: "bar"},
			{Typ: lexer.TypStr, DatS: "baz"},
			{Typ: lexer.TypInvoke, DatS: "concat", Cap: 3},
		})
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBytes, addr.Type())
		assert.Equal(t, []byte("foobarbaz"), addr.Bytes())

		assert.Equal(t, "concat(bytes=\"foo\", bytes=\"bar\", bytes=\"baz\") -> bytes=\"foobarbaz\"\n", tracing.String())
	})

	t.Run("func equals int", func(t *testing.T) {
		tracing.Reset()

		bcode := comp.Compile([]lexer.Node{
			{Typ: lexer.TypInt, DatN: 1},
			{Typ: lexer.TypInt, DatN: 1},
			{Typ: lexer.TypInvoke, DatS: "equals", Cap: 2},
		})
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBool, addr.Type())
		assert.True(t, addr.Bool())

		assert.Equal(t, "equals(int64=1, int64=1) -> bool=1\n", tracing.String())
	})

	t.Run("func equals str", func(t *testing.T) {
		tracing.Reset()

		bcode := comp.Compile([]lexer.Node{
			{Typ: lexer.TypStr, DatS: "foo"},
			{Typ: lexer.TypStr, DatS: "foo"},
			{Typ: lexer.TypInvoke, DatS: "equals", Cap: 2},
		})
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBool, addr.Type())
		assert.True(t, addr.Bool())

		assert.Equal(t, "equals(bytes=\"foo\", bytes=\"foo\") -> bool=1\n", tracing.String())
	})

	t.Run("func equals vector", func(t *testing.T) {
		tracing.Reset()

		bcode := comp.Compile([]lexer.Node{
			{Typ: lexer.TypStr, DatS: "foo"},
			{Typ: lexer.TypStr, DatS: "bar"},
			{Typ: lexer.TypVector, Cap: 2},
			{Typ: lexer.TypStr, DatS: "foo"},
			{Typ: lexer.TypStr, DatS: "bar"},
			{Typ: lexer.TypVector, Cap: 2},
			{Typ: lexer.TypInvoke, DatS: "equals", Cap: 2},
		})
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBool, addr.Type())
		assert.True(t, addr.Bool())

		assert.Equal(t, "equals(vector=[bytes=\"foo\", bytes=\"bar\"], vector=[bytes=\"foo\", bytes=\"bar\"]) -> bool=1\n", tracing.String())
	})

	t.Run("func equals concat", func(t *testing.T) {
		tracing.Reset()

		bcode := comp.Compile([]lexer.Node{
			{Typ: lexer.TypStr, DatS: "foobarbaz"},
			{Typ: lexer.TypStr, DatS: "foo"},
			{Typ: lexer.TypStr, DatS: "bar"},
			{Typ: lexer.TypStr, DatS: "baz"},
			{Typ: lexer.TypInvoke, DatS: "concat", Cap: 3},
			{Typ: lexer.TypInvoke, DatS: "equals", Cap: 2},
		})
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBool, addr.Type())
		assert.True(t, addr.Bool())

		assert.Equal(t, "concat(bytes=\"foo\", bytes=\"bar\", bytes=\"baz\") -> bytes=\"foobarbaz\"\nequals(bytes=\"foobarbaz\", bytes=\"foobarbaz\") -> bool=1\n", tracing.String())
	})

	t.Run("func contains", func(t *testing.T) {
		tracing.Reset()

		bcode := comp.Compile([]lexer.Node{
			{Typ: lexer.TypStr, DatS: "a"},
			{Typ: lexer.TypStr, DatS: "b"},
			{Typ: lexer.TypStr, DatS: "c"},
			{Typ: lexer.TypVector, Cap: 3},
			{Typ: lexer.TypStr, DatS: "a"},
			{Typ: lexer.TypInvoke, DatS: "contains", Cap: 2},
		})
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBool, addr.Type())
		assert.True(t, addr.Bool())

		assert.Equal(t, "contains(vector=[bytes=\"a\", bytes=\"b\", bytes=\"c\"], bytes=\"a\") -> bool=1\n", tracing.String())
	})

	t.Run("func intersects", func(t *testing.T) {
		tracing.Reset()

		bcode := comp.Compile([]lexer.Node{
			{Typ: lexer.TypInt, DatN: 1},
			{Typ: lexer.TypInt, DatN: 2},
			{Typ: lexer.TypVector, Cap: 2},
			{Typ: lexer.TypInt, DatN: 2},
			{Typ: lexer.TypVector, Cap: 1},
			{Typ: lexer.TypInvoke, DatS: "intersects", Cap: 2},
		})
		addr, err := exec.Exec(bcode)
		require.NoError(t, err)
		require.NotNil(t, addr)

		require.Equal(t, memory.TypeBool, addr.Type())
		assert.True(t, addr.Bool())

		assert.Equal(t, "intersects(vector=[int64=1, int64=2], vector=[int64=2]) -> bool=1\n", tracing.String())
	})

	t.Run("delegator not exists", func(t *testing.T) {
		tracing.Reset()

		bcode := comp.Compile([]lexer.Node{
			{Typ: lexer.TypInvoke, DatS: "NONAME"},
		})
		addr, err := exec.Exec(bcode)
		require.EqualError(t, err, "failed to exec frame at 0: delegator <NONAME> not exists")
		assert.Equal(t, memory.Nil, addr)
	})

	t.Run("math expected int", func(t *testing.T) {
		bcode := comp.Compile([]lexer.Node{
			{Typ: lexer.TypStr, DatS: "foo"},
			{Typ: lexer.TypInt, DatN: 1},
			{Typ: lexer.TypOpAdd},
		})
		addr, err := exec.Exec(bcode)
		require.Equal(t, memory.Nil, addr)
		assert.EqualError(t, err, "failed to exec frame at 2: unexpected type bytes instead of int64")
	})

	t.Run("math no args", func(t *testing.T) {
		for _, typ := range [...]lexer.Typ{lexer.TypOpAdd, lexer.TypOpSub, lexer.TypOpMul, lexer.TypOpDiv} {
			t.Run(typ.String(), func(t *testing.T) {
				bcode := comp.Compile([]lexer.Node{
					{Typ: typ},
				})
				addr, err := exec.Exec(bcode)
				require.Equal(t, memory.Nil, addr)
				assert.EqualError(t, err, "failed to exec frame at 0: stack tail shorter than 2")
			})
		}
	})

	t.Run("math", func(t *testing.T) {
		cases := []struct {
			name     string
			nodes    []lexer.Node
			expected int64
		}{
			{
				"1+2",
				[]lexer.Node{
					{Typ: lexer.TypInt, DatN: 1},
					{Typ: lexer.TypInt, DatN: 2},
					{Typ: lexer.TypOpAdd},
				},
				3,
			},
			{
				"2-3",
				[]lexer.Node{
					{Typ: lexer.TypInt, DatN: 2},
					{Typ: lexer.TypInt, DatN: 3},
					{Typ: lexer.TypOpSub},
				},
				-1,
			},
			{
				"2*3",
				[]lexer.Node{
					{Typ: lexer.TypInt, DatN: 2},
					{Typ: lexer.TypInt, DatN: 3},
					{Typ: lexer.TypOpMul},
				},
				6,
			},
			{
				"6/2",
				[]lexer.Node{
					{Typ: lexer.TypInt, DatN: 6},
					{Typ: lexer.TypInt, DatN: 2},
					{Typ: lexer.TypOpDiv},
				},
				3,
			},
			{
				"1+2*3",
				[]lexer.Node{
					{Typ: lexer.TypInt, DatN: 1},
					{Typ: lexer.TypInt, DatN: 2},
					{Typ: lexer.TypInt, DatN: 3},
					{Typ: lexer.TypOpMul},
					{Typ: lexer.TypOpAdd},
				},
				7,
			},
			{
				"1*2+3",
				[]lexer.Node{
					{Typ: lexer.TypInt, DatN: 1},
					{Typ: lexer.TypInt, DatN: 2},
					{Typ: lexer.TypOpMul},
					{Typ: lexer.TypInt, DatN: 3},
					{Typ: lexer.TypOpAdd},
				},
				5,
			},
		}
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				bcode := comp.Compile(c.nodes)
				addr, err := exec.Exec(bcode)
				require.NoError(t, err)
				require.NotNil(t, addr)

				require.Equal(t, memory.TypeInt64, addr.Type())
				assert.Equal(t, c.expected, addr.Int64())
			})
		}
	})

	t.Run("math div by zero", func(t *testing.T) {
		bcode := comp.Compile([]lexer.Node{
			{Typ: lexer.TypInt, DatN: 6},
			{Typ: lexer.TypInt, DatN: 0},
			{Typ: lexer.TypOpDiv},
		})
		addr, err := exec.Exec(bcode)
		require.Equal(t, memory.Nil, addr)
		require.EqualError(t, err, "failed to exec frame at 2: division by zero")
	})
}
