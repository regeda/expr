package tokenz_test

import (
	"testing"

	"github.com/regeda/expr/tokenz"
	"github.com/stretchr/testify/assert"
)

func TestTokenz_Parse(t *testing.T) {
	cases := []struct {
		expr     string
		expected []tokenz.Token
	}{
		{
			"true",
			[]tokenz.Token{{tokenz.TkTrue, nil}},
		},
		{
			"false",
			[]tokenz.Token{{tokenz.TkFalse, nil}},
		},
		{
			`""`,
			[]tokenz.Token{{tokenz.TkStr, []byte(`""`)}},
		},
		{
			`"foo"`,
			[]tokenz.Token{{tokenz.TkStr, []byte(`"foo"`)}},
		},
		{
			`"foo bar"`,
			[]tokenz.Token{{tokenz.TkStr, []byte(`"foo bar"`)}},
		},
		{
			`"\\\"foo\""`,
			[]tokenz.Token{{tokenz.TkStr, []byte(`"\\\"foo\""`)}},
		},
		{
			"000123",
			[]tokenz.Token{{tokenz.TkInt, []byte("000123")}},
		},
		{
			"-1",
			[]tokenz.Token{{tokenz.TkInt, []byte("-1")}},
		},
		{
			"[]",
			[]tokenz.Token{
				{tokenz.TkPunct, []byte{'['}},
				{tokenz.TkPunct, []byte{']'}},
			},
		},
		{
			"[ ]",
			[]tokenz.Token{
				{tokenz.TkPunct, []byte{'['}},
				{tokenz.TkPunct, []byte{']'}},
			},
		},
		{
			"[1,2]",
			[]tokenz.Token{
				{tokenz.TkPunct, []byte{'['}},
				{tokenz.TkInt, []byte("1")},
				{tokenz.TkPunct, []byte{','}},
				{tokenz.TkInt, []byte("2")},
				{tokenz.TkPunct, []byte{']'}},
			},
		},
		{
			"foo",
			[]tokenz.Token{
				{tokenz.TkIdent, []byte("foo")},
			},
		},
		{
			"foo()",
			[]tokenz.Token{
				{tokenz.TkIdent, []byte("foo")},
				{tokenz.TkPunct, []byte{'('}},
				{tokenz.TkPunct, []byte{')'}},
			},
		},
		{
			"foo ()",
			[]tokenz.Token{
				{tokenz.TkIdent, []byte("foo")},
				{tokenz.TkPunct, []byte{'('}},
				{tokenz.TkPunct, []byte{')'}},
			},
		},
		{
			"foo (  )",
			[]tokenz.Token{
				{tokenz.TkIdent, []byte("foo")},
				{tokenz.TkPunct, []byte{'('}},
				{tokenz.TkPunct, []byte{')'}},
			},
		},
		{
			"foo(1)",
			[]tokenz.Token{
				{tokenz.TkIdent, []byte("foo")},
				{tokenz.TkPunct, []byte{'('}},
				{tokenz.TkInt, []byte("1")},
				{tokenz.TkPunct, []byte{')'}},
			},
		},
		{
			"foo( 1 )",
			[]tokenz.Token{
				{tokenz.TkIdent, []byte("foo")},
				{tokenz.TkPunct, []byte{'('}},
				{tokenz.TkInt, []byte("1")},
				{tokenz.TkPunct, []byte{')'}},
			},
		},
		{
			"foo ( 1 )",
			[]tokenz.Token{
				{tokenz.TkIdent, []byte("foo")},
				{tokenz.TkPunct, []byte{'('}},
				{tokenz.TkInt, []byte("1")},
				{tokenz.TkPunct, []byte{')'}},
			},
		},
		{
			"foo(1, 2)",
			[]tokenz.Token{
				{tokenz.TkIdent, []byte("foo")},
				{tokenz.TkPunct, []byte{'('}},
				{tokenz.TkInt, []byte("1")},
				{tokenz.TkPunct, []byte{','}},
				{tokenz.TkInt, []byte("2")},
				{tokenz.TkPunct, []byte{')'}},
			},
		},
		{
			"foo( 1 , 2 )",
			[]tokenz.Token{
				{tokenz.TkIdent, []byte("foo")},
				{tokenz.TkPunct, []byte{'('}},
				{tokenz.TkInt, []byte("1")},
				{tokenz.TkPunct, []byte{','}},
				{tokenz.TkInt, []byte("2")},
				{tokenz.TkPunct, []byte{')'}},
			},
		},
	}

	var tkz tokenz.Tokenz

	for _, c := range cases {
		t.Run(c.expr, func(t *testing.T) {
			tokens, err := tkz.Parse([]byte(c.expr))
			assert.NoError(t, err)
			assert.Equal(t, c.expected, tokens)
		})
	}
}

func TestTokenz_CouldNotParse(t *testing.T) {
	cases := []struct {
		expr, err string
	}{
		{`"foo`, "token parsing error at 0"},
		{`foo"`, "token parsing error at 3"},
		{"x{", "token parsing error at 1"},
	}

	var tkz tokenz.Tokenz

	for _, c := range cases {
		t.Run(c.expr, func(t *testing.T) {
			tokens, err := tkz.Parse([]byte(c.expr))
			assert.EqualError(t, err, c.err)
			assert.Empty(t, tokens)
		})
	}
}
