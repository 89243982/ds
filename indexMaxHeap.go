package sort

import (
	"fmt"
	"math"
)

type IndexMaxHeap struct {
	/**real data**/
	datas []int
	/**
	  binary heap of index for data
	  **/
	indexs []int
	revs   []int
}

func NewIndexMaxHeap(cap int) *IndexMaxHeap {
	heap := &IndexMaxHeap{make([]int, 0, cap+1), make([]int, 0, cap+1),
		make([]int, 0, cap+1)}
	heap.datas = append(heap.datas, math.MinInt32)
	heap.indexs = append(heap.indexs, math.MinInt32)
	heap.revs = append(heap.revs, math.MinInt32)
	return heap
}

func (heap *IndexMaxHeap) Size() int {
	return len(heap.indexs) - 1
}

func (heap *IndexMaxHeap) Insert(v int) {
	heap.datas = append(heap.datas, v)
	heap.indexs = append(heap.indexs, len(heap.datas)-1)
	heap.revs = append(heap.revs, len(heap.indexs)-1)
	heap.shiftUp(heap.Size())
}

func (heap *IndexMaxHeap) Extract() int {
	v := heap.datas[heap.indexs[1]]
	heap.indexs[1], heap.indexs[heap.Size()] =
		heap.indexs[heap.Size()], heap.indexs[1]
	heap.revs[heap.indexs[1]] = 1
	heap.revs[heap.indexs[heap.Size()]] = 0
	heap.indexs = heap.indexs[:heap.Size()]
	heap.shiftDown(1)
	return v
}

func (heap *IndexMaxHeap) shiftUp(k int) {
	for k > 1 && heap.datas[heap.indexs[k/2]] < heap.datas[heap.indexs[k]] {
		heap.indexs[k], heap.indexs[k/2] = heap.indexs[k/2], heap.indexs[k]
		heap.revs[heap.indexs[k]] = k
		heap.revs[heap.indexs[k/2]] = k / 2
		k /= 2
	}
}

func (heap *IndexMaxHeap) shiftDown(k int) {
	for 2*k < len(heap.indexs) {
		j := 2 * k
		if j+1 < len(heap.indexs) &&
			heap.datas[heap.indexs[j]] < heap.datas[heap.indexs[j+1]] {
			j++
		}
		if heap.datas[heap.indexs[j]] <= heap.datas[heap.indexs[k]] {
			break
		}
		heap.indexs[k], heap.indexs[j] = heap.indexs[j], heap.indexs[k]
		heap.revs[heap.indexs[k]] = k
		heap.revs[heap.indexs[j]] = j
		k = j
	}
}

func (heap *IndexMaxHeap) Empty() bool {
	return len(heap.indexs) == 1
}

func (heap *IndexMaxHeap) Contain(i int) bool {
	i++
	if i >= len(heap.datas) {
		return false
	}
	return heap.revs[i] != 0
}

func (heap *IndexMaxHeap) Modify(i int, v int) {
	i++
	heap.datas[i] = v
	// j := heap.revs[i]
	heap.shiftUp(heap.revs[i])
	heap.shiftDown(heap.revs[i])
}

func (heap *IndexMaxHeap) TestPrintln() {
	fmt.Printf("datas:%v\n", heap.datas)
	fmt.Printf("indexs:%v\n", heap.indexs)
	fmt.Printf("revs:%v\n", heap.revs)
}
