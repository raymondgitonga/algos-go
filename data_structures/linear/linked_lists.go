package linear

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

// We're using the pointer receiver as we're going to modify
// the linked list itself and not a copy
// Note: Use of variadic functions here
func (l *linkedList) prepend(node ...*node) {

	for _, n := range node {
		second := l.head
		newHead := n
		l.head = newHead
		newHead.next = second
		l.length++
	}

}

// We're using a value receiver as we're just going to print the list and not modify it
func (l linkedList) printLinkedList() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Println(toPrint.data)
		toPrint = l.head.next
		l.length--
	}
}

func (l *linkedList) deleteWithValue(value int) {
	currentHead := l.head

	if l.length < 1 {
		return
	}

	if currentHead.data == value {
		l.head = currentHead.next
		l.length--
		return
	}

	// If the item next after the current head
	for currentHead.next.data != value {
		if currentHead.next == nil {
			return
		}
		currentHead = currentHead.next
	}

	currentHead.next = currentHead.next.next
	l.length--

}
