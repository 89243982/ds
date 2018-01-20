package sort

func SelectionSort(arr Sortable) {
	n := arr.Len()
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr.Less(j, minIndex) {
				minIndex = j
			}
		}
		arr.Swap(i, minIndex)
	}
}
