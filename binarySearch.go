package search

func BinarySearch(datas []int, target int) int {
	l, r := 0, len(datas)-1
	//在datas[l,r]这个区间查找target
	for l <= r {
		mid := l + (r-l)/2
		if datas[mid] == target {
			return mid
		} else if target < datas[mid] {
			r = mid - 1
		} else if datas[mid] < target {
			l = mid + 1
		}
	}
	return -1
}
