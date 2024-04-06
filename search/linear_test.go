package search

import (
	"testing"
)

func TestLinear_ElementPresent(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	value := 3

	actual := Linear(items, value)
	if actual != 2 {
		t.Errorf("Linear search(%v, %d) = %d: want 2", items, value, actual)
	}
}

func TestLinear_ElementNotPresent(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	value := 6

	actual := Linear(items, value)
	if actual != -1 {
		t.Errorf("Linear search(%v, %d) = %d: want -1", items, value, actual)
	}
}

func BenchmarkLinear(b *testing.B) {
	items := []int{1, 2, 3, 4, 5}
	value := 3

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Linear(items, value)
	}
}