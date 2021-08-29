package queuelist

import (
	"data-structures-golang/list"
	"data-structures-golang/list/singlylinkedlist"
	"data-structures-golang/queue"
)

type LinearQueue struct {
	elements list.List
}

func NewLinearQueueEmp() queue.Queue {
	return &LinearQueue{
		elements: singlylinkedlist.NewSinglyLinkedListEmp(),
	}
}

func NewLinearQueueEle(element interface{}) queue.Queue {
	return &LinearQueue{
		elements: singlylinkedlist.NewSinglyLinkedListEle(element),
	}
}

func (lq *LinearQueue) Enqueue(element interface{}) error {
	lq.elements.AddAtEnd(element)
	return nil
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
