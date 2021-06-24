package expr_test

import (
	"testing"

	"github.com/regeda/expr/compiler"
	"github.com/regeda/expr/delegate"
	"github.com/regeda/expr/exec"
	"github.com/regeda/expr/lexer"
	"github.com/regeda/expr/memory"
	"github.com/regeda/expr/stdlib"
)

func BenchmarkExec(b *testing.B) {
	var comp compiler.Compiler

	ex := exec.New(
		exec.WithRegistry(delegate.Import(stdlib.Compare, stdlib.Strings)),
		exec.WithStackSize(0xff),
		exec.WithMemory(
			memory.New(
				memory.PreallocHeap(0xff),
				memory.PreallocGrid(0xff),
			),
		),
	)

	bcode := comp.Compile([]lexer.Node{
		{Typ: lexer.TypStr, DatS: "foo,bar,baz"},
		{Typ: lexer.TypStr, DatS: ","},
		{Typ: lexer.TypStr, DatS: "foo"},
		{Typ: lexer.TypStr, DatS: "bar"},
		{Typ: lexer.TypStr, DatS: "baz"},
		{Typ: lexer.TypVector, Cap: 3},
		{Typ: lexer.TypInvoke, DatS: "join", Cap: 2},
		{Typ: lexer.TypInvoke, DatS: "equals", Cap: 2},
	})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := ex.Exec(bcode)
		if err != nil {
			b.FailNow()
		}
	}
}
