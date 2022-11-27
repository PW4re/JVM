package attributes

import (
	"jvm/src/class_file/cp"
	"jvm/src/utils"
)

type RuntimeInvisibleAnnotations struct {
	commonInfo
	NumAnnotations uint16
	Annotations    []Annotation
}

func (a *RuntimeInvisibleAnnotations) fillSpecificInfo(reader *utils.BytesReader, _ cp.ConstantPool) {
	a.NumAnnotations = reader.ReadUint16()
	a.Annotations = make([]Annotation, a.NumAnnotations)
	for i := range a.Annotations {
		a.Annotations[i] = Annotation{
			TypeIndex: cp.Index(reader.ReadUint16()),
		}
	}
}
