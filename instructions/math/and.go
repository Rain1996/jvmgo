package math

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type IAND struct{ base.NoOperandsInstruction }

func (instruction *IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

type LAND struct{ base.NoOperandsInstruction }

func (instruction *LAND) Execute(frame *rtda.Frame) {
	stcak := frame.OperandStack()
	v2 := stcak.PopLong()
	v1 := stcak.PopLong()

	result := v1 & v2
	stcak.PushLong(result)
}
