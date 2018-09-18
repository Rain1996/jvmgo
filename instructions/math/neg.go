package math

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type DNEG struct{ base.NoOperandsInstruction }

func (instruction *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

type FNEG struct{ base.NoOperandsInstruction }

func (instruction *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

type INEG struct{ base.NoOperandsInstruction }

func (instruction *INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

type LNEG struct{ base.NoOperandsInstruction }

func (instruction *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}
