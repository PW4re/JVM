package attributes

import (
	"jvm/src/class_file/cp"
	"jvm/src/utils"
)

type SourceFile struct {
	commonInfo
	SourceFileIndex cp.Index
}

func (sf *SourceFile) fillSpecificInfo(reader *utils.BytesReader, _ cp.ConstantPool) {
	sf.SourceFileIndex = cp.Index(reader.ReadUint16())
}
