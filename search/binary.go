package search

func BinarySearch(arr []int, target int) int {
	// 2, 5, 7, 9, 12, 69, 74
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
