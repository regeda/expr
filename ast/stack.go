package ast

type stack []*Node

func (s *stack) reset() {
	*s = (*s)[:0]
}

func (s *stack) empty() bool {
	return s.len() == 0
}

func (s stack) len() int {
	return len(s)
}

func (s stack) top() *Node {
	return s[len(s)-1]
}

func (s *stack) push(n *Node) {
	*s = append(*s, n)
}

func (s *stack) popN(n int) []*Node {
	l := len(*s) - n
	tail := (*s)[l:]
	*s = (*s)[:l]
	return tail
}

func (s *stack) pop() *Node {
	return s.popN(1)[0]
}

func (s *stack) nest(n *Node) *Node {
	s.top().Nest(n)
	return n
}
