package queuearray

import (
	"data-structures-golang/queue"
	"errors"
)

type CircularQueue struct {
	elements []interface{}
	front    int16
	back     int16
	size     uint16
}

func NewCircularQueue(size uint16) queue.Queue {
	return &CircularQueue{
		elements: make([]interface{}, size),
		front:    int16(-1),
		back:     int16(-1),
		size:     uint16(0),
	}
}

func (cq *CircularQueue) Enqueue(element interface{}) error {
	if cq.IsFull() {
		return errors.New("queue is full")
	}

	if cq.front == -1 && cq.back == -1 {
		cq.front = int16(0)
		cq.back = int16(0)
	} else {
		cq.back = (cq.back + 1) % int16(cap(cq.elements))
	}

	cq.elements[cq.back] = element
	cq.size++
	return nil
}

func (cq *CircularQueue) Dequeue() interface{} {
	if cq.IsEmpty() {
		return nil
	}

	temp := cq.elements[cq.front]
	if cq.front == cq.back {
		cq.front = -1
		cq.back = -1
	} else {
		cq.front = (cq.front + 1) % int16(cap(cq.elements))
	}

	cq.size--
	return temp
}

func (cq *CircularQueue) IsEmpty() bool {
	return cq.front == -1 && cq.back == -1
}

func (cq *CircularQueue) IsFull() bool {
	if cq.IsEmpty() {
		return false
	}
	mod := (cq.back + 1) % int16(cap(cq.elements))
	return cq.front == mod
}

func (cq *CircularQueue) Peek() interface{} {
	if cq.IsEmpty() {
		return nil
	}
	return cq.elements[cq.front]
}

func (cq *CircularQueue) Size() uint16 {
	return cq.size
}
