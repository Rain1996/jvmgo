package comparisons

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

func _dcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

type DCMPG struct{ base.NoOperandsInstruction }

func (instruction *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, false)
}

type DCMPL struct{ base.NoOperandsInstruction }

func (instruction *DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}
