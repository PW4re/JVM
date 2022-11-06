package attributes

import (
	"jvm/src/class_file/cp"
	"jvm/src/util"
)

type SourceFile struct {
	commonInfo
	SourceFileIndex cp.Index
}

func (sf *SourceFile) fillSpecificInfo(reader *util.BytesReader) {
	sf.SourceFileIndex = cp.Index(reader.ReadUint16())
}

func (sf *SourceFile) GetValue() any {
	//TODO implement me
	panic("implement me")
}
