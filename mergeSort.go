package sort

import "math"

func MergeSort(datas []int) {
	mergeSort(datas, 0, len(datas)-1)
}

/**
  递归对[l...r]区间进行归并排序
**/
func mergeSort(datas []int, l int, r int) {
	if r-l <= 15 {
		InsertSortSection(datas, l, r)
		return
	}
	mid := (l + r) / 2
	mergeSort(datas, l, mid)
	mergeSort(datas, mid+1, r)
	if datas[mid] > datas[mid+1] {
		merge(datas, l, mid, r)
	}
}

func merge(datas []int, l int, m int, r int) {
	arr := make([]int, r-l+1)
	copy(arr, datas[l:r+1])
	for i, j, k := l, m+1, l; k <= r; k++ {
		if i >= m+1 {
			datas[k] = arr[j-l]
			j++
		} else if j > r {
			datas[k] = arr[i-l]
			i++
		} else if arr[i-l] <= arr[j-l] {
			datas[k] = arr[i-l]
			i++
		} else {
			datas[k] = arr[j-l]
			j++
		}
	}
}

func MergeSortBU(datas []int) {
	for sz := 1; sz < len(datas); sz += sz {
		for i := 0; i+sz < len(datas); i += sz + sz {
			if datas[i+sz-1] > datas[i+sz] {
				ends := int(math.Min(float64(i+sz+sz-1), float64(len(datas)-1)))
				if sz+sz <= 16 {
					InsertSortSection(datas, i, ends)
				} else {
					merge(datas, i, i+sz-1, ends)
				}
			}
		}
	}
}
