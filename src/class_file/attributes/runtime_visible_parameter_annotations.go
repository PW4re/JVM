package attributes

import (
	"jvm/src/class_file/cp"
	"jvm/src/utils"
)

type RuntimeVisibleParameterAnnotations struct {
	commonInfo
	NumParameters        uint16
	ParameterAnnotations []ParameterAnnotation
}

func (a *RuntimeVisibleParameterAnnotations) fillSpecificInfo(reader *utils.BytesReader, _ cp.ConstantPool) {
	a.NumParameters = reader.ReadUint16()
	a.ParameterAnnotations = make([]ParameterAnnotation, a.NumParameters)
	for i := range a.ParameterAnnotations {
		numAnnotations := reader.ReadUint16()
		annotations := make([]Annotation, numAnnotations)
		a.ParameterAnnotations[i] = ParameterAnnotation{
			NumAnnotations: numAnnotations,
			Annotations:    annotations,
		}
	}
}

type ParameterAnnotation struct {
	NumAnnotations uint16
	Annotations    []Annotation
}
