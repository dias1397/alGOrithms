package sort

func Bubble(items []int) []int {
	for i := 0; i < len(items); i++ {
		for j := 0; j < len(items)-1-i; j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}

	return items
}
