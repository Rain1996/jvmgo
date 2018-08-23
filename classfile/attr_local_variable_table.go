package classfile

type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

func (lvt *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	lvt.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range lvt.localVariableTable {
		lvt.localVariableTable[i] = &LocalVariableTableEntry{
			startPc:         reader.readUint16(),
			length:          reader.readUint16(),
			nameIndex:       reader.readUint16(),
			descriptorIndex: reader.readUint16(),
			index:           reader.readUint16(),
		}
	}
}
