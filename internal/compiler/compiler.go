package compiler

import (
	"fmt"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/regeda/expr/internal/ast"
	"github.com/regeda/expr/internal/bytecode"
)

const (
	MinorVersion byte = 1
	MajorVersion byte = 0
)

type Compiler struct {
	b      *flatbuffers.Builder
	frames []flatbuffers.UOffsetT
}

func New() *Compiler {
	return &Compiler{
		b: flatbuffers.NewBuilder(1024),
	}
}

func (c *Compiler) reset() {
	c.b.Reset()
	c.frames = c.frames[:0]
}

func (c *Compiler) writeVersion(major, minor byte) flatbuffers.UOffsetT {
	bytecode.VersionStart(c.b)
	bytecode.VersionAddMajor(c.b, major)
	bytecode.VersionAddMinor(c.b, minor)
	return bytecode.VersionEnd(c.b)
}

func (c *Compiler) writeProgram(version, frames flatbuffers.UOffsetT) flatbuffers.UOffsetT {
	bytecode.ProgramStart(c.b)
	bytecode.ProgramAddVer(c.b, version)
	bytecode.ProgramAddFrames(c.b, frames)
	return bytecode.ProgramEnd(c.b)
}

func (c *Compiler) writeOpPushInt(v int64) flatbuffers.UOffsetT {
	bytecode.OpPushIntStart(c.b)
	bytecode.OpPushIntAddVal(c.b, v)
	return bytecode.OpPushIntEnd(c.b)
}

func (c *Compiler) writeOpPushBool(v bool) flatbuffers.UOffsetT {
	bytecode.OpPushBoolStart(c.b)
	bytecode.OpPushBoolAddVal(c.b, v)
	return bytecode.OpPushBoolEnd(c.b)
}

func (c *Compiler) writeOpPushStr(v string) flatbuffers.UOffsetT {
	offset := c.b.CreateSharedString(v)

	bytecode.OpPushStrStart(c.b)
	bytecode.OpPushStrAddVal(c.b, offset)
	return bytecode.OpPushStrEnd(c.b)
}

func (c *Compiler) writeOpPushVector(elems int) flatbuffers.UOffsetT {
	bytecode.OpPushVectorStart(c.b)
	bytecode.OpPushVectorAddElems(c.b, uint16(elems))
	return bytecode.OpPushVectorEnd(c.b)
}

func (c *Compiler) writeOpRet() flatbuffers.UOffsetT {
	bytecode.OpRetStart(c.b)
	return bytecode.OpRetEnd(c.b)
}

func (c *Compiler) writeOpSysCall(fn string, args int) flatbuffers.UOffsetT {
	fnOffset := c.b.CreateSharedString(fn)

	bytecode.OpSysCallStart(c.b)
	bytecode.OpSysCallAddArgs(c.b, uint16(args))
	bytecode.OpSysCallAddName(c.b, fnOffset)
	return bytecode.OpSysCallEnd(c.b)
}

func (c *Compiler) writeFrames(node *ast.Node) flatbuffers.UOffsetT {
	c.discoverFrames(node)

	num := len(c.frames)
	bytecode.ProgramStartFramesVector(c.b, num)
	for i := num - 1; i >= 0; i-- {
		c.b.PrependUOffsetT(c.frames[i])
	}
	return c.b.EndVector(num)
}

func (c *Compiler) writeFrame(offset flatbuffers.UOffsetT, op bytecode.Op) flatbuffers.UOffsetT {
	bytecode.FrameStart(c.b)
	bytecode.FrameAddOpType(c.b, op)
	bytecode.FrameAddOp(c.b, offset)
	return bytecode.FrameEnd(c.b)
}

func (c *Compiler) pushFrame(f flatbuffers.UOffsetT) {
	c.frames = append(c.frames, f)
}

func (c *Compiler) discoverFrames(nodes ...*ast.Node) {
	for _, node := range nodes {
		switch node.Token {
		case ast.Node_EXIT:
			c.discoverFrames(node.Nested...)
			c.pushFrame(c.writeFrame(
				c.writeOpRet(),
				bytecode.OpOpRet,
			))
		case ast.Node_CALL:
			c.discoverFrames(node.Nested...)
			c.pushFrame(c.writeFrame(
				c.writeOpSysCall(node.GetS(), len(node.Nested)),
				bytecode.OpOpSysCall,
			))
		case ast.Node_STR:
			c.pushFrame(c.writeFrame(
				c.writeOpPushStr(node.GetS()),
				bytecode.OpOpPushStr,
			))
		case ast.Node_INT:
			c.pushFrame(c.writeFrame(
				c.writeOpPushInt(node.GetI()),
				bytecode.OpOpPushInt,
			))
		case ast.Node_BOOL:
			c.pushFrame(c.writeFrame(
				c.writeOpPushBool(node.GetB()),
				bytecode.OpOpPushBool,
			))
		case ast.Node_ARR:
			c.discoverFrames(node.Nested...)
			c.pushFrame(c.writeFrame(
				c.writeOpPushVector(len(node.Nested)),
				bytecode.OpOpPushVector,
			))
		default:
			panic(fmt.Sprintf("unknown token <%s>, see ast.Node_Token", node.Token))
		}
	}
}

func (c *Compiler) Compile(node *ast.Node) []byte {
	c.reset()

	c.b.Finish(c.writeProgram(
		c.writeVersion(MajorVersion, MinorVersion),
		c.writeFrames(node),
	))

	return c.b.FinishedBytes()
}
