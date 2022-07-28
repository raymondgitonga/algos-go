package main

func depthFirstValues(root *Node) []string {
	if root == nil {
		return []string{}
	}
	visited := make([]string, 0)
	unvisitedStack := make([]*Node, 0)
	unvisitedStack = append(unvisitedStack, root)

	for len(unvisitedStack) > 0 {
		visitedNode := unvisitedStack[len(unvisitedStack)-1]

		visited = append(visited, visitedNode.Key)
		//pop
		unvisitedStack = unvisitedStack[:len(unvisitedStack)-1]

		if visitedNode.Right != nil {
			unvisitedStack = append(unvisitedStack, visitedNode.Right)
		}
		if visitedNode.Left != nil {
			unvisitedStack = append(unvisitedStack, visitedNode.Left)
		}
	}
	return visited
}

func depthFirstValuesRecursion(root *Node) []string {
	visited := make([]string, 0)

	return traverse(root, visited)
}

func traverse(root *Node, visited []string) []string {
	visited = append(visited, root.Key)

	if root.Left != nil {
		visited = traverse(root.Left, visited)
	}

	if root.Right != nil {
		visited = traverse(root.Right, visited)
	}

	return visited
}
