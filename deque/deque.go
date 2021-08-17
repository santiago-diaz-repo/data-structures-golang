package deque

type Deque interface {
	AddFront(interface{}) error
	AddBack(interface{}) error
	RemoveFront() interface{}
	RemoveBack() interface{}
	PeekFirst() interface{}
	PeekLast() interface{}
	IsFull() bool
	IsEmpty() bool
	Size() uint16
}
