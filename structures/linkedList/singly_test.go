package linkedlist

import (
	"testing"
)

func TestSingly(t *testing.T) {
	list := Singly[int]{}

	list.Add(1)
	list.Add(2)
	list.Add(3)

	if list.Size() != 3 {
		t.Errorf("Expected size: 3, got: %d", list.Size())
	}

	position0, _ := list.Get(0)
	position1, _ := list.Get(1)
	position2, _ := list.Get(2)
	if position0 != 1 || position1 != 2 || position2 != 3 {
		t.Errorf("Expected values: 1 2 3, got: %d %d %d", position0, position1, position2)
	}

	list.AddAt(4, 1)

	position1, _ = list.Get(1)
	if list.Size() != 4 || position1 != 4 {
		t.Errorf("Expected size: 4, got: %d\nExpected value at index 1: 4, got: %d", list.Size(), position1)
	}

	list.Remove()

	position3, _ := list.Get(3)
	if list.Size() != 3 || position3 != 0 {
		t.Errorf("Expected size: 3, got: %d\nExpected value at index 3: nil, got: %d", list.Size(), position3)
	}

	list.RemoveAt(1)

	position1, _ = list.Get(1)
	if list.Size() != 2 || position1 != 2 {
		t.Errorf("Expected size: 2, got: %d\nExpected value at index 1: 2, got: %d", list.Size(), position1)
	}
}
