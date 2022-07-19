package non_linear

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearchTree_Insert(t *testing.T) {
	node := &Node{Key: 100}

	node.insert(120)
	node.insert(70)

	assert.Equal(t, 100, node.Key)
	assert.Equal(t, 120, node.Right.Key)
	assert.Equal(t, 70, node.Left.Key)
}

func TestBinarySearchTree_Search(t *testing.T) {
	node := &Node{Key: 100}
	node.insert(120)

	exists := node.search(120)
	notExists := node.search(10)

	assert.Equal(t, true, exists)
	assert.Equal(t, false, notExists)
}
