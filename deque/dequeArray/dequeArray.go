package dequeArray

import (
	"data-structures-golang/deque"
	"errors"
)

type DequeArray struct {
	elements []interface{}
	front    int16
	back     int16
	size     uint16
	capacity uint16
}

func NewDequeArray(capacity uint16) deque.Deque {
	return &DequeArray{
		elements: make([]interface{}, int(capacity)),
		front:    int16(-1),
		back:     int16(-1),
		size:     uint16(0),
		capacity: capacity,
	}
}

func (da *DequeArray) AddFront(data interface{}) error {
	if da.IsFull() {
		return errors.New("Deque is full")
	}

	if da.front == -1 && da.back == -1 {
		da.front++
		da.back++
	} else {
		if da.front == 0 {
			da.front = int16(da.capacity) - 1
		} else {
			da.front--
		}
	}
	da.elements[da.front] = data
	da.size++
	return nil
}

func (da *DequeArray) AddBack(data interface{}) error {
	if da.IsFull() {
		return errors.New("Deque is full")
	}

	if da.front == -1 && da.back == -1 {
		da.front++
		da.back++
	} else {
		if da.back == int16(da.capacity)-1 {
			da.back = int16(0)
		} else {
			da.back++
		}
	}

	da.elements[da.front] = data
	da.size++
	return nil
}

func (da *DequeArray) RemoveFront() interface{} {
	if da.IsEmpty() {
		return nil
	}

	var data interface{}
	if da.front == da.back {
		data = da.elements[da.front]
		da.front = -1
		da.back = -1
	} else {
		if da.front == int16(da.capacity)-1 {
			data = da.elements[da.front]
			da.front = 0
		} else {
			data = da.elements[da.front]
			da.front++
		}
	}
	da.size--
	return data
}

func (da *DequeArray) RemoveBack() interface{} {
	if da.IsEmpty() {
		return nil
	}

	var data interface{}
	if da.front == da.back {
		data = da.elements[da.back]
		da.front = -1
		da.back = -1
	} else {
		if da.back == 0 {
			data = da.elements[da.back]
			da.back = int16(da.capacity) - 1
		} else {
			data = da.elements[da.back]
			da.back--
		}
	}
	da.size--
	return data
}

func (da *DequeArray) PeekFirst() interface{} {
	if da.IsEmpty() {
		return nil
	}

	return da.elements[da.front]
}

func (da *DequeArray) PeekLast() interface{} {
	if da.IsEmpty() {
		return nil
	}

	return da.elements[da.back]
}

func (da *DequeArray) IsFull() bool {
	return da.back+1 == da.front || (da.front == 0 && da.back == int16(da.capacity)-1)
}

func (da *DequeArray) IsEmpty() bool {
	return da.front == -1 && da.back == -1
}

func (da *DequeArray) Size() uint16 {
	return da.size
}
