package attributes

import (
	"jvm/src/class_file/cp"
	"jvm/src/utils"
)

type Exceptions struct {
	commonInfo
	NumberOfExceptions  uint16
	ExceptionIndexTable []cp.Index
}

func (a *Exceptions) fillSpecificInfo(reader *utils.BytesReader, _ cp.ConstantPool) {
	a.NumberOfExceptions = reader.ReadUint16()
	a.ExceptionIndexTable = make([]cp.Index, a.NumberOfExceptions)
	for i, _ := range a.ExceptionIndexTable {
		a.ExceptionIndexTable[i] = cp.Index(reader.ReadUint16())
	}
}
