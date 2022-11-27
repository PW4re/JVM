package asm

import (
	"jvm/src/runtime"
	"jvm/src/utils"
)

type Instruction interface {
	FetchOperands(reader *utils.BytesReader)
	Execute(fr *runtime.Frame)
}

type NoOperandsInstruction struct {
	// empty
}

func (instr *NoOperandsInstruction) FetchOperands(reader *utils.BytesReader) {
	// nothing to do
}

type BranchInstruction struct {
	Offset int
}

func (instr *BranchInstruction) FetchOperands(reader *utils.BytesReader) {
	instr.Offset = int(reader.ReadUint16())
}

type Index8Instruction struct {
	Index uint
}

func (instr *Index8Instruction) FetchOperands(reader *utils.BytesReader) {
	instr.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (instr *Index16Instruction) FetchOperands(reader *utils.BytesReader) {
	instr.Index = uint(reader.ReadUint16())
}
