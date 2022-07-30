package main

// Return a combination of numbers that sum to the target

func howSum(target int, numbers []int) []int {
	if target == 0 {
		return []int{}
	}

	if target < 0 {
		return nil
	}

	for _, num := range numbers {
		remainder := target - num
		remainderResult := howSum(remainder, numbers)

		if remainderResult != nil {
			return append(remainderResult, num)
		}
	}
	return nil
}

func howSumMemoization(target int, numbers []int) []int {
	memo := make(map[int][]int)

	return memoizationHowSum(target, numbers, memo)
}

func memoizationHowSum(target int, numbers []int, memo map[int][]int) []int {
	if _, ok := memo[target]; ok {
		return memo[target]
	}

	if target == 0 {
		return []int{}
	}

	if target < 0 {
		return nil
	}

	for _, num := range numbers {
		remainder := target - num
		memo[remainder] = memoizationHowSum(remainder, numbers, memo)
		if memo[remainder] != nil {
			return append(memo[remainder], num)
		}
	}
	return nil
}
