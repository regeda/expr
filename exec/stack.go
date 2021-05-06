package exec

import (
	"fmt"

	"github.com/regeda/expr/memory"
)

type stack []memory.Addr

func (s *stack) reset() {
	*s = (*s)[:0]
}

func (s *stack) push(a memory.Addr) {
	*s = append(*s, a)
}

func (s *stack) pop(n uint32) ([]memory.Addr, error) {
	l := uint32(len(*s))
	if n > l {
		return nil, fmt.Errorf("stack tail shorter than %d", n)
	}
	offset := l - n
	a := (*s)[offset:l]
	*s = (*s)[:offset]
	return a, nil
}
