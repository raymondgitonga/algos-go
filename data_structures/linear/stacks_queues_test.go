package linear

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := Stack{}

	stack.Push(12, 10, 15, 17)

	assert.Equal(t, 4, len(stack.items))
	assert.Equal(t, 17, stack.items[len(stack.items)-1])
}

func TestStack_Pop(t *testing.T) {
	stack := Stack{}

	stack.Push(12, 10, 15, 17)
	removedItem := stack.Pop()

	assert.Equal(t, 3, len(stack.items))
	assert.Equal(t, 15, stack.items[len(stack.items)-1])
	assert.Equal(t, 17, removedItem)
}

func TestQueue_Enqueue(t *testing.T) {
	queue := Queue{}

	queue.Enqueue(1, 2, 3, 4)

	assert.Equal(t, 4, len(queue.items))
}

func TestQueue_Dequeue(t *testing.T) {
	queue := Queue{}

	queue.Enqueue(1, 2, 3, 4)
	removedItem := queue.Dequeue()

	assert.Equal(t, 1, removedItem)
	assert.Equal(t, 3, len(queue.items))
}
