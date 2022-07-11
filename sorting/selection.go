package sorting

func SelectionSort(arr []int) []int {
	//9, 4, 1, 3, 0
	//set minimum as the first element in the array
	minVal := 0
	minIndex := 0

	for i := 0; i < len(arr); i++ {
		minVal = arr[i]
		minIndex = i
		// Get an element smaller than the minimum in the whole array
		for j := i; j < len(arr); j++ {
			if arr[j] < minVal {
				minVal = arr[j]
				minIndex = j
			} else {
				continue
			}
		}
		temp := arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = temp
	}

	return arr
}
