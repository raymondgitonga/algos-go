package recursion

func MergeSortedLinkedList(a *Node, b *Node) *Node {
	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	if a.val < b.val {
		a.next = MergeSortedLinkedList(a.next, b)
		return a
	} else {
		b.next = MergeSortedLinkedList(a, b.next)
		return b
	}

}
