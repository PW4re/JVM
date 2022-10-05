package attributes

import "jvm/src/class_file/constant_pool"

type commonInfo struct { // todo: migrate (two fields) -> (embedded commonInfo)
	AttributeNameIndex constant_pool.Index
	AttributeLength    uint32
}
