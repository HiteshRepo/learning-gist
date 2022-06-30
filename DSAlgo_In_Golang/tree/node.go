package tree

import "fmt"

type Node struct {
	value    int
	Children []*Node
}

func GetNewTreeNode(val int) *Node {
	return &Node{value: val, Children: make([]*Node, 0)}
}

func (n *Node) Add(node *Node) {
	n.Children = append(n.Children, node)
}

func (n *Node) Remove(i int) {
	n.Children = append(n.Children[0:i], n.Children[i+1:]...)
}

type Tree struct {
	Root *Node
}

func GetNewTree(node *Node) *Tree {
	return &Tree{Root: node}
}

func (t *Tree) Display() {
	t.Traverse(t.Root)
	fmt.Println()
}

func (t *Tree) Traverse(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf("%d, ", node.value)
	for _,c := range node.Children {
		t.Traverse(c)
	}
}

func (t *Tree) Size(node *Node)  int{
	if node == nil {
		return 0
	}

	s := 0
	for _,c := range node.Children {
		s += t.Size(c)
	}

	return s + 1
}

func (t *Tree) Max(node *Node)  int{
	if node == nil {
		return 0
	}

	m := -1
	for _,c := range node.Children {
		cm := t.Max(c)
		m = max(m, cm)
	}

	return max(node.value, m)
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func (t *Tree) Height(node *Node)  int{
	if node == nil {
		return 0
	}

	h := -1
	for _,c := range node.Children {
		ch := t.Height(c)
		h = max(h, ch)
	}

	return h+1
}

func (t *Tree) PreOrderTraversal(node *Node) {
	fmt.Printf("%d, ", node.value)
	for _, c := range node.Children{
		t.PreOrderTraversal(c)
	}
}

func (t *Tree) PostOrderTraversal(node *Node) {
	for _, c := range node.Children{
		t.PostOrderTraversal(c)
	}
	fmt.Printf("%d, ", node.value)
}

func (t *Tree) LevelOrderTraversal() {
	st := NewTreeNodeStack()
	st.Push(t.Root)

	for st.Len() > 0 {
		removed := st.PopLeft()

		fmt.Printf("%d, ", removed.value)

		for _,c := range removed.Children {
			st.Push(c)
		}
	}
}

func (t *Tree) Mirror(node *Node) {
	if node == nil {
		return
	}

	childrenCount := len(node.Children)
	for i:=0; i < childrenCount/2; i++ {
		val := node.Children[i]
		node.Children[i] = node.Children[childrenCount-i-1]
		node.Children[childrenCount-i-1] = val
	}

	for _,ch := range node.Children {
		t.Mirror(ch)
	}
}

func (t *Tree) RemoveLeaves(node *Node) {
	if node == nil {
		return
	}

	for i:=len(node.Children)-1; i >= 0 ; i-- {
		child := node.Children[i]
		if len(child.Children) == 0 {
			node.Remove(i)
		}
	}

	for _, ch := range node.Children {
		t.RemoveLeaves(ch)
	}
}

func (t *Tree) Linearize(node *Node) {
	for _, ch := range node.Children {
		t.Linearize(ch)
	}

	for len(node.Children) > 1 {
		lastChild := node.Children[len(node.Children) - 1]
		secondLastChild := node.Children[len(node.Children) - 2]
		node.Remove(len(node.Children) - 1)
		secondLastChildTail := t.getTail(secondLastChild)
		secondLastChildTail.Add(lastChild)
	}
}

func (t *Tree) getTail(node *Node) *Node {
	tmp := node
	for len(tmp.Children) > 0 {
		tmp = tmp.Children[0]
	}

	return tmp
}

func (t *Tree) Find(data int) bool {
	return t.find(t.Root, data)
}

func (t *Tree) find(node *Node, data int) bool {
	if node == nil {
		return false
	}

	if node.value == data {
		return true
	}

	for _, ch := range node.Children {
		if t.find(ch, data) {
			return true
		}
	}

	return false
}

func (t *Tree) NodeToRoot(data int) []int {
	return t.nodeToRoot(t.Root, data)
}

func (t *Tree) nodeToRoot(node *Node, data int) []int {
	if node.value == data {
		return []int{node.value}
	}

	for _, ch := range node.Children {
		chPath := t.nodeToRoot(ch, data)
		if len(chPath) > 0 {
			chPath = append(chPath, node.value)
			return chPath
		}
	}

	return []int{}
}

func (t *Tree) GetSuccessorAndPredecessor(data int) (int, int) {
	state := 0
	successor := -1
	predecessor := -1
	t.getSuccessorAndPredecessor(t.Root, data, &state, &successor, &predecessor)

	return successor, predecessor
}

func (t *Tree) getSuccessorAndPredecessor(node *Node, data int, state *int, successor, predecessor *int) {

	if *state == 1 {
		*successor = node.value
		*state = 2
	}

	if *state == 0 && node.value == data {
		*state = 1
	}

	if *state == 0 {
		*predecessor = node.value
	}

	for _,ch := range node.Children {
		t.getSuccessorAndPredecessor(ch, data, state, successor, predecessor)
	}
}

func (t *Tree) GetCeilAndFloor(data int) (int, int) {
	ceil := -1
	floor := -1
	t.getCeilAndFloor(t.Root, data, &ceil, &floor)

	return ceil, floor
}

func (t *Tree) getCeilAndFloor(node *Node, data int, ceil, floor *int) {

	if node.value > data && (*ceil == -1 || node.value < *ceil) {
		*ceil = node.value
	}

	if node.value < data && node.value > *floor {
		*floor = node.value
	}

	for _,ch := range node.Children {
		t.getCeilAndFloor(ch, data, ceil, floor)
	}
}

func (t *Tree) getFloor(node *Node, data int, floor *int) {

	if node.value < data && node.value > *floor {
		*floor = node.value
	}

	for _,ch := range node.Children {
		t.getFloor(ch, data, floor)
	}
}

func (t *Tree) KthLargest(k int) int {
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)

	currData := MaxInt
	currFloor := -1
	for i:=0; i<k; i++ {
		t.getFloor(t.Root, currData, &currFloor)
		currData = currFloor
		currFloor = -1
	}

	return currData
}

func (t *Tree) GetNodeWithMaxSubtreeSum() int {
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	maxSum := MinInt
	ans := MinInt
	var maxSumNode = &ans
	t.traverseAndGetMaxSumSubtree(t.Root, &maxSum, maxSumNode)
	return ans
}

func (t *Tree) traverseAndGetMaxSumSubtree(node *Node, maxSum *int, maxSumNode *int) int {
	sum := node.value

	for _,ch := range node.Children {
		sum += t.traverseAndGetMaxSumSubtree(ch, maxSum, maxSumNode)
	}

	if sum > *maxSum {
		*maxSum = sum
		*maxSumNode = node.value
	}

	return sum
}

func (t *Tree) Diameter() int {
	dia := -1
	t.calculateChildHeight(t.Root, &dia)
	return dia
}

func (t *Tree) calculateChildHeight(node *Node, dia *int) int {
	deepestChildHeight := -1
	secondDeepestChildHeight := -1

	for _,ch := range node.Children {
		childHeight := t.calculateChildHeight(ch, dia)
		if childHeight > deepestChildHeight {
			secondDeepestChildHeight = deepestChildHeight
			deepestChildHeight = childHeight
		} else if childHeight > secondDeepestChildHeight {
			secondDeepestChildHeight = childHeight
		}
	}

	if (deepestChildHeight + secondDeepestChildHeight + 2) > *dia {
		*dia = deepestChildHeight + secondDeepestChildHeight + 2
	}

	deepestChildHeight += 1
	return deepestChildHeight
}