package search

func Linear(items []int, value int) int {
	for index, item := range items {
		if item == value {
			return index
		}
	}

	return -1
}
