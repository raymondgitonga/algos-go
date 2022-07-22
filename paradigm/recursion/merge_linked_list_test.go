package recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSortedLinkedList(t *testing.T) {
	a := &LinkedList{}
	b := &LinkedList{}

	//A 1 -> 6 -> 9
	//B 2 -> 7 -> 8

	a.prepend(&Node{val: 1}, &Node{val: 6}, &Node{val: 9})
	b.prepend(&Node{val: 2}, &Node{val: 7}, &Node{val: 8})

	MergeSortedLinkedList(a.head, b.head)

	assert.Equal(t, 1, a.head.val)
}
