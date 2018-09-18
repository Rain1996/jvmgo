package conversations

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type D2F struct{ base.NoOperandsInstruction }

func (instruction *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	stack.PushFloat(float32(d))
}

type D2I struct{ base.NoOperandsInstruction }

func (instruction *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	stack.PushInt(int32(d))
}

type D2L struct{ base.NoOperandsInstruction }

func (instruction *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	stack.PushLong(int64(d))
}
