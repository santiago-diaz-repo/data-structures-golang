package doublylinkedlist

import (
	"data-structures-golang/list"
	"fmt"
)

type node struct {
	element interface{}
	next    *node
	prev    *node
}

type DoublyLinkedList struct {
	head *node
	tail *node
	size uint16
}

func NewDoublyLinkedList(element interface{}) list.List {
	node := &node{
		element: element,
		next:    nil,
		prev:    nil,
	}

	return &DoublyLinkedList{
		head: node,
		tail: node,
		size: 1,
	}
}

func (dll *DoublyLinkedList) AddAt(pos uint16, element interface{}) bool {
	node := &node{
		element: element,
		next:    nil,
		prev:    nil,
	}

	if dll.IsEmpty() {
		if pos == 1 {
			dll.head = node
			dll.tail = node
			dll.size++
			return true
		} else {
			return false
		}
	}

	if dll.size < pos-1 {
		return false
	}

	if pos == 1 {
		dll.AddAtHead(element)
		return true
	}

	if pos == dll.size+1 {
		dll.AddAtEnd(element)
		return true
	}

	temp := dll.head
	count := uint16(1)
	for count < pos-1 {
		temp = temp.next
		count++
	}

	node.next = temp.next
	node.prev = temp
	node.next.prev = node
	temp.next = node

	dll.size++
	return true
}

func (dll *DoublyLinkedList) AddAtEnd(element interface{}) {
	node := &node{
		element: element,
		next:    nil,
		prev:    nil,
	}

	if dll.IsEmpty() {
		dll.head = node
		dll.tail = node
		dll.size++
		return
	}

	node.prev = dll.tail
	dll.tail.next = node
	dll.tail = node
	dll.size++
}

func (dll *DoublyLinkedList) AddAtHead(element interface{}) {
	node := &node{
		element: element,
		next:    nil,
		prev:    nil,
	}

	if dll.IsEmpty() {
		dll.head = node
		dll.tail = node
		dll.size++
		return
	}

	node.next = dll.head
	dll.head.prev = node
	dll.head = node
	dll.size++
}

func (dll *DoublyLinkedList) Contains(element interface{}) bool {
	if dll.IsEmpty() {
		return false
	}

	if dll.size == 1 && dll.head.element != element {
		return false
	}

	frontTmp := dll.head
	backTmp := dll.tail

	for {
		if frontTmp.element == element || backTmp.element == element {
			return true
		}

		if frontTmp.next == backTmp || frontTmp == backTmp {
			break
		}
		frontTmp = frontTmp.next
		backTmp = backTmp.prev
	}
	return false
}

func (dll DoublyLinkedList) GetFirstElement() interface{} {
	if dll.IsEmpty() {
		return nil
	}

	return dll.head.element
}

func (dll DoublyLinkedList) Get(index uint16) interface{} {
	if dll.IsEmpty() {
		return nil
	}

	if dll.size < index {
		return nil
	}

	tmp := dll.head
	for i := uint16(1); i < index; i++ {
		tmp = tmp.next
	}

	return tmp.element
}

func (dll DoublyLinkedList) IndexOf(element interface{}) int16 {
	if dll.IsEmpty() {
		return -1
	}

	pos := int16(1)
	temp := dll.head
	for temp != nil {
		if temp.element == element {
			return pos
		}
		temp = temp.next
		pos++
	}
	return -1
}

func (dll *DoublyLinkedList) IsEmpty() bool {
	if dll.size == 0 {
		return true
	}
	return false
}

func (dll DoublyLinkedList) GetLastElement() interface{} {
	if dll.IsEmpty() {
		return nil
	}

	return dll.tail.element
}

func (dll DoublyLinkedList) Print() {
	tmp := dll.head
	for tmp != nil {
		fmt.Printf("%v-> ", tmp.element)
		tmp = tmp.next
	}
	fmt.Println()
}

func (dll DoublyLinkedList) RemoveAt(pos uint16) interface{} {
	if dll.IsEmpty() {
		return nil
	}

	if dll.size < pos {
		return nil
	}

	if pos == 1 {
		return dll.RemoveAtHead()
	}

	if dll.size == pos {
		return dll.RemoveAtEnd()
	}

	temp := dll.head
	count := uint16(1)

	for count < pos-1 {
		temp = temp.next
		count++
	}

	old := temp.next
	temp.next = temp.next.next
	temp.next.prev = temp
	old.next = nil
	old.prev = nil
	return old.element

}

func (dll DoublyLinkedList) RemoveAtEnd() interface{} {
	if dll.IsEmpty() {
		return nil
	}

	tmp := dll.tail
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
		return tmp.element
	}

	dll.tail = dll.tail.prev
	dll.tail.next = nil
	dll.size--
	return tmp.element
}

func (dll DoublyLinkedList) RemoveAtHead() interface{} {
	if dll.IsEmpty() {
		return nil
	}

	tmp := dll.head
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
		return tmp.element
	}

	dll.head = dll.head.next
	dll.head.prev = nil
	dll.size--
	return tmp.element
}

func (dll DoublyLinkedList) Replace(pos uint16, element interface{}) interface{} {
	if dll.IsEmpty() {
		return nil
	}

	if dll.size < pos {
		return nil
	}

	temp := dll.head
	count := uint16(1)
	for count < pos {
		temp = temp.next
		count++
	}

	old := temp.element
	temp.element = element
	return old
}

func (dll *DoublyLinkedList) Size() uint16 {
	return dll.size
}
