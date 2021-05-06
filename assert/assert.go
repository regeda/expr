package assert

import (
	"github.com/pkg/errors"
	"github.com/regeda/expr/memory"
)

var (
	errWrongArgsNumber = errors.New("wrong arguments number")
	errWrongArgType    = errors.New("wrong argument type")
)

type Asserter interface {
	Assert([]memory.Addr) error
}

type AsserterFunc func([]memory.Addr) error

func (a AsserterFunc) Assert(argv []memory.Addr) error {
	return a(argv)
}

func Len(l int) AsserterFunc {
	return func(argv []memory.Addr) error {
		if len(argv) != l {
			return errWrongArgsNumber
		}
		return nil
	}
}

func TypeAt(i int, t memory.Type) AsserterFunc {
	return func(argv []memory.Addr) error {
		if argv[i].Type() != t {
			return errWrongArgType
		}
		return nil
	}
}

func VectorAt(i int, a Asserter) AsserterFunc {
	return func(argv []memory.Addr) error {
		return a.Assert(argv[i].Vector())
	}
}

func Type(t memory.Type) AsserterFunc {
	return func(argv []memory.Addr) error {
		for _, arg := range argv {
			if arg.Type() != t {
				return errWrongArgType
			}
		}
		return nil
	}
}

func Any(a ...Asserter) AsserterFunc {
	return func(argv []memory.Addr) error {
		var lastErr error
		for _, aa := range a {
			err := aa.Assert(argv)
			if err == nil {
				return nil
			}
			lastErr = err
		}
		return errors.Wrap(lastErr, "all assertions failed")
	}
}

func Every(a ...Asserter) AsserterFunc {
	return func(argv []memory.Addr) error {
		for _, aa := range a {
			if err := aa.Assert(argv); err != nil {
				return err
			}
		}
		return nil
	}
}
