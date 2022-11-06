package class_file

import (
	"jvm/src/class_file/attributes"
	"jvm/src/class_file/cp"
	"jvm/src/util"
)

type CommonFieldsAndMethods struct {
	AccessFlags     AccessFlags
	NameIndex       cp.Index
	DescriptorIndex cp.Index
	AttributesCount uint16
	Attributes      []attributes.Attribute
}

func ParseCommonFieldsAndMethodsStructure(clsReader *util.BytesReader) CommonFieldsAndMethods {
	attributesCount := clsReader.ReadUint16()
	return CommonFieldsAndMethods{
		AccessFlags:     AccessFlags(clsReader.ReadUint16()),
		NameIndex:       cp.Index(clsReader.ReadUint16()),
		DescriptorIndex: cp.Index(clsReader.ReadUint16()),
		AttributesCount: attributesCount,
		Attributes:      attributes.Parse(clsReader, attributesCount),
	}
}
