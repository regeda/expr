package ast_test

import (
	"testing"

	"github.com/regeda/expr/ast"
	"github.com/regeda/expr/tokenz"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuilder_Build(t *testing.T) {
	cases := []struct {
		name     string
		tokens   []tokenz.Token
		expected *ast.Node
	}{
		{
			"int",
			[]tokenz.Token{
				{Tk: tokenz.TkInt, Dat: []byte("1")},
			},
			ast.Exit().Nest(ast.Int(1)),
		},
		{
			"str",
			[]tokenz.Token{
				{Tk: tokenz.TkStr, Dat: []byte(`"foo"`)},
			},
			ast.Exit().Nest(ast.Str("foo")),
		},
		{
			"true",
			[]tokenz.Token{
				{Tk: tokenz.TkTrue},
			},
			ast.Exit().Nest(ast.True),
		},
		{
			"false",
			[]tokenz.Token{
				{Tk: tokenz.TkFalse},
			},
			ast.Exit().Nest(ast.False),
		},
		{
			"empty arr",
			[]tokenz.Token{
				{Tk: tokenz.TkPunct, Dat: []byte("[")},
				{Tk: tokenz.TkPunct, Dat: []byte("]")},
			},
			ast.Exit().Nest(ast.Arr()),
		},
		{
			"arr with 1 elem",
			[]tokenz.Token{
				{Tk: tokenz.TkPunct, Dat: []byte("[")},
				{Tk: tokenz.TkTrue},
				{Tk: tokenz.TkPunct, Dat: []byte("]")},
			},
			ast.Exit().Nest(
				ast.Arr().Nest(
					ast.True,
				),
			),
		},
		{
			"arr with 2 elems",
			[]tokenz.Token{
				{Tk: tokenz.TkPunct, Dat: []byte("[")},
				{Tk: tokenz.TkTrue},
				{Tk: tokenz.TkPunct, Dat: []byte(",")},
				{Tk: tokenz.TkInt, Dat: []byte("1")},
				{Tk: tokenz.TkPunct, Dat: []byte("]")},
			},
			ast.Exit().Nest(
				ast.Arr().Nest(
					ast.True,
					ast.Int(1),
				),
			),
		},
		{
			"arr of arr",
			[]tokenz.Token{
				{Tk: tokenz.TkPunct, Dat: []byte("[")},
				{Tk: tokenz.TkPunct, Dat: []byte("[")},
				{Tk: tokenz.TkPunct, Dat: []byte("]")},
				{Tk: tokenz.TkPunct, Dat: []byte("]")},
			},
			ast.Exit().Nest(ast.Arr().Nest(ast.Arr())),
		},
		{
			"func no params",
			[]tokenz.Token{
				{Tk: tokenz.TkIdent, Dat: []byte("foo")},
				{Tk: tokenz.TkPunct, Dat: []byte("(")},
				{Tk: tokenz.TkPunct, Dat: []byte(")")},
			},
			ast.Exit().Nest(ast.Call("foo")),
		},
		{
			"func with 1 scalar param",
			[]tokenz.Token{
				{Tk: tokenz.TkIdent, Dat: []byte("foo")},
				{Tk: tokenz.TkPunct, Dat: []byte("(")},
				{Tk: tokenz.TkTrue},
				{Tk: tokenz.TkPunct, Dat: []byte(")")},
			},
			ast.Exit().Nest(
				ast.Call("foo").Nest(ast.True),
			),
		},
		{
			"func with 2 scalar params",
			[]tokenz.Token{
				{Tk: tokenz.TkIdent, Dat: []byte("foo")},
				{Tk: tokenz.TkPunct, Dat: []byte("(")},
				{Tk: tokenz.TkTrue},
				{Tk: tokenz.TkPunct, Dat: []byte(",")},
				{Tk: tokenz.TkInt, Dat: []byte("1")},
				{Tk: tokenz.TkPunct, Dat: []byte(")")},
			},
			ast.Exit().Nest(
				ast.Call("foo").Nest(
					ast.True,
					ast.Int(1),
				),
			),
		},
		{
			"func with func",
			[]tokenz.Token{
				{Tk: tokenz.TkIdent, Dat: []byte("foo")},
				{Tk: tokenz.TkPunct, Dat: []byte("(")},
				{Tk: tokenz.TkIdent, Dat: []byte("bar")},
				{Tk: tokenz.TkPunct, Dat: []byte("(")},
				{Tk: tokenz.TkPunct, Dat: []byte(")")},
				{Tk: tokenz.TkPunct, Dat: []byte(")")},
			},
			ast.Exit().Nest(
				ast.Call("foo").Nest(
					ast.Call("bar"),
				),
			),
		},
	}

	var b ast.Builder

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			n, err := b.Build(c.tokens)

			require.NoError(t, err)
			assert.Equal(t, c.expected, n)
		})
	}
}

func TestBuilder_BuildError(t *testing.T) {
	cases := []struct {
		name   string
		tokens []tokenz.Token
		err    string
	}{
		{
			"ident with no digit prefix",
			[]tokenz.Token{
				{Tk: tokenz.TkInt, Dat: []byte("1")},
				{Tk: tokenz.TkIdent, Dat: []byte("foo")},
			},
			"unexpected ident after int=1",
		},
		{
			"parse int64",
			[]tokenz.Token{
				{Tk: tokenz.TkInt, Dat: []byte("9223372036854775808")},
			},
			"strconv.ParseInt: parsing \"9223372036854775808\": value out of range",
		},
		{
			"int after str",
			[]tokenz.Token{
				{Tk: tokenz.TkStr, Dat: []byte(`"foo"`)},
				{Tk: tokenz.TkInt, Dat: []byte("1")},
			},
			"unexpected integer after str=\"foo\"",
		},
		{
			"unquote str",
			[]tokenz.Token{
				{Tk: tokenz.TkStr, Dat: []byte(`"hello \ world"`)},
			},
			`strconv.Unquote "hello \ world": invalid syntax`,
		},
		{
			"str after str",
			[]tokenz.Token{
				{Tk: tokenz.TkStr, Dat: []byte(`"foo"`)},
				{Tk: tokenz.TkStr, Dat: []byte(`"bar"`)},
			},
			"unexpected string after str=\"foo\"",
		},
		{
			"true after true",
			[]tokenz.Token{
				{Tk: tokenz.TkTrue},
				{Tk: tokenz.TkTrue},
			},
			"unexpected TRUE after true",
		},
		{
			"false after false",
			[]tokenz.Token{
				{Tk: tokenz.TkFalse},
				{Tk: tokenz.TkFalse},
			},
			"unexpected FALSE after false",
		},
		{
			"unknown token",
			[]tokenz.Token{
				{Tk: tokenz.Tk(0xff)},
			},
			"unexpected token _unknown_255",
		},
		{
			"wrong punct len",
			[]tokenz.Token{
				{Tk: tokenz.TkPunct, Dat: []byte("xyz")},
			},
			"The token Punct should contain 1 byte of data",
		},
		{
			"wrong punct",
			[]tokenz.Token{
				{Tk: tokenz.TkPunct, Dat: []byte("X")},
			},
			"unexpected punct X after none",
		},
		{
			"wrong array closing",
			[]tokenz.Token{
				{Tk: tokenz.TkPunct, Dat: []byte("]")},
			},
			"unexpected array closing after none",
		},
		{
			"array after int",
			[]tokenz.Token{
				{Tk: tokenz.TkInt, Dat: []byte("1")},
				{Tk: tokenz.TkPunct, Dat: []byte("[")},
			},
			"unexpected array after int=1",
		},
		{
			"wrong invokation closing",
			[]tokenz.Token{
				{Tk: tokenz.TkPunct, Dat: []byte(")")},
			},
			"unexpected invokation closing after none",
		},
		{
			"invokation after int",
			[]tokenz.Token{
				{Tk: tokenz.TkInt, Dat: []byte("1")},
				{Tk: tokenz.TkPunct, Dat: []byte("(")},
			},
			"unexpected invokation after int=1",
		},
		{
			"wrong comma",
			[]tokenz.Token{
				{Tk: tokenz.TkPunct, Dat: []byte(",")},
			},
			"unexpected comma after none",
		},
		{
			"comma after array open",
			[]tokenz.Token{
				{Tk: tokenz.TkPunct, Dat: []byte("[")},
				{Tk: tokenz.TkPunct, Dat: []byte(",")},
			},
			"unexpected comma after punct=[",
		},
		{
			"comma after ident",
			[]tokenz.Token{
				{Tk: tokenz.TkIdent, Dat: []byte("foo")},
				{Tk: tokenz.TkPunct, Dat: []byte(",")},
			},
			"unexpected comma after ident=foo",
		},
		{
			"comma after int",
			[]tokenz.Token{
				{Tk: tokenz.TkInt, Dat: []byte("1")},
				{Tk: tokenz.TkPunct, Dat: []byte(",")},
			},
			"unexpected comma after int=1",
		},
		{
			"unexpected end of call",
			[]tokenz.Token{
				{Tk: tokenz.TkIdent, Dat: []byte("foo")},
				{Tk: tokenz.TkPunct, Dat: []byte("(")},
			},
			"unexpected stack length 2",
		},
		{
			"unexpected end of array",
			[]tokenz.Token{
				{Tk: tokenz.TkPunct, Dat: []byte("[")},
			},
			"unexpected stack length 2",
		},
		{
			"unexpected array closing",
			[]tokenz.Token{
				{Tk: tokenz.TkPunct, Dat: []byte("[")},
				{Tk: tokenz.TkIdent, Dat: []byte("foo")},
				{Tk: tokenz.TkPunct, Dat: []byte("(")},
				{Tk: tokenz.TkPunct, Dat: []byte("]")},
			},
			"stack error: expected array, got CALL",
		},
		{
			"unexpected func closing",
			[]tokenz.Token{
				{Tk: tokenz.TkIdent, Dat: []byte("foo")},
				{Tk: tokenz.TkPunct, Dat: []byte("(")},
				{Tk: tokenz.TkPunct, Dat: []byte("[")},
				{Tk: tokenz.TkPunct, Dat: []byte(")")},
			},
			"stack error: expected invokation, got ARR",
		},
	}

	var b ast.Builder

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			n, err := b.Build(c.tokens)

			require.Nil(t, n)
			assert.EqualError(t, err, c.err)
		})
	}
}
