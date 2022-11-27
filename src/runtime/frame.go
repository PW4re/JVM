package runtime

import (
	"jvm/src/runtime/heap"
	"jvm/src/runtime/thread"
)

type Frame struct {
	OperandStack
	LocalVariables
	Method    heap.Method
	Thread    *thread.Thread
	prev      *Frame
	maxLocals uint
	maxStack  uint
	NextPC    int
}
