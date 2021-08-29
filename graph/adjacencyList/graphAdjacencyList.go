package adjacencyList

import (
	"data-structures-golang/graph"
	"data-structures-golang/list"
	"data-structures-golang/list/doublylinkedlist"
	"data-structures-golang/queue/queuelist"
	"data-structures-golang/stack"
	"data-structures-golang/stack/stacklist"
	"fmt"
)

type GraphAdjacencyList struct {
	elements []list.List
	size     uint16 // number of edges
	order    uint16 //number of vertices
}

func NewGraphAdjacencyList(order uint16) graph.Graph {
	elements := make([]list.List, order)
	for i := range elements {
		elements[i] = doublylinkedlist.NewDoublyLinkedListEmp()
	}
	return &GraphAdjacencyList{
		elements: elements,
		size:     0,
		order:    order,
	}
}

func (g GraphAdjacencyList) Elements() []list.List {
	return g.elements
}

func (g GraphAdjacencyList) AddEdge(node1 uint16, node2 uint16) error {
	g.elements[node1].AddAtEnd(node2)
	g.elements[node2].AddAtEnd(node1)
	g.size++
	return nil
}

func (g GraphAdjacencyList) AddEdgeDirected(node1 uint16, node2 uint16) error {
	g.elements[node1].AddAtEnd(node2)
	g.size++
	return nil
}

// BreathFirstTraverse assumes the first element is the root
func (g GraphAdjacencyList) BreathFirstTraverse(root uint16) {
	visited := make([]bool, cap(g.elements))
	queue := queuelist.NewLinearQueueEmp()
	visited[root] = true
	queue.Enqueue(root)
	for !queue.IsEmpty() {
		data := queue.Dequeue()
		fmt.Printf("%v -> ", data)
		for i := 1; i <= int(g.elements[data.(uint16)].Size()); i++ {
			value := g.elements[data.(uint16)].Get(uint16(i)).(uint16)
			if !visited[value] {
				queue.Enqueue(value)
				visited[value] = true
			}

		}
	}
}

func (g GraphAdjacencyList) DepthFirstSearch(root uint16) {
	visited := make([]bool, g.order)
	stack := stacklist.NewStackListEmp()
	stack.Push(root)
	var val uint16
	for !stack.IsEmpty() {
		val = stack.Pop().(uint16)
		if !visited[val] {
			visited[val] = true
			fmt.Printf("%v ->", val)
			for i := 1; i <= int(g.elements[val].Size()); i++ {
				pos := g.elements[val].Get(uint16(i)).(uint16)
				if !visited[pos] {
					stack.Push(pos)
				}
			}
		}
	}
}

func (g GraphAdjacencyList) DepthFirstSearchRecursive() {
	visited := make([]bool, g.order)
	for i := 0; i < int(g.order); i++ {
		if !visited[i] {
			g.DepthFirstSearchRecursiveCall(i, visited)
		}
	}
}

func (g GraphAdjacencyList) DepthFirstSearchRecursiveCall(vertex int, visited []bool) {
	visited[vertex] = true
	fmt.Printf("%v ->", vertex)
	for i := 1; i <= int(g.elements[vertex].Size()); i++ {
		val := g.elements[vertex].Get(uint16(i)).(uint16)
		if !visited[int(val)] {
			g.DepthFirstSearchRecursiveCall(int(val), visited)
		}
	}
}

func (g GraphAdjacencyList) CycleDetectionUndirected() bool {
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

func (g GraphAdjacencyList) cycleDetectionUndirected(vertex int, visited []bool, parent int) bool {
	visited[vertex] = true
	for i := 1; i <= int(g.elements[vertex].Size()); i++ {
		val := g.elements[vertex].Get(uint16(i)).(uint16)
		if !visited[val] {
			if g.cycleDetectionUndirected(int(val), visited, vertex) {
				return true
			}
		} else if int(val) != parent {
			return true
		}
	}
	return false
}

func (g GraphAdjacencyList) CycleDetectionDirected() bool {
	visited := make([]bool, g.order)
	curStack := make([]bool, g.order)
	for i := 0; i < int(g.order); i++ {
		if g.cycleDetectionDirected(i, visited, curStack) {
			return true
		}
	}
	return false
}

func (g GraphAdjacencyList) cycleDetectionDirected(vertex int, visited []bool, curStack []bool) bool {
	if curStack[vertex] {
		return true
	}

	if visited[vertex] {
		return false
	}

	curStack[vertex] = true
	visited[vertex] = true

	for i := 1; i <= int(g.elements[vertex].Size()); i++ {
		val := g.elements[vertex].Get(uint16(i)).(uint16)
		if g.cycleDetectionDirected(int(val), visited, curStack) {
			return true
		}
	}
	curStack[vertex] = false
	return false
}

func (g GraphAdjacencyList) TopologicalSort() []int {
	visited := make([]bool, g.order)
	stack := stacklist.NewStackListEmp()

	for i := 0; i < int(g.order); i++ {
		if !visited[i] {
			g.topologicalSort(i, visited, stack)
		}
	}

	order := make([]int, stack.Size())
	for i := 0; i < len(order); i++ {
		order[i] = stack.Pop().(int)
	}

	return order
}

func (g GraphAdjacencyList) topologicalSort(vertex int, visited []bool, stack stack.Stack) {
	visited[vertex] = true

	for i := 1; i <= int(g.elements[vertex].Size()); i++ {
		val := g.elements[vertex].Get(uint16(i)).(uint16)
		if !visited[val] {
			g.topologicalSort(int(val), visited, stack)
		}
	}

	stack.Push(vertex)
}

func (g GraphAdjacencyList) Transpose() graph.Graph {
	transposed := NewGraphAdjacencyList(g.order)
	for i := 0; i < int(g.order); i++ {
		for j := 1; j <= int(g.elements[i].Size()); j++ {
			val := g.elements[i].Get(uint16(j)).(uint16)
			transposed.AddEdgeDirected(val, uint16(i))
		}
	}
	return transposed
}

func (g GraphAdjacencyList) StronglyConnectedComponents() {
	visited := make([]bool, g.order)
	stack := stacklist.NewStackListEmp()
	for i := 0; i < int(g.order); i++ {
		if !visited[i] {
			g.topologicalSort(i, visited, stack)
		}
	}

	transposed := g.Transpose()

	for i := 0; i < int(g.order); i++ {
		visited[i] = false
	}

	for !stack.IsEmpty() {
		vertex := stack.Pop().(int)
		if !visited[vertex] {
			transposed.DepthFirstSearchRecursiveCall(vertex, visited)
			fmt.Println()
		}
	}
}

func (g GraphAdjacencyList) MotherGraph() int {
	visited := make([]bool, g.order)
	var lastVisited int
	for i := 0; i < int(g.order); i++ {
		if !visited[i] {
			g.DepthFirstSearchRecursiveCall(i, visited)
			lastVisited = i
		}
	}

	for i := 0; i < int(g.order); i++ {
		visited[i] = false
	}

	g.DepthFirstSearchRecursiveCall(lastVisited, visited)
	for i := 0; i < len(visited); i++ {
		if !visited[i] {
			return -1
		}
	}

	return lastVisited
}

func (g GraphAdjacencyList) NumberEdges(directed bool) int {
	sum := 0
	for i := 0; i < int(g.order); i++ {
		sum += int(g.elements[i].Size())
	}

	if directed {
		return sum
	}

	return sum / 2
}

func (g GraphAdjacencyList) PathExists(source int, target int) bool {
	if source == target {
		return true
	}

	visited := make([]bool, g.order)
	stack := stacklist.NewStackListEmp()

	stack.Push(source)
	visited[source] = true

	for !stack.IsEmpty() {
		vertex := stack.Pop().(int)

		for i := 1; i <= int(g.elements[vertex].Size()); i++ {
			val := g.elements[vertex].Get(uint16(i)).(uint16)
			if !visited[val] {
				if int(val) == target {
					return true
				}

				stack.Push(int(val))
				visited[vertex] = true
			}
		}
	}
	return false
}

func (g GraphAdjacencyList) HasOneParent(root int) bool {
	visited := make([]bool, g.order)

	for i := 0; i < int(g.order); i++ {
		for j := 1; j <= int(g.elements[i].Size()); j++ {
			val := g.elements[i].Get(uint16(j)).(uint16)
			if visited[val] {
				return false
			}
			visited[val] = true
		}
	}

	// assuming there's only one root
	for i := 0; i < int(g.order); i++ {
		if i == root && visited[i] {
			return false
		} else if i != root && !visited[i] {
			return false
		}
	}

	return true
}

func (g GraphAdjacencyList) ShortestPathEdges(source int, target int) int {

	visited := make([]bool, g.order)
	distance := make([]int, g.order)
	queue := queuelist.NewLinearQueueEmp()

	queue.Enqueue(source)

	for !queue.IsEmpty() {
		vertex := queue.Dequeue().(int)

		for i := 1; i <= int(g.elements[vertex].Size()); i++ {
			val := g.elements[vertex].Get(uint16(i)).(uint16)
			if !visited[val] {
				queue.Enqueue(int(val))
				visited[val] = true
				distance[val] = distance[vertex] + 1
			}

			if int(val) == target {
				return distance[val]
			}
		}
	}
	return -1
}

func (g GraphAdjacencyList) IsBipartite(root int) bool {

	coloured := make([]int, g.order)
	queue := queuelist.NewLinearQueueEmp()

	queue.Enqueue(root)
	coloured[root] = 1

	for !queue.IsEmpty() {
		parent := queue.Dequeue().(int)

		for i := 1; i <= int(g.elements[parent].Size()); i++ {
			child := g.elements[parent].Get(uint16(i)).(uint16)
			if coloured[child] == 0 {
				coloured[child] = coloured[parent] * -1
				queue.Enqueue(int(child))
			} else if coloured[child] == coloured[parent] {
				return false
			}
		}
	}
	return true

}
