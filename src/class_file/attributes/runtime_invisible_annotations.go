package attributes

type RuntimeInvisibleAnnotations struct {
	commonInfo
	NumAnnotations uint16
	Annotations    []Annotation
}
