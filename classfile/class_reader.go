package classfile

import (
	"encoding/binary"
)

// ClassReader class file binary date
type ClassReader struct {
	data []byte
}

// read jvm u1 type
func (cr *ClassReader) readUint8() uint8 {
	val := cr.data[0]
	cr.data = cr.data[1:]
	return val
}

// read jvm u2 type, big end
func (cr *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(cr.data)
	cr.data = cr.data[2:]
	return val
}

// read jvm u4 type, big end
func (cr *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(cr.data)
	cr.data = cr.data[4:]
	return val
}

func (cr *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(cr.data)
	cr.data = cr.data[8:]
	return val
}

// read jvm table
func (cr *ClassReader) readUint16s() []uint16 {
	n := cr.readUint16()   // table head
	s := make([]uint16, n) // table item
	for i := range s {
		s[i] = cr.readUint16()
	}
	return s
}

// read bytes
func (cr *ClassReader) readBytes(n uint32) []byte {
	bytes := cr.data[:n]
	cr.data = cr.data[n:]
	return bytes
}
