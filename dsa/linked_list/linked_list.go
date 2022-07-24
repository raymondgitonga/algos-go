package main

import "fmt"

// Node is a recursive type hence the recursive type needs to be a pointer
type Node struct {
	val  int
	next *Node
}

func printLinkedList(head *Node) {
	for head != nil {
		fmt.Println(head.val)
		head = head.next
	}
}

func printLinkedListRecursively(head *Node) {
	if head == nil {
		return
	}
	fmt.Println(head.val)
	printLinkedListRecursively(head.next)
}

func addValuesLinkedList(head *Node) []int {
	list := make([]int, 0)

	for head != nil {
		list = append(list, head.val)
		head = head.next
	}

	return list
}

func addValuesLinkedListRecursively(head *Node, list []int) []int {
	if head == nil {
		return list
	}

	list = append(list, head.val)

	return addValuesLinkedListRecursively(head.next, list)
}

func sumLinkedList(head *Node) int {
	sum := 0
	for head != nil {
		sum += head.val
		head = head.next
	}
	return sum
}

/**
The recursive method will keep adding the head value
to the sum till it reaches to the top of the stack, it will then pop the stack
layers while 'carrying' the sum with it till it reaches to the end of the stack
where it then returns the sum
 **/
func sumLinkedListRecursive(head *Node, sum int) int {
	if head == nil {
		return sum
	}
	sum += head.val
	return sumLinkedListRecursive(head.next, sum)
}

// space complexity is 0(1), time complexity is 0(n)
func linkedListFind(head *Node, target int) bool {
	for head != nil {
		if head.val == target {
			return true
		}

		head = head.next
	}

	return false
}

// Uses a linear amount of space because of the callstack, time complexity is 0(n)
func linkedListFindRecursive(head *Node, target int) bool {
	if head == nil {
		return false
	}
	if head.val == target {
		return true
	}
	return linkedListFind(head.next, target)
}

func getNodeValueAtIndex(head *Node, index int) int {
	count := 0
	for head != nil {
		if count == index {
			return head.val
		}
		count += 1
		head = head.next
	}
	return -1
}

func reverseLinkedListRecursive(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	d := reverseLinkedListRecursive(head.next)
	head.next.next = head
	head.next = nil
	return d
}

func reverseLinkedList(head *Node) *Node {
	prevNode := &Node{}
	current := head
	for current != nil {
		next := current.next
		current.next = prevNode
		prevNode = current
		current = next
	}
	return prevNode
}

func zipperList(node1 *Node, node2 *Node) *Node {
	combinedNode := &Node{}
	head := combinedNode
	flag := 0

	for node1 != nil && node2 != nil {
		if flag == 0 {
			combinedNode.next = node1
			node1 = node1.next
			flag = 1
		} else {
			combinedNode.next = node2
			node2 = node2.next
			flag = 0
		}
		combinedNode = combinedNode.next
	}

	if node1 != nil {
		combinedNode.next = node1
	}

	if node2 != nil {
		combinedNode.next = node2
	}

	return head.next
}

func main() {
}
