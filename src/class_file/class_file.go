package class_file

import (
	"jvm/src/class_file/acc"
	"jvm/src/class_file/attributes"
	"jvm/src/class_file/cp"
	"jvm/src/class_file/fields"
	"jvm/src/class_file/methods"
)

const MAGIC = 0xCAFEBABE

type ClassFile struct {
	MinorVersion, MajorVersion uint16
	CpCount                    uint16
	ConstantPool               cp.ConstantPool
	AccessFlags                acc.AccessFlags
	ThisClass, SuperClass      cp.Index
	InterfacesCount            uint16
	Interfaces                 []cp.Index
	FieldCount                 uint16
	Fields                     []fields.Info
	MethodsCount               uint16
	Methods                    []methods.Info
	AttributesCount            uint16
	Attributes                 []attributes.Attribute
}
