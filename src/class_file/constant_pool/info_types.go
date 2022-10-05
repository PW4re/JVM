package constant_pool

// embedding - наше всё!
type info struct {
	ClassInfo
	RefInfo
	StringInfo
	NumericInfo
	NameAndTypeInfo
	Utf8Info
	MethodHandleInfo
	MethodTypeInfo
	InvokeDynamicInfo
}

type ClassInfo struct{ NameIndex uint16 }

type RefInfo struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

type StringInfo struct {
	StringIndex uint16
}

type NumericInfo struct {
	Integer int32
	Float   float32
	Long    int64
	Double  float64
}

type NameAndTypeInfo struct {
	NameIndex       uint16
	DescriptorIndex uint16
}

type Utf8Info struct {
	BytesLength uint16
	Value       string
}

type MethodHandleInfo struct {
	ReferenceKind  uint8 // TODO: https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-5.html#jvms-5.4.3.5
	ReferenceIndex uint16
}

type MethodTypeInfo struct {
	DescriptorIndex uint16
}

type InvokeDynamicInfo struct {
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}
