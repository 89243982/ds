package sort

import (
	"math"
)

type MaxHeap struct {
	datas []int
}

func NewMaxHeap(cap int) *MaxHeap {
	heap := &MaxHeap{make([]int, 0, cap)}
	heap.datas = append(heap.datas, math.MinInt32)
	return heap
}

func NewMaxHeapByHeapify(datas []int) *MaxHeap {
	heap := &MaxHeap{make([]int, len(datas)+1)}
	heap.datas[0] = math.MinInt32
	copy(heap.datas[1:], datas)
	for k := heap.Size() / 2; k > 0; k-- {
		heap.shiftDown(k)
	}
	return heap
}

func (heap *MaxHeap) Insert(v int) {
	heap.datas = append(heap.datas, v)
	heap.shiftUp(len(heap.datas) - 1)
}

func (heap *MaxHeap) Extract() int {
	e := heap.datas[1]
	heap.datas[1], heap.datas[len(heap.datas)-1] =
		heap.datas[len(heap.datas)-1], heap.datas[1]
	heap.datas = heap.datas[:len(heap.datas)-1]
	heap.shiftDown(1)
	return e
}

func (heap *MaxHeap) shiftUp(k int) {
	for k > 1 && heap.datas[k/2] < heap.datas[k] {
		heap.datas[k], heap.datas[k/2] = heap.datas[k/2], heap.datas[k]
		k /= 2
	}
}
func (heap *MaxHeap) shiftDown(k int) {
	for 2*k < len(heap.datas) {
		j := 2 * k
		if j+1 < len(heap.datas) && heap.datas[j] < heap.datas[j+1] {
			j++
		}
		if heap.datas[j] <= heap.datas[k] {
			break
		}
		heap.datas[k], heap.datas[j] = heap.datas[j], heap.datas[k]
		k = j
	}
}

func (heap *MaxHeap) Empty() bool {
	return len(heap.datas) == 1
}

func (heap *MaxHeap) Size() int {
	return len(heap.datas) - 1
}

// func (heap *MaxHeap) Println() {
// 	fmt.Println(heap.datas)
// }
