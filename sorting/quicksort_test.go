package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	unsorted := []int{9, 4, 1, 3, 0}

	sorted := QuickSort(unsorted, 0, len(unsorted)-1)

	assert.Equal(t, sorted[0], 0)
	assert.Equal(t, sorted[len(sorted)-1], 9)
}
