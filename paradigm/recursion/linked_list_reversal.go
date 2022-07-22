package recursion

type LinkedList struct {
	head *Node
}

type Node struct {
	val  int
	next *Node
}

func reverseLinkedList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	d := reverseLinkedList(head.next)

	head.next.next = head
	head.next = nil

	return d
}
