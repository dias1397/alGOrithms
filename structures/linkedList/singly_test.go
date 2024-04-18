package linkedlist

import (
	"testing"
)

func TestSingly(t *testing.T) {
	list := Singly[int]{}

	list.AddAt(1, 0)
	list.Display()
	list.Add(3)
	list.Display()
	got, error := list.Get(0)

	if got.value != 1 {
		t.Errorf("Unable to get element at index 0: %v(%v)", got, error)
	}

	list.AddAt(2, 1)
	list.Display()
	got, error = list.Get(1)

	if got.value != 2 {
		t.Errorf("Unable to get element at index 1: %v(%v)", got, error)
	}
}
