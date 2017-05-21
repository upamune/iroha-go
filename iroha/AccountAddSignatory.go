// automatically generated by the FlatBuffers compiler, do not modify

package iroha

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type AccountAddSignatory struct {
	_tab flatbuffers.Table
}

func GetRootAsAccountAddSignatory(buf []byte, offset flatbuffers.UOffsetT) *AccountAddSignatory {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AccountAddSignatory{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *AccountAddSignatory) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AccountAddSignatory) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *AccountAddSignatory) Account() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AccountAddSignatory) Signatory(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *AccountAddSignatory) SignatoryLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func AccountAddSignatoryStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func AccountAddSignatoryAddAccount(builder *flatbuffers.Builder, account flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(account), 0)
}
func AccountAddSignatoryAddSignatory(builder *flatbuffers.Builder, signatory flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(signatory), 0)
}
func AccountAddSignatoryStartSignatoryVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func AccountAddSignatoryEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
