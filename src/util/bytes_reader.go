package util

import (
	"encoding/binary"
	"math"
)

// BytesReader One reader must correspond to one class file
type BytesReader struct {
	byteOrder binary.ByteOrder
	data      []byte
	position  int
}

func NewBEAtStart(data []byte) BytesReader {
	return New(binary.BigEndian, data, 0)
}

func New(byteOrder binary.ByteOrder, data []byte, position int) BytesReader {
	return BytesReader{
		byteOrder: byteOrder,
		data:      data,
		position:  position,
	}
}

func (rd *BytesReader) Position() int {
	return rd.position
}

func (rd *BytesReader) ReadUint8() uint8 {
	i := rd.data[rd.position]
	rd.position++
	return i
}

func (rd *BytesReader) ReadUint16() uint16 {
	i := rd.byteOrder.Uint16(rd.data[rd.position:])
	rd.position += 2
	return i
}

func (rd *BytesReader) ReadUint32() uint32 {
	i := rd.byteOrder.Uint32(rd.data[rd.position:])
	rd.position += 4
	return i
}

func (rd *BytesReader) ReadUint64() uint64 {
	i := rd.byteOrder.Uint64(rd.data[rd.position:])
	rd.position += 8
	return i
}

func (rd *BytesReader) ReadFloat() float32 {
	return math.Float32frombits(rd.ReadUint32())
}

func (rd *BytesReader) ReadDouble() float64 {
	return math.Float64frombits(rd.ReadUint64())
}

func (rd *BytesReader) ReadBytes(n int) []byte {
	bytes := rd.data[rd.position : rd.position+n]
	rd.position += n
	return bytes
}
