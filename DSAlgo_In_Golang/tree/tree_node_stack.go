package tree

type Stack struct {
	items []*Node
}

func NewTreeNodeStack() *Stack {
	return &Stack{items: make([]*Node, 0)}
}

func (s *Stack) Push(item *Node) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() *Node {
	item := s.items[s.Len()-1]
	s.items = s.items[0: s.Len()-1]
	return item
}

func (s *Stack) PopLeft() *Node {
	item := s.items[0]
	s.items = s.items[1:]
	return item
}

func (s *Stack) Len() int {
	return len(s.items)
}

func (s *Stack) Peek() *Node {
	return s.items[s.Len()-1]
}
