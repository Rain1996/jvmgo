package stack

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type SWAP struct{ base.NoOperandsInstruction }

func (instruction *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()

	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
