package main

import "fmt"

// Time complexity O(2^n+m), Space Complexity of O(n + m)
func gridTraveller(m int, n int) int {
	if m == 1 && n == 1 {
		return 1
	}

	if m == 0 || n == 0 {
		return 0
	}

	return gridTraveller(m-1, n) + gridTraveller(m, n-1)
}

// Time complexity O(n*m), Space Complexity of O(n + m)
func gridTravellerWithMemoization(m int, n int) int {
	memo := make(map[string]int)

	return memoizationGrid(m, n, memo)
}

func memoizationGrid(m int, n int, memo map[string]int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if m == 1 && n == 1 {
		return 1
	}

	key := fmt.Sprintf("%d%d", m, n)

	if _, ok := memo[key]; ok {
		return memo[key]
	}

	memo[key] = memoizationGrid(m-1, n, memo) + memoizationGrid(m, n-1, memo)

	return memo[key]
}
