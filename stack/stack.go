package stack

// Stack todo
type Stack struct {
	data []interface{}
}

// Push todo
func (s *Stack) Push(item interface{}) {
	s.data = append(s.data, item)
}

// Pop todo
func (s *Stack) Pop() interface{} {
	last := len(s.data) - 1
	item := s.data[last]
	s.data = s.data[:last]
	return item
}

// Size todo
func (s *Stack) Size() int {
	return len(s.data)
}

// IsEmpty todo
func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}
