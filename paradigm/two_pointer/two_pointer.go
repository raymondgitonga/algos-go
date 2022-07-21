package two_pointer

func isPalindrome(str string) bool {
	start := 0
	end := len(str) - 1

	for start < end {
		if str[start] == str[end] {
			start++
			end--
		} else {
			return false
		}
	}

	return true
}
