package control

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type GOTO struct{ base.BranchInstruction }

func (instruction *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, instruction.Offset)
}
