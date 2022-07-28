package main

import "fmt"

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
