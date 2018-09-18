package loads

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

func _fload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

type FLOAD struct{ base.Index8Instruction }

func (instruction *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, instruction.Index)
}

type FLOAD_0 struct{ base.NoOperandsInstruction }

func (instruction *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}

type FLOAD_1 struct{ base.NoOperandsInstruction }

func (instruction *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

type FLOAD_2 struct{ base.NoOperandsInstruction }

func (instruction *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

type FLOAD_3 struct{ base.NoOperandsInstruction }

func (instruction *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}