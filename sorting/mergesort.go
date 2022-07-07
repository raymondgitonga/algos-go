package sorting

func MergeSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	arrLength := len(arr)
	midPoint := arrLength / 2
	leftArr := make([]int, midPoint)
	rightArr := make([]int, arrLength-midPoint)

	for i := 0; i < midPoint; i++ {
		leftArr[i] = arr[i]
	}

	for i := midPoint; i < arrLength; i++ {
		rightArr[i-midPoint] = arr[i]
	}

	MergeSort(leftArr)
	MergeSort(rightArr)

	sort(arr, leftArr, rightArr)
	return arr
}

func sort(originalArr []int, leftArr []int, rightArr []int) {
	leftFlag, rightFlag, originalFlag := 0, 0, 0

	for leftFlag < len(leftArr) && rightFlag < len(rightArr) {
		if leftArr[leftFlag] <= rightArr[rightFlag] {
			originalArr[originalFlag] = leftArr[leftFlag]

			leftFlag++
		} else {
			originalArr[originalFlag] = rightArr[rightFlag]

			rightFlag++
		}

		originalFlag++
	}

	for leftFlag < len(leftArr) {
		originalArr[originalFlag] = leftArr[leftFlag]

		originalFlag++
		leftFlag++
	}

	for rightFlag < len(rightArr) {
		originalArr[originalFlag] = rightArr[rightFlag]

		originalFlag++
		rightFlag++
	}
}
