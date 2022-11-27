package attributes

import (
	"jvm/src/class_file/cp"
	"jvm/src/utils"
)

type Synthetic struct {
	commonInfo
}

func (s *Synthetic) fillSpecificInfo(reader *utils.BytesReader, _ cp.ConstantPool) {
}
