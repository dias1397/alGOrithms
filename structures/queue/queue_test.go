package queue

import (
	"testing"
)

func Test_queue(t *testing.T) {
	queue := Queue[int]{}

	got := queue.Size()
	if got != 0 {
		t.Errorf("Size want 0; got %d", got)
	}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	got, _ = queue.Dequeue()
	if got != 1 {
		t.Errorf("Dequeue want 1; got %d", got)
	}

	got = queue.Size()
	if got != 2 {
		t.Errorf("Size want 2; got %d", got)
	}

	got = queue.Peek()
	if got != 2 {
		t.Errorf("Peek want 2; got %d", got)
	}

	queue.Dequeue()
	queue.Dequeue()

	if !queue.IsEmpty() {
		t.Errorf("Stack should be empty")
	}
}
