package compiler

import (
	"fmt"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/regeda/expr/bytecode"
	"github.com/regeda/expr/lexer"
)

const (
	MinorVersion byte = 2
	MajorVersion byte = 0
)

var Default = Compiler{}

func Compile(nodes []lexer.Node) []byte {
	return Default.Compile(nodes)
}

type Compiler struct {
	b      flatbuffers.Builder
	frames []flatbuffers.UOffsetT
}

func (c *Compiler) Compile(nodes []lexer.Node) []byte {
	c.reset()

	c.b.Finish(c.writeProgram(
		c.writeFrames(nodes),
	))

	return c.b.FinishedBytes()
}

func (c *Compiler) reset() {
	c.b.Reset()
	c.frames = c.frames[:0]
}

func (c *Compiler) writeVersion(major, minor byte) flatbuffers.UOffsetT {
	return bytecode.CreateVersion(&c.b, minor, major)
}

func (c *Compiler) writeProgram(frames flatbuffers.UOffsetT) flatbuffers.UOffsetT {
	bytecode.ProgramStart(&c.b)
	bytecode.ProgramAddVer(&c.b, c.writeVersion(MajorVersion, MinorVersion))
	bytecode.ProgramAddFrames(&c.b, frames)
	return bytecode.ProgramEnd(&c.b)
}

func (c *Compiler) writeOpPushInt(v int64) flatbuffers.UOffsetT {
	bytecode.OpPushIntStart(&c.b)
	bytecode.OpPushIntAddVal(&c.b, v)
	return bytecode.OpPushIntEnd(&c.b)
}

func (c *Compiler) writeOpPushTrue() flatbuffers.UOffsetT {
	bytecode.OpPushTrueStart(&c.b)
	return bytecode.OpPushTrueEnd(&c.b)
}

func (c *Compiler) writeOpPushFalse() flatbuffers.UOffsetT {
	bytecode.OpPushFalseStart(&c.b)
	return bytecode.OpPushFalseEnd(&c.b)
}

func (c *Compiler) writeOpAdd() flatbuffers.UOffsetT {
	bytecode.OpAddStart(&c.b)
	return bytecode.OpAddEnd(&c.b)
}

func (c *Compiler) writeOpSub() flatbuffers.UOffsetT {
	bytecode.OpSubStart(&c.b)
	return bytecode.OpSubEnd(&c.b)
}

func (c *Compiler) writeOpMul() flatbuffers.UOffsetT {
	bytecode.OpMulStart(&c.b)
	return bytecode.OpMulEnd(&c.b)
}

func (c *Compiler) writeOpDiv() flatbuffers.UOffsetT {
	bytecode.OpDivStart(&c.b)
	return bytecode.OpDivEnd(&c.b)
}

func (c *Compiler) writeOpPushStr(v string) flatbuffers.UOffsetT {
	offset := c.b.CreateSharedString(v)

	bytecode.OpPushStrStart(&c.b)
	bytecode.OpPushStrAddVal(&c.b, offset)
	return bytecode.OpPushStrEnd(&c.b)
}

func (c *Compiler) writeOpPushVector(elems uint) flatbuffers.UOffsetT {
	bytecode.OpPushVectorStart(&c.b)
	bytecode.OpPushVectorAddElems(&c.b, uint16(elems))
	return bytecode.OpPushVectorEnd(&c.b)
}

func (c *Compiler) writeOpInvoke(fn string, args uint) flatbuffers.UOffsetT {
	fnOffset := c.b.CreateSharedString(fn)

	bytecode.OpInvokeStart(&c.b)
	bytecode.OpInvokeAddArgs(&c.b, uint16(args))
	bytecode.OpInvokeAddName(&c.b, fnOffset)
	return bytecode.OpInvokeEnd(&c.b)
}

func (c *Compiler) writeFrames(nodes []lexer.Node) flatbuffers.UOffsetT {
	c.discoverFrames(nodes)

	num := len(c.frames)
	bytecode.ProgramStartFramesVector(&c.b, num)
	for i := num - 1; i >= 0; i-- {
		c.b.PrependUOffsetT(c.frames[i])
	}
	return c.b.EndVector(num)
}

func (c *Compiler) writeFrame(offset flatbuffers.UOffsetT, op bytecode.Op) flatbuffers.UOffsetT {
	bytecode.FrameStart(&c.b)
	bytecode.FrameAddOpType(&c.b, op)
	bytecode.FrameAddOp(&c.b, offset)
	return bytecode.FrameEnd(&c.b)
}

func (c *Compiler) pushFrame(f flatbuffers.UOffsetT) {
	c.frames = append(c.frames, f)
}

func (c *Compiler) discoverFrames(nodes []lexer.Node) {
	for _, node := range nodes {
		switch node.Typ {
		case lexer.TypInvoke:
			c.pushFrame(c.writeFrame(
				c.writeOpInvoke(node.DatS, node.Cap),
				bytecode.OpOpInvoke,
			))
		case lexer.TypStr:
			c.pushFrame(c.writeFrame(
				c.writeOpPushStr(node.DatS),
				bytecode.OpOpPushStr,
			))
		case lexer.TypInt:
			c.pushFrame(c.writeFrame(
				c.writeOpPushInt(node.DatN),
				bytecode.OpOpPushInt,
			))
		case lexer.TypTrue:
			c.pushFrame(c.writeFrame(
				c.writeOpPushTrue(),
				bytecode.OpOpPushTrue,
			))
		case lexer.TypFalse:
			c.pushFrame(c.writeFrame(
				c.writeOpPushFalse(),
				bytecode.OpOpPushFalse,
			))
		case lexer.TypVector:
			c.pushFrame(c.writeFrame(
				c.writeOpPushVector(node.Cap),
				bytecode.OpOpPushVector,
			))
		case lexer.TypOpAdd:
			c.pushFrame(c.writeFrame(
				c.writeOpAdd(),
				bytecode.OpOpAdd,
			))
		case lexer.TypOpSub:
			c.pushFrame(c.writeFrame(
				c.writeOpSub(),
				bytecode.OpOpSub,
			))
		case lexer.TypOpMul:
			c.pushFrame(c.writeFrame(
				c.writeOpMul(),
				bytecode.OpOpMul,
			))
		case lexer.TypOpDiv:
			c.pushFrame(c.writeFrame(
				c.writeOpDiv(),
				bytecode.OpOpDiv,
			))
		default:
			panic(fmt.Sprintf("unknown type <%d>, see lexer.Typ", node.Typ))
		}
	}
}
