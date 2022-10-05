package class_file

import (
	"jvm/src/class_file/constant_pool"
	"jvm/src/reader"
)

func parse() {

}

func Parse(class []byte) (ClassFile, error) {
	clsReader := reader.NewBEAtStart(&class)
	if clsReader.ReadUint32() != MAGIC {
		return ClassFile{}, VerificationError{}
	}

	minorVersion, majorVersion := clsReader.ReadUint16(), clsReader.ReadUint16()
	constantPoolCount := clsReader.ReadUint16()
	constantPool := parseConstantPool(&clsReader, constantPoolCount)
	accessFlags := AccessFlags(clsReader.ReadUint16())
	thisClass, superClass := clsReader.ReadUint16(), clsReader.ReadUint16()
	interfacesCount := clsReader.ReadUint16()

	return ClassFile{
		minorVersion:    minorVersion,
		majorVersion:    majorVersion,
		cpCount:         constantPoolCount,
		constantPool:    constantPool,
		accessFlags:     accessFlags,
		thisClass:       thisClass,
		superClass:      superClass,
		interfacesCount: interfacesCount,
	}, nil
}

func parseConstantPool(clsReader *reader.BytesReader, cpCount uint16) constant_pool.ConstantPool {
	pool := constant_pool.ConstantPool{}
	for entryIndex := uint16(0); entryIndex < cpCount; {
		tag := constant_pool.ConstantType(clsReader.ReadUint8())
		parse := constant_pool.ParseMap[tag]
		pool[tag] = parse(tag, clsReader)
	}
	return pool
}
