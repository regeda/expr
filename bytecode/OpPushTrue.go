// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package bytecode

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type OpPushTrue struct {
	_tab flatbuffers.Table
}

func GetRootAsOpPushTrue(buf []byte, offset flatbuffers.UOffsetT) *OpPushTrue {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &OpPushTrue{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *OpPushTrue) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *OpPushTrue) Table() flatbuffers.Table {
	return rcv._tab
}

func OpPushTrueStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func OpPushTrueEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
