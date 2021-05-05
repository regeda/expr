package memory_test

import (
	"testing"

	"github.com/regeda/expr/memory"
	"github.com/stretchr/testify/assert"
)

func Test_Type_String(t *testing.T) {
	unknown := -1

	cases := []struct {
		typ      memory.Type
		expected string
	}{
		{memory.TypeBytes, "bytes"},
		{memory.TypeInt64, "int64"},
		{memory.TypeBool, "bool"},
		{memory.TypeVector, "vector"},
		{memory.Type(unknown), "unknown"},
	}

	for _, c := range cases {
		t.Run(c.expected, func(t *testing.T) {
			assert.Equal(t, c.expected, c.typ.String())
		})
	}
}
