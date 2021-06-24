package exec

import (
	"fmt"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/pkg/errors"
	"github.com/regeda/expr/bytecode"
	"github.com/regeda/expr/compiler"
	"github.com/regeda/expr/delegate"
	"github.com/regeda/expr/memory"
)

type VM struct {
	globals Globals
	memory  memory.Memory
	stack   stack
	prog    bytecode.Program
}

type Opt func(*VM)

func WithMemory(m memory.Memory) Opt {
	return func(vm *VM) {
		vm.memory = m
	}
}

func WithStackSize(n uint32) Opt {
	return func(vm *VM) {
		vm.stack = make(stack, 0, n)
	}
}

func WithRegistry(reg delegate.Registry) Opt {
	return func(vm *VM) {
		vm.globals.SetRegistry(reg)
	}
}

func New(opts ...Opt) *VM {
	vm := &VM{}

	for _, opt := range opts {
		opt(vm)
	}

	return vm
}

func (v *VM) reset() {
	v.memory.Reset()
	v.stack.reset()
}

func (v *VM) checkVersion() error {
	var ver bytecode.Version

	if v.prog.Ver(&ver) == nil {
		return errNoVersion
	}

	majorVer := ver.Major()

	if majorVer > compiler.MajorVersion {
		return fmt.Errorf("bytecode version %d is not supported, compiler version %d", majorVer, compiler.MajorVersion)
	}

	return nil
}

func (v *VM) terminate() (memory.Addr, error) {
	val, err := v.stack.pop(1)
	if err != nil {
		return memory.Nil, err
	}
	return val[0], nil
}

func (v *VM) Exec(bcode []byte) (memory.Addr, error) {
	v.reset()

	v.prog.Init(bcode, flatbuffers.GetUOffsetT(bcode))

	if err := v.checkVersion(); err != nil {
		return memory.Nil, err
	}

	framesLen := v.prog.FramesLength()
	if framesLen == 0 {
		return memory.Nil, errNoFrames
	}

	var frame bytecode.Frame
	for i := 0; i < framesLen; i++ {
		if !v.prog.Frames(&frame, i) {
			return memory.Nil, errUnexpectedEOF
		}
		addr, err := v.execFrame(&frame)
		if err != nil {
			return memory.Nil, errors.Wrapf(err, "failed to exec frame at %d", i)
		}
		v.stack.push(addr)
	}

	return v.terminate()
}

func (v *VM) popType(n uint32, t memory.Type) ([]memory.Addr, error) {
	argv, err := v.stack.pop(n)
	if err != nil {
		return nil, err
	}
	for _, arg := range argv {
		if !arg.TypeOf(t) {
			return nil, fmt.Errorf("unexpected type %v instead of %v", arg.Type(), t)
		}
	}
	return argv, nil
}

func (v *VM) execFrame(frame *bytecode.Frame) (memory.Addr, error) {
	opType := frame.OpType()
	switch opType {
	case bytecode.OpOpPushTrue:
		return memory.True, nil
	case bytecode.OpOpPushFalse:
		return memory.False, nil
	case bytecode.OpOpAdd:
		argv, err := v.popType(2, memory.TypeInt64)
		if err != nil {
			return memory.Nil, err
		}
		return v.memory.AllocInt64(argv[0].Int64() + argv[1].Int64()), nil
	case bytecode.OpOpSub:
		argv, err := v.popType(2, memory.TypeInt64)
		if err != nil {
			return memory.Nil, err
		}
		return v.memory.AllocInt64(argv[0].Int64() - argv[1].Int64()), nil
	case bytecode.OpOpMul:
		argv, err := v.popType(2, memory.TypeInt64)
		if err != nil {
			return memory.Nil, err
		}
		return v.memory.AllocInt64(argv[0].Int64() * argv[1].Int64()), nil
	case bytecode.OpOpDiv:
		argv, err := v.popType(2, memory.TypeInt64)
		if err != nil {
			return memory.Nil, err
		}
		div := argv[1].Int64()
		if div == 0 {
			return memory.Nil, errDivByZero
		}
		return v.memory.AllocInt64(argv[0].Int64() / div), nil
	case bytecode.OpOpPushStr:
		var tab flatbuffers.Table
		if !frame.Op(&tab) {
			return memory.Nil, errNoOperation
		}
		var op bytecode.OpPushStr
		op.Init(tab.Bytes, tab.Pos)
		return v.memory.AllocBytesAddr(op.Val()), nil
	case bytecode.OpOpPushInt:
		var tab flatbuffers.Table
		if !frame.Op(&tab) {
			return memory.Nil, errNoOperation
		}
		var op bytecode.OpPushInt
		op.Init(tab.Bytes, tab.Pos)
		return v.memory.AllocInt64(op.Val()), nil
	case bytecode.OpOpPushVector:
		var tab flatbuffers.Table
		if !frame.Op(&tab) {
			return memory.Nil, errNoOperation
		}
		var op bytecode.OpPushVector
		op.Init(tab.Bytes, tab.Pos)
		elems, err := v.stack.pop(uint32(op.Elems()))
		if err != nil {
			return memory.Nil, err
		}
		return v.memory.CopyVector(elems...), nil
	case bytecode.OpOpInvoke:
		var tab flatbuffers.Table
		if !frame.Op(&tab) {
			return memory.Nil, errNoOperation
		}
		var op bytecode.OpInvoke
		op.Init(tab.Bytes, tab.Pos)
		fn := op.Name()
		if fn == nil {
			return memory.Nil, errEmptyDelegatorName
		}
		delegator, ok := v.globals.delegators.Get(string(fn))
		if !ok {
			return memory.Nil, fmt.Errorf("delegator <%s> not exists", fn)
		}
		args := op.Args()
		argv, err := v.stack.pop(uint32(args))
		if err != nil {
			return memory.Nil, err
		}
		return delegator.Delegate(&v.memory, argv)
	default:
		return memory.Nil, fmt.Errorf("unexpected frame type %s", opType)
	}
}
