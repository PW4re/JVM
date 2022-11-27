package fields

import (
	"jvm/src/class_file/attributes"
	"jvm/src/class_file/cp"
	"jvm/src/utils"
)

type Info attributes.CommonFieldsAndMethods

func Parse(clsReader *utils.BytesReader, count uint16, pool cp.ConstantPool) []Info {
	res := make([]Info, count)
	for i := range res {
		res[i] = Info(attributes.ParseCommonFieldsAndMethodsStructure(clsReader, pool))
	}
	return res
}
