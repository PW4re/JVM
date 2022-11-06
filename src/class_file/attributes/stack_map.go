package attributes

import (
	"jvm/src/util"
)

type StackMapTable struct {
	commonInfo
	NumberOfEntries uint16
	// TODO: здесь должна быть сама таблица, но пока пропускаю - шибко сложно
}

func (a *StackMapTable) fillSpecificInfo(reader *util.BytesReader) {
	//TODO implement me
	panic("implement me")
}
