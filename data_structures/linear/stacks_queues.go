package linear

//Stack LIFO
type Stack struct {
	items []int
}

func (s *Stack) Push(val ...int) {
	for _, v := range val {
		s.items = append(s.items, v)
	}
}

func (s *Stack) Pop() int {
	stackSize := len(s.items)
	removedVal := s.items[stackSize-1]
	s.items = s.items[:stackSize-1]
	return removedVal
}

//Queue FIFO
type Queue struct {
	items []int
}

func (q *Queue) Enqueue(val ...int) {
	for _, v := range val {
		q.items = append(q.items, v)
	}
}

func (q *Queue) Dequeue() int {
	itemToRemove := q.items[0]
	q.items = q.items[1:]
	return itemToRemove
}
