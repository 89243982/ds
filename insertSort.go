package sort

func InsertSort(arr Sortable) {
	for i := 1; i < arr.Len(); i++ {
		for j := i; j > 0 && arr.Less(j, j-1); j-- {
			arr.Swap(j, j-1)
		}
	}
}

func InsertSortByOptim(arr Sortable) {
	sia := arr.(SortAndIndexAble)
	for i := 1; i < sia.Len(); i++ {
		e := sia.Get(i)
		j := i
		for ; j > 0 && sia.Get(j-1) > e; j-- {
			sia.Set(j, sia.Get(j-1))
		}
		sia.Set(j, e)
	}

}

/**对区间arr[l,r]的数据进行插入排序**/
func InsertSortSection(arr []int, l, r int) {
	for i := l + 1; i <= r; i++ {
		e := arr[i]
		j := i
		for ; j > l && arr[j-1] > e; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}
