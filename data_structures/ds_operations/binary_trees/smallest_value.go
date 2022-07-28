package main

func smallestValue(root *NodeInt) int {
	// use Breadth first FIFO
	smallest := root.Key
	unvisited := make([]*NodeInt, 0)
	unvisited = append(unvisited, root)

	for len(unvisited) > 0 {
		current := unvisited[0]

		if current.Key < smallest {
			smallest = current.Key
		}
		//dequeque
		unvisited = unvisited[1:]

		if current.Left != nil {
			unvisited = append(unvisited, current.Left)
		}

		if current.Right != nil {
			unvisited = append(unvisited, current.Right)
		}
	}

	return smallest
}
