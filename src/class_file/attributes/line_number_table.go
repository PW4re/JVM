package attributes

import (
	"jvm/src/util"
)

type LineNumberTable struct {
	commonInfo
	LineNumberTableLength uint16
	Table                 []LineNumberTableEntry
}

func (a *LineNumberTable) fillSpecificInfo(reader *util.BytesReader) {
	a.LineNumberTableLength = reader.ReadUint16()
	a.Table = make([]LineNumberTableEntry, a.LineNumberTableLength)
	for i, _ := range a.Table {
		a.Table[i] = LineNumberTableEntry{
			StartPc:    reader.ReadUint16(),
			LineNumber: reader.ReadUint16(),
		}
	}
}

func (a *LineNumberTable) GetValue() any {
	//TODO implement me
	panic("implement me")
}

type LineNumberTableEntry struct {
	StartPc    uint16
	LineNumber uint16
}
