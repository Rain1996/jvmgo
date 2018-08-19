package classfile

type ConstantMemberRefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (cmr *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	cmr.classIndex = reader.readUint16()
	cmr.nameAndTypeIndex = reader.readUint16()
}

func (cmr *ConstantMemberRefInfo) ClassName() string {
	return cmr.cp.getClassName(cmr.classIndex)
}

func (cmr *ConstantMemberRefInfo) nameAndDescriptor() (string, string) {
	return cmr.cp.getNameAndType(cmr.nameAndTypeIndex)
}

type ConstantFieldRefInfo struct{ ConstantMemberRefInfo }
type ConstantMethodRefInfo struct{ ConstantMemberRefInfo }
type ConstantInterfaceMethodRefInfo struct{ ConstantMemberRefInfo }
