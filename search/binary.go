package search

func Binary(items []int, value int) int {
	return binary(items, value, 0, len(items)-1)
}

func binary(items []int, value, low, high int) int {
	if low > high {
		return -1
	}

	mid := int(low + ((high - low) / 2))
	if items[mid] == value {
		return mid
	} else if items[mid] > value {
		return binary(items, value, low, mid-1)
	} else {
		return binary(items, value, mid+1, high)
	}
}

func BinaryIterative(items []int, value int) int {
	low := 0
	high := len(items) - 1

	for low <= high {
		mid := int(low + ((high - low) / 2))

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
