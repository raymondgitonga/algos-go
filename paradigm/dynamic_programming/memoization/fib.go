package main

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fibWithMemoization(n int) int {
	memo := make(map[int]int)
	return memoization(n, memo)
}

func memoization(n int, memo map[int]int) int {
	if _, ok := memo[n]; ok {
		return memo[n]
	}
	if n <= 2 {
		return 1
	}
	memo[n] = memoization(n-1, memo) + memoization(n-2, memo)
	return memo[n]
}
