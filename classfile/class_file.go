package classfile

import (
	"fmt"
)

// ClassFile class file structure
type ClassFile struct {
	// magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
}

func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	if magic := reader.readUint32(); magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	cf.minorVersion = reader.readUint16()
	cf.majorVersion = reader.readUint16()

	switch {
	case cf.majorVersion == 45:
		return
	case cf.majorVersion > 45 && cf.majorVersion <= 54:
		if cf.minorVersion == 0 {
			return
		}
	default:
		panic("java.lang.UnsupportedClassVersionError!")
	}
}

func (cf *ClassFile) read(reader *ClassReader) {
	cf.readAndCheckMagic(reader)
	cf.readAndCheckVersion(reader)
}

// Parse parse class data
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", err)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}
