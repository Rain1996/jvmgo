package main

import (
	"fmt"
	"jvmgo/classfile"
	"jvmgo/instructions"
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	byteCode := codeAttr.Code()

	thread := rtda.NewThread(rtda.DefaultMaxStackSize)
	frame := rtda.NewFrame(thread, maxLocals, maxStack)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, byteCode)
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars: %v\n", frame.LocalVars())
		fmt.Printf("OperandStack: %v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}

	for {
		pc := frame.NextPC()
		thread.SetPc(pc)

		// decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		fmt.Printf("pc: %2d inst: %T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
