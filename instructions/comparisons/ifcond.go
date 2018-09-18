package comparisons

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type IFEQ struct{ base.BranchInstruction }

func (instruction *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, instruction.Offset)
	}
}

type IFNE struct{ base.BranchInstruction }

func (instruction *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, instruction.Offset)
	}
}

type IFLT struct{ base.BranchInstruction }

func (instruction *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, instruction.Offset)
	}
}

type IFLE struct{ base.BranchInstruction }

func (instruction *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, instruction.Offset)
	}
}

type IFGT struct{ base.BranchInstruction }

func (instruction *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, instruction.Offset)
	}
}

type IFGE struct{ base.BranchInstruction }

func (instruction *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, instruction.Offset)
	}
}
