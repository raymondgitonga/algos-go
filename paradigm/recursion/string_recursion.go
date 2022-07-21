package recursion

func reverseString(str string) string {
	if str == "" {
		return ""
	}
	return reverseString(str[1:]) + string(str[0])
}

func isPalindrome(str string) bool {
	if len(str) <= 1 {
		return true
	}

	if str[0] == str[len(str)-1] {
		return isPalindrome(str[1 : len(str)-1])
	}

	return false
}
