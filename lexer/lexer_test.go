package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLexer_Parse(t *testing.T) {
	cases := []struct {
		input    string
		expected []Node
	}{
		{
			"true",
			[]Node{
				{Typ: TypTrue},
			},
		},
		{
			"false",
			[]Node{
				{Typ: TypFalse},
			},
		},
		{
			`"str"`,
			[]Node{
				{Typ: TypStr, DatS: "str"},
			},
		},
		{
			`"foo\"bar"`,
			[]Node{
				{Typ: TypStr, DatS: `foo"bar`},
			},
		},
		{
			"000123",
			[]Node{
				{Typ: TypInt, DatN: 123},
			},
		},
		{
			"-123",
			[]Node{
				{Typ: TypInt, DatN: -123},
			},
		},
		{
			"+123",
			[]Node{
				{Typ: TypInt, DatN: 123},
			},
		},
		{
			"[]",
			[]Node{
				{Typ: TypVector},
			},
		},
		{
			"[ ]",
			[]Node{
				{Typ: TypVector},
			},
		},
		{
			" [ ]",
			[]Node{
				{Typ: TypVector},
			},
		},
		{
			" [ ] ",
			[]Node{
				{Typ: TypVector},
			},
		},
		{
			"[ foo () ]",
			[]Node{
				{Typ: TypInvoke, DatS: "foo"},
				{Typ: TypVector, Cap: 1},
			},
		},
		{
			"[[]]",
			[]Node{
				{Typ: TypVector},
				{Typ: TypVector, Cap: 1},
			},
		},
		{
			`[255]`,
			[]Node{
				{Typ: TypInt, DatN: 255},
				{Typ: TypVector, Cap: 1},
			},
		},
		{
			"[1,2]",
			[]Node{
				{Typ: TypInt, DatN: 1},
				{Typ: TypInt, DatN: 2},
				{Typ: TypVector, Cap: 2},
			},
		},
		{
			"[1 * 2]",
			[]Node{
				{Typ: TypInt, DatN: 1},
				{Typ: TypInt, DatN: 2},
				{Typ: TypOpMul},
				{Typ: TypVector, Cap: 1},
			},
		},
		{
			"[1 + 2, 3 * 4]",
			[]Node{
				{Typ: TypInt, DatN: 1},
				{Typ: TypInt, DatN: 2},
				{Typ: TypOpAdd},
				{Typ: TypInt, DatN: 3},
				{Typ: TypInt, DatN: 4},
				{Typ: TypOpMul},
				{Typ: TypVector, Cap: 2},
			},
		},
		{
			"foo()",
			[]Node{
				{Typ: TypInvoke, DatS: "foo"},
			},
		},
		{
			"foo ()",
			[]Node{
				{Typ: TypInvoke, DatS: "foo"},
			},
		},
		{
			"foo ( )",
			[]Node{
				{Typ: TypInvoke, DatS: "foo"},
			},
		},
		{
			"foo(bar())",
			[]Node{
				{Typ: TypInvoke, DatS: "bar"},
				{Typ: TypInvoke, DatS: "foo", Cap: 1},
			},
		},
		{
			"foo(123)",
			[]Node{
				{Typ: TypInt, DatN: 123},
				{Typ: TypInvoke, DatS: "foo", Cap: 1},
			},
		},
		{
			"foo([1])",
			[]Node{
				{Typ: TypInt, DatN: 1},
				{Typ: TypVector, Cap: 1},
				{Typ: TypInvoke, DatS: "foo", Cap: 1},
			},
		},
		{
			"foo(1, 2)",
			[]Node{
				{Typ: TypInt, DatN: 1},
				{Typ: TypInt, DatN: 2},
				{Typ: TypInvoke, DatS: "foo", Cap: 2},
			},
		},
		{
			"1+1",
			[]Node{
				{Typ: TypInt, DatN: 1},
				{Typ: TypInt, DatN: 1},
				{Typ: TypOpAdd},
			},
		},
		{
			"1+-1",
			[]Node{
				{Typ: TypInt, DatN: 1},
				{Typ: TypInt, DatN: -1},
				{Typ: TypOpAdd},
			},
		},
		{
			"1-1",
			[]Node{
				{Typ: TypInt, DatN: 1},
				{Typ: TypInt, DatN: 1},
				{Typ: TypOpSub},
			},
		},
		{
			"1-+1",
			[]Node{
				{Typ: TypInt, DatN: 1},
				{Typ: TypInt, DatN: 1},
				{Typ: TypOpSub},
			},
		},
		{
			"1*1",
			[]Node{
				{Typ: TypInt, DatN: 1},
				{Typ: TypInt, DatN: 1},
				{Typ: TypOpMul},
			},
		},
		{
			"1/1",
			[]Node{
				{Typ: TypInt, DatN: 1},
				{Typ: TypInt, DatN: 1},
				{Typ: TypOpDiv},
			},
		},
		{
			"1 + 2 * 3",
			[]Node{
				{Typ: TypInt, DatN: 1},
				{Typ: TypInt, DatN: 2},
				{Typ: TypInt, DatN: 3},
				{Typ: TypOpMul},
				{Typ: TypOpAdd},
			},
		},
		{
			"(1 + 2) * 3",
			[]Node{
				{Typ: TypInt, DatN: 1},
				{Typ: TypInt, DatN: 2},
				{Typ: TypOpAdd},
				{Typ: TypInt, DatN: 3},
				{Typ: TypOpMul},
			},
		},
		{
			" ( 1 ) ",
			[]Node{
				{Typ: TypInt, DatN: 1},
			},
		},
		{
			" ( ( 1 + 2 ) * 3 ) ",
			[]Node{
				{Typ: TypInt, DatN: 1},
				{Typ: TypInt, DatN: 2},
				{Typ: TypOpAdd},
				{Typ: TypInt, DatN: 3},
				{Typ: TypOpMul},
			},
		},
		{
			"1 * 2 + 3",
			[]Node{
				{Typ: TypInt, DatN: 1},
				{Typ: TypInt, DatN: 2},
				{Typ: TypOpMul},
				{Typ: TypInt, DatN: 3},
				{Typ: TypOpAdd},
			},
		},
		{
			"1 * (2 + (-3))",
			[]Node{
				{Typ: TypInt, DatN: 1},
				{Typ: TypInt, DatN: 2},
				{Typ: TypInt, DatN: -3},
				{Typ: TypOpAdd},
				{Typ: TypOpMul},
			},
		},
		{
			"(100)+(200)",
			[]Node{
				{Typ: TypInt, DatN: 100},
				{Typ: TypInt, DatN: 200},
				{Typ: TypOpAdd},
			},
		},
		{
			"foo((1))",
			[]Node{
				{Typ: TypInt, DatN: 1},
				{Typ: TypInvoke, DatS: "foo", Cap: 1},
			},
		},
		{
			"foo(1 + 2)",
			[]Node{
				{Typ: TypInt, DatN: 1},
				{Typ: TypInt, DatN: 2},
				{Typ: TypOpAdd},
				{Typ: TypInvoke, DatS: "foo", Cap: 1},
			},
		},
		{
			"foo(1 + 2 * 3, [true])",
			[]Node{
				{Typ: TypInt, DatN: 1},
				{Typ: TypInt, DatN: 2},
				{Typ: TypInt, DatN: 3},
				{Typ: TypOpMul},
				{Typ: TypOpAdd},
				{Typ: TypTrue},
				{Typ: TypVector, Cap: 1},
				{Typ: TypInvoke, DatS: "foo", Cap: 2},
			},
		},
		{
			"1 + foo(true) * 3",
			[]Node{
				{Typ: TypInt, DatN: 1},
				{Typ: TypTrue},
				{Typ: TypInvoke, DatS: "foo", Cap: 1},
				{Typ: TypInt, DatN: 3},
				{Typ: TypOpMul},
				{Typ: TypOpAdd},
			},
		},
	}

	var lex Lexer

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			tokens, err := lex.Parse([]byte(c.input))

			require.NoError(t, err)
			assert.Equal(t, c.expected, tokens)
		})
	}
}

func TestLexer_ParseFailed(t *testing.T) {
	cases := []struct {
		input, errmsg string
	}{
		{
			"",
			"token parsing error at 0",
		},
		{
			"9223372036854775808",
			"parsing error at 0: strconv.ParseInt: parsing \"9223372036854775808\": value out of range",
		},
		{
			`"hello`,
			"token parsing error at 0",
		},
		{
			`"hello \ world"`,
			`parsing error at 0: strconv.Unquote "hello \ world": invalid syntax`,
		},
		{
			"1 2",
			"token parsing error at 0",
		},
		{
			"tru",
			"token parsing error at 0",
		},
		{
			"(",
			"stack parsing error at 0",
		},
		{
			"()",
			"stack parsing error at 0",
		},
		{
			"func([)",
			"stack parsing error at 0",
		},
		{
			"[func(]",
			"stack parsing error at 1",
		},
		{
			"*1",
			"token parsing error at 0",
		},
		{
			"/1",
			"token parsing error at 0",
		},
		{
			"1*",
			"token parsing error at 1",
		},
		{
			"1/",
			"token parsing error at 1",
		},
		{
			"1+",
			"token parsing error at 1",
		},
		{
			"1-",
			"token parsing error at 1",
		},
		{
			"(1-",
			"stack parsing error at 2",
		},
		{
			"1()",
			"token parsing error at 0",
		},
		{
			"1+++1",
			"token parsing error at 2",
		},
		{
			"1)",
			"token parsing error at 0",
		},
		{
			`1 + ""`,
			"token parsing error at 2",
		},
		{
			"true + false",
			"token parsing error at 0",
		},
		{
			"func(1 2)",
			"stack parsing error at 5",
		},
		{
			"func(1, 2,)",
			"stack parsing error at 8",
		},
		{
			"[1 2]",
			"stack parsing error at 1",
		},
		{
			"[1, 2,]",
			"stack parsing error at 4",
		},
	}

	var lex Lexer

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			tokens, err := lex.Parse([]byte(c.input))

			require.Nil(t, tokens)
			assert.EqualError(t, err, c.errmsg)
		})
	}
}
