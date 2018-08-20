package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (lnt *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	LineNumberTableLength := reader.readUint16()
	lnt.lineNumberTable = make([]*LineNumberTableEntry, LineNumberTableLength)
	for i := range lnt.lineNumberTable {
		lnt.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
