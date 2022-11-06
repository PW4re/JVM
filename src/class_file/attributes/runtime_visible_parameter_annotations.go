package attributes

import (
	"jvm/src/util"
)

type RuntimeVisibleParameterAnnotations struct {
	commonInfo
	NumParameters        uint16
	ParameterAnnotations []ParameterAnnotation
}

func (a *RuntimeVisibleParameterAnnotations) fillSpecificInfo(reader *util.BytesReader) {
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
