package comparisons

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

// LCMP compare long
type LCMP struct{ base.NoOperandsInstruction }

func (instruction *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}