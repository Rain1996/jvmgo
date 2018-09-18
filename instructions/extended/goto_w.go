package extended

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type GOTO_W struct {
	offset int
}

func (instruction *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	instruction.offset = int(reader.ReadInt32())
}

func (instruction *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, instruction.offset)
}
