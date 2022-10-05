package attributes

import (
	"jvm/src/class_file"
	"jvm/src/class_file/constant_pool"
)

type MethodParameters struct {
	commonInfo
	ParametersCount uint8
	Parameters      []Parameter
}

type Parameter struct {
	NameIndex   constant_pool.Index
	AccessFlags class_file.AccessFlags
}
