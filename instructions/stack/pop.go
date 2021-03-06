package stack

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type POP struct{ base.NoOperandsInstruction }

func (instruction *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

type POP2 struct{ base.NoOperandsInstruction }

func (instruction *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
