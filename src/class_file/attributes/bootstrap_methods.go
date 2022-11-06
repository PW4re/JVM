package attributes

import "jvm/src/class_file/cp"

type BootstrapMethods struct {
	commonInfo
	NumBootstrapMethods uint16
	//Methods TODO 4.7.23
}

type BootstrapMethod struct {
	BootstrapMethodRef    cp.Index
	NumBootstrapArguments uint16
	BootstrapArguments    []cp.Index
}
