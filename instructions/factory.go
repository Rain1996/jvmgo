package instructions

import (
	"fmt"
	"jvmgo/instructions/base"
	"jvmgo/instructions/comparisons"
	"jvmgo/instructions/constants"
	"jvmgo/instructions/control"
	"jvmgo/instructions/conversations"
	"jvmgo/instructions/extended"
	"jvmgo/instructions/loads"
	"jvmgo/instructions/math"
	"jvmgo/instructions/stack"
	"jvmgo/instructions/stores"
)

var (
	nop        = &constants.NOP{}
	aconstNull = &constants.ACONST_NULL{}

	iconstM1 = &constants.ICONST_M1{}
	iconst0  = &constants.ICONST_0{}
	iconst1  = &constants.ICONST_1{}
	iconst2  = &constants.ICONST_2{}
	iconst3  = &constants.ICONST_3{}
	iconst4  = &constants.ICONST_4{}
	iconst5  = &constants.ICONST_5{}

	lconst0 = &constants.LCONST_0{}
	lconst1 = &constants.LCONST_1{}

	fconst0 = &constants.FCONST_0{}
	fconst1 = &constants.FCONST_1{}
	fconst2 = &constants.FCONST_2{}

	dconst0 = &constants.DCONST_0{}
	dconst1 = &constants.DCONST_1{}

	iload0 = &loads.ILOAD_0{}
	iload1 = &loads.ILOAD_1{}
	iload2 = &loads.ILOAD_2{}
	iload3 = &loads.ILOAD_3{}

	lload0 = &loads.LLOAD_0{}
	lload1 = &loads.LLOAD_1{}
	lload2 = &loads.LLOAD_2{}
	lload3 = &loads.LLOAD_3{}

	fload0 = &loads.FLOAD_0{}
	fload1 = &loads.FLOAD_1{}
	fload2 = &loads.FLOAD_2{}
	fload3 = &loads.FLOAD_3{}

	dload0 = &loads.DLOAD_0{}
	dload1 = &loads.DLOAD_1{}
	dload2 = &loads.DLOAD_2{}
	dload3 = &loads.DLOAD_3{}

	aload0 = &loads.ALOAD_0{}
	aload1 = &loads.ALOAD_1{}
	aload2 = &loads.ALOAD_2{}
	aload3 = &loads.ALOAD_3{}

	istore0 = &stores.ISTORE_0{}
	istore1 = &stores.ISTORE_1{}
	istore2 = &stores.ISTORE_2{}
	istore3 = &stores.ISTORE_3{}

	lstore0 = &stores.LSTORE_0{}
	lstore1 = &stores.LSTORE_1{}
	lstore2 = &stores.LSTORE_2{}
	lstore3 = &stores.LSTORE_3{}

	fstore0 = &stores.FSTORE_0{}
	fstore1 = &stores.FSTORE_1{}
	fstore2 = &stores.FSTORE_2{}
	fstore3 = &stores.FSTORE_3{}

	dstore0 = &stores.DSTORE_0{}
	dstore1 = &stores.DSTORE_1{}
	dstore2 = &stores.DSTORE_2{}
	dstore3 = &stores.DSTORE_3{}

	astore0 = &stores.ASTORE_0{}
	astore1 = &stores.ASTORE_1{}
	astore2 = &stores.ASTORE_2{}
	astore3 = &stores.ASTORE_3{}

	pop  = &stack.POP{}
	pop2 = &stack.POP2{}

	dup    = &stack.DUP{}
	dupX1  = &stack.DUP_X1{}
	dupX2  = &stack.DUP_X2{}
	dup2   = &stack.DUP2{}
	dup2X1 = &stack.DUP2_X1{}
	dup2X2 = &stack.DUP2_X2{}

	swap = &stack.SWAP{}

	iadd = &math.IADD{}
	ladd = &math.LADD{}
	fadd = &math.FADD{}
	dadd = &math.DADD{}

	isub = &math.ISUB{}
	lsub = &math.LSUB{}
	fsub = &math.FSUB{}
	dsub = &math.DSUB{}

	imul = &math.IMUL{}
	lmul = &math.LMUL{}
	fmul = &math.FMUL{}
	dmul = &math.DMUL{}

	idiv = &math.IDIV{}
	ldiv = &math.LDIV{}
	fdiv = &math.FDIV{}
	ddiv = &math.DDIV{}

	irem = &math.IREM{}
	lrem = &math.LREM{}
	frem = &math.FREM{}
	drem = &math.DREM{}

	ineg = &math.INEG{}
	lneg = &math.LNEG{}
	fneg = &math.FNEG{}
	dneg = &math.DNEG{}

	ishl  = &math.ISHL{}
	lshl  = &math.LSHL{}
	ishr  = &math.ISHR{}
	lshr  = &math.LSHR{}
	iushr = &math.IUSHR{}
	lushr = &math.LUSHR{}

	iand = &math.IADD{}
	land = &math.LADD{}
	ior  = &math.IOR{}
	lor  = &math.LOR{}
	ixor = &math.IXOR{}
	lxor = &math.LXOR{}

	i2l = &conversations.I2L{}
	i2f = &conversations.I2F{}
	i2d = &conversations.I2D{}

	l2i = &conversations.L2I{}
	l2f = &conversations.L2F{}
	l2d = &conversations.L2D{}

	f2i = &conversations.F2I{}
	f2l = &conversations.F2L{}
	f2d = &conversations.F2D{}

	d2i = &conversations.D2I{}
	d2l = &conversations.D2L{}
	d2f = &conversations.D2F{}

	i2b = &conversations.I2B{}
	i2c = &conversations.I2C{}
	i2s = &conversations.I2S{}

	lcmp  = &comparisons.LCMP{}
	fcmpl = &comparisons.FCMPL{}
	fcmpg = &comparisons.FCMPG{}
	dcmpl = &comparisons.DCMPL{}
	dcmpg = &comparisons.DCMPG{}
)

// NewInstruction return a instruction by opcode
func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return nop
	case 0x01:
		return aconstNull
	case 0x02:
		return iconstM1
	case 0x03:
		return iconst0
	case 0x04:
		return iconst1
	case 0x05:
		return iconst2
	case 0x06:
		return iconst3
	case 0x07:
		return iconst4
	case 0x08:
		return iconst5
	case 0x09:
		return lconst0
	case 0x0a:
		return lconst1
	case 0x0b:
		return fconst0
	case 0x0c:
		return fconst1
	case 0x0d:
		return fconst2
	case 0x0e:
		return dconst0
	case 0x0f:
		return dconst1
	case 0x10:
		return &constants.BIPUSH{}
	case 0x11:
		return &constants.SIPUSH{}
	case 0x15:
		return &loads.ILOAD{}
	case 0x16:
		return &loads.LLOAD{}
	case 0x17:
		return &loads.FLOAD{}
	case 0x18:
		return &loads.DLOAD{}
	case 0x19:
		return &loads.ALOAD{}
	case 0x1a:
		return iload0
	case 0x1b:
		return iload1
	case 0x1c:
		return iload2
	case 0x1d:
		return iload3
	case 0x1e:
		return lload0
	case 0x1f:
		return lload1
	case 0x20:
		return lload2
	case 0x21:
		return lload3
	case 0x22:
		return fload0
	case 0x23:
		return fload1
	case 0x24:
		return fload2
	case 0x25:
		return fload3
	case 0x26:
		return dload0
	case 0x27:
		return dload1
	case 0x28:
		return dload2
	case 0x29:
		return dload3
	case 0x2a:
		return aload0
	case 0x2b:
		return aload1
	case 0x2c:
		return aload2
	case 0x2d:
		return aload3
	case 0x36:
		return &stores.ISTORE{}
	case 0x37:
		return &stores.LSTORE{}
	case 0x38:
		return &stores.FSTORE{}
	case 0x39:
		return &stores.DSTORE{}
	case 0x3a:
		return &stores.ASTORE{}
	case 0x3b:
		return istore0
	case 0x3c:
		return istore1
	case 0x3d:
		return istore2
	case 0x3e:
		return istore3
	case 0x3f:
		return lstore0
	case 0x40:
		return lstore1
	case 0x41:
		return lstore2
	case 0x42:
		return lstore3
	case 0x43:
		return fstore0
	case 0x44:
		return fstore1
	case 0x45:
		return fstore2
	case 0x46:
		return fstore3
	case 0x47:
		return dstore0
	case 0x48:
		return dstore1
	case 0x49:
		return dstore2
	case 0x4a:
		return dstore3
	case 0x4b:
		return astore0
	case 0x4c:
		return astore1
	case 0x4d:
		return astore2
	case 0x4e:
		return astore3
	case 0x57:
		return pop
	case 0x58:
		return pop2
	case 0x59:
		return dup
	case 0x5a:
		return dupX1
	case 0x5b:
		return dupX2
	case 0x5c:
		return dup2
	case 0x5d:
		return dup2X1
	case 0x5e:
		return dup2X2
	case 0x5f:
		return swap
	case 0x60:
		return iadd
	case 0x61:
		return ladd
	case 0x62:
		return fadd
	case 0x63:
		return dadd
	case 0x64:
		return isub
	case 0x65:
		return lsub
	case 0x66:
		return fsub
	case 0x67:
		return dsub
	case 0x68:
		return imul
	case 0x69:
		return lmul
	case 0x6a:
		return fmul
	case 0x6b:
		return dmul
	case 0x6c:
		return idiv
	case 0x6d:
		return ldiv
	case 0x6e:
		return fdiv
	case 0x6f:
		return ddiv
	case 0x70:
		return irem
	case 0x71:
		return lrem
	case 0x72:
		return frem
	case 0x73:
		return drem
	case 0x74:
		return ineg
	case 0x75:
		return lneg
	case 0x76:
		return fneg
	case 0x77:
		return dneg
	case 0x78:
		return ishl
	case 0x79:
		return lshl
	case 0x7a:
		return ishr
	case 0x7b:
		return lshr
	case 0x7c:
		return iushr
	case 0x7d:
		return lushr
	case 0x7e:
		return iand
	case 0x7f:
		return land
	case 0x80:
		return ior
	case 0x81:
		return lor
	case 0x82:
		return ixor
	case 0x83:
		return lxor
	case 0x84:
		return &math.IINC{}
	case 0x85:
		return i2l
	case 0x86:
		return i2f
	case 0x87:
		return i2d
	case 0x88:
		return l2i
	case 0x89:
		return l2f
	case 0x8a:
		return l2d
	case 0x8b:
		return f2i
	case 0x8c:
		return f2l
	case 0x8d:
		return f2d
	case 0x8e:
		return d2i
	case 0x8f:
		return d2l
	case 0x90:
		return d2f
	case 0x91:
		return i2b
	case 0x92:
		return i2c
	case 0x93:
		return i2s
	case 0x94:
		return lcmp
	case 0x95:
		return fcmpl
	case 0x96:
		return fcmpg
	case 0x97:
		return dcmpl
	case 0x98:
		return dcmpg
	case 0x99:
		return &comparisons.IFEQ{}
	case 0x9a:
		return &comparisons.IFNE{}
	case 0x9b:
		return &comparisons.IFLT{}
	case 0x9c:
		return &comparisons.IFGE{}
	case 0x9d:
		return &comparisons.IFGT{}
	case 0x9e:
		return &comparisons.IFLE{}
	case 0x9f:
		return &comparisons.IF_ICMPEQ{}
	case 0xa0:
		return &comparisons.IF_ICMPNE{}
	case 0xa1:
		return &comparisons.IF_ICMPLT{}
	case 0xa2:
		return &comparisons.IF_ICMPGE{}
	case 0xa3:
		return &comparisons.IF_ICMPGT{}
	case 0xa4:
		return &comparisons.IF_ICMPLE{}
	case 0xa5:
		return &comparisons.IF_ACMPEQ{}
	case 0xa6:
		return &comparisons.IF_ACMPNE{}
	case 0xa7:
		return &control.GOTO{}
	case 0xaa:
		return &control.TABLE_SWITCH{}
	case 0xab:
		return &control.LOOKUP_SWITCH{}
	case 0xc4:
		return &extended.IFNULL{}
	case 0xc7:
		return &extended.IFNONULL{}
	case 0xc8:
		return &extended.GOTO_W{}
	default:
		panic(fmt.Errorf("Unsported opcode: 0x%x", opcode))
	}
}
