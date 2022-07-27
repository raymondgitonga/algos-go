package main

func hasPathDFS(g GraphItem, source string, dest string) bool {
	// stack LIFO
	visited := make([]string, 0)
	unvisited := make([]string, 0)

	unvisited = append(unvisited, source)

	for len(unvisited) > 0 {
		vst := unvisited[len(unvisited)-1]

		unvisited = unvisited[:len(unvisited)-1]
		visited = append(visited, vst)
		for _, v := range g.graph[vst] {
			unvisited = append(unvisited, v)
		}
	}

	for _, v := range visited {
		if dest == v {
			return true
		}
	}

	return false
}

func hasPathDFSRecursive(g GraphItem, source string, dest string) bool {
	if source == dest {
		return true
	}

	neighbours := g.graph[source]

	for _, v := range neighbours {
		if hasPathDFSRecursive(g, v, dest) {
			return true
		}
	}
	return false
}

func hasPathBFS(g GraphItem, source string, dest string) bool {
	visited := make([]string, 0)
	unvisited := make([]string, 0)

	unvisited = append(unvisited, source)

	for len(unvisited) > 0 {
		visitedNode := unvisited[0]
		unvisited = unvisited[1:]
		visited = append(visited, visitedNode)

		for _, v := range g.graph[visitedNode] {
			unvisited = append(unvisited, v)
		}
	}

	for _, v := range visited {
		if v == dest {
			return true
		}
	}

	return false
}
