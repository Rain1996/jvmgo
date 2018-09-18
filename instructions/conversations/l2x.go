package conversations

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type L2D struct{ base.NoOperandsInstruction }

func (instruction *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	stack.PushDouble(float64(l))
}

type L2F struct{ base.NoOperandsInstruction }

func (instruction *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	stack.PushFloat(float32(l))
}

type L2I struct{ base.NoOperandsInstruction }

func (instruction *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	stack.PushInt(int32(l))
}
