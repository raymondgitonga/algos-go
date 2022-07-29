package main

// Time complexity of O(n^m), Space complexity O(m) m = targetSum n =array length
func canSum(targetSum int, numbers []int) bool {
	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}

	for _, num := range numbers {

		newTarget := targetSum - num

		if canSum(newTarget, numbers) == true {
			return true
		}
	}

	return false
}

// Time complexity of O(n* m), Space complexity O(m)
func canSumMemoization(targetSum int, numbers []int) bool {
	memo := make(map[int]bool)

	return memoizationSum(targetSum, numbers, memo)
}

func memoizationSum(targetSum int, numbers []int, memo map[int]bool) bool {
	if _, ok := memo[targetSum]; ok {
		return memo[targetSum]
	}

	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}

	for _, num := range numbers {

		newTarget := targetSum - num

		if memo[newTarget] = memoizationSum(newTarget, numbers, memo); true {
			memo[newTarget] = true
			return memo[newTarget]
		}
	}

	memo[targetSum] = false
	return memo[targetSum]
}
