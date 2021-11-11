package main

type Stack struct {
	elements []*Pair
}

func NewStack() *Stack {
	return &Stack{elements: make([]*Pair, 0)}
}

func (stack *Stack) Push(pair *Pair) {
	stack.elements = append(stack.elements, pair)
}

func (stack *Stack) Pop() *Pair {
	length := len(stack.elements)
	last := stack.elements[length - 1]
	stack.elements = stack.elements[0: length-1]
	return last
}

func (stack *Stack) Peek() *Pair {
	length := len(stack.elements)
	last := stack.elements[length - 1]
	return last
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.elements) == 0
}

func (stack *Stack) Size() int {
	return len(stack.elements)
}
