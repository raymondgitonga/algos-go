package sorting

import (
	"fmt"
	"strconv"
)

func BubbleSort(arr []int) []int {
	//9, 4, 1, 3, 0
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				p := strconv.Itoa(arr[j])
				a := strconv.Itoa(arr[j+1])

				fmt.Println("Switching :" + p + " with " + a)
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}
