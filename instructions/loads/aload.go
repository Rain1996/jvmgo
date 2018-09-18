package loads

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

func _aload(frame *rtda.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}

type ALOAD struct{ base.Index8Instruction }

func (instruction *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, instruction.Index)
}

type ALOAD_0 struct{ base.NoOperandsInstruction }

func (instruction *ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

type ALOAD_1 struct{ base.NoOperandsInstruction }

func (instruction *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

type ALOAD_2 struct{ base.NoOperandsInstruction }

func (instruction *ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

type ALOAD_3 struct{ base.NoOperandsInstruction }

func (instruction *ALOAD_3) Execute(frame *rtda.Frame) {
}
