package recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringRecursion_Reverse(t *testing.T) {
	type testCase struct {
		input  string
		result string
	}

	testCases := []testCase{
		{
			input:  "hello",
			result: "olleh",
		},
		{
			input:  "car",
			result: "rac",
		},
		{
			input:  "",
			result: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			reversedString := reverseString(tc.input)
			assert.Equal(t, tc.result, reversedString)
		})
	}
}

func TestStringRecursion_Palindrome(t *testing.T) {
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
