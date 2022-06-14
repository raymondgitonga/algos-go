package sorting

func InsertionSort(arr []int) []int {
	//9, 4, 1, 3, 0

	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				temp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
			}
			j = j - 1
		}
	}
	return arr
}
