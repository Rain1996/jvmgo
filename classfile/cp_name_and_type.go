package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex        uint16
	descriptionIndex uint16
}

func (cnt *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	cnt.nameIndex = reader.readUint16()
	cnt.descriptionIndex = reader.readUint16()
}
