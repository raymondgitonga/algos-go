package sorting

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				temp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
			}
			// step back till you reach to the end of the sorted pile
			j = j - 1
		}
	}
	return arr
}
