package constants

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

// BIPUSH push byte
type BIPUSH struct{ val int8 }

func (instruction *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	instruction.val = reader.ReadInt8()
}

func (instruction *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(instruction.val)
	frame.OperandStack().PushInt(i)
}

// SIPUSH push short
type SIPUSH struct{ val int16 }

func (instruction *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	instruction.val = reader.ReadInt16()
}

func (instruction *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(instruction.val)
	frame.OperandStack().PushInt(i)
}
