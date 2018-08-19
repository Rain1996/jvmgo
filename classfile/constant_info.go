package classfile

const (
	ConstantClass              = 7
	ConstantFieldRef           = 9
	ConstantMethodRef          = 10
	ConstantInterfaceMethodRef = 11
	ConstantString             = 8
	ConstantInteger            = 3
	ConstantFloat              = 4
	ConstantLong               = 5
	ConstantDouble             = 6
	ConstantNameAndType        = 12
	ConstantUtf8               = 1
	ConstantMethodHandle       = 15
	ConstantMethodType         = 16
	ConstantInvokeDynamic      = 17
)

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	// switch tag {
	// default:
	// 	painc("java.lang.ClassFormatError: constant pool tag!")
	// }
	return nil
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}
