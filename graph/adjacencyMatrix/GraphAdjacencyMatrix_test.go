package adjacencyMatrix

import (
	"fmt"
	"testing"
)

func getEmptyGraph() GraphAdjacencyMatrix {
	elements := make([][]interface{}, 5)
	for i := range elements {
		elements[i] = make([]interface{}, 5)
	}
	return GraphAdjacencyMatrix{
		elements: elements,
		order:    5,
		size:     0,
	}
}

func TestGraphAdjacencyMatrix_AddEdge(t *testing.T) {
	graph := getEmptyGraph()
	graph.AddEdge(2, 4)
	fmt.Printf("")
}

func TestGraphAdjacencyList_BreathFirstTraverse(t *testing.T) {
	elements := getEmptyGraph()
	elements.AddEdge(0, 2)
	elements.AddEdge(0, 3)
	elements.AddEdge(2, 4)
	elements.AddEdge(3, 4)

	elements.BreathFirstTraverse(0)
}

func TestGraphAdjacencyList_DepthFirstSearch(t *testing.T) {
	elements := getEmptyGraph()
	elements.AddEdge(0, 2)
	elements.AddEdge(0, 3)
	elements.AddEdge(2, 4)
	elements.AddEdge(3, 4)

	elements.DepthFirstSearch(0)
}

func TestGraphAdjacencyList_DepthFirstSearchRecursive(t *testing.T) {
	elements := getEmptyGraph()
	elements.AddEdge(0, 2)
	elements.AddEdge(0, 3)
	elements.AddEdge(2, 4)
	elements.AddEdge(3, 1)

	elements.DepthFirstSearchRecursive()
}

func TestGraphAdjacencyList_CycleDetectionUndirected(t *testing.T) {
	elements := getEmptyGraph()
	elements.AddEdge(0, 1)
	elements.AddEdge(0, 2)
	elements.AddEdge(1, 3)
	elements.AddEdge(3, 4)
	want := false
	got := elements.CycleDetectionUndirected()

	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}

	elements.AddEdge(4, 0)
	want = true
	got = elements.CycleDetectionUndirected()

	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}
}
