package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	sortedArr := []int{2, 5, 7, 9, 12, 69, 74}

	position := BinarySearch(sortedArr, 12)

	assert.Equal(t, 4, position)
}
