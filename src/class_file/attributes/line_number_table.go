package attributes

import (
	"jvm/src/class_file/cp"
	"jvm/src/utils"
)

type LineNumberTable struct {
	commonInfo
	LineNumberTableLength uint16
	Table                 []LineNumberTableEntry
}

func (a *LineNumberTable) fillSpecificInfo(reader *utils.BytesReader, _ cp.ConstantPool) {
	a.LineNumberTableLength = reader.ReadUint16()
	a.Table = make([]LineNumberTableEntry, a.LineNumberTableLength)
	for i, _ := range a.Table {
		a.Table[i] = LineNumberTableEntry{
			StartPc:    reader.ReadUint16(),
			LineNumber: reader.ReadUint16(),
		}
	}
}

type LineNumberTableEntry struct {
	StartPc    uint16
	LineNumber uint16
}
