package delegate

import (
	"fmt"
	"io"

	"github.com/regeda/expr/memory"
)

type Tracing struct {
	k string
	d Delegator
	w io.Writer
}

func WithTracing(key string, d Delegator, w io.Writer) Tracing {
	return Tracing{
		k: key,
		d: d,
		w: w,
	}
}

func (t Tracing) write(f string, argv ...interface{}) {
	fmt.Fprintf(t.w, f, argv...)
}

func (t Tracing) Delegate(mem *memory.Memory, argv []memory.Addr) (memory.Addr, error) {
	t.write("%s(", t.k)
	for i, a := range argv {
		if i > 0 {
			t.write(", ")
		}
		a.Print(t.w)
	}
	addr, err := t.d.Delegate(mem, argv)
	if err != nil {
		t.write(") FATAL %v\n", err)
		return addr, err
	}
	t.write(") -> ")
	addr.Print(t.w)
	t.write("\n")
	return addr, nil
}
