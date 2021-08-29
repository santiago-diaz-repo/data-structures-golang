package adjacencyList

import (
	"data-structures-golang/list"
	"data-structures-golang/list/doublylinkedlist"
	"fmt"
	"testing"
)

func getEmptyGraph() GraphAdjacencyList {
	elements := make([]list.List, 5)
	for i := range elements {
		elements[i] = doublylinkedlist.NewDoublyLinkedListEmp()
	}
	return GraphAdjacencyList{
		elements: elements,
		order:    5,
		size:     0,
	}
}

func TestGraphAdjacencyList_AddEdge(t *testing.T) {
	elements := getEmptyGraph()
	elements.AddEdge(5, 6)
	elements.AddEdge(3, 5)
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

func TestGraphAdjacencyList_AddEdgeDirected(t *testing.T) {
	elements := getEmptyGraph()
	elements.AddEdgeDirected(0, 3)
	elements.AddEdgeDirected(3, 4)
	elements.AddEdgeDirected(2, 0)
	elements.AddEdgeDirected(2, 1)

	want := false
	got := elements.CycleDetectionDirected()

	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}

	elements.AddEdgeDirected(1, 2)
	want = true
	got = elements.CycleDetectionDirected()

	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}
}

func TestGraphAdjacencyList_TopologicalSort(t *testing.T) {
	elements := getEmptyGraph()
	elements.AddEdgeDirected(2, 3)
	elements.AddEdgeDirected(2, 1)
	elements.AddEdgeDirected(1, 0)
	elements.AddEdgeDirected(3, 4)

	want := []int{2, 3, 4, 1, 0}
	got := elements.TopologicalSort()

	for i := 0; i < 5; i++ {
		if got[i] != want[i] {
			t.Errorf("It expected %v, but got %v", want[i], got[i])
		}
	}
}

func TestGraphAdjacencyList_Transpose(t *testing.T) {
	elements := getEmptyGraph()
	elements.AddEdgeDirected(2, 3)
	elements.AddEdgeDirected(2, 1)
	elements.AddEdgeDirected(1, 0)
	elements.AddEdgeDirected(3, 4)

	want := getEmptyGraph()
	want.AddEdgeDirected(0, 1)
	want.AddEdgeDirected(1, 2)
	want.AddEdgeDirected(3, 2)
	want.AddEdgeDirected(4, 3)

	got := elements.Transpose().Elements()

	for i := 0; i < 5; i++ {
		if want.elements[i].Size() != got[i].Size() {
			t.Errorf("It expected size %v, but got size %v", want.elements[i].Size(), got[i].Size())
		}
		for j := 1; j <= int(want.elements[i].Size()); j++ {
			if want.elements[i].Get(uint16(j)) != got[i].Get(uint16(j)) {
				t.Errorf("It expected %v, but got %v", want.elements[i].Get(uint16(j)), got[i].Get(uint16(j)))
			}
		}
	}
}

func TestGraphAdjacencyList_StronglyConnectedComponents(t *testing.T) {
	elements := getEmptyGraph()
	elements.AddEdgeDirected(2, 3)
	elements.AddEdgeDirected(2, 1)
	elements.AddEdgeDirected(1, 0)
	elements.AddEdgeDirected(3, 4)
	elements.AddEdgeDirected(4, 2)

	elements.StronglyConnectedComponents()
}

func TestGraphAdjacencyList_MotherGraph(t *testing.T) {
	elements := getEmptyGraph()
	elements.AddEdgeDirected(2, 3)
	elements.AddEdgeDirected(2, 1)
	elements.AddEdgeDirected(1, 0)
	elements.AddEdgeDirected(3, 4)

	want := 2
	got := elements.MotherGraph()

	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}
}

func TestGraphAdjacencyList_NumberEdges(t *testing.T) {
	elements := getEmptyGraph()
	elements.AddEdgeDirected(2, 3)
	elements.AddEdgeDirected(2, 1)
	elements.AddEdgeDirected(1, 0)
	elements.AddEdgeDirected(3, 4)

	want := 4
	got := elements.NumberEdges(true)

	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}
}

func TestGraphAdjacencyList_PathExists(t *testing.T) {
	elements := getEmptyGraph()
	elements.AddEdgeDirected(2, 3)
	elements.AddEdgeDirected(2, 1)
	elements.AddEdgeDirected(1, 0)
	elements.AddEdgeDirected(3, 4)

	want := true
	got := elements.PathExists(2, 4)

	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}
}

func TestGraphAdjacencyList_HasOneParent(t *testing.T) {
	elements := getEmptyGraph()
	elements.AddEdgeDirected(2, 3)
	elements.AddEdgeDirected(2, 1)
	elements.AddEdgeDirected(1, 0)
	elements.AddEdgeDirected(3, 4)

	want := true
	got := elements.HasOneParent(2)

	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}

	elements.AddEdgeDirected(3, 1)
	want = false
	got = elements.HasOneParent(2)

	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}
}

func TestGraphAdjacencyList_ShortestPathEdges(t *testing.T) {
	elements := getEmptyGraph()
	elements.AddEdgeDirected(2, 3)
	elements.AddEdgeDirected(2, 1)
	elements.AddEdgeDirected(1, 0)
	elements.AddEdgeDirected(3, 4)

	want := 2
	got := elements.ShortestPathEdges(2, 4)

	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}
}

func TestGraphAdjacencyList_IsBipartite(t *testing.T) {
	elements := getEmptyGraph()
	elements.AddEdgeDirected(2, 3)
	elements.AddEdgeDirected(2, 1)
	elements.AddEdgeDirected(1, 0)
	elements.AddEdgeDirected(3, 4)

	want := true
	got := elements.IsBipartite(2)

	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}

	elements.AddEdgeDirected(2, 4)

	want = false
	got = elements.IsBipartite(2)

	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}
}
