package doublylinkedlist

import (
	"reflect"
	"testing"
)

func getEmptyList() DoublyLinkedList {
	return DoublyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func getOneNodeLongList() DoublyLinkedList {
	node := &node{
		element: uint16(0),
		next:    nil,
		prev:    nil,
	}

	return DoublyLinkedList{
		head: node,
		tail: node,
		size: 1,
	}
}

func getOddNumberLongList() DoublyLinkedList {
	list := getOneNodeLongList()
	var nod *node

	for i := 1; i < 3; i++ {
		nod = new(node)
		nod.element = uint16(i)
		nod.next = nil
		list.tail.next = nod
		nod.prev = list.tail
		list.tail = nod
		list.size++
	}

	return list
}

func getEvenNumberLongList() DoublyLinkedList {
	list := getOneNodeLongList()
	var nod *node

	for i := 1; i < 4; i++ {
		nod = new(node)
		nod.element = uint16(i)
		nod.next = nil
		list.tail.next = nod
		nod.prev = list.tail
		list.tail = nod
		list.size++
	}

	return list
}

func TestNewDoublyLinkedListEmp(t *testing.T) {
	want := "*doublylinkedlist.DoublyLinkedList"
	subject := NewDoublyLinkedListEmp()
	got := reflect.TypeOf(subject)
	if got.String() != want {
		t.Errorf("It expected %v but got %v", want, got)
	}
}

func TestNewDoublyLinkedListEle(t *testing.T) {
	want := "*doublylinkedlist.DoublyLinkedList"
	subject := NewDoublyLinkedListEle(5)
	got := reflect.TypeOf(subject)
	if got.String() != want {
		t.Errorf("It expected %v but got %v", want, got)
	}
}

func TestDoublyLinkedList_Size(t *testing.T) {
	dataTable := []struct {
		name string
		data DoublyLinkedList
		want uint16
	}{
		{"EmptyList", getEmptyList(), 0},
		{"OneLongSize", getOneNodeLongList(), 1},
		{"ThreeLongSize", getOddNumberLongList(), 3},
		{"FourLongSize", getEvenNumberLongList(), 4},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.Size()
			if got != v.want {
				t.Errorf("It expected %v, but got %v", v.want, got)
			}
		})
	}
}

func TestDoublyLinkedList_IsEmpty(t *testing.T) {
	dataTable := []struct {
		name string
		data DoublyLinkedList
		want bool
	}{
		{"EmptyList", getEmptyList(), true},
		{"NotEmptyList", getOneNodeLongList(), false},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.IsEmpty()
			if got != v.want {
				t.Errorf("It expected %v, but got %v", v.want, got)
			}
		})
	}
}

func TestDoublyLinkedList_Get(t *testing.T) {

	dataTable := []struct {
		name   string
		data   DoublyLinkedList
		getPos uint16
		want   interface{}
	}{
		{"EmptyList", getEmptyList(), uint16(1), nil},
		{"GraterThanCurrentSize", getOneNodeLongList(), uint16(2), nil},
		{"GetFirstElement", getOddNumberLongList(), uint16(1), uint16(0)},
		{"GetLastElement", getOddNumberLongList(), uint16(3), uint16(2)},
		{"GetInBetweenElement", getEvenNumberLongList(), uint16(3), uint16(2)},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.Get(v.getPos)
			if got != v.want {
				t.Errorf("It expected %v, but got %v", v.want, got)
			}
		})
	}
}

func TestDoublyLinkedList_AddAtEnd(t *testing.T) {

	dataTable := []struct {
		name string
		data DoublyLinkedList
		want uint16
	}{
		{"EmptyList", getEmptyList(), uint16(4)},
		{"NonEmptyList", getOneNodeLongList(), uint16(4)},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			v.data.AddAtEnd(v.want)
			got := v.data.tail.element
			if got != v.want {
				t.Errorf("It expected %v, but got %v", v.want, got)
			}
		})
	}
}

func TestDoublyLinkedList_AddAtHead(t *testing.T) {

	dataTable := []struct {
		name string
		data DoublyLinkedList
		want uint16
	}{
		{"EmptyList", getEmptyList(), uint16(4)},
		{"OneElementLong", getOneNodeLongList(), uint16(4)},
		{"MoreThenOneElementLong", getOddNumberLongList(), uint16(4)},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			v.data.AddAtHead(v.want)
			got := v.data.head.element
			if got != v.want {
				t.Errorf("It expected %v, but got %v", v.want, got)
			}
		})
	}
}

func TestDoublyLinkedList_RemoveAtEndAtEnd(t *testing.T) {

	dataTable := []struct {
		name string
		data DoublyLinkedList
		want interface{}
	}{
		{"EmptyList", getEmptyList(), nil},
		{"OneElementLong", getOneNodeLongList(), uint16(0)},
		{"MoreThenOneElementLong", getOddNumberLongList(), uint16(2)},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.RemoveAtEnd()
			if got != v.want {
				t.Errorf("It expected %v, but got %v", v.want, got)
			}
		})
	}
}

func TestDoublyLinkedList_RemoveAtEndAtHead(t *testing.T) {
	dataTable := []struct {
		name string
		data DoublyLinkedList
		want interface{}
	}{
		{"EmptyList", getEmptyList(), nil},
		{"OneElementLong", getOneNodeLongList(), uint16(0)},
		{"MoreThenOneElementLong", getOddNumberLongList(), uint16(0)},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.RemoveAtHead()
			if got != v.want {
				t.Errorf("It expected %v, but got %v", v.want, got)
			}
		})
	}
}

func TestDoublyLinkedList_Contains(t *testing.T) {
	dataTable := []struct {
		name  string
		data  DoublyLinkedList
		value uint16
		want  bool
	}{
		{"EmptyList", getEmptyList(), uint16(9), false},
		{"OneElementLongNotSameValue", getOneNodeLongList(), uint16(9), false},
		{"OneElementLongSameValue", getOneNodeLongList(), uint16(0), true},
		{"MoreThanOneElementLong", getOddNumberLongList(), uint16(2), true},
		{"Notexist", getOddNumberLongList(), uint16(3), false},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.Contains(v.value)
			if got != v.want {
				t.Errorf("It expected %v, but got %v", v.want, got)
			}
		})
	}
}

func TestDoublyLinkedList_GetFirstElement(t *testing.T) {
	dataTable := []struct {
		name string
		data DoublyLinkedList
		want interface{}
	}{
		{"EmptyList", getEmptyList(), nil},
		{"OneElementLong", getOneNodeLongList(), uint16(0)},
		{"MoreThenOneElementLong", getOddNumberLongList(), uint16(0)},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.GetFirstElement()
			if got != v.want {
				t.Errorf("It expected %v, but got %v", v.want, got)
			}
		})
	}
}

func TestDoublyLinkedList_GetLastElement(t *testing.T) {
	dataTable := []struct {
		name string
		data DoublyLinkedList
		want interface{}
	}{
		{"EmptyList", getEmptyList(), nil},
		{"OneElementLong", getOneNodeLongList(), uint16(0)},
		{"MoreThenOneElementLong", getOddNumberLongList(), uint16(2)},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.GetLastElement()
			if got != v.want {
				t.Errorf("It expected %v, but got %v", v.want, got)
			}
		})
	}
}

func TestDoublyLinkedList_IndexOf(t *testing.T) {
	dataTable := []struct {
		name    string
		data    DoublyLinkedList
		element interface{}
		want    int16
	}{
		{"EmptyList", getEmptyList(), uint16(5), -1},
		{"OneElementLong", getOneNodeLongList(), uint16(0), 1},
		{"MoreThenOneElementLong", getEvenNumberLongList(), uint16(2), 3},
		{"MoreThenOneElementLongNotExist", getEvenNumberLongList(), uint16(9), -1},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.IndexOf(v.element)
			if got != v.want {
				t.Errorf("It expected %v, but got %v", v.want, got)
			}
		})
	}
}

func TestDoublyLinkedList_AddAt(t *testing.T) {
	dataTable := []struct {
		name        string
		data        DoublyLinkedList
		pos         uint16
		wantElement interface{}
		wantResult  bool
	}{
		{"EmptyListPosOne", getEmptyList(), uint16(1), uint16(4), true},
		{"EmptyListPosNotOne", getEmptyList(), uint16(2), uint16(4), false},
		{"GreaterThanSize", getEvenNumberLongList(), uint16(6), uint16(4), false},
		{"PosOne", getEvenNumberLongList(), uint16(1), uint16(4), true},
		{"SizePlusOne", getEvenNumberLongList(), uint16(5), uint16(4), true},
		{"InBetween", getEvenNumberLongList(), uint16(3), uint16(4), true},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.AddAt(v.pos, v.wantElement)
			if got != v.wantResult {
				t.Errorf("It expected %v, but got %v", v.wantResult, got)
			}

			if v.wantResult {
				cont := uint16(1)
				for cont < v.pos {
					v.data.head = v.data.head.next
					cont++
				}

				if v.data.head.element != v.wantElement {
					t.Errorf("It expected %v, but got %v", v.wantElement, v.data.head.element)
				}
			}
		})
	}
}

func TestDoublyLinkedList_RemoveAt(t *testing.T) {
	dataTable := []struct {
		name string
		data DoublyLinkedList
		pos  uint16
		want interface{}
	}{
		{"EmptyList", getEmptyList(), uint16(1), nil},
		{"GreaterThanOneElement", getOneNodeLongList(), uint16(2), nil},
		{"FirstPos", getEvenNumberLongList(), uint16(1), uint16(0)},
		{"LastPos", getEvenNumberLongList(), uint16(4), uint16(3)},
		{"InBetween", getEvenNumberLongList(), uint16(3), uint16(2)},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.RemoveAt(v.pos)
			if got != v.want {
				t.Errorf("It expected %v, but got %v", v.want, got)
			}
		})
	}
}

func TestDoublyLinkedList_Replace(t *testing.T) {
	dataTable := []struct {
		name    string
		data    DoublyLinkedList
		pos     uint16
		element interface{}
		want    interface{}
	}{
		{"EmptyList", getEmptyList(), uint16(1), uint16(7), nil},
		{"GreaterThanOneElement", getOneNodeLongList(), uint16(2), uint16(7), nil},
		{"FirstPos", getEvenNumberLongList(), uint16(1), uint16(7), uint16(0)},
		{"LastPos", getEvenNumberLongList(), uint16(4), uint16(7), uint16(3)},
		{"InBetween", getEvenNumberLongList(), uint16(2), uint16(7), uint16(1)},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.Replace(v.pos, v.element)
			if got != v.want {
				t.Errorf("It expected %v, but got %v", v.want, got)
			}

			if got != nil {
				cont := uint16(1)
				for cont < v.pos {
					v.data.head = v.data.head.next
					cont++
				}

				if v.data.head.element != v.element {
					t.Errorf("It expected %v, but got %v", v.element, v.data.head.element)
				}
			}
		})
	}
}
