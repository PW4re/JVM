package attributes

import (
	"jvm/src/class_file/acc"
	"jvm/src/class_file/cp"
	"jvm/src/utils"
)

type CommonFieldsAndMethods struct {
	AccessFlags     acc.AccessFlags
	NameIndex       cp.Index
	DescriptorIndex cp.Index
	AttributesCount uint16
	Attributes      []Attribute
}

func ParseCommonFieldsAndMethodsStructure(clsReader *utils.BytesReader, pool cp.ConstantPool) CommonFieldsAndMethods {
	res := CommonFieldsAndMethods{
		AccessFlags:     acc.AccessFlags(clsReader.ReadUint16()),
		NameIndex:       cp.Index(clsReader.ReadUint16()),
		DescriptorIndex: cp.Index(clsReader.ReadUint16()),
		AttributesCount: clsReader.ReadUint16(),
	}
	res.Attributes = Parse(clsReader, res.AttributesCount, pool)
	return res
}
