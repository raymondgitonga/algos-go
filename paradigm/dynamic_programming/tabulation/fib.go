package main

func fib(n int) int {
	table := make([]int, n+1)
	table[1] = 1
	for i := range table {
		if i+2 >= len(table) {
			break
		}
		table[i+1] += table[i]
		table[i+2] += table[i]
	}
	return table[n]
}
