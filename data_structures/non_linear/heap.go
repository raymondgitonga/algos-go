package non_linear

// TODO Add extract and heapify down

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) insert(key int) {
	h.array = append(h.array, key)
	h.MaxHeapifyUp(len(h.array) - 1)
}

// MaxHeapifyUp will heapify bottom to top
func (h *MaxHeap) MaxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}
func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}
