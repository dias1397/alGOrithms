package search

import (
	"testing"
)

func TestLinear(t *testing.T) {
	testCases := []struct {
		name  string
		items []int
		value int
		want  int
	}{
		{"Element at the beginning", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0},
		{"Element in the middle", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4},
		{"Element at the end", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9},
		{"Element not present", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 25, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Linear(tc.items, tc.value)

			if got != tc.want {
				t.Errorf("Linear search(%v, %d) = %d: want %d", tc.items, tc.value, got, tc.want)
			}
		})
	}
}

func BenchmarkLinear(b *testing.B) {
	testCases := []struct {
		name  string
		items []int
		value int
	}{
		{"Element at the beginning", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1},
		{"Element in the middle", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
		{"Element at the end", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10},
		{"Element not present", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 25},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Linear(tc.items, tc.value)
			}
		})
	}
}
