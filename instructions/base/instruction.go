package base

import (
	"jvmgo/rtda"
)

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct{}

func (instruction NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// noting to do
}

type BranchInstruction struct {
	Offset int
}

func (instruction *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	instruction.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint
}

func (instruction *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	instruction.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (instruction *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	instruction.Index = uint(reader.ReadUint16())
}
