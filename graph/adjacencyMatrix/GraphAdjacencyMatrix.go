package adjacencyMatrix

import (
	"data-structures-golang/graph"
	"data-structures-golang/queue/queuelist"
	"data-structures-golang/stack/stacklist"
	"fmt"
)

type GraphAdjacencyMatrix struct {
	elements [][]interface{}
	size     uint16 // number of edges
	order    uint16 //number of vertices
}

func NewGraphAdjacencyMatrix(order uint16) graph.Graph {
	elements := make([][]interface{}, order)
	for i := range elements {
		elements[i] = make([]interface{}, order)
	}
	return &GraphAdjacencyMatrix{
		elements: elements,
		size:     0,
		order:    order,
	}
}

func (g GraphAdjacencyMatrix) AddEdge(node1 uint16, node2 uint16) error {
	g.elements[node1][node2] = 1
	g.elements[node2][node1] = 1
	g.size++
	return nil
}

func (g GraphAdjacencyMatrix) BreathFirstTraverse(root uint16) {
	visited := make([]bool, g.order)
	queue := queuelist.NewLinearQueueEmp()
	visited[root] = true
	queue.Enqueue(root)
	var vis uint16
	for !queue.IsEmpty() {
		vis = queue.Dequeue().(uint16)
		fmt.Printf("%v ->", vis)
		for i := 0; i < int(g.order); i++ {
			if g.elements[vis][i] == 1 && !visited[i] {
				visited[i] = true
				queue.Enqueue(uint16(i))
			}
		}
	}
}

func (g GraphAdjacencyMatrix) DepthFirstSearch(root uint16) {
	visited := make([]bool, g.order)
	stack := stacklist.NewStackListEmp()
	stack.Push(root)
	var val uint16
	for !stack.IsEmpty() {
		val = stack.Pop().(uint16)
		if !visited[val] {
			visited[val] = true
			fmt.Printf("%v ->", val)
			for i := 0; i < int(g.order); i++ {
				if g.elements[val][i] == 1 && !visited[i] {
					stack.Push(uint16(i))
				}
			}
		}
	}
}

func (g GraphAdjacencyMatrix) DepthFirstSearchRecursive() {
	visited := make([]bool, g.order)
	for i := 0; i < int(g.order); i++ {
		if !visited[i] {
			g.depthFirstSearchRecursive(i, visited)
		}
	}
}

func (g GraphAdjacencyMatrix) depthFirstSearchRecursive(vertex int, visited []bool) {
	visited[vertex] = true
	fmt.Printf("%v ->", vertex)
	for i := 0; i < int(g.order); i++ {
		if g.elements[vertex][i] == 1 && !visited[i] {
			g.depthFirstSearchRecursive(i, visited)
		}
	}
}

func (g GraphAdjacencyMatrix) CycleDetectionUndirected() bool {
	visited := make([]bool, g.order)
	for i := 0; i < int(g.order); i++ {
		if !visited[i] {
			if g.cycleDetectionUndirected(i, visited, -1) {
				return true
			}
		}
	}
	return false
}

func (g GraphAdjacencyMatrix) cycleDetectionUndirected(vertex int, visited []bool, parent int) bool {
	visited[vertex] = true
	for i := 0; i < int(g.order); i++ {
		if g.elements[vertex][i] == 1 {
			if !visited[i] {
				if g.cycleDetectionUndirected(i, visited, vertex) {
					return true
				}
			} else if parent != -1 && i != parent {
				return true
			}
		}
	}
	return false
}
