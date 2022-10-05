package attributes

type RuntimeVisibleParameterAnnotations struct {
	commonInfo
	NumParameters        uint16
	ParameterAnnotations []ParameterAnnotation
}

type ParameterAnnotation struct {
	NumAnnotations uint16
	Annotations    []Annotation
}
