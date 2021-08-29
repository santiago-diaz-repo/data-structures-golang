package singlylinkedlist

import (
	"data-structures-golang/list"
	"fmt"
)

type node struct {
	element interface{}
	next    *node
}

type SinglyLinkedList struct {
	head *node
	tail *node
	size uint16
}

func NewSinglyLinkedListEmp() list.List {
	return &SinglyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func NewSinglyLinkedListEle(element interface{}) list.List {
	node := &node{
		element: element,
		next:    nil,
	}

	return &SinglyLinkedList{
		head: node,
		tail: node,
		size: 1,
	}
}

func (sll *SinglyLinkedList) AddAt(pos uint16, element interface{}) bool {
	ele := &node{
		element: element,
		next:    nil,
	}

	if sll.IsEmpty() {
		if pos == 1 {
			sll.head = ele
			sll.tail = ele
			sll.size++
			return true
		} else {
			return false
		}
	}

	if sll.size < pos-1 {
		return false
	}

	if pos == 1 {
		sll.AddAtHead(element)
		return true
	}

	if pos == sll.size+1 {
		sll.AddAtEnd(element)
		return true
	}

	temp := sll.head
	count := uint16(1)
	for count < pos-1 {
		temp = temp.next
		count++
	}

	ele.next = temp.next
	temp.next = ele

	sll.size++
	return true
}

func (sll *SinglyLinkedList) AddAtEnd(element interface{}) {
	node := &node{
		element: element,
		next:    nil,
	}

	if sll.IsEmpty() {
		sll.head = node
		sll.tail = node
		sll.size++
		return
	}

	sll.tail.next = node
	sll.tail = node
	sll.size++
}

func (sll *SinglyLinkedList) AddAtHead(element interface{}) {

	node := &node{
		element: element,
		next:    nil,
	}

	if sll.IsEmpty() {
		sll.head = node
		sll.tail = node
		sll.size++
		return
	}

	if sll.head == sll.tail {
		node.next = sll.tail
		sll.head = node
	} else {
		node.next = sll.head
		sll.head = node
	}

	sll.size++
}

func (sll *SinglyLinkedList) Contains(element interface{}) bool {
	if sll.IsEmpty() {
		return false
	}

	if sll.size == 1 && sll.head.element != element {
		return false
	}

	tmp := sll.head

	for tmp != nil {
		if tmp.element == element {
			return true
		}
		tmp = tmp.next
	}
	return false
}

func (sll *SinglyLinkedList) GetFirstElement() interface{} {
	if sll.IsEmpty() {
		return nil
	}

	return sll.head.element
}

func (sll *SinglyLinkedList) Get(index uint16) interface{} {
	if sll.IsEmpty() {
		return nil
	}

	if sll.size < index {
		return nil
	}

	if index == 1 {
		return sll.head.element
	}

	if sll.size == index {
		return sll.tail.element
	}

	tmp := sll.head
	for i := uint16(1); i < index; i++ {
		tmp = tmp.next
	}

	return tmp.element
}

func (sll *SinglyLinkedList) IndexOf(element interface{}) int16 {
	if sll.IsEmpty() {
		return -1
	}

	pos := int16(1)
	temp := sll.head
	for temp != nil {
		if temp.element == element {
			return pos
		}
		temp = temp.next
		pos++
	}
	return -1
}

func (sll *SinglyLinkedList) IsEmpty() bool {
	if sll.size == 0 {
		return true
	}
	return false
}

func (sll *SinglyLinkedList) GetLastElement() interface{} {
	if sll.IsEmpty() {
		return nil
	}

	return sll.tail.element
}

func (sll *SinglyLinkedList) Print() {

	tmp := sll.head
	for tmp != nil {
		fmt.Printf("%v-> ", tmp.element)
		tmp = tmp.next
	}
	fmt.Println()
}

func (sll *SinglyLinkedList) RemoveAt(pos uint16) interface{} {
	if sll.IsEmpty() {
		return nil
	}

	if sll.size < pos {
		return nil
	}

	if pos == 1 {
		return sll.RemoveAtHead()
	}

	if pos == sll.size {
		return sll.RemoveAtEnd()
	}

	temp := sll.head
	count := uint16(1)

	for count < pos-1 {
		temp = temp.next
		count++
	}

	old := temp.next
	temp.next = temp.next.next
	return old.element

}

func (sll *SinglyLinkedList) RemoveAtEnd() interface{} {
	if sll.IsEmpty() {
		return nil
	}

	tmp := sll.head
	if sll.head == sll.tail {
		sll.head = nil
		sll.tail = nil
	} else {
		for tmp.next != sll.tail {
			tmp = tmp.next
		}

		sll.tail = tmp
		tmp = sll.tail.next
		sll.tail = nil
	}

	sll.size--
	return tmp.element
}

func (sll *SinglyLinkedList) RemoveAtHead() interface{} {
	if sll.IsEmpty() {
		return nil
	}
	tmp := sll.head
	sll.head = sll.head.next
	sll.size--
	return tmp.element
}

func (sll *SinglyLinkedList) Replace(pos uint16, element interface{}) interface{} {
	if sll.IsEmpty() {
		return nil
	}

	if sll.size < pos {
		return nil
	}

	temp := sll.head
	count := uint16(1)
	for count < pos {
		temp = temp.next
		count++
	}

	old := temp.element
	temp.element = element
	return old
}

func (sll *SinglyLinkedList) Size() uint16 {
	return sll.size
}
