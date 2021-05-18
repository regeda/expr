package delegate

import (
	"github.com/regeda/expr/memory"
)

type Delegator interface {
	Delegate(*memory.Memory, []memory.Addr) (memory.Addr, error)
}

type DelegatorFunc func(*memory.Memory, []memory.Addr) (memory.Addr, error)

func (f DelegatorFunc) Delegate(mem *memory.Memory, argv []memory.Addr) (memory.Addr, error) {
	return f(mem, argv)
}
