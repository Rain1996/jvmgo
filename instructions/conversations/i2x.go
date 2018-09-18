package conversations

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type I2B struct{ base.NoOperandsInstruction }

func (instruction *I2B) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	b := int32(int8(i))
	stack.PushInt(b)
}

type I2C struct{ base.NoOperandsInstruction }

func (instruction *I2C) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	c := int32(uint16(i))
	stack.PushInt(c)
}

type I2S struct{ base.NoOperandsInstruction }

func (instruction *I2S) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	s := int32(int16(i))
	stack.PushInt(s)
}

type I2L struct{ base.NoOperandsInstruction }

func (instruction *I2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	stack.PushLong(int64(i))
}

type I2F struct{ base.NoOperandsInstruction }

func (instruction *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	stack.PushFloat(float32(i))
}

type I2D struct{ base.NoOperandsInstruction }

func (instruction *I2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	stack.PushDouble(float64(i))
}
