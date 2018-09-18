package math

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type IINC struct {
	Index uint
	Const int32
}

func (instruction *IINC) FetchOperands(reader *base.BytecodeReader) {
	instruction.Index = uint(reader.ReadUint8())
	instruction.Const = int32(reader.ReadUint8())
}

func (instruction *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(instruction.Index)
	val += instruction.Const
	localVars.SetInt(instruction.Index, val)
}
