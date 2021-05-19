package stdlib_test

import (
	"testing"

	"github.com/regeda/expr/memory"
	"github.com/regeda/expr/stdlib"
	"github.com/stretchr/testify/require"
)

func TestStrings_Error(t *testing.T) {
	mem := memory.New()

	cases := map[string][]struct {
		name   string
		argv   []memory.Addr
		errMsg string
	}{
		"concat": {
			{
				"expected bytes",
				[]memory.Addr{memory.Nil},
				"concat: expected bytes",
			},
		},
		"join": {
			{
				"wrong args num",
				[]memory.Addr{},
				"join: expected 2 args",
			},
			{
				"expected bytes at 0",
				[]memory.Addr{memory.Nil, memory.Nil},
				"join: expected bytes at 0",
			},
			{
				"expected array at 0",
				[]memory.Addr{memory.NoBytes, memory.Nil},
				"join: expected array of bytes at 1",
			},
			{
				"expected array of bytes at 0",
				[]memory.Addr{memory.NoBytes, mem.CopyVector(memory.Nil)},
				"join: expected array of bytes at 1",
			},
		},
	}

	for dname, unit := range cases {
		for _, c := range unit {
			t.Run(dname+"_"+c.name, func(t *testing.T) {
				addr, err := stdlib.Strings[dname].Delegate(&mem, c.argv)

				require.Equal(t, memory.Nil, addr)
				require.EqualError(t, err, c.errMsg)
			})
		}
	}
}
