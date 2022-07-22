package recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func (l *LinkedList) prepend(node ...*node) {
	for _, n := range node {
		second := l.head
		newHead := n
		l.head = newHead
		newHead.next = second
	}

}

func TestLinkedListReversal(t *testing.T) {
	l := LinkedList{}
	l.prepend(&node{val: 1}, &node{val: 2}, &node{val: 3})

	reverseLinkedList(l.head)

	assert.Equal(t, 3, l.head.val)
}
