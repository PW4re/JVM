package runtime

type OperandStack struct {
	maxDepth uint
	curr     uint
	stack    []Value
}

func newEmptyOperandStack(maxDepth uint) OperandStack {
	return OperandStack{
		maxDepth: maxDepth,
		curr:     0,
		stack:    make([]Value, maxDepth, maxDepth),
	}
}

func newOperandStack(elements []Value) OperandStack {
	depth := uint(len(elements))
	return OperandStack{
		maxDepth: depth,
		curr:     depth - 1,
		stack:    elements,
	}
}

func (s *OperandStack) Push(elem Value) {
	if s.curr+1 == s.maxDepth {
		panic("StackOverflow: operand stack overflow")
	}
	s.curr++
	s.stack[s.curr] = elem
}

func (s *OperandStack) Pop() Value {
	peek := s.Peek()
	s.stack[s.curr] = Empty{}
	s.curr--
	return peek
}

func (s *OperandStack) Peek() Value {
	return s.stack[s.curr]
}

func (s *OperandStack) IsEmpty() bool {
	return len(s.stack) == 0
}
