package main

import "fmt"

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
