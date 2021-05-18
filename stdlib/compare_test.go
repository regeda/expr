package stdlib_test

import (
	"testing"

	"github.com/regeda/expr/memory"
	"github.com/regeda/expr/stdlib"
	"github.com/stretchr/testify/require"
)

func TestCompare_Error(t *testing.T) {
	mem := memory.New()

	cases := map[string][]struct {
		name   string
		argv   []memory.Addr
		errMsg string
	}{
		"equals": {
			{
				"wrong args num",
				[]memory.Addr{},
				"equals: expected 2 args",
			},
		},
		"contains": {
			{
				"wrong args num",
				[]memory.Addr{},
				"contains: expected 2 args",
			},
			{
				"first arg should be array",
				[]memory.Addr{memory.Nil, memory.Nil},
				"contains: expected array at 0",
			},
			{
				"second arg is scalar",
				[]memory.Addr{mem.AllocVector(1), mem.AllocVector(1)},
				"contains: expected scalar at 1",
			},
		},
		"intersects": {
			{
				"wrong args num",
				[]memory.Addr{},
				"intersects: expected 2 args",
			},
			{
				"expected arrays only",
				[]memory.Addr{memory.Nil, memory.Nil},
				"intersects: expected arrays",
			},
		},
	}

	for dname, unit := range cases {
		for _, c := range unit {
			t.Run(dname+"_"+c.name, func(t *testing.T) {
				addr, err := stdlib.Compare[dname].Delegate(&mem, c.argv)

				require.Equal(t, memory.Nil, addr)
				require.EqualError(t, err, c.errMsg)
			})
		}
	}
}
