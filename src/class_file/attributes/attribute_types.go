package attributes

type AttributesInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	Info
}

type Info struct {
}

type ConstantValue struct {
	AttributeNameIndex uint16
	AttributesLength   uint32 // must be two
	ConstantValueIndex uint16
}

type Code struct {
	AttributeNameIndex   uint16
	AttributeLength      uint32
	MaxStack             uint16
	MaxLocals            uint16
	CodeLength           uint32
	Code                 []byte // jvm-bytecode
	ExceptionTableLength uint16
	ExceptionTable       ExceptionTable
	AttributesCount      uint16
	Attributes           []AttributesInfo
}

type ExceptionTable []ExceptionTableEntry

type ExceptionTableEntry struct {
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType uint16
}

type StackMapTable struct {
	AttributeNameIndex uint16
	AttributesLength   uint32
	NumberOfEntries    uint16
	// TODO: здесь должна быть сама таблица, но пока пропускаю - шибко сложно
}

type Exceptions struct {
	AttributeNameIndex  uint16
	AttributeLength     uint32
	NumberOfExceptions  uint16
	ExceptionIndexTable []uint16
}
