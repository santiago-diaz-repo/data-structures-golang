package graph

import "data-structures-golang/list"

type Graph interface {
	Elements() []list.List
	AddEdge(uint16, uint16) error
	AddEdgeDirected(uint16, uint16) error
	DepthFirstSearchRecursiveCall(int, []bool)
}
