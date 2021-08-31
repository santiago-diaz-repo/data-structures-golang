package binarySearchTree

import (
	"testing"
)

func TestBST_AddIterative(t *testing.T) {
	subject := NewBST()

	want := true
	got := subject.AddIterative(5)

	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}

	subject.AddIterative(2)
	subject.AddIterative(8)

	want2 := 3
	got2 := subject.Size()

	if got2 != want2 {
		t.Errorf("It expected %v, but got %v", want2, got2)
	}
}

func TestBST_AddRecursive(t *testing.T) {
	subject := NewBST()

	want := true
	got := subject.AddRecursive(5)

	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}

	subject.AddRecursive(2)
	subject.AddRecursive(8)

	want2 := 3
	got2 := subject.Size()

	if got2 != want2 {
		t.Errorf("It expected %v, but got %v", want2, got2)
	}
}

func TestBST_SearchIterative(t *testing.T) {
	subject := NewBST()
	subject.AddIterative(9)
	subject.AddIterative(4)
	subject.AddIterative(15)
	subject.AddIterative(3)
	subject.AddIterative(10)
	want := true
	got := subject.SearchIterative(3)

	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}

	want = false
	got = subject.SearchIterative(90)

	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}
}

func TestBST_SearchRecursive(t *testing.T) {
	subject := NewBST()
	subject.AddRecursive(9)
	subject.AddRecursive(4)
	subject.AddRecursive(15)
	subject.AddRecursive(3)
	subject.AddRecursive(10)
	want := true
	got := subject.SearchRecursive(3)

	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}

	want = false
	got = subject.SearchRecursive(90)

	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}
}

func TestBST_Delete(t *testing.T) {
	subject := NewBST()

	want := false
	got := subject.Delete(5)
	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}

	subject.AddRecursive(9)
	want = true
	got = subject.Delete(9)
	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}

	wantSize := 0
	gotSize := subject.Size()
	if gotSize != wantSize {
		t.Errorf("It expected %v, but got %v", want, got)
	}

	subject.AddRecursive(9)
	subject.AddRecursive(4)
	subject.AddRecursive(15)
	subject.AddRecursive(11)
	subject.AddRecursive(5)
	subject.AddRecursive(3)
	subject.AddRecursive(18)

	want = false
	got = subject.Delete(90)
	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}

	want = true
	got = subject.Delete(5)
	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}
	wantSize = 6
	gotSize = subject.Size()
	if gotSize != wantSize {
		t.Errorf("It expected %v, but got %v", want, got)
	}

	want = true
	got = subject.Delete(4)
	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}
	wantSize = 5
	gotSize = subject.Size()
	if gotSize != wantSize {
		t.Errorf("It expected %v, but got %v", want, got)
	}

	want = true
	got = subject.Delete(15)
	if got != want {
		t.Errorf("It expected %v, but got %v", want, got)
	}
	wantSize = 4
	gotSize = subject.Size()
	if gotSize != wantSize {
		t.Errorf("It expected %v, but got %v", want, got)
	}
}

func TestBST_PreOrderTraversal(t *testing.T) {
	subject := NewBST()
	subject.AddRecursive(9)
	subject.AddRecursive(4)
	subject.AddRecursive(15)
	subject.AddRecursive(11)
	subject.AddRecursive(5)
	subject.AddRecursive(3)
	subject.AddRecursive(18)

	want := []int{9, 4, 3, 5, 15, 11, 18}
	got := subject.PreOrderTraversal()

	for i := 0; i < len(*got); i++ {
		if (*got)[i] != want[i] {
			t.Errorf("It expected %v, but got %v", want, got)
		}
	}
}

func TestBST_InOrderTraversal(t *testing.T) {
	subject := NewBST()
	subject.AddRecursive(9)
	subject.AddRecursive(4)
	subject.AddRecursive(15)
	subject.AddRecursive(11)
	subject.AddRecursive(5)
	subject.AddRecursive(3)
	subject.AddRecursive(18)

	want := []int{3, 4, 5, 9, 11, 15, 18}
	got := subject.InOrderTraversal()

	for i := 0; i < len(*got); i++ {
		if (*got)[i] != want[i] {
			t.Errorf("It expected %v, but got %v", want, got)
		}
	}
}

func TestBST_PostOrderTraversal(t *testing.T) {
	subject := NewBST()
	subject.AddRecursive(9)
	subject.AddRecursive(4)
	subject.AddRecursive(15)
	subject.AddRecursive(11)
	subject.AddRecursive(5)
	subject.AddRecursive(3)
	subject.AddRecursive(18)

	want := []int{3, 5, 4, 11, 18, 15, 9}
	got := subject.PostOrderTraversal()

	for i := 0; i < len(*got); i++ {
		if (*got)[i] != want[i] {
			t.Errorf("It expected %v, but got %v", want, got)
		}
	}
}
