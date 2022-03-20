package clone

import "github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"

func CopyRandomListOptimized(head *linkedlist.Node) *linkedlist.Node {
	tmpHead := head

	for tmpHead != nil {
		newNode := &linkedlist.Node{Val: tmpHead.Val}
		tmp := tmpHead.Next
		tmpHead.Next = newNode
		newNode.Next = tmp
		tmpHead = newNode.Next
	}

	tmpHead = head

	for tmpHead != nil && tmpHead.Next != nil {
		if tmpHead.Random == nil {
			tmpHead.Next.Random = nil
		} else {
			tmpHead.Next.Random = tmpHead.Random.Next
		}
		tmpHead = tmpHead.Next.Next
	}

	tmpHead = head
	frontHead := tmpHead.Next.Next
	newHead := head.Next

	for frontHead != nil {
		tmp := tmpHead.Next.Next
		tmpHead.Next.Next = frontHead.Next
		tmpHead = tmp
		frontHead = tmpHead.Next.Next
	}

	return newHead
}

func CopyRandomList(head *linkedlist.Node) *linkedlist.Node {
	store := make(map[*linkedlist.Node]*linkedlist.Node)
	tmp := head

	for tmp != nil {
		newNode := &linkedlist.Node{Val: tmp.Val}
		store[tmp] = newNode
		tmp = tmp.Next
	}

	for k,v := range store {
		v.Next = store[k.Next]
		v.Random = store[k.Random]
	}

	return store[head]
}
