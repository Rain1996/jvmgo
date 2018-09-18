package control

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

/*
tableswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
lowbyte1
lowbyte2
lowbyte3
lowbyte4
highbyte1
highbyte2
highbyte3
highbyte4
jump offsets...
*/

type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (instruction *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	instruction.defaultOffset = reader.ReadInt32()
	instruction.low = reader.ReadInt32()
	instruction.high = reader.ReadInt32()
	jumpOffsetsCount := instruction.high - instruction.low + 1
	instruction.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (instruction *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= instruction.low && index <= instruction.high {
		offset = int(instruction.jumpOffsets[index-instruction.low])
	} else {
		offset = int(instruction.defaultOffset)
	}

	base.Branch(frame, offset)
}
