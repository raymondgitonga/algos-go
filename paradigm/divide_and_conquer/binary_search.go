package divide_and_conquer

func BinarySearch(arr []int, start int, end int, target int) int {
	if start > end {
		return -1
	}

	midPoint := (start + end) / 2

	if target == arr[midPoint] {
		return midPoint
	}

	if target < arr[midPoint] {
		return BinarySearch(arr, start, midPoint-1, target)
	}

	return BinarySearch(arr, midPoint+1, end, target)
}
