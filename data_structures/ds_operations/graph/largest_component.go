package main

//"0": {"8", "1", "5"},
//"1": {"0"},
//"5": {"0", "8"},
//"8": {"0", "5"},
//"2": {"3", "4"},
//"3": {"2", "4"},
//"4": {"3", "2"},

func largestComponentsCount(graphItem GraphItem) int {
	longest := 0
	//lookup and insertion will be 0(1)
	visited := make(map[string]string)
	for node := range graphItem.graph {
		size := exploreComponent(graphItem, node, visited)

		if size > longest {
			longest = size
		}
	}
	return longest
}

func exploreComponent(item GraphItem, src string, visited map[string]string) int {
	if _, ok := visited[src]; ok {
		return 0
	}

	visited[src] = ""

	// Represents current node were in
	size := 1

	for _, v := range item.graph[src] {
		size += exploreComponent(item, v, visited)
	}

	return size
}
