package search

import (
	"testing"
)

func TestBinary_ElementPresent(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	value := 8

	actual := Binary(items, value)
	if actual != 7 {
		t.Errorf("Binary search(%v, %d) = %d: want 2", items, value, actual)
	}
}

func TestBinary_ElementNotPresent(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	value := 25

	actual := Binary(items, value)
	if actual != -1 {
		t.Errorf("Linear search(%v, %d) = %d: want -1", items, value, actual)
	}
}

func BenchmarkBinary(b *testing.B) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	value := 8

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Binary(items, value)
	}
}
