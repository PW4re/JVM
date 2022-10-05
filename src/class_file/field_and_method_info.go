package class_file

import "jvm/src/class_file/attributes"

type commonStructure struct {
	AccessFlags     AccessFlags
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes      attributes.AttributesInfo
}

type FieldInfo commonStructure

type MethodInfo commonStructure
