package stores

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

func _astore(frame *rtda.Frame, index uint) {
	ref := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, ref)
}

// ASTORE store reference into local variable
type ASTORE struct{ base.Index8Instruction }

func (instruction *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, uint(instruction.Index))
}

type ASTORE_0 struct{ base.NoOperandsInstruction }

func (instruction *ASTORE_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}

type ASTORE_1 struct{ base.NoOperandsInstruction }

func (instruction *ASTORE_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}

type ASTORE_2 struct{ base.NoOperandsInstruction }

func (instruction *ASTORE_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}

type ASTORE_3 struct{ base.NoOperandsInstruction }

func (instruction *ASTORE_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}
