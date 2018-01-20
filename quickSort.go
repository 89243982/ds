package sort

import (
	"math/rand"
	"time"
)

func QuickSort(datas []int) {
	rand.Seed(time.Now().Unix())
	quickSort(datas, 0, len(datas)-1)
}

func quickSort(datas []int, l int, r int) {
	if r-l <= 15 {
		InsertSortSection(datas, l, r)
		return
	}
	/**
	p := partition2(datas, l, r)
	quickSort(datas, l, p-1)
	quickSort(datas, p+1, r)
	**/
	lt, gt := partition3Ways(datas, l, r)
	quickSort(datas, l, lt)
	quickSort(datas, gt, r)
}

func partition(datas []int, l, r int) int {

	swapIndex := rand.Int()%(r-l+1) + l
	datas[l], datas[swapIndex] = datas[swapIndex], datas[l]
	e := datas[l]
	/**datas[l+1,j]<e<datas[j+1,i)**/
	j := l
	for i := l + 1; i <= r; i++ {
		if datas[i] < e {
			datas[j+1], datas[i] = datas[i], datas[j+1]
			j++
		}
	}
	datas[l], datas[j] = datas[j], datas[l]
	return j
}
func partition2(datas []int, l, r int) int {
	swapIndex := rand.Int()%(r-l+1) + l
	datas[l], datas[swapIndex] = datas[swapIndex], datas[l]
	e := datas[l]
	/**
	datas[l+1,i)<=e<=datas(j,r]
	**/
	i := l + 1
	j := r
	for {
		for i <= r && datas[i] < e {
			i++
		}
		for j >= l+1 && datas[j] > e {
			j--
		}
		if i > j {
			break
		}
		datas[i], datas[j] = datas[j], datas[i]
		i++
		j--
	}
	datas[l], datas[j] = datas[j], datas[l]
	return j
}

func partition3Ways(datas []int, l, r int) (int, int) {
	swapIndex := rand.Int()%(r-l+1) + l
	datas[l], datas[swapIndex] = datas[swapIndex], datas[l]
	e := datas[l]
	/**
	datas[l+1,lt]<e   datas[lt+1,i)=e datas[gt,r]>e
	**/
	lt := l
	i := lt + 1
	gt := r + 1
	for i < gt {
		if datas[i] == e {
			i++
		} else if datas[i] < e {
			datas[lt+1], datas[i] = datas[i], datas[lt+1]
			lt++
			i++
		} else if datas[i] > e {
			datas[gt-1], datas[i] = datas[i], datas[gt-1]
			gt--
		}
	}
	datas[l], datas[lt] = datas[lt], datas[l]
	lt--
	return lt, gt
}
