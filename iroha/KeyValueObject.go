// automatically generated by the FlatBuffers compiler, do not modify

package iroha

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/////////////////////////////////////
/// Key-Value object
type KeyValueObject struct {
	_tab flatbuffers.Table
}

func GetRootAsKeyValueObject(buf []byte, offset flatbuffers.UOffsetT) *KeyValueObject {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KeyValueObject{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *KeyValueObject) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KeyValueObject) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *KeyValueObject) Key() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *KeyValueObject) Value(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *KeyValueObject) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *KeyValueObject) ValueBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func KeyValueObjectStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func KeyValueObjectAddKey(builder *flatbuffers.Builder, key flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(key), 0)
}
func KeyValueObjectAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(value), 0)
}
func KeyValueObjectStartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func KeyValueObjectEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
