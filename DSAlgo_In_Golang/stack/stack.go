package stack

type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{items: make([]int, 0)}
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	item := s.items[s.Len()-1]
	s.items = s.items[0: s.Len()-1]
	return item
}

func (s *Stack) Len() int {
	return len(s.items)
}

func (s *Stack) Peek() int {
	return s.items[s.Len()-1]
}
