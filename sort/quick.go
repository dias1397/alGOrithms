package sort

func Quick(items []int) []int {
	qs(items, 0, len(items)-1)
	return items
}

func qs(items []int, lo int, hi int) {
	if lo >= hi {
		return
	}

	pivotIndex := partition(items, lo, hi)
	qs(items, lo, pivotIndex-1)
	qs(items, pivotIndex+1, hi)
}

func partition(items []int, lo int, hi int) int {
	pivot := items[hi]

	index := lo - 1

	for i := lo; i < hi; i++ {
		if items[i] <= pivot {
			index++
			items[index], items[i] = items[i], items[index]
		}
	}

	index++
	items[index], items[hi] = items[hi], items[index]

	return index
}
