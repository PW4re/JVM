package attributes

import (
	"jvm/src/class_file/cp"
	"jvm/src/utils"
)

type LocalVariableTable struct {
	commonInfo
	LocalVariableTableLength uint16
	Table                    []LocalVariableTableEntry
}

func (a *LocalVariableTable) fillSpecificInfo(reader *utils.BytesReader, _ cp.ConstantPool) {
	a.LocalVariableTableLength = reader.ReadUint16()
	a.Table = make([]LocalVariableTableEntry, a.LocalVariableTableLength)
	for i := range a.Table {
		a.Table[i] = LocalVariableTableEntry{
			StartPc:         reader.ReadUint16(),
			Length:          reader.ReadUint16(),
			NameIndex:       cp.Index(reader.ReadUint16()),
			DescriptorIndex: cp.Index(reader.ReadUint16()),
			Index:           reader.ReadUint16(),
		}
	}
}

type LocalVariableTableEntry struct {
	StartPc, Length uint16
	NameIndex       cp.Index
	DescriptorIndex cp.Index
	// Index Note: If the local variable at index is of type double or long, it occupies both index and index + 1
	Index uint16
}
