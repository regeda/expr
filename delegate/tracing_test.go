package delegate_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/regeda/expr/delegate"
	"github.com/regeda/expr/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTracing(t *testing.T) {
	mem := memory.New()

	t.Run("result", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)

		f := delegate.DelegatorFunc(func(mem *memory.Memory, argv []memory.Addr) (memory.Addr, error) {
			return mem.AllocInt64(argv[0].Int64() + argv[1].Int64()), nil
		})

		d := delegate.WithTracing("sum", f, buf)

		_, err := d.Delegate(&mem, []memory.Addr{mem.AllocInt64(2), mem.AllocInt64(3)})
		require.NoError(t, err)

		assert.Equal(t, "sum(int64=2, int64=3) -> int64=5\n", buf.String())
	})

	t.Run("error", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)

		f := delegate.DelegatorFunc(func(*memory.Memory, []memory.Addr) (memory.Addr, error) {
			return memory.Nil, errors.New("die!")
		})

		d := delegate.WithTracing("err", f, buf)

		_, err := d.Delegate(&mem, []memory.Addr{})
		require.Error(t, err)

		assert.Equal(t, "err() FATAL die!\n", buf.String())
	})
}
