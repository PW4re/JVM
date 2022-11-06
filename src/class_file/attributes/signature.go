package attributes

import (
	"jvm/src/class_file/cp"
	"jvm/src/util"
)

type Signature struct {
	commonInfo
	SignatureIndex cp.Index
}

func (s *Signature) fillSpecificInfo(reader *util.BytesReader) {
	s.SignatureIndex = cp.Index(reader.ReadUint16())
}

func (s *Signature) GetValue() any {
	//TODO implement me
	panic("implement me")
}
