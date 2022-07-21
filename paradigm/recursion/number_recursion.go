package recursion

import "strconv"

func decimalToBinary(decimal int, result string) string {
	if decimal == 0 {
		return result
	}
	result = strconv.Itoa(decimal%2) + result
	return decimalToBinary(decimal/2, result)
}

func sumOfNaturalNumbers(num int, sum int) int {
	if num == 0 {
		return sum
	}
	sum = num + sum
	return sumOfNaturalNumbers(num-1, sum)
}
