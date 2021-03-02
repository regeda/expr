package asm

import (
	"github.com/regeda/expr/internal/compiler"
	"github.com/regeda/expr/internal/lexer"
)

type ASM struct {
	lex  *lexer.Lexer
	comp *compiler.Compiler
}

func New() *ASM {
	return &ASM{
		lex:  lexer.New(),
		comp: compiler.New(),
	}
}

func (a *ASM) Assemble(code []byte) ([]byte, error) {
	ast, err := a.lex.Parse(code)
	if err != nil {
		return nil, err
	}
	return a.comp.Compile(ast), nil
}
