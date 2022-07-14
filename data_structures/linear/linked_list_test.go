package linear

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrependNode(t *testing.T) {
	list := linkedList{}
	node := &node{
		data: 5,
		next: nil,
	}

	list.prepend(node)

	assert.Equal(t, 1, list.length)
	assert.Equal(t, 5, list.head.data)
}

func TestDeleteWithValue(t *testing.T) {
	list := linkedList{}
	node1 := &node{data: 5, next: nil}
	node2 := &node{data: 9, next: nil}
	node3 := &node{data: 12, next: nil}

	list.prepend(node1, node2, node3)
	list.deleteWithValue(9)

	assert.Equal(t, 2, list.length)

}
