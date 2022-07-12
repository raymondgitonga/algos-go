package sorting

func QuickSort(arr []int, lowIndex int, highIndex int) []int {

	if lowIndex >= highIndex {
		return arr
	}
	pivot := arr[highIndex]
	leftPointer := lowIndex
	rightPointer := highIndex

	// While left pointer and right pointer are not in the same place
	for leftPointer < rightPointer {

		// While element in left pointer is smaller than pivot and right pointer hasn't caught up
		for arr[leftPointer] <= pivot && leftPointer < rightPointer {
			leftPointer++
		}

		// While element in right pointer is larger than pivot and left pointer hasn't caught up
		for arr[rightPointer] >= pivot && leftPointer < rightPointer {
			rightPointer--
		}

		// Swap elements in right and left pointer
		swap(arr, leftPointer, rightPointer)
	}

	// Swap elements in both pointers position(In same position) with pivot(last element)
	swap(arr, leftPointer, highIndex)

	// Quick sort left array i.e from start of the array to index before where pivot has been inserted
	QuickSort(arr, lowIndex, leftPointer-1)
	// Quick sort right array i.e from index before where pivot has been inserted to end of the array
	QuickSort(arr, leftPointer+1, highIndex)

	return arr
}

func swap(arr []int, index int, index2 int) {
	temp := arr[index]
	arr[index] = arr[index2]
	arr[index2] = temp
}
