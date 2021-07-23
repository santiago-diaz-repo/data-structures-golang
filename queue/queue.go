package queue

type Queue interface {
	Enqueue(interface{})
	Dequeue() interface{}
	IsEmpty() bool
	IsFull() bool
	Peek() interface{}
	Size() uint16
}
