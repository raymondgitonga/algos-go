package main

import "fmt"

type GraphItem struct {
	graph map[string][]string
}

func main() {
	// Directed graph
	_ = GraphItem{graph: map[string][]string{
		"f": {"g", "i"},
		"g": {"h"},
		"h": {},
		"i": {"g", "k"},
		"j": {"i"},
	}}

	//Undirected graph
	nodes := [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}

	x := undirectedPath(nodes, "k", "m")

	fmt.Println(x)
}

func undirectedPath(edges [][]string, src string, dest string) bool {
	graphItem := buildGraph(edges)

	lookupMap := make(map[string]string)
	return hasPath(graphItem, src, dest, lookupMap)
}

func hasPath(graphItem GraphItem, src string, dest string, lookup map[string]string) bool {
	if src == dest {
		return true
	}
	if _, ok := lookup[src]; ok {
		return false
	}

	lookup[src] = src
	neighbours := graphItem.graph[src]

	for _, v := range neighbours {
		if hasPath(graphItem, v, dest, lookup) {
			return true
		}
	}
	return false
}

func buildGraph(edges [][]string) GraphItem {
	graphItem := GraphItem{}
	graphItem.graph = make(map[string][]string)

	for _, v := range edges {
		a, b := v[0], v[1]

		if _, ok := graphItem.graph[a]; !ok {
			graphItem.graph[a] = []string{}
		}

		if _, ok := graphItem.graph[b]; !ok {
			graphItem.graph[b] = []string{}
		}

		graphItem.graph[a] = append(graphItem.graph[a], b)
		graphItem.graph[b] = append(graphItem.graph[b], a)
	}
	return graphItem
}
func depthFirstTraversal(g GraphItem, source string) {
	unvisited := make([]string, 0)

	unvisited = append(unvisited, source)

	for len(unvisited) > 0 {
		// Get the last item in the slice
		visited := unvisited[len(unvisited)-1]
		// Print it
		fmt.Println(visited)
		//  Remove the last item from the slice (LIFO)
		unvisited = unvisited[:len(unvisited)-1]

		// Add all adjacent neighbours of the removed item into the unvisited list
		for _, v := range g.graph[visited] {
			unvisited = append(unvisited, v)
		}
	}
}

func depthFirstTraversalRecursively(g GraphItem, source string) {
	fmt.Println(source)

	neighbours := g.graph[source]

	for _, v := range neighbours {
		depthFirstTraversalRecursively(g, v)
	}
}

func breadthFirstTraversal(g GraphItem, source string) {
	unvisited := make([]string, 0)

	unvisited = append(unvisited, source)

	for len(unvisited) > 0 {
		visited := unvisited[0]

		fmt.Println(visited)

		// Dequeue
		unvisited = unvisited[1:]

		for _, neighbours := range g.graph[visited] {
			unvisited = append(unvisited, neighbours)
		}
	}
}

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
