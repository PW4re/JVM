package attributes

import (
	cp "jvm/src/class_file/cp"
	"jvm/src/utils"
)

type RuntimeVisibleAnnotations struct {
	commonInfo
	NumAnnotations uint16
	Annotations    []Annotation
}

func (a *RuntimeVisibleAnnotations) fillSpecificInfo(reader *utils.BytesReader, _ cp.ConstantPool) {
	a.NumAnnotations = reader.ReadUint16()
	a.Annotations = make([]Annotation, a.NumAnnotations)
	for i := range a.Annotations {
		a.Annotations[i] = Annotation{
			TypeIndex: cp.Index(reader.ReadUint16()),
		}
	}
}

type Annotation struct {
	TypeIndex            cp.Index
	NumElementValuePairs uint16
	ElementValuePairs    []ElementValuePair
}

type ElementValuePair struct {
	ElementNameIndex cp.Index
	Value            ElementValue
}

type ElementValue struct {
	Tag byte
	// Value must contain one and only one non-nil field
	Value struct {
		ConstValueIndex cp.Index
		//EnumConstValue  struct {
		//	TypeNameIndex  cp.Index
		//	ConstNameIndex cp.Index
		//}
		//ClassInfoIndex  cp.Index
		//AnnotationValue Annotation
		//ArrayType       struct {
		//	NumValues uint16
		//	Values    []ElementValue
		//}
	}
}

// Table 4.7.16.1-A. Interpretation of tag values as types
// tag 		Type 	    		value 	 			Constant type
var possibleElementTags = []byte{
	'B', //	byte			 	const_value_index 	CONSTANT_Integer
	'C', //	char			 	const_value_index 	CONSTANT_Integer
	'D', // double			 	const_value_index 	CONSTANT_Double
	'F', // float			 	const_value_index 	CONSTANT_Float
	'I', // int				 	const_value_index 	CONSTANT_Integer
	'J', // long 				const_value_index 	CONSTANT_Long
	'S', //	short			 	const_value_index 	CONSTANT_Integer
	'Z', //	boolean 			const_value_index 	CONSTANT_Integer
	's', //	String 				const_value_index 	CONSTANT_Utf8
	'e', //	Enum type 			enum_const_value 	Not applicable
	'c', //	Class 				class_info_index 	Not applicable
	'@', //	Annotation type		annotation_value 	Not applicable
	'[', // Array type 	        array_value			Not applicable
}

func parseAnnotation(reader *utils.BytesReader) (ann Annotation) {
	ann.TypeIndex = cp.Index(reader.ReadUint16())
	ann.NumElementValuePairs = reader.ReadUint16()
	ann.ElementValuePairs = make([]ElementValuePair, ann.NumElementValuePairs)
	for i := range ann.ElementValuePairs {
		elementNameIndex := cp.Index(reader.ReadUint16())
		elementValueIndex := reader.ReadUint8()
		// TODO: check elementValueIndex in possibleElementTags (от него зависит как парсить Value)
		ann.ElementValuePairs[i] = ElementValuePair{
			ElementNameIndex: elementNameIndex,
			Value: ElementValue{
				Tag: elementValueIndex,
				Value: struct {
					ConstValueIndex cp.Index
				}{
					ConstValueIndex: cp.Index(reader.ReadUint16()),
				},
			},
		}
	}

	return ann
}
