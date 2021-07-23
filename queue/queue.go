package queue

type Queue interface {
	Enqueue(interface{})
	Dequeue() interface{}
	IsEmpty() bool
	Top() interface{}
	Size() uint16
}
