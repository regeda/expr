package asm

import (
	"github.com/regeda/expr/ast"
	"github.com/regeda/expr/compiler"
	"github.com/regeda/expr/tokenz"
)

type ASM struct {
	comp compiler.Compiler
	tkz  tokenz.Tokenz
	astb ast.Builder
}

func (a *ASM) Assemble(code []byte) ([]byte, error) {
	tokens, err := a.tkz.Parse(code)
	if err != nil {
		return nil, err
	}
	ast, err := a.astb.Build(tokens)
	if err != nil {
		return nil, err
	}
	return a.comp.Compile(ast), nil
}
