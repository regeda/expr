package exec

import (
	"fmt"

	"github.com/regeda/expr/memory"
)

const (
	stacklimit = 0xff
)

type stack struct {
	p   uint32
	seq [stacklimit]*memory.Addr
}

func (s *stack) reset() {
	s.p = 0
}

func (s *stack) push(a *memory.Addr) error {
	if s.p == uint32(len(s.seq)) {
		return errStackOverflow
	}
	s.seq[s.p] = a
	s.p++
	return nil
}

func (s *stack) tail(n uint32) bool {
	return s.p-n >= 0
}

func (s *stack) pop(n uint32) ([]*memory.Addr, error) {
	if !s.tail(n) {
		return nil, fmt.Errorf("stack tail shorter than %d", n)
	}
	offset := s.p - n
	a := s.seq[offset:s.p]
	s.p = offset
	return a, nil
}
