package class_file

import (
	"jvm/src/class_file/acc"
	"jvm/src/class_file/attributes"
	"jvm/src/class_file/cp"
	"jvm/src/class_file/fields"
	"jvm/src/class_file/methods"
	"jvm/src/utils"
)

func ParseClassFile(class []byte) (ClassFile, error) {
	clsReader := utils.NewBEAtStart(class)
	if clsReader.ReadUint32() != MAGIC {
		return ClassFile{}, attributes.VerificationError{}
	}

	minorVersion, majorVersion := clsReader.ReadUint16(), clsReader.ReadUint16()
	constantPoolCount := clsReader.ReadUint16()
	constantPool := parseConstantPool(&clsReader, constantPoolCount)
	accessFlags := acc.AccessFlags(clsReader.ReadUint16())
	thisClass, superClass := cp.Index(clsReader.ReadUint16()), cp.Index(clsReader.ReadUint16())
	interfacesCount := clsReader.ReadUint16()
	interfaces := parseInterfaces(&clsReader, interfacesCount, constantPool)
	fieldsCount := clsReader.ReadUint16()
	_fields := fields.Parse(&clsReader, fieldsCount, constantPool)
	methodsCount := clsReader.ReadUint16()
	_methods := methods.Parse(&clsReader, methodsCount, constantPool)
	attributesCount := clsReader.ReadUint16()
	_attributes := attributes.Parse(&clsReader, attributesCount, constantPool)

	return ClassFile{
		MinorVersion:    minorVersion,
		MajorVersion:    majorVersion,
		CpCount:         constantPoolCount,
		ConstantPool:    constantPool,
		AccessFlags:     accessFlags,
		ThisClass:       thisClass,
		SuperClass:      superClass,
		InterfacesCount: interfacesCount,
		Interfaces:      interfaces,
		FieldCount:      fieldsCount,
		Fields:          _fields,
		MethodsCount:    methodsCount,
		Methods:         _methods,
		AttributesCount: attributesCount,
		Attributes:      _attributes,
	}, nil
}

func parseConstantPool(clsReader *utils.BytesReader, cpCount uint16) cp.ConstantPool {
	pool := cp.NewConstantPool()
	for entryIndex := uint16(0); entryIndex < cpCount-1; entryIndex++ {
		tag := cp.ConstantType(clsReader.ReadUint8())
		parse := cp.ParseMap[tag]
		pool = append(pool, parse(tag, clsReader))
	}
	return pool
}

func parseInterfaces(clsReader *utils.BytesReader, count uint16, constantPool cp.ConstantPool) []cp.Index {
	res := make([]cp.Index, count)
	for i := uint16(0); i < count; i++ {
		res[i] = cp.Index(clsReader.ReadUint16())
		if constantPool[res[i]].Tag != cp.CONSTANT_Class {
			logger.Panicf("Incorrect type at CONSTANT POOL index %d. CONSTANT_Class_info expected.", res[i])
		}
	}
	return res
}
