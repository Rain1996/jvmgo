package extended

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type IFNULL struct{ base.BranchInstruction }

func (instruction *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, instruction.Offset)
	}
}

type IFNONULL struct{ base.BranchInstruction }

func (instruction *IFNONULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, instruction.Offset)
	}
}
