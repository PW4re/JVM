package fields

import (
	"jvm/src/class_file"
	"jvm/src/util"
)

type Info class_file.CommonFieldsAndMethods

func Parse(clsReader *util.BytesReader, count uint16) []Info {
	res := make([]Info, count)
	for i := range res {
		res[i] = Info(class_file.ParseCommonFieldsAndMethodsStructure(clsReader))
	}
	return res
}
