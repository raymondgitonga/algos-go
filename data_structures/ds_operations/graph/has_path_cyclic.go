package main

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
