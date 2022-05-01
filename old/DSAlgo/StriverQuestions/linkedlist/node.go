package linkedlist

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func CreateNodeFromList(list [][]*int) *Node {
	arr := make([]int, 0)
	randoms := make([]*int, 0)

	for _, l := range list {
		arr = append(arr, *l[0])
		randoms = append(randoms, l[1])
	}

	head := createNodeFromArray(arr)
	tmp := head

	count := 0
	for tmp != nil {
		if randoms[count] == nil {
			tmp.Random = nil
			count += 1
			tmp = tmp.Next
			continue
		}

		random := getNode(head, *randoms[count])
		tmp.Random = random

		count += 1
		tmp = tmp.Next
	}

	return head
}

func createNodeFromArray(arr []int) *Node {
	var head *Node

	var prev *Node
	for i, n := range arr {
		curr := &Node{
			Val:  n,
			Next: nil,
		}

		if i > 0 {
			prev.Next = curr
		}

		prev = curr
		if i == 0 {
			head = prev
		}
	}

	return head
}

func getNode(head *Node, pos int) *Node {
	var node *Node

	tmp := head
	count := 0
	for tmp != nil {
		if count == pos {
			node = tmp
			break
		}
		count += 1
		tmp = tmp.Next
	}

	return node
}