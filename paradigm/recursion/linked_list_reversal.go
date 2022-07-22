package recursion

type LinkedList struct {
	head *node
}

type node struct {
	val  int
	next *node
}

func reverseLinkedList(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}
	d := reverseLinkedList(head.next)

	head.next.next = head
	head.next = nil

	return d
}
