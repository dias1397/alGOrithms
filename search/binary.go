package search

func Binary(items []int, value int) int {
	low := 0
	high := len(items) - 1

	for low <= high {
		mid := int(low + (high-low)/2)

		if items[mid] == value {
			return mid
		} else if items[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
