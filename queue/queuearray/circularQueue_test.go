package queuearray

import (
	"errors"
	"reflect"
	"testing"
)

func getEmptyCircularQueue() CircularQueue {
	arr := make([]interface{}, 5)
	return CircularQueue{
		elements: arr,
		front:    int16(-1),
		back:     int16(-1),
		size:     uint16(0),
	}
}

func getOneElementLongCircularQueue() CircularQueue {
	arr := make([]interface{}, 5)
	arr[0] = 0
	return CircularQueue{
		elements: arr,
		front:    int16(0),
		back:     int16(0),
		size:     uint16(1),
	}
}

func getManyElementLongCircularQueue() CircularQueue {
	arr := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		arr[i] = i
	}
	return CircularQueue{
		elements: arr,
		front:    int16(0),
		back:     int16(14),
		size:     uint16(15),
	}
}

func TestNewCircularQueue(t *testing.T) {
	want := "*queuearray.CircularQueue"
	subject := NewCircularQueue(6)
	got := reflect.TypeOf(subject)
	if got.String() != want {
		t.Errorf("It expected %v but got %v", want, got)
	}
}

func TestLinearQueue_Dequeue(t *testing.T) {
	dataTable := []struct {
		name string
		data CircularQueue
		want interface{}
	}{
		{"EmptyCircularQueue", getEmptyCircularQueue(), nil},
		{"OneElementLongCircularQueue", getOneElementLongCircularQueue(), 0},
		{"ManyElementsLongCircularQueue", getManyElementLongCircularQueue(), 0},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.Dequeue()
			if got != v.want {
				t.Errorf("It expected %v, but got %v", v.want, got)
			}
		})
	}
}

func TestLinearQueue_Enqueue(t *testing.T) {
	dataTable := []struct {
		name    string
		data    CircularQueue
		element interface{}
		want    error
	}{
		{"EmptyCircularQueue", getEmptyCircularQueue(), 2, nil},
		{"OneElementLongCircularQueue", getOneElementLongCircularQueue(), 2, nil},
		{"FullCircularQueue", getManyElementLongCircularQueue(), 2, errors.New("test")},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.Enqueue(v.element)
			if v.want == nil {
				if got != v.want {
					t.Errorf("It expected %v, but got %v", v.want, got)
				}
			} else {
				if got == nil {
					t.Errorf("It expected %v, but got %v", nil, got)
				}
			}
		})
	}
}

func TestLinearQueue_IsEmpty(t *testing.T) {
	dataTable := []struct {
		name string
		data CircularQueue
		want bool
	}{
		{"EmptyCircularQueue", getEmptyCircularQueue(), true},
		{"OneElementLongCircularQueue", getOneElementLongCircularQueue(), false},
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

func TestLinearQueue_IsFull(t *testing.T) {
	dataTable := []struct {
		name string
		data CircularQueue
		want bool
	}{
		{"EmptyLinearQueue", getEmptyCircularQueue(), false},
		{"NonFullCircularQueue", getOneElementLongCircularQueue(), false},
		{"FullCircularQueue", getManyElementLongCircularQueue(), true},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.IsFull()
			if got != v.want {
				t.Errorf("It expected %v, but got %v", v.want, got)
			}
		})
	}
}

func TestLinearQueue_Size(t *testing.T) {
	dataTable := []struct {
		name string
		data CircularQueue
		want uint16
	}{
		{"EmptyCircularQueue", getEmptyCircularQueue(), uint16(0)},
		{"OneElementLongCircularQueue", getOneElementLongCircularQueue(), uint16(1)},
		{"ManyElementsLongCircularQueue", getManyElementLongCircularQueue(), uint16(15)},
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

func TestLinearQueue_Peek(t *testing.T) {
	dataTable := []struct {
		name string
		data CircularQueue
		want interface{}
	}{
		{"EmptyCircularQueue", getEmptyCircularQueue(), nil},
		{"OneElementLongCircularQueue", getOneElementLongCircularQueue(), 0},
		{"ManyElementsLongCircularQueue", getManyElementLongCircularQueue(), 0},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			v.data.Peek()
			got := v.data.Peek()
			if got != v.want {
				t.Errorf("It expected %v, but got %v", v.want, got)
			}
		})
	}
}
