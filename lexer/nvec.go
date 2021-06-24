package lexer

type nvec []Node

func (s *nvec) reset() {
	*s = (*s)[:0]
}

func (s *nvec) empty() bool {
	return s.size() == 0
}

func (s *nvec) size() int {
	return len(*s)
}

func (s nvec) top() Node {
	return s[len(s)-1]
}

func (s *nvec) push(t Node) {
	*s = append(*s, t)
}

func (s *nvec) popN(n int) []Node {
	l := len(*s) - n
	tail := (*s)[l:]
	*s = (*s)[:l]
	return tail
}

func (s *nvec) pop() Node {
	return s.popN(1)[0]
}
