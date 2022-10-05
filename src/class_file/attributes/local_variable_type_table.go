package attributes

import "jvm/src/class_file/constant_pool"

type LocalVariableTypeTable struct {
	commonInfo
	LocalVariableTypeTableLength uint16
	Table                        []LocalVariableTypeTableEntry
}

type LocalVariableTypeTableEntry struct {
	StartPc, Length uint16
	NameIndex       constant_pool.Index
	SignatureIndex  constant_pool.Index
	Index           uint16 // TODO: If the local variable at index is of type double or long, it occupies both index and index + 1.
}
