package attributes

import "jvm/src/class_file/constant_pool"

type BootstrapMethods struct {
	commonInfo
	NumBootstrapMethods uint16
	//Methods TODO 4.7.23
}

type BootstrapMethod struct {
	BootstrapMethodRef    constant_pool.Index
	NumBootstrapArguments uint16
	BootstrapArguments    []constant_pool.Index
}
