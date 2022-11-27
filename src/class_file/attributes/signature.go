package attributes

import (
	"jvm/src/class_file/cp"
	"jvm/src/utils"
)

type Signature struct {
	commonInfo
	SignatureIndex cp.Index
}

func (s *Signature) fillSpecificInfo(reader *utils.BytesReader, _ cp.ConstantPool) {
	s.SignatureIndex = cp.Index(reader.ReadUint16())
}
