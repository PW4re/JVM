package attributes

import "jvm/src/util"

type Synthetic struct {
	commonInfo
}

func (s *Synthetic) fillSpecificInfo(reader *util.BytesReader) {
}

func (s *Synthetic) GetValue() any {
	//TODO implement me
	panic("implement me")
}
