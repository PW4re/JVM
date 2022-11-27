package attributes

import (
	"jvm/src/class_file/cp"
	"jvm/src/utils"
)

type SourceDebugExtension struct {
	commonInfo
	// TODO: debug extension here
}

func (sde *SourceDebugExtension) fillSpecificInfo(reader *utils.BytesReader, _ cp.ConstantPool) {
	//TODO implement me
	panic("implement me")
}
