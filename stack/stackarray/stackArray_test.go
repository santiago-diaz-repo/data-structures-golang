package stackarray

import (
	"reflect"
	"testing"
)

func getEmptyStackArray() StackArray {
	arr := make([]interface{}, 5)
	return StackArray{
		elements: arr,
		front:    int16(-1),
		size:     uint16(0),
	}
}

func getOneElementLongStackArray() StackArray {
	arr := make([]interface{}, 5)
	arr[0] = 0
	return StackArray{
		elements: arr,
		front:    int16(0),
		size:     uint16(1),
	}
}

func getManyElementLongStackArray() StackArray {
	arr := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		arr[i] = i
	}
	return StackArray{
		elements: arr,
		front:    int16(14),
		size:     uint16(15),
	}
}

func TestNewStackArray(t *testing.T) {
	want := "*stackarray.StackArray"
	subject := NewStackArray(2)
	got := reflect.TypeOf(subject)
	if got.String() != want {
		t.Errorf("It expected %v but got %v", want, got)
	}
}

func TestStackArray_Push(t *testing.T) {
	dataTable := []struct {
		name string
		data StackArray
		want interface{}
	}{
		{"EmptyStackArray", getEmptyStackArray(), 4},
		{"OneElementLongStackArray", getOneElementLongStackArray(), 0},
		{"ManyElementsLongStackArray", getManyElementLongStackArray(), 0},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			v.data.Push(v.want)
			got := v.data.elements[0]
			if got != v.want {
				t.Errorf("It expected %v, but got %v", v.want, got)
			}
		})
	}
}

func TestStackArray_Pop(t *testing.T) {
	dataTable := []struct {
		name string
		data StackArray
		want interface{}
	}{
		{"EmptyStackArray", getEmptyStackArray(), nil},
		{"OneElementLongStackArray", getOneElementLongStackArray(), 0},
		{"ManyElementsLongStackArray", getManyElementLongStackArray(), 14},
	}

	for _, v := range dataTable {
		t.Run(v.name, func(t *testing.T) {
			got := v.data.Pop()
			if got != v.want {
				t.Errorf("It expected %v, but got %v", v.want, got)
			}
		})
	}
}

func TestStackArray_Top(t *testing.T) {
	dataTable := []struct {
		name string
		data StackArray
		want interface{}
	}{
		{"EmptyStackArray", getEmptyStackArray(), nil},
		{"OneElementLongStackArray", getOneElementLongStackArray(), 0},
		{"ManyElementsLongStackArray", getManyElementLongStackArray(), 14},
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

func TestStackArray_Size(t *testing.T) {
	dataTable := []struct {
		name string
		data StackArray
		want uint16
	}{
		{"EmptyStackArray", getEmptyStackArray(), uint16(0)},
		{"OneElementLongStackArray", getOneElementLongStackArray(), uint16(1)},
		{"ManyElementsLongStackArray", getManyElementLongStackArray(), uint16(15)},
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

func TestStackArray_IsEmpty(t *testing.T) {
	dataTable := []struct {
		name string
		data StackArray
		want bool
	}{
		{"EmptyStackArray", getEmptyStackArray(), true},
		{"OneElementLongStackArray", getOneElementLongStackArray(), false},
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

func TestStackArray_IsFull(t *testing.T) {
	dataTable := []struct {
		name string
		data StackArray
		want bool
	}{
		{"EmptyStackArray", getEmptyStackArray(), false},
		{"OneElementLongStackArray", getOneElementLongStackArray(), false},
		{"FullStackArray", getManyElementLongStackArray(), true},
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
