package expr_test

import (
	"testing"

	"github.com/regeda/expr/exec"
	"github.com/regeda/expr/internal/ast/value"
	"github.com/regeda/expr/internal/compiler"
	"github.com/regeda/expr/memory"
	"github.com/regeda/expr/stdlib"
)

func BenchmarkExec(b *testing.B) {
	comp := compiler.New()
	vm := exec.New(stdlib.Registry(),
		exec.WithStackSize(0xff),
		exec.WithMemory(
			memory.New(
				memory.PreallocHeap(0xff),
				memory.PreallocGrid(0xff),
			),
		),
	)

	bcode := comp.Compile(value.Nest(
		value.Exit(),
		value.Nest(
			value.Call("equals"),
			value.Str("foo,bar,baz"),
			value.Nest(
				value.Call("join"),
				value.Str(","),
				value.Nest(
					value.Arr(),
					value.Str("foo"),
					value.Str("bar"),
					value.Str("baz"),
				),
			),
		),
	))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := vm.Exec(bcode)
		if err != nil {
			b.FailNow()
		}
	}
}
