package class_file

import "jvm/src/class_file/constant_pool"

const MAGIC = 0xCAFEBABE

type ClassFile struct {
	minorVersion, majorVersion uint16
	cpCount                    uint16
	constantPool               constant_pool.ConstantPool
	accessFlags                AccessFlags
	thisClass, superClass      uint16
	interfacesCount            uint16
	interfaces                 []uint16
	fieldCount                 uint16
	// TODO: fields here
	methodsCount uint16
	// TODO methods here
	attributesCount uint16
	// TODO: attributes here
}
