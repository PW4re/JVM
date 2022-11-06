package attributes

import (
	"jvm/src/util"
)

type RuntimeInvisibleParameterAnnotations struct {
	commonInfo
	NumParameters        uint16
	ParameterAnnotations []ParameterAnnotation
}

func (a *RuntimeInvisibleParameterAnnotations) fillSpecificInfo(reader *util.BytesReader) {
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
