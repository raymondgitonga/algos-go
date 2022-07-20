package non_linear

import "fmt"

// Graph represents an adjacency list graph, it will be directional, cyclic and unweighted
type Graph struct {
	vertices []*Vertex
}

// Vertex represents a graph vertex
type Vertex struct {
	key      int
	adjacent []*Vertex
}

func (g *Graph) AddVertex(k int) {

	if contains(g.vertices, k) {
		return
	}
	g.vertices = append(g.vertices, &Vertex{key: k})
}

func (g *Graph) getVertex(k int) *Vertex {
	for i := range g.vertices {
		if g.vertices[i].key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(from, to int) {
	// get vertex of from and to
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		fmt.Println("Error getting vertex")
		return
	} else if contains(fromVertex.adjacent, to) {
		fmt.Printf("Existing edge from %v to %v", from, to)
		return
	}
	// add edge
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
}

func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if v.key == k {
			return true
		}
	}
	return false
}
