package runtime

func newJVMStack(maxSize uint) *JVMStack {
	return &JVMStack{
		_head:   nil,
		maxSize: maxSize,
		size:    0,
	}
}

type JVMStack struct {
	_head   *Frame
	maxSize uint
	size    uint
}

func (s *JVMStack) push(elem *Frame) {
	if s.size >= s.maxSize {
		panic("StackOverflow")
	}
	if s._head != nil {
		elem.prev = s._head
	}
	s._head = elem
	s.size++
}

func (s *JVMStack) pop() *Frame {
	head := s._head
	s._head = s._head.prev
	s.size--
	return head
}

func (s *JVMStack) peek() *Frame {
	return s._head
}

func (s *JVMStack) isEmpty() bool {
	return s.size == 0
}
