package sort

func HeapSort(datas []int) {
	// sortedDatas := make([]int, len(datas))
	heap := NewMaxHeapByHeapify(datas)
	for i := len(datas) - 1; i >= 0; i-- {
		datas[i] = heap.Extract()
	}
}

func HeapSortByLocal(datas []int) {
	//heapify，datas[0]存在数据
	for i := ((len(datas) - 1) - 1) / 2; i >= 0; i-- {
		shiftDown(datas, i)
	}

	for len(datas) > 1 {
		datas[0], datas[len(datas)-1] = datas[len(datas)-1], datas[0]
		datas = datas[:len(datas)-1]
		shiftDown(datas, 0)
	}
}

func shiftDown(datas []int, k int) {
	for 2*k+1 < len(datas) {
		j := 2*k + 1
		if j+1 < len(datas) && datas[j] < datas[j+1] {
			j++
		}
		if datas[j] <= datas[k] {
			break
		}
		datas[k], datas[j] = datas[j], datas[k]
		k = j
	}
}
