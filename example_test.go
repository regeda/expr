package expr_test

import (
	"fmt"

	"github.com/regeda/expr/compiler"
	"github.com/regeda/expr/delegate"
	"github.com/regeda/expr/exec"
	"github.com/regeda/expr/lexer"
	"github.com/regeda/expr/stdlib"
)

func ExampleExpr() {
	code := `join(",", ["a", "b"])`

	tokens, err := lexer.Parse([]byte(code))
	if err != nil {
		panic(err)
	}

	bytecode := compiler.Compile(tokens)

	ex := exec.New(
		exec.WithRegistry(delegate.Import(stdlib.Compare, stdlib.Strings)),
	)
	addr, err := ex.Exec(bytecode)
	if err != nil {
		panic(err)
	}

	fmt.Println(addr.Type(), addr.Bytes())
}
