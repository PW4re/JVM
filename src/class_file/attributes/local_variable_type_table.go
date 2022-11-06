package attributes

import (
	"jvm/src/class_file/cp"
	"jvm/src/util"
)

type LocalVariableTypeTable struct {
	commonInfo
	LocalVariableTypeTableLength uint16
	Table                        []LocalVariableTypeTableEntry
}

func (a *LocalVariableTypeTable) fillSpecificInfo(reader *util.BytesReader) {
	a.LocalVariableTypeTableLength = reader.ReadUint16()
	a.Table = make([]LocalVariableTypeTableEntry, a.LocalVariableTypeTableLength)
	for i := range a.Table {
		a.Table[i] = LocalVariableTypeTableEntry{
			StartPc:        reader.ReadUint16(),
			Length:         reader.ReadUint16(),
			NameIndex:      cp.Index(reader.ReadUint16()),
			SignatureIndex: cp.Index(reader.ReadUint16()),
			Index:          reader.ReadUint16(),
		}
	}
}

type LocalVariableTypeTableEntry struct {
	StartPc, Length uint16
	NameIndex       cp.Index
	SignatureIndex  cp.Index
	// Index Note: If the local variable at index is of type double or long, it occupies both index and index + 1
	Index uint16
}
