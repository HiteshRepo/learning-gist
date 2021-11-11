package main

type Stack struct {
	elements []int
}

func NewStack() *Stack {
	return &Stack{elements: make([]int, 0)}
}

func (stack *Stack) Push(data int) {
	stack.elements = append(stack.elements, data)
}

func (stack *Stack) Pop() int {
	length := len(stack.elements)
	last := stack.elements[length - 1]
	stack.elements = stack.elements[0: length-1]
	return last
}

func (stack *Stack) Peek() int {
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
