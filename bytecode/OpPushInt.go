// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package bytecode

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type OpPushInt struct {
	_tab flatbuffers.Table
}

func GetRootAsOpPushInt(buf []byte, offset flatbuffers.UOffsetT) *OpPushInt {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &OpPushInt{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *OpPushInt) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *OpPushInt) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *OpPushInt) Val() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *OpPushInt) MutateVal(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func OpPushIntStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func OpPushIntAddVal(builder *flatbuffers.Builder, val int64) {
	builder.PrependInt64Slot(0, val, 0)
}
func OpPushIntEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
