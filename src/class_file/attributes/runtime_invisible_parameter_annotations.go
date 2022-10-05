package attributes

type RuntimeInvisibleParameterAnnotations struct {
	commonInfo
	NumParameters        uint16
	ParameterAnnotations []ParameterAnnotation
}
