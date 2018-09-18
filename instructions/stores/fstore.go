package stores

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

func _fstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}

// FSTORE store float into local varibale
type FSTORE struct{ base.Index8Instruction }

func (instruction *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame, uint(instruction.Index))
}

type FSTORE_0 struct{ base.NoOperandsInstruction }

func (instruction *FSTORE_0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}

type FSTORE_1 struct{ base.NoOperandsInstruction }

func (instruction *FSTORE_1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}

type FSTORE_2 struct{ base.NoOperandsInstruction }

func (instruction *FSTORE_2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}

type FSTORE_3 struct{ base.NoOperandsInstruction }

func (instruction *FSTORE_3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}
