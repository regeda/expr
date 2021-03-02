package delegate

import (
	"github.com/regeda/expr/assert"
	"github.com/regeda/expr/memory"
)

type Delegator interface {
	Delegate(*memory.Memory, []*memory.Addr) (*memory.Addr, error)
}

type DelegatorFunc func(*memory.Memory, []*memory.Addr) (*memory.Addr, error)

func (f DelegatorFunc) Delegate(memory *memory.Memory, argv []*memory.Addr) (*memory.Addr, error) {
	return f(memory, argv)
}

func (f DelegatorFunc) Assert(a assert.Asserter) DelegatorFunc {
	return func(memory *memory.Memory, argv []*memory.Addr) (*memory.Addr, error) {
		if err := a.Assert(argv); err != nil {
			return nil, err
		}
		return f(memory, argv)
	}
}
