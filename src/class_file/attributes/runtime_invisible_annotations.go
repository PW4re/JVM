package attributes

import (
	"jvm/src/class_file/cp"
	"jvm/src/util"
)

type RuntimeInvisibleAnnotations struct {
	commonInfo
	NumAnnotations uint16
	Annotations    []Annotation
}

func (a *RuntimeInvisibleAnnotations) fillSpecificInfo(reader *util.BytesReader) {
	a.NumAnnotations = reader.ReadUint16()
	a.Annotations = make([]Annotation, a.NumAnnotations)
	for i := range a.Annotations {
		a.Annotations[i] = Annotation{
			TypeIndex: cp.Index(reader.ReadUint16()),
		}
	}
}
