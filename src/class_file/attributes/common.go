package attributes

import (
	"jvm/src/class_file/cp"
	"jvm/src/utils"
)

var nameMap = map[string]func() Attribute{
	"ConstantValue":                        func() Attribute { return &ConstantValue{} },
	"Code":                                 func() Attribute { return &Code{} },
	"StackMapTable":                        func() Attribute { return &StackMapTable{} },
	"Exceptions":                           func() Attribute { return &Exceptions{} },
	"InnerClasses":                         func() Attribute { return &InnerClasses{} },
	"EnclosingMethod":                      func() Attribute { return &EnclosingMethod{} },
	"Synthetic":                            func() Attribute { return &Synthetic{} },
	"Signature":                            func() Attribute { return &Signature{} },
	"SourceFile":                           func() Attribute { return &SourceFile{} },
	"SourceDebugExtension":                 func() Attribute { return &SourceDebugExtension{} },
	"LineNumberTable":                      func() Attribute { return &LineNumberTable{} },
	"LocalVariableTable":                   func() Attribute { return &LocalVariableTable{} },
	"LocalVariableTypeTable":               func() Attribute { return &LocalVariableTypeTable{} },
	"Deprecated":                           func() Attribute { return &Deprecated{} },
	"RuntimeVisibleAnnotations":            func() Attribute { return &RuntimeVisibleAnnotations{} },
	"RuntimeInvisibleAnnotations":          func() Attribute { return &RuntimeInvisibleAnnotations{} },
	"RuntimeVisibleParameterAnnotations":   func() Attribute { return &RuntimeVisibleParameterAnnotations{} },
	"RuntimeInvisibleParameterAnnotations": func() Attribute { return &RuntimeInvisibleParameterAnnotations{} },
	"RuntimeVisibleTypeAnnotations":        func() Attribute { return &RuntimeVisibleTypeAnnotations{} },
	//"AnnotationDefault":&AnnotationDefault{},
	//"BootstrapMethods":&BootstrapMethods{},
	//"MethodParameters":&MethodParameters{},
}

type Attribute interface {
	// GetName returns attribute name that identifies it in JVM
	GetName(constantPool cp.ConstantPool) string
	// fillSpecificInfo fills in all attribute-specific fields. Doesn't fill in commonInfo
	fillSpecificInfo(reader *utils.BytesReader, constantPool cp.ConstantPool)
	//GetValue() any
	// fillCommonInfo fills in only commonInfo.AttributeNameIndex and commonInfo.AttributeLength
	fillCommonInfo(nameIndex cp.Index, length uint32)
}

type commonInfo struct {
	AttributeNameIndex cp.Index
	AttributeLength    uint32
}

func (info *commonInfo) GetName(constantPool cp.ConstantPool) string {
	cpEntry := constantPool[info.AttributeNameIndex]
	if cpEntry.Tag != cp.CONSTANT_Utf8 {
		logger.Panicf("Incorrect type at CONSTANT POOL index %d. CONSTANT_Utf8_info expected.", info.AttributeNameIndex)
	}
	return cpEntry.Utf8Info.Value
}

func (info *commonInfo) fillCommonInfo(nameIndex cp.Index, length uint32) {
	info.AttributeNameIndex = nameIndex
	info.AttributeLength = length
}

// readAttribute returns completed Attribute
func readAttribute(reader *utils.BytesReader, constantPool cp.ConstantPool) (attr Attribute) {
	index := cp.Index(reader.ReadUint16())
	cpEntry := constantPool[index]
	length := reader.ReadUint32()
	if cpEntry.Tag != cp.CONSTANT_Utf8 {
		logger.Panicf("Incorrect type at CONSTANT POOL index %d. CONSTANT_Utf8_info expected.", index)
		return nil
	}
	if f, ok := nameMap[cpEntry.Utf8Info.Value]; ok {
		attribute := f()
		attribute.fillCommonInfo(index, length)
		attribute.fillSpecificInfo(reader, constantPool)
		return attribute
	}
	return nil // TODO: strange situation, panic
}

func Parse(reader *utils.BytesReader, count uint16, constantPool cp.ConstantPool) []Attribute {
	res := make([]Attribute, count)
	for i := uint16(0); i < count; i++ {
		res[i] = readAttribute(reader, constantPool)
	}
	return res
}
