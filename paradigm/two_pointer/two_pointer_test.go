package two_pointer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoPointer_IsPalindrome(t *testing.T) {
	type testCase struct {
		input  string
		result bool
	}

	testCases := []testCase{
		{
			input:  "john",
			result: false,
		},
		{
			input:  "racecar",
			result: true,
		},
		{
			input:  "",
			result: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			isPalindrome := isPalindrome(tc.input)
			assert.Equal(t, tc.result, isPalindrome)
		})
	}

}
