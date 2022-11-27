package attributes

import (
	"jvm/src/class_file/cp"
	"jvm/src/utils"
)

type Code struct {
	commonInfo
	MaxStack             uint16
	MaxLocals            uint16
	CodeLength           uint32
	Code                 []byte // jvm-bytecode
	ExceptionTableLength uint16
	ExceptionTable       ExceptionTable
	AttributesCount      uint16
	Attributes           []Attribute
}

func (c *Code) fillSpecificInfo(reader *utils.BytesReader, constantPool cp.ConstantPool) {
	c.MaxStack = reader.ReadUint16()
	c.MaxLocals = reader.ReadUint16()
	codeLen := reader.ReadUint32()
	if 0 >= codeLen && codeLen >= 65536 {
		logger.Panicf("code_length of Code attribute must be in range (0, 65536)")
	}
	c.CodeLength = codeLen
	c.Code = reader.ReadBytes(int(codeLen))
	c.ExceptionTableLength = reader.ReadUint16()
	c.ExceptionTable = parseExceptionTable(reader, c.ExceptionTableLength)
	c.AttributesCount = reader.ReadUint16()
	c.Attributes = Parse(reader, c.AttributesCount, constantPool)
}

type ExceptionTable []ExceptionTableEntry

type ExceptionTableEntry struct {
	// StartPc inclusive index of the first opcode of an instruction at which the exception handler is active
	StartPc uint16
	// EndPc exclusive index of the last opcode of an instruction at which the exception handler is active
	EndPc     uint16
	HandlerPc uint16
	// CatchType The constant_pool entry at that index must be a CONSTANT_Class_info. 0 for all exceptions handling
	CatchType cp.Index
}

func parseExceptionTable(reader *utils.BytesReader, length uint16) ExceptionTable {
	var table ExceptionTable = make([]ExceptionTableEntry, length)
	for i := uint16(0); i < length; i++ {
		startPc := reader.ReadUint16()
		endPc := reader.ReadUint16()
		if startPc >= endPc {
			// TODO: logger.Panicf()
		}
		table[i] = ExceptionTableEntry{
			StartPc:   startPc,
			EndPc:     endPc,
			HandlerPc: reader.ReadUint16(),
			CatchType: cp.Index(reader.ReadUint16()),
		}
	}
	return table
}
