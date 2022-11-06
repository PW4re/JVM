package attributes

import (
	"jvm/src/class_file"
	"jvm/src/class_file/cp"
	"jvm/src/util"
)

type InnerClasses struct {
	commonInfo
	NumberOfClasses uint16
	Classes         []Class
}

func (ic *InnerClasses) fillSpecificInfo(reader *util.BytesReader) {
	ic.NumberOfClasses = reader.ReadUint16()
	classes := make([]Class, ic.NumberOfClasses)
	for i, _ := range classes {
		classes[i] = Class{
			InnerClassInfoIndex:   cp.Index(reader.ReadUint16()),
			OuterClassInfoIndex:   cp.Index(reader.ReadUint16()),
			InnerNameIndex:        cp.Index(reader.ReadUint16()),
			InnerClassAccessFlags: class_file.AccessFlags(reader.ReadUint16()),
		}
	}
	ic.Classes = classes
}

func (ic *InnerClasses) GetValue() any {
	//TODO implement me
	panic("implement me")
}

type Class struct {
	InnerClassInfoIndex   cp.Index
	OuterClassInfoIndex   cp.Index
	InnerNameIndex        cp.Index
	InnerClassAccessFlags class_file.AccessFlags
}
