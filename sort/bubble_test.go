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
		{"Unordered array", []int{1, 3, 5, 4, 2}, []int{1, 2, 3, 4, 5}},
		//{"Empty array", []int{}, []int{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Bubble(tc.items)

			if slices.Equal(got, tc.want) {
				t.Errorf("Bubble sort(%v) = %v: want %v", tc.items, got, tc.want)
			}
		})
	}
}
