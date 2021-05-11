package exec

import (
	"github.com/regeda/expr/delegate"
)

type Globals struct {
	delegators delegate.Registry
}

func (g *Globals) SetRegistry(r delegate.Registry) {
	g.delegators = r
}
