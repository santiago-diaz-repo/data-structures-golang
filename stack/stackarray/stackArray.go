package stackarray

import (
	"data-structures-golang/stack"
	"errors"
)

type StackArray struct {
	elements []interface{}
	front    int16
	size     uint16
}

func NewStackArray(size uint16) stack.Stack {
	return &StackArray{
		elements: make([]interface{}, size),
		front:    int16(-1),
		size:     uint16(0),
	}
}

func (sa *StackArray) Push(element interface{}) error {
	if sa.IsFull() {
		return errors.New("stack is full")
	}

	sa.front++
	sa.elements[sa.front] = element
	sa.size++
	return nil
}

func (sa *StackArray) Pop() interface{} {
	if sa.IsEmpty() {
		return nil
	}

	sa.front--
	sa.size--
	return sa.elements[sa.front+1]
}

func (sa *StackArray) Size() uint16 {
	return sa.size
}

func (sa *StackArray) IsEmpty() bool {
	return sa.front == int16(-1)
}

func (sa *StackArray) IsFull() bool {
	return sa.front == int16(len(sa.elements)-1)
}

func (sa *StackArray) Peek() interface{} {
	if sa.IsEmpty() {
		return nil
	}

	return sa.elements[sa.front]
}
