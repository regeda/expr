package expr_test

import (
	"fmt"

	"github.com/regeda/expr/asm"
	"github.com/regeda/expr/delegate"
	"github.com/regeda/expr/exec"
	"github.com/regeda/expr/stdlib"
)

func ExampleExpr() {
	code := `join(",", ["a", "b"])`

	var a asm.ASM
	bytecode, err := a.Assemble([]byte(code))
	if err != nil {
		panic(err)
	}

	ex := exec.New(
		exec.WithRegistry(delegate.Import(stdlib.Compare, stdlib.Strings)),
	)
	addr, err := ex.Exec(bytecode)
	if err != nil {
		panic(err)
	}

	fmt.Println(addr.Type(), addr.Bytes())
}
