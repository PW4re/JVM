package attributes

import (
	"jvm/src/class_file/acc"
	"jvm/src/class_file/cp"
	"jvm/src/utils"
)

type InnerClasses struct {
	commonInfo
	NumberOfClasses uint16
	Classes         []Class
}

func (ic *InnerClasses) fillSpecificInfo(reader *utils.BytesReader, _ cp.ConstantPool) {
	ic.NumberOfClasses = reader.ReadUint16()
	classes := make([]Class, ic.NumberOfClasses)
	for i, _ := range classes {
		classes[i] = Class{
			InnerClassInfoIndex:   cp.Index(reader.ReadUint16()),
			OuterClassInfoIndex:   cp.Index(reader.ReadUint16()),
			InnerNameIndex:        cp.Index(reader.ReadUint16()),
			InnerClassAccessFlags: acc.AccessFlags(reader.ReadUint16()),
		}
	}
	ic.Classes = classes
}

type Class struct {
	InnerClassInfoIndex   cp.Index
	OuterClassInfoIndex   cp.Index
	InnerNameIndex        cp.Index
	InnerClassAccessFlags acc.AccessFlags
}
