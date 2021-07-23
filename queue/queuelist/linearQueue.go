package queuelist

import (
	"data-structures-golang/list"
	"data-structures-golang/list/singlylinkedlist"
	"data-structures-golang/queue"
)

type LinearQueue struct {
	elements list.List
}

func NewLinearQueue(element interface{}) queue.Queue {
	return &LinearQueue{
		elements: singlylinkedlist.NewSinglyLinkedList(element),
	}
}

func (lq *LinearQueue) Enqueue(element interface{}) {
	lq.elements.AddAtEnd(element)
}

func (lq *LinearQueue) Dequeue() interface{} {
	return lq.elements.RemoveAtHead()
}

func (lq *LinearQueue) IsEmpty() bool {
	return lq.elements.Size() == 0
}

func (lq *LinearQueue) IsFull() bool {
	return false
}

func (lq *LinearQueue) Peek() interface{} {
	return lq.elements.GetFirstElement()
}

func (lq *LinearQueue) Size() uint16 {
	return lq.elements.Size()
}
