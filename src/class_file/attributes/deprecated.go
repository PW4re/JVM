package attributes

import (
	"jvm/src/util"
)

type Deprecated struct {
	commonInfo
}

func (a *Deprecated) fillSpecificInfo(reader *util.BytesReader) {
	// nothing to do
}
