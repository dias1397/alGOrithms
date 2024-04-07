package search

import (
	"testing"
)

func TestBinary(t *testing.T) {
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
			got := Binary(tc.items, tc.value)

			if got != tc.want {
				t.Errorf("Binary search(%v, %d) = %d: want %d", tc.items, tc.value, got, tc.want)
			}
		})
	}
}

func TestBinaryIteractive(t *testing.T) {
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
		{"Empty array", []int{}, 1, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BinaryIterative(tc.items, tc.value)

			if got != tc.want {
				t.Errorf("Iteractive binary search(%v, %d) = %d: want %d", tc.items, tc.value, got, tc.want)
			}
		})
	}
}

func BenchmarkBinary(b *testing.B) {
	testCases := []struct {
		name  string
		items []int
		value int
	}{
		{"Element at the beginning", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1},
		{"Element in the middle", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
		{"Element at the end", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10},
		{"Element not present", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 25},
		{"Empty array", []int{}, 1},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Binary(tc.items, tc.value)
			}
		})
	}
}

func BenchmarkBinaryIteractive(b *testing.B) {
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
				_ = BinaryIterative(tc.items, tc.value)
			}
		})
	}
}
