package expr_test

import (
	"testing"

	"github.com/regeda/expr/ast"
	"github.com/regeda/expr/compiler"
	"github.com/regeda/expr/delegate"
	"github.com/regeda/expr/exec"
	"github.com/regeda/expr/memory"
	"github.com/regeda/expr/stdlib"
)

func BenchmarkExec(b *testing.B) {
	var comp compiler.Compiler

	vm := exec.New(
		exec.WithRegistry(delegate.Import(stdlib.Compare, stdlib.Strings)),
		exec.WithStackSize(0xff),
		exec.WithMemory(
			memory.New(
				memory.PreallocHeap(0xff),
				memory.PreallocGrid(0xff),
			),
		),
	)

	bcode := comp.Compile(
		ast.Exit().Nest(
			ast.Call("equals").Nest(
				ast.Str("foo,bar,baz"),
				ast.Call("join").Nest(
					ast.Str(","),
					ast.Arr().Nest(
						ast.Str("foo"),
						ast.Str("bar"),
						ast.Str("baz"),
					),
				),
			),
		),
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := vm.Exec(bcode)
		if err != nil {
			b.FailNow()
		}
	}
}
