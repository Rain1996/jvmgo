package loads

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

func _lload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}

type LLOAD struct{ base.Index8Instruction }

func (instruction *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, instruction.Index)
}

type LLOAD_0 struct{ base.NoOperandsInstruction }

func (instruction *LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}

type LLOAD_1 struct{ base.NoOperandsInstruction }

func (instruction *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}

type LLOAD_2 struct{ base.NoOperandsInstruction }

func (instruction *LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}

type LLOAD_3 struct{ base.NoOperandsInstruction }

func (instruction *LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}
