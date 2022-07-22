package recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberRecursion_DecimalToBinary(t *testing.T) {
	result := ""
	decimal := 233

	binary := decimalToBinary(decimal, result)

	assert.Equal(t, "11101001", binary)
}

func TestNumberRecursion_SumOfNaturalNumbers(t *testing.T) {
	num := 5
	sum := 0

	result := sumOfNaturalNumbers(num, sum)

	assert.Equal(t, 15, result)
}
