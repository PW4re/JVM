package attributes

type LineNumberTable struct {
	commonInfo
	LineNumberTableLength uint16
	Table                 []LineNumberTableEntry
}

type LineNumberTableEntry struct {
	StartPc    uint16
	LineNumber uint16
}
