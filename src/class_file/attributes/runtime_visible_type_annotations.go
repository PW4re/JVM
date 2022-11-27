package attributes

import (
	"jvm/src/class_file/cp"
	"jvm/src/utils"
)

type RuntimeVisibleTypeAnnotations struct {
	commonInfo
	NumAnnotations uint16
	//Annotations  TODO: шибко сложно - потом §4.7.20.
}

type TargetPath struct {
}

type TypeAnnotation struct {
	TargetType byte
	//TargetInfo
	TargetPath           TargetPath
	TypeIndex            cp.Index
	NumElementValuePairs uint16
	ElementValuePairs    []ElementValuePair
}

func (a *RuntimeVisibleTypeAnnotations) fillSpecificInfo(reader *utils.BytesReader, _ cp.ConstantPool) {

}
