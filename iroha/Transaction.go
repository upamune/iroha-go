// automatically generated by the FlatBuffers compiler, do not modify

package iroha

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Transaction struct {
	_tab flatbuffers.Table
}

func GetRootAsTransaction(buf []byte, offset flatbuffers.UOffsetT) *Transaction {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Transaction{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Transaction) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Transaction) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Transaction) CreatorPubKey() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Transaction) CommandType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Transaction) MutateCommandType(n byte) bool {
	return rcv._tab.MutateByteSlot(6, n)
}

func (rcv *Transaction) Command(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func (rcv *Transaction) Signatures(obj *Signature, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Transaction) SignaturesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Transaction) Hash(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Transaction) HashLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Transaction) HashBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Transaction) Timestamp() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Transaction) MutateTimestamp(n uint64) bool {
	return rcv._tab.MutateUint64Slot(14, n)
}

func (rcv *Transaction) Attachment(obj *Attachment) *Attachment {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Attachment)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func TransactionStart(builder *flatbuffers.Builder) {
	builder.StartObject(7)
}
func TransactionAddCreatorPubKey(builder *flatbuffers.Builder, creatorPubKey flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(creatorPubKey), 0)
}
func TransactionAddCommandType(builder *flatbuffers.Builder, commandType byte) {
	builder.PrependByteSlot(1, commandType, 0)
}
func TransactionAddCommand(builder *flatbuffers.Builder, command flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(command), 0)
}
func TransactionAddSignatures(builder *flatbuffers.Builder, signatures flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(signatures), 0)
}
func TransactionStartSignaturesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TransactionAddHash(builder *flatbuffers.Builder, hash flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(hash), 0)
}
func TransactionStartHashVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func TransactionAddTimestamp(builder *flatbuffers.Builder, timestamp uint64) {
	builder.PrependUint64Slot(5, timestamp, 0)
}
func TransactionAddAttachment(builder *flatbuffers.Builder, attachment flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(attachment), 0)
}
func TransactionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
