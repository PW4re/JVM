package attributes

import (
	"jvm/src/class_file/cp"
	"jvm/src/utils"
)

type Deprecated struct {
	commonInfo
}

func (a *Deprecated) fillSpecificInfo(reader *utils.BytesReader, _ cp.ConstantPool) {
	// nothing to do
}
