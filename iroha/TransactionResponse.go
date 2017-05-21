// automatically generated by the FlatBuffers compiler, do not modify

package iroha

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TransactionResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsTransactionResponse(buf []byte, offset flatbuffers.UOffsetT) *TransactionResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TransactionResponse{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *TransactionResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TransactionResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TransactionResponse) Message() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *TransactionResponse) Index() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TransactionResponse) MutateIndex(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func (rcv *TransactionResponse) Code() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TransactionResponse) MutateCode(n byte) bool {
	return rcv._tab.MutateByteSlot(8, n)
}

func (rcv *TransactionResponse) Transactions(obj *Transaction, j int) bool {
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

func (rcv *TransactionResponse) TransactionsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func TransactionResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func TransactionResponseAddMessage(builder *flatbuffers.Builder, message flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(message), 0)
}
func TransactionResponseAddIndex(builder *flatbuffers.Builder, index uint64) {
	builder.PrependUint64Slot(1, index, 0)
}
func TransactionResponseAddCode(builder *flatbuffers.Builder, code byte) {
	builder.PrependByteSlot(2, code, 0)
}
func TransactionResponseAddTransactions(builder *flatbuffers.Builder, transactions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(transactions), 0)
}
func TransactionResponseStartTransactionsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TransactionResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
