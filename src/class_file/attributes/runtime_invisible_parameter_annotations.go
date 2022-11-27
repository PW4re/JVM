package attributes

import (
	"jvm/src/class_file/cp"
	"jvm/src/utils"
)

type RuntimeInvisibleParameterAnnotations struct {
	commonInfo
	NumParameters        uint16
	ParameterAnnotations []ParameterAnnotation
}

func (a *RuntimeInvisibleParameterAnnotations) fillSpecificInfo(reader *utils.BytesReader, _ cp.ConstantPool) {
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
