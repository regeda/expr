package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/regeda/expr/compiler"
	"github.com/regeda/expr/delegate"
	"github.com/regeda/expr/exec"
	"github.com/regeda/expr/lexer"
	"github.com/regeda/expr/stdlib"
)

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	tokens, err := lexer.Parse(bytes)
	if err != nil {
		panic(err)
	}

	bcode := compiler.Compile(tokens)

	ex := exec.New(
		exec.WithRegistry(delegate.Import(stdlib.Compare, stdlib.Strings)),
	)

	addr, err := ex.Exec(bcode)
	if err != nil {
		log.Fatal(err)
	}

	addr.Print(os.Stdout)
	fmt.Fprintln(os.Stdout)
}
