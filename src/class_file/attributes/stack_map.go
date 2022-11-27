package attributes

import (
	"jvm/src/class_file/cp"
	"jvm/src/utils"
)

// StackMapTable TODO
type StackMapTable struct {
	commonInfo
	NumberOfEntries uint16
	Entries         []StackMapFrame
}

func (a *StackMapTable) fillSpecificInfo(reader *utils.BytesReader, _ cp.ConstantPool) {
	//TODO implement me
	panic("implement me")
}

type StackMapFrame struct {
}

const ( // verification tag
	ITEM_Top               = 0
	ITEM_Integer           = 1
	ITEM_Float             = 2
	ITEM_Null              = 5
	ITEM_UninitializedThis = 6
	ITEM_Object            = 7
	ITEM_Uninitialized     = 8
	ITEM_Long              = 4
	ITEM_Double            = 3
)

type verificationTypeInfo struct {
}
