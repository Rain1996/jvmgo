package conversations

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type F2D struct{ base.NoOperandsInstruction }

func (instruction *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	stack.PushDouble(float64(f))
}

type F2I struct{ base.NoOperandsInstruction }

func (instruction *F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	stack.PushInt(int32(f))
}

type F2L struct{ base.NoOperandsInstruction }

func (instruction *F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	stack.PushLong(int64(f))
}
