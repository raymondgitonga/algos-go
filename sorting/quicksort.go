package sorting

func QuickSort(arr []int, lowIndex int, highIndex int) []int {

	if lowIndex >= highIndex {
		return arr
	}
	pivot := arr[highIndex]
	leftPointer := lowIndex
	rightPointer := highIndex

	for leftPointer < rightPointer {
		for arr[leftPointer] <= pivot && leftPointer < rightPointer {
			leftPointer++
		}

		for arr[rightPointer] >= pivot && leftPointer < rightPointer {
			rightPointer--
		}

		swap(arr, leftPointer, rightPointer)
	}

	swap(arr, leftPointer, highIndex)

	QuickSort(arr, lowIndex, leftPointer-1)
	QuickSort(arr, leftPointer+1, highIndex)

	return arr
}

func swap(arr []int, index int, index2 int) {
	temp := arr[index]
	arr[index] = arr[index2]
	arr[index2] = temp
}
