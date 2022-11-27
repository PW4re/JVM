package acc

type AccessFlags uint16

const (
	ACC_PUBLIC     AccessFlags = 0x0001
	ACC_FINAL      AccessFlags = 0x0010
	ACC_SUPER      AccessFlags = 0x0020
	ACC_INTERFACE  AccessFlags = 0x0200
	ACC_ABSTRACT   AccessFlags = 0x0400
	ACC_SYNTHETIC  AccessFlags = 0x1000
	ACC_ANNOTATION AccessFlags = 0x2000
	ACC_ENUM       AccessFlags = 0x4000
)
