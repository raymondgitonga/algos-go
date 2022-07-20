package non_linear

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGraph_AddVertex(t *testing.T) {
	graph := Graph{}
	graph.AddVertex(1)
	assert.Equal(t, true, contains(graph.vertices, 1))
}

func TestGraph_GetVertex(t *testing.T) {
	graph := Graph{}
	graph.AddVertex(1)

	vertex := graph.getVertex(1)
	assert.NotNil(t, vertex)
}

func TestGraph_AddEdge(t *testing.T) {
	graph := Graph{}
	graph.AddVertex(1)
	graph.AddVertex(2)

	graph.AddEdge(1, 2)

	edgeKey := graph.vertices[0].adjacent[0].key

	assert.Equal(t, 2, edgeKey)
}
