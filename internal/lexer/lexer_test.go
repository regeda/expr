package lexer_test

import (
	"testing"

	"github.com/regeda/expr/internal/ast"
	"github.com/regeda/expr/internal/lexer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLexer_Parse(t *testing.T) {
	cases := []struct {
		expr, expected string
	}{
		{"true", "nested:<token:BOOL b:true > "},
		{"false", "nested:<token:BOOL b:false > "},
		{`""`, "nested:<token:STR s:\"\" > "},
		{`"foo"`, "nested:<token:STR s:\"foo\" > "},
		{`"\"a\" \"b\" \"c\""`, `nested:<token:STR s:"\"a\" \"b\" \"c\"" > `},
		{"foo()", "nested:<token:CALL s:\"foo\" > "},
		{"[1,2]", "nested:<token:ARR nested:<token:INT i:1 > nested:<token:INT i:2 > > "},
		{"foo([1])", "nested:<token:CALL s:\"foo\" nested:<token:ARR nested:<token:INT i:1 > > > "},
		{"foo(bar(true))", "nested:<token:CALL s:\"foo\" nested:<token:CALL s:\"bar\" nested:<token:BOOL b:true > > > "},
		{`foo("bar", baz(1))`, "nested:<token:CALL s:\"foo\" nested:<token:STR s:\"bar\" > nested:<token:CALL s:\"baz\" nested:<token:INT i:1 > > > "},
		{`foo ( "bar" , baz ( 1 ) )`, "nested:<token:CALL s:\"foo\" nested:<token:STR s:\"bar\" > nested:<token:CALL s:\"baz\" nested:<token:INT i:1 > > > "},
		{`foo ( baz ( 1 ) , "bar" )`, "nested:<token:CALL s:\"foo\" nested:<token:CALL s:\"baz\" nested:<token:INT i:1 > > nested:<token:STR s:\"bar\" > > "},
		{"foo ( baz ( 1 ) , bar ( true ) )", "nested:<token:CALL s:\"foo\" nested:<token:CALL s:\"baz\" nested:<token:INT i:1 > > nested:<token:CALL s:\"bar\" nested:<token:BOOL b:true > > > "},
		{"footrue()", "nested:<token:CALL s:\"footrue\" > "},
		{"falsefoo(true)", "nested:<token:CALL s:\"falsefoo\" nested:<token:BOOL b:true > > "},
		{"foo([])", "nested:<token:CALL s:\"foo\" nested:<token:ARR > > "},
		{"foo([1,2,3])", "nested:<token:CALL s:\"foo\" nested:<token:ARR nested:<token:INT i:1 > nested:<token:INT i:2 > nested:<token:INT i:3 > > > "},
		{`foo ( [ 1 ] , [ "bar" , true ] )`, "nested:<token:CALL s:\"foo\" nested:<token:ARR nested:<token:INT i:1 > > nested:<token:ARR nested:<token:STR s:\"bar\" > nested:<token:BOOL b:true > > > "},
		{"foo([1, 2, 3], [3, 1])", "nested:<token:CALL s:\"foo\" nested:<token:ARR nested:<token:INT i:1 > nested:<token:INT i:2 > nested:<token:INT i:3 > > nested:<token:ARR nested:<token:INT i:3 > nested:<token:INT i:1 > > > "},
	}

	l := lexer.New()

	for _, c := range cases {
		t.Run(c.expr, func(t *testing.T) {
			node, err := l.Parse([]byte(c.expr))
			require.NoError(t, err)

			assert.Equal(t, ast.Node_EXIT, node.Token)
			assert.Equal(t, c.expected, node.String())
		})
	}
}

func TestLexer_CouldNotParse(t *testing.T) {
	cases := []struct {
		expr, err string
	}{
		{"", "token parse error: "},
		{" ", "token parse error:  "},
		{"tru", "token parse error: tru"},
		{"1.1", "token parse error: 1.1"},
		{"true()", "token parse error: true()"},
		{"1()", "token parse error: 1()"},
		{`"foo"()`, `token parse error: "foo"()`},
		{"foo(", "stack parse error: foo("},
		{"[1,2", "token parse error: [1,2"},
		{`"foo`, `token parse error: "foo`},
		{`"foo\"`, `token parse error: "foo\"`},
		{"foo(bar(1)", "stack parse error: foo(bar(1)"},
		{"[foo()]", "token parse error: [foo()]"},
		{"[1 1]", "token parse error: [1 1]"},
	}

	l := lexer.New()

	for _, c := range cases {
		t.Run(c.expr, func(t *testing.T) {
			node, err := l.Parse([]byte(c.expr))
			require.EqualError(t, err, c.err)
			assert.Nil(t, node)
		})
	}
}
