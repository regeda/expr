// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package bytecode

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type OpRet struct {
	_tab flatbuffers.Table
}

func GetRootAsOpRet(buf []byte, offset flatbuffers.UOffsetT) *OpRet {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &OpRet{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *OpRet) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *OpRet) Table() flatbuffers.Table {
	return rcv._tab
}

func OpRetStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func OpRetEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}