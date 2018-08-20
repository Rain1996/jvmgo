package classfile

type MarkerAttribute struct{}

func (ma *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}

type DeprecatedAttribute struct{ MarkerAttribute }
type SyntheticAttribute struct{ MarkerAttribute }
