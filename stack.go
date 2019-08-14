package golang_stack

type Item interface{}
type Stack interface {
	Top() 	Item
	Pop() 	Item
	Push(Item)
	Size() 	int
}

type stack struct {
	data	[]Item
}

func (s *stack) Top() Item {
	return s.data[len(s.data)-1]
}

func (s *stack) Pop() Item {
	res := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1:len(s.data)-1]
	return res
}

func (s *stack) Push(newItem Item) {
	s.data = append(s.data, newItem)
}

func (s *stack) Size() int {
	return len(s.data)
}

func New() Stack {
	return &stack{}
}


