package attributes

type LocalVariableTable struct {
	commonInfo
	LocalVariableTableLength uint16
	Table                    []LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	StartPc, Length uint16
	NameIndex       uint16
	DescriptorIndex uint16
	Index           uint16 // TODO: If the local variable at index is of type double or long, it occupies both index and index + 1.
}
