package main

func connectedComponentsCount(graphItem GraphItem) int {
	// 1. Step 1: Traverse the graph till you get to the end then go to the next connected nodes
	// 2. Step 2: Make sure you store the visited to prevent endless loo[
	// 3. Step 3: Keep track of count for every terminated traversal
	visited := make(map[string]string)
	count := 0

	// We are getting the key value
	for i := range graphItem.graph {
		if explore(graphItem, i, visited) {
			count += 1
		}
	}
	return count
}

func explore(item GraphItem, src string, visited map[string]string) bool {
	if _, ok := visited[src]; ok {
		return false
	}

	visited[src] = ""
	for _, v := range item.graph[src] {
		explore(item, v, visited)
	}

	return true
}
