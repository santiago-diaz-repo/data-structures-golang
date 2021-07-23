package list

type List interface {
	AddAt(uint16, interface{}) bool
	AddAtEnd(interface{})
	AddAtHead(interface{})
	Contains(interface{}) bool
	GetFirstElement() interface{}
	Get(uint16) interface{}
	IndexOf(interface{}) int16
	IsEmpty() bool
	GetLastElement() interface{}
	Print()
	RemoveAt(uint16) interface{}
	RemoveAtEnd() interface{}
	RemoveAtHead() interface{}
	Replace(uint16, interface{}) interface{}
	Size() uint16
}
