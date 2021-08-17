package dequeArray

import (
	"errors"
	"reflect"
	"testing"
)

func getEmptyDequeArray() DequeArray {
	arr := make([]interface{}, 5)
	return DequeArray{
		elements: arr,
		front:    int16(-1),
		back:     int16(-1),
		size:     uint16(0),
		capacity: uint16(5),
	}
}

func getOneElementDequeArray() DequeArray {
	arr := make([]interface{}, 5)
	arr[0] = 0
	return DequeArray{
		elements: arr,
		front:    int16(0),
		back:     int16(0),
		size:     uint16(1),
		capacity: uint16(5),
	}
}

func getTwoElementDequeArray() DequeArray {
	arr := make([]interface{}, 5)
	arr[0] = 0
	arr[4] = 3
	return DequeArray{
		elements: arr,
		front:    int16(4),
		back:     int16(0),
		size:     uint16(2),
		capacity: uint16(5),
	}
}

func getThreeElementDequeArray() DequeArray {
	arr := make([]interface{}, 10)
	arr[4] = 1
	arr[5] = 2
	arr[6] = 3
	return DequeArray{
		elements: arr,
		front:    int16(4),
		back:     int16(6),
		size:     uint16(3),
		capacity: uint16(10),
	}
}

func getFourElementDequeArray() DequeArray {
	arr := make([]interface{}, 10)
	arr[7] = 1
	arr[8] = 2
	arr[9] = 3
	return DequeArray{
		elements: arr,
		front:    int16(7),
		back:     int16(9),
		size:     uint16(3),
		capacity: uint16(10),
	}
}

func getManyElementLongDequeArray() DequeArray {
	arr := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		arr[i] = i
	}
	return DequeArray{
		elements: arr,
		front:    int16(0),
		back:     int16(14),
		size:     uint16(15),
		capacity: uint16(15),
	}
}

func TestNewDequeArray(t *testing.T) {
	want := "*dequeArray.DequeArray"
	subject := NewDequeArray(5)
	got := reflect.TypeOf(subject)
	if got.String() != want {
		t.Errorf("It expected %v but got %v", want, got)
	}
}

func TestDequeArray_AddFront(t *testing.T) {
	dataTable := []struct {
		name string
		data DequeArray
		want interface{}
	}{
		{"EmptyDequeArray", getEmptyDequeArray(), nil},
		{"OneElement", getOneElementDequeArray(), nil},
		{"ThreeElements", getThreeElementDequeArray(), nil},
		{"FullDequeArray", getManyElementLongDequeArray(), errors.New("Deque is full")},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.AddFront(5)
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

func TestDequeArray_AddBack(t *testing.T) {
	dataTable := []struct {
		name string
		data DequeArray
		want interface{}
	}{
		{"EmptyDequeArray", getEmptyDequeArray(), nil},
		{"OneElement", getOneElementDequeArray(), nil},
		{"ThreeElements", getThreeElementDequeArray(), nil},
		{"FourElements", getFourElementDequeArray(), nil},
		{"FullDequeArray", getManyElementLongDequeArray(), errors.New("Deque is full")},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.AddBack(5)
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

func TestDequeArray_RemoveFront(t *testing.T) {
	dataTable := []struct {
		name string
		data DequeArray
		want interface{}
	}{
		{"EmptyDequeArray", getEmptyDequeArray(), nil},
		{"OneElement", getOneElementDequeArray(), 0},
		{"TwoElements", getTwoElementDequeArray(), 3},
		{"ThreeElements", getThreeElementDequeArray(), 1},
		{"FullDequeArray", getManyElementLongDequeArray(), 0},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.RemoveFront()
			if got != v.want {
				t.Errorf("It expected %v but got %v", v.want, got)
			}
		})
	}
}

func TestDequeArray_RemoveBack(t *testing.T) {
	dataTable := []struct {
		name string
		data DequeArray
		want interface{}
	}{
		{"EmptyDequeArray", getEmptyDequeArray(), nil},
		{"OneElement", getOneElementDequeArray(), 0},
		{"TwoElements", getTwoElementDequeArray(), 0},
		{"ThreeElements", getThreeElementDequeArray(), 3},
		{"FullDequeArray", getManyElementLongDequeArray(), 14},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.RemoveBack()
			if got != v.want {
				t.Errorf("It expected %v but got %v", v.want, got)
			}
		})
	}
}

func TestDequeArray_IsEmpty(t *testing.T) {
	dataTable := []struct {
		name string
		data DequeArray
		want bool
	}{
		{"EmptyDequeArray", getEmptyDequeArray(), true},
		{"OneElement", getOneElementDequeArray(), false},
		{"ThreeElements", getThreeElementDequeArray(), false},
		{"FullDequeArray", getManyElementLongDequeArray(), false},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.IsEmpty()
			if got != v.want {
				t.Errorf("It expected %v but got %v", v.want, got)
			}
		})
	}
}

func TestDequeArray_IsFull(t *testing.T) {
	dataTable := []struct {
		name string
		data DequeArray
		want bool
	}{
		{"EmptyDequeArray", getEmptyDequeArray(), false},
		{"OneElement", getOneElementDequeArray(), false},
		{"ThreeElements", getThreeElementDequeArray(), false},
		{"FullDequeArray", getManyElementLongDequeArray(), true},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.IsFull()
			if got != v.want {
				t.Errorf("It expected %v but got %v", v.want, got)
			}
		})
	}
}

func TestDequeArray_PeekFirst(t *testing.T) {
	dataTable := []struct {
		name string
		data DequeArray
		want interface{}
	}{
		{"EmptyDequeArray", getEmptyDequeArray(), nil},
		{"OneElement", getOneElementDequeArray(), 0},
		{"ThreeElements", getThreeElementDequeArray(), 1},
		{"FullDequeArray", getManyElementLongDequeArray(), 0},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.PeekFirst()
			if got != v.want {
				t.Errorf("It expected %v but got %v", v.want, got)
			}
		})
	}
}

func TestDequeArray_PeekLast(t *testing.T) {
	dataTable := []struct {
		name string
		data DequeArray
		want interface{}
	}{
		{"EmptyDequeArray", getEmptyDequeArray(), nil},
		{"OneElement", getOneElementDequeArray(), 0},
		{"ThreeElements", getThreeElementDequeArray(), 3},
		{"FullDequeArray", getManyElementLongDequeArray(), 14},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.PeekLast()
			if got != v.want {
				t.Errorf("It expected %v but got %v", v.want, got)
			}
		})
	}
}

func TestDequeArray_Size(t *testing.T) {
	dataTable := []struct {
		name string
		data DequeArray
		want uint16
	}{
		{"EmptyDequeArray", getEmptyDequeArray(), 0},
		{"OneElement", getOneElementDequeArray(), 1},
		{"ThreeElements", getThreeElementDequeArray(), 3},
		{"FullDequeArray", getManyElementLongDequeArray(), 15},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.Size()
			if got != v.want {
				t.Errorf("It expected %v but got %v", v.want, got)
			}
		})
	}
}
