package n_ary

import (
	"data-structures-golang/list"
	"data-structures-golang/list/doublylinkedlist"
)

type treeNode struct {
	value    interface{}
	children list.List
}

type NAryTree struct {
	root *treeNode
}

func NewNAryTreeEmp() *NAryTree {
	return &NAryTree{
		root: nil,
	}
}

func NewNAryTreeValue(value interface{}) *NAryTree {
	root := treeNode{
		value:    value,
		children: doublylinkedlist.NewDoublyLinkedListEmp(),
	}

	return &NAryTree{
		root: &root,
	}
}

func NewAryTreeValueChildren(value interface{}, children list.List) *NAryTree {
	root := treeNode{
		value:    value,
		children: children,
	}

	return &NAryTree{
		root: &root,
	}
}
