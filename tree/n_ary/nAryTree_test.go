package n_ary

import (
	"data-structures-golang/list/doublylinkedlist"
	"testing"
)

func TestNewNAryTreeValue(t *testing.T) {
	want := 5
	subject := NewNAryTreeValue(want)

	got := subject.root.value

	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}
}

func TestNewAryTreeValueChildren(t *testing.T) {

	children := doublylinkedlist.NewDoublyLinkedListEmp()
	children.AddAtEnd(1)
	children.AddAtEnd(4)
	children.AddAtEnd(9)
	children.AddAtEnd(8)

	want := 5
	subject := NewAryTreeValueChildren(want, children)

	got := subject.root.value

	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}

	want = 4
	got = int(subject.root.children.Size())
	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}
}
