package attributes

import (
	"jvm/src/class_file/acc"
	"jvm/src/class_file/cp"
)

type MethodParameters struct {
	commonInfo
	ParametersCount uint8
	Parameters      []Parameter
}

type Parameter struct {
	NameIndex   cp.Index
	AccessFlags acc.AccessFlags
}
