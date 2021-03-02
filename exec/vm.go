package exec

import (
	"fmt"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/pkg/errors"
	"github.com/regeda/expr/delegate"
	"github.com/regeda/expr/internal/bytecode"
	"github.com/regeda/expr/internal/compiler"
	"github.com/regeda/expr/memory"
)

type VM struct {
	delegators map[string]delegate.Delegator

	memory memory.Memory
	stack  stack
	// buffered variables
	prog bytecode.Program
}

func New(delegators map[string]delegate.Delegator) *VM {
	return &VM{
		delegators: delegators,
	}
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

func (v *VM) terminate() (*memory.Addr, error) {
	val, err := v.stack.pop(1)
	if err != nil {
		return nil, err
	}
	return val[0], nil
}

func (v *VM) Exec(bcode []byte) (*memory.Addr, error) {
	v.reset()

	v.prog.Init(bcode, flatbuffers.GetUOffsetT(bcode))

	if err := v.checkVersion(); err != nil {
		return nil, err
	}

	framesLen := v.prog.FramesLength()
	if framesLen == 0 {
		return nil, errNoFrames
	}

	for i := 0; i < framesLen; i++ {
		addr, err := v.execFrame(i)
		if err != nil {
			if err == errOpRet {
				return v.terminate()
			}
			return nil, errors.Wrapf(err, "failed to exec frame at %d", i)
		}
		v.stack.push(addr)
	}

	return nil, errUnexpectedEOP
}

func (v *VM) execFrame(i int) (*memory.Addr, error) {
	var frame bytecode.Frame
	if !v.prog.Frames(&frame, i) {
		return nil, errUnexpectedEOF
	}

	var tab flatbuffers.Table
	if !frame.Op(&tab) {
		return nil, errNoOperation
	}

	opType := frame.OpType()
	switch opType {
	case bytecode.OpOpPushBool:
		var op bytecode.OpPushBool
		op.Init(tab.Bytes, tab.Pos)
		if op.Val() {
			return memory.ConstTrue, nil
		}
		return memory.ConstFalse, nil
	case bytecode.OpOpPushStr:
		var op bytecode.OpPushStr
		op.Init(tab.Bytes, tab.Pos)
		return v.memory.AllocBytesAddr(op.Val())
	case bytecode.OpOpPushInt:
		var op bytecode.OpPushInt
		op.Init(tab.Bytes, tab.Pos)
		return v.memory.AllocInt64(op.Val())
	case bytecode.OpOpPushVector:
		var op bytecode.OpPushVector
		op.Init(tab.Bytes, tab.Pos)
		elems, err := v.stack.pop(uint32(op.Elems()))
		if err != nil {
			return nil, err
		}
		return v.memory.CopyVector(elems...)
	case bytecode.OpOpSysCall:
		var op bytecode.OpSysCall
		op.Init(tab.Bytes, tab.Pos)
		fn := op.Name()
		if fn == nil {
			return nil, errEmptyDelegatorName
		}
		delegator, ok := v.delegators[string(fn)]
		if !ok {
			return nil, fmt.Errorf("delegator <%s> not exists", fn)
		}
		args := op.Args()
		argv, err := v.stack.pop(uint32(args))
		if err != nil {
			return nil, err
		}
		return delegator.Delegate(&v.memory, argv)
	case bytecode.OpOpRet:
		return nil, errOpRet
	default:
		return nil, fmt.Errorf("unexpected frame type %s", opType)
	}
}
