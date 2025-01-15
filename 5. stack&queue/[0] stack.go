package stack_queue

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{data: []int{}}
}

func (s *Stack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		return -1
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

func (s *Stack) Peek() int {
	return s.data[len(s.data)-1]
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
