package stack

import (
	"github.com/regeda/expr/internal/ast"
)

type Stack []*ast.Node

func (s *Stack) Reset() {
	*s = (*s)[:0]
}

func (s *Stack) Empty() bool {
	return s.Len() == 0
}

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) Top() *ast.Node {
	return s[len(s)-1]
}

func (s *Stack) Push(n *ast.Node) {
	*s = append(*s, n)
}

func (s *Stack) PopN(n int) []*ast.Node {
	l := len(*s) - n
	tail := (*s)[l:]
	*s = (*s)[:l]
	return tail
}

func (s *Stack) Pop() *ast.Node {
	return s.PopN(1)[0]
}

func (s *Stack) Nest(n *ast.Node) *ast.Node {
	t := s.Top()
	t.Nested = append(t.Nested, n)
	return n
}
