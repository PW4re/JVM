package attributes

import (
	"jvm/src/class_file/cp"
	"jvm/src/utils"
)

type ConstantValue struct {
	commonInfo         // length must be two
	ConstantValueIndex cp.Index
}

// GetValue TODO: подумать о нужности метода
//func (v *ConstantValue) GetValue() any {
//	cpEntry := class_file.ConstantPool[v.ConstantValueIndex]
//	switch cpEntry.Tag {
//	case cp.CONSTANT_Long:
//		return cpEntry.Long
//	case cp.CONSTANT_Float:
//		return cpEntry.Float
//	case cp.CONSTANT_Double:
//		return cpEntry.Double
//	case cp.CONSTANT_Integer:
//		return cpEntry.Integer
//	case cp.CONSTANT_String:
//		index := cpEntry.StringIndex
//		cpEntry = class_file.ConstantPool[index]
//		if cpEntry.Tag != cp.CONSTANT_Utf8 {
//			logger.Panicf("Incorrect type at CONSTANT POOL index %d. CONSTANT_Utf8_info expected.", index)
//		}
//		return cpEntry.Value
//	default:
//		logger.Panicf("Incorrect type at CONSTANT POOL index %d. One of %v expected.",
//			v.ConstantValueIndex,
//			[]cp.ConstantType{cp.CONSTANT_Long, cp.CONSTANT_Float, cp.CONSTANT_Double, cp.CONSTANT_Integer, cp.CONSTANT_String})
//		return nil
//	}
//}

func (v *ConstantValue) fillSpecificInfo(reader *utils.BytesReader, _ cp.ConstantPool) {
	v.ConstantValueIndex = cp.Index(reader.ReadUint16())
}
