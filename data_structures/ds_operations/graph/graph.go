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

	//Undirected/cyclic graph
	_ = [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}

	//Undirected/cyclic graph check connections
	connections := GraphItem{map[string][]string{
		"0": {"8", "1", "5"},
		"1": {"0"},
		"5": {"0", "8"},
		"8": {"0", "5"},
		"2": {"3", "4"},
		"3": {"2", "4"},
		"4": {"3", "2"},
	}}

	x := connectedComponentsCount(connections)
	fmt.Println(x)
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
