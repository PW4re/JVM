package constant_pool

import (
	"jvm/src/reader"
	"strings"
)

const (
	CONSTANT_Class              ConstantType = 7
	CONSTANT_Fieldref           ConstantType = 9
	CONSTANT_Methodref          ConstantType = 10
	CONSTANT_InterfaceMethodref ConstantType = 11
	CONSTANT_String             ConstantType = 8
	CONSTANT_Integer            ConstantType = 3
	CONSTANT_Float              ConstantType = 4
	CONSTANT_Long               ConstantType = 5
	CONSTANT_Double             ConstantType = 6
	CONSTANT_NameAndType        ConstantType = 12
	CONSTANT_Utf8               ConstantType = 1
	CONSTANT_MethodHandle       ConstantType = 15
	CONSTANT_MethodType         ConstantType = 16
	CONSTANT_InvokeDynamic      ConstantType = 18
)

type ConstantPool []ConstantInfo

type Index uint16 // TODO: migrate uint16 -> cp.Index

type ConstantType uint8

type ConstantInfo struct {
	Tag ConstantType
	info
}

type ParserFunc func(tag ConstantType, reader *reader.BytesReader) ConstantInfo

var ParseMap = map[ConstantType]ParserFunc{
	CONSTANT_Class:              parseClass,
	CONSTANT_Fieldref:           parseRef,
	CONSTANT_Methodref:          parseRef,
	CONSTANT_InterfaceMethodref: parseRef,
	CONSTANT_String:             parseString,
	CONSTANT_Integer:            parseInt,
	CONSTANT_Float:              parseFloat,
	CONSTANT_Long:               parseLong,
	CONSTANT_Double:             parseDouble,
	CONSTANT_NameAndType:        parseNameAndType,
	CONSTANT_Utf8:               parseUtf8,
	CONSTANT_MethodHandle:       parseMethodHandle,
	CONSTANT_MethodType:         parseMethodType,
	CONSTANT_InvokeDynamic:      parseInvokeDynamic,
}

func parseClass(_ ConstantType, reader *reader.BytesReader) ConstantInfo {
	return ConstantInfo{
		CONSTANT_Class,
		info{ClassInfo: ClassInfo{NameIndex: reader.ReadUint16()}},
	}
}

func parseRef(tag ConstantType, reader *reader.BytesReader) ConstantInfo {
	return ConstantInfo{
		tag,
		info{RefInfo: RefInfo{ClassIndex: reader.ReadUint16(), NameAndTypeIndex: reader.ReadUint16()}},
	}
}

func parseString(_ ConstantType, reader *reader.BytesReader) ConstantInfo {
	return ConstantInfo{
		CONSTANT_String,
		info{StringInfo: StringInfo{StringIndex: reader.ReadUint16()}},
	}
}

func parseInt(_ ConstantType, reader *reader.BytesReader) ConstantInfo {
	return ConstantInfo{
		CONSTANT_Integer,
		info{NumericInfo: NumericInfo{Integer: int32(reader.ReadUint32())}},
	}
}

func parseFloat(_ ConstantType, reader *reader.BytesReader) ConstantInfo {
	return ConstantInfo{
		CONSTANT_Float,
		info{NumericInfo: NumericInfo{Float: reader.ReadFloat()}},
	}
}

func parseLong(_ ConstantType, reader *reader.BytesReader) ConstantInfo {
	return ConstantInfo{
		CONSTANT_Long,
		info{NumericInfo: NumericInfo{Long: int64(reader.ReadUint64())}},
	}
}

func parseDouble(_ ConstantType, reader *reader.BytesReader) ConstantInfo {
	return ConstantInfo{
		CONSTANT_Double,
		info{NumericInfo: NumericInfo{Double: reader.ReadDouble()}},
	}
}

func parseNameAndType(_ ConstantType, reader *reader.BytesReader) ConstantInfo {
	return ConstantInfo{
		CONSTANT_NameAndType,
		info{NameAndTypeInfo: NameAndTypeInfo{NameIndex: reader.ReadUint16(), DescriptorIndex: reader.ReadUint16()}},
	}
}

func parseUtf8(_ ConstantType, reader *reader.BytesReader) ConstantInfo {
	length := reader.ReadUint16()
	bytes := reader.ReadBytes(int(length))
	var builder strings.Builder
	var x, y, z byte
	var surrogateIndicator byte = 0b11101101
	for idx := uint16(0); idx < length; {
		x = bytes[idx]
		y = bytes[idx+1]
		z = bytes[idx+2]
		if x>>7 == 0 {
			builder.WriteByte(x)
			idx++
			continue
		} else if x>>5 == 0b110 && y>>6 == 0b10 {
			builder.WriteRune((rune(x&0x1f) << 6) + rune(y&0x3f))
			idx += 2
			continue
		} else if x>>4 == 0b1110 && y>>6 == 0b10 && z>>6 == 0b10 {
			builder.WriteRune((rune(x&0xf) << 12) + (rune(y&0x3f) << 6) + rune(z&0x3f))
			idx += 3
			continue
		} else if x == surrogateIndicator && bytes[idx+3] == surrogateIndicator {
			builder.WriteRune(_processSupplementaryCharacters(x, y, z, bytes[idx+3], bytes[idx+4], bytes[idx+5]))
			idx += 6
		} else {
			logger.Panicf("Malformed byte sequence detected.\n Start index: %d", idx)
		}
	}

	return ConstantInfo{
		CONSTANT_Utf8,
		info{Utf8Info: Utf8Info{BytesLength: length, Value: builder.String()}}}
}

func _processSupplementaryCharacters(u, v, w, x, y, z byte) rune {
	return 0x10000 + (rune(v&0x0f) << 16) + (rune(w&0x3f) << 10) +
		(rune(y&0x0f) << 6) + rune(z&0x3f)
}

func parseMethodHandle(_ ConstantType, reader *reader.BytesReader) ConstantInfo {
	refKind := reader.ReadUint8()
	if 1 < refKind && refKind > 9 {
		logger.Panicf("Incorrect reference_kind value in CONSTANT_MethodHandle_info")
	}
	return ConstantInfo{
		CONSTANT_MethodHandle,
		info{MethodHandleInfo: MethodHandleInfo{ReferenceKind: refKind}},
	}
}

func parseMethodType(_ ConstantType, reader *reader.BytesReader) ConstantInfo {
	return ConstantInfo{
		CONSTANT_MethodType,
		info{MethodTypeInfo: MethodTypeInfo{DescriptorIndex: reader.ReadUint16()}},
	}
}

func parseInvokeDynamic(_ ConstantType, reader *reader.BytesReader) ConstantInfo {
	return ConstantInfo{
		CONSTANT_InvokeDynamic,
		info{InvokeDynamicInfo: InvokeDynamicInfo{
			BootstrapMethodAttrIndex: reader.ReadUint16(),
			NameAndTypeIndex:         reader.ReadUint16(),
		}},
	}
}
