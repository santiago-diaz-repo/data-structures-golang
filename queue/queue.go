package queue

type Queue interface {
	Enqueue(interface{}) error
	Dequeue() interface{}
	IsEmpty() bool
	IsFull() bool
	Peek() interface{}
	Size() uint16
}
