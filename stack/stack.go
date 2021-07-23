package stack

type Stack interface {
	Push(interface{})
	Pop() interface{}
	Size() uint16
	IsEmpty() bool
	IsFull() bool
	Peek() interface{}
}
