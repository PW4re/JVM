package attributes

import (
	"jvm/src/util"
)

type SourceDebugExtension struct {
	commonInfo
	// TODO: debug extension here
}

func (sde *SourceDebugExtension) fillSpecificInfo(reader *util.BytesReader) {
	//TODO implement me
	panic("implement me")
}

func (sde *SourceDebugExtension) GetValue() any {
	//TODO implement me
	panic("implement me")
}
