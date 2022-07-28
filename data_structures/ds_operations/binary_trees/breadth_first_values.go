package main

func breadthFirstValues(root *Node) []string {
	if root == nil {
		return []string{}
	}

	visited := make([]string, 0)
	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		current := queue[0]

		queue = queue[1:]

		visited = append(visited, current.Key)

		if current.Right != nil {
			queue = append(queue, current.Right)
		}

		if current.Left != nil {
			queue = append(queue, current.Left)
		}

	}

	return visited
}
