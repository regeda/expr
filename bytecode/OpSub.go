// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package bytecode

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type OpSub struct {
	_tab flatbuffers.Table
}

func GetRootAsOpSub(buf []byte, offset flatbuffers.UOffsetT) *OpSub {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &OpSub{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *OpSub) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *OpSub) Table() flatbuffers.Table {
	return rcv._tab
}

func OpSubStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func OpSubEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
