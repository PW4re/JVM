package attributes

import "jvm/src/class_file/constant_pool"

type RuntimeVisibleAnnotations struct {
	commonInfo
	NumAnnotations uint16
	Annotations    []Annotation
}

type Annotation struct {
	TypeIndex            constant_pool.Index
	NumElementValuePairs uint16
}

type ElementValuePair struct {
	ElementNameIndex constant_pool.Index
	Value            ElementValue
}

type ElementValue struct {
	// TODO: 4.7.16.1.
}
