package divide_and_conquer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type testCase struct {
		name     string
		value    int
		expected int
	}

	arr := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	testCases := []testCase{
		{
			name:     "exists",
			value:    10,
			expected: 8,
		},
		{
			name:     "does not exist",
			value:    14,
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := BinarySearch(arr, 0, len(arr)-1, tc.value)
			assert.Equal(t, tc.expected, result)
		})
	}

}
