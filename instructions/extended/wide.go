package extended

import (
	"jvmgo/instructions/base"
	"jvmgo/instructions/loads"
	"jvmgo/instructions/math"
	"jvmgo/instructions/stores"
	"jvmgo/rtda"
)

type WIDE struct {
	modifiedInstrution base.Instruction
}

func (instruction *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15:
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		instruction.modifiedInstrution = inst
	case 0x16:
		inst := &loads.LLOAD{}
		inst.Index = uint(reader.ReadUint16())
		instruction.modifiedInstrution = inst
	case 0x17:
		inst := &loads.FLOAD{}
		inst.Index = uint(reader.ReadUint16())
		instruction.modifiedInstrution = inst
	case 0x18:
		inst := &loads.DLOAD{}
		inst.Index = uint(reader.ReadUint16())
		instruction.modifiedInstrution = inst
	case 0x19:
		inst := &loads.ALOAD{}
		inst.Index = uint(reader.ReadUint16())
		instruction.modifiedInstrution = inst
	case 0x36:
		inst := &stores.ISTORE{}
		inst.Index = uint(reader.ReadUint16())
		instruction.modifiedInstrution = inst
	case 0x37:
		inst := &stores.LSTORE{}
		inst.Index = uint(reader.ReadUint16())
		instruction.modifiedInstrution = inst
	case 0x38:
		inst := &stores.FSTORE{}
		inst.Index = uint(reader.ReadUint16())
		instruction.modifiedInstrution = inst
	case 0x39:
		inst := &stores.DSTORE{}
		inst.Index = uint(reader.ReadUint16())
		instruction.modifiedInstrution = inst
	case 0x3a:
		inst := &stores.ASTORE{}
		inst.Index = uint(reader.ReadUint16())
		instruction.modifiedInstrution = inst
	case 0x84:
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadUint16())
		instruction.modifiedInstrution = inst
	case 0xa9:
		panic("Unsupported opcpde: 0xa9!")
	}
}

func (instruction *WIDE) Execute(frame *rtda.Frame) {
	instruction.modifiedInstrution.Execute(frame)
}
