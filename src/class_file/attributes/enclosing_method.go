package attributes

import (
	"jvm/src/class_file/cp"
	"jvm/src/util"
)

type EnclosingMethod struct {
	commonInfo
	ClassIndex  cp.Index
	MethodIndex cp.Index
}

func (em *EnclosingMethod) fillSpecificInfo(reader *util.BytesReader) {
	em.ClassIndex = cp.Index(reader.ReadUint16())
	em.MethodIndex = cp.Index(reader.ReadUint16())
}

func (em *EnclosingMethod) GetValue() any {
	//TODO implement me
	panic("implement me")
}
