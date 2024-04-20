package stack

import (
	"testing"
)

func Test_stack(t *testing.T) {
	stack := Stack[int]{}

	got := stack.Size()
	if got != 0 {
		t.Errorf("Size want 0; got %d", got)
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	got, _ = stack.Pop()
	if got != 3 {
		t.Errorf("Pop want 3; got %d", got)
	}

	got = stack.Size()
	if got != 2 {
		t.Errorf("Size want 2; got %d", got)
	}

	got = stack.Peek()
	if got != 2 {
		t.Errorf("Peek want 2; got %d", got)
	}

	stack.Pop()
	stack.Pop()

	if !stack.IsEmpty() {
		t.Errorf("Stack should be empty")
	}
}
