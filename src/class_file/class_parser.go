package class_file

import (
	"jvm/src/class_file/cp"
	"jvm/src/class_file/fields"
	"jvm/src/util"
)

var ConstantPool cp.ConstantPool

func Parse(class []byte) (ClassFile, error) {
	clsReader := util.NewBEAtStart(class)
	if clsReader.ReadUint32() != MAGIC {
		return ClassFile{}, VerificationError{}
	}

	minorVersion, majorVersion := clsReader.ReadUint16(), clsReader.ReadUint16()
	constantPoolCount := clsReader.ReadUint16()
	ConstantPool = parseConstantPool(&clsReader, constantPoolCount)
	accessFlags := AccessFlags(clsReader.ReadUint16())
	thisClass, superClass := cp.Index(clsReader.ReadUint16()), cp.Index(clsReader.ReadUint16())
	interfacesCount := clsReader.ReadUint16()
	interfaces := parseInterfaces(&clsReader, interfacesCount)
	fieldsCount := clsReader.ReadUint16()

	return ClassFile{
		MinorVersion:    minorVersion,
		MajorVersion:    majorVersion,
		CpCount:         constantPoolCount,
		ConstantPool:    ConstantPool,
		AccessFlags:     accessFlags,
		thisClass:       thisClass,
		SuperClass:      superClass,
		InterfacesCount: interfacesCount,
		Interfaces:      interfaces,
		FieldCount:      fieldsCount,
	}, nil
}

func parseConstantPool(clsReader *util.BytesReader, cpCount uint16) cp.ConstantPool {
	pool := cp.NewConstantPool()
	for entryIndex := uint16(0); entryIndex < cpCount; {
		tag := cp.ConstantType(clsReader.ReadUint8())
		parse := cp.ParseMap[tag]
		pool = append(pool, parse(tag, clsReader))
	}
	return pool
}

func parseInterfaces(clsReader *util.BytesReader, count uint16) []cp.Index {
	res := make([]cp.Index, count)
	for i := uint16(0); i < count; i++ {
		res[i] = cp.Index(clsReader.ReadUint16())
		if ConstantPool[res[i]].Tag != cp.CONSTANT_Class {
			logger.Panicf("Incorrect type at CONSTANT POOL index %d. CONSTANT_Class_info expected.", res[i])
		}
	}
	return res
}

func parseFields(clsReader *util.BytesReader, count uint16) []fields.Info {
	res := make([]fields.Info, count)
	for i := uint16(0); i < count; i++ {

	}
	return res
}
