package main

import "fmt"

func main() {
	//x := gridTravellerWithMemoization(2, 3)
	//
	//fmt.Println(x)

	arr := []int{2, 3}

	x := canSumMemoization(7, arr)

	fmt.Println(x)
}
