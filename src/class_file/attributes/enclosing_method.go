package attributes

import (
	"jvm/src/class_file/cp"
	"jvm/src/utils"
)

type EnclosingMethod struct {
	commonInfo
	ClassIndex  cp.Index
	MethodIndex cp.Index
}

func (em *EnclosingMethod) fillSpecificInfo(reader *utils.BytesReader, _ cp.ConstantPool) {
	em.ClassIndex = cp.Index(reader.ReadUint16())
	em.MethodIndex = cp.Index(reader.ReadUint16())
}
