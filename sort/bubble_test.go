package sort

import (
	"slices"
	"testing"
)

func TestBubble(t *testing.T) {
	testCases := []struct {
		name  string
		items []int
		want  []int
	}{
		{"Sorted unsigned", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"Reversed unsigned", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"Sorted signed", []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}, []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
		{"Reversed signed", []int{5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5}, []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
		{"Random signed", []int{3, 5, -1, -2, 1, 0, 2, -5, -3, 4, -4}, []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
		{"Singleton", []int{1}, []int{1}},
		{"Empty array", []int{}, []int{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Bubble(tc.items)

			if !slices.Equal(got, tc.want) {
				t.Errorf("Bubble sort(%v) = %v: want %v", tc.items, got, tc.want)
			}
		})
	}
}

func BenchmarkBubble(b *testing.B) {
	testCases := []struct {
		name  string
		items []int
		want  []int
	}{
		{"Sorted unsigned", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"Reversed unsigned", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"Sorted signed", []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}, []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
		{"Reversed signed", []int{5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5}, []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
		{"Random signed", []int{3, 5, -1, -2, 1, 0, 2, -5, -3, 4, -4}, []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
		{"Singleton", []int{1}, []int{1}},
		{"Empty array", []int{}, []int{}},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Bubble(tc.items)
			}
		})
	}
}
