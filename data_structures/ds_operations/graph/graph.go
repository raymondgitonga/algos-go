package main

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

	//Undirected/cyclic graph check component count/check connections
	_ = GraphItem{map[string][]string{
		"0": {"8", "1", "5"},
		"1": {"0"},
		"5": {"0", "8"},
		"8": {"0", "5"},
		"2": {"3", "4"},
		"3": {"2", "4"},
		"4": {"3", "2"},
	}}

	_ = [][]string{
		{"w", "x"},
		{"x", "y"},
		{"z", "y"},
		{"z", "v"},
		{"w", "v"},
	}
}
