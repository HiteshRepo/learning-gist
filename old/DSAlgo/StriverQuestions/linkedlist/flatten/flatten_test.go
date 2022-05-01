package flatten_test

import (
	"github.com/HiteshRepo/learninggist/DSAlgo/StriverQuestions/linkedlist"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Flatten(t *testing.T) {
	l1 := []int{5, 7, 8, 10}
	l2 := []int{10, 20}
	l3 := []int{19, 22, 50}
	l4 := []int{28, 35, 40, 45}
	expected := []int{5,7,8,10,10,19,20,22,28,35,40,45,50}

	head1 := linkedlist.CreateDoubleListNodeFromArray(l1)
	head2 := linkedlist.CreateDoubleListNodeFromArray(l2)
	head3 := linkedlist.CreateDoubleListNodeFromArray(l3)
	head4 := linkedlist.CreateDoubleListNodeFromArray(l4)

	head1.Next = head2
	head2.Next = head3
	head3.Next = head4

	actualHead := Flatten(head1)
	assert.Equal(t, expected, actualHead.GetArrayFromLinkedList())
}
