package attributes

import "jvm/src/class_file"

type InnerClasses struct {
	commonInfo
	NumberOfClasses uint16
	Classes         []Class
}

type Class struct {
	InnerClassInfoIndex   uint16
	OuterClassInfoIndex   uint16
	InnerNameIndex        uint16
	InnerClassAccessFlags class_file.AccessFlags
}
