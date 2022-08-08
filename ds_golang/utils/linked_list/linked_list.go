package linked_list

type Node struct {
	data int
	next *Node
}

func GetNewNode(data int) *Node {
	return &Node{data: data, next: nil}
}

func (n *Node) GetData() int {
	return n.data
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) SetNext(node *Node) {
	n.next = node
}

type LinkedList struct {
	head *Node
}

func GetNewLinkedList(data int) *LinkedList {
	return &LinkedList{head: GetNewNode(data)}
}

func (l *LinkedList) GetHead() *Node {
	return l.head
}

func (l *LinkedList) AddFromArray(arr []int) {
	node := l.head
	for _,n := range arr {
		newNode := GetNewNode(n)
		node.next = newNode
		node = node.next
	}
}
