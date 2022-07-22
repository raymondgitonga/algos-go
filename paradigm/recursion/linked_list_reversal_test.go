package recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func (l *LinkedList) prepend(node ...*Node) {
	for _, n := range node {
		second := l.head
		newHead := n
		l.head = newHead
		newHead.next = second
	}

}

func TestLinkedListReversal(t *testing.T) {
	l := LinkedList{}
	l.prepend(&Node{val: 1}, &Node{val: 2}, &Node{val: 3})

	reverseLinkedList(l.head)

	assert.Equal(t, 3, l.head.val)
}
