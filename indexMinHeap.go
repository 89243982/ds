package search

import (
	"math"
)

type IndexMinHeap struct {
	data    []int
	indexes []int
	reverse []int
}

func NewIndexMinHeap(cap int) *IndexMinHeap {
	heap := &IndexMinHeap{make([]int, cap+1), make([]int, 0, cap+1),
		make([]int, cap+1)}
	// heap.data = append(heap.data, math.MinInt32)

	heap.indexes = append(heap.indexes, math.MinInt32)
	// for i, _ := range heap.data {
	// 	heap.data[i] = math.MinInt32
	// 	// heap.reverse[i] = math.MinInt32
	// }
	// heap.reverse = append(heap.reverse, math.MinInt32)
	return heap
}

func (heap *IndexMinHeap) IsEmpty() bool {
	return len(heap.indexes) == 1
}

func (heap *IndexMinHeap) Size() int {
	return len(heap.indexes) - 1
}

func (heap *IndexMinHeap) Insert(index int, weight int) {
	index++
	heap.data[index] = weight
	heap.indexes = append(heap.indexes, index)
	heap.reverse[index] = len(heap.indexes) - 1
	heap.shiftUp(heap.Size())
}

func (heap *IndexMinHeap) ExtractMin() int {
	v := heap.data[heap.indexes[1]]
	heap.indexes[1], heap.indexes[heap.Size()] =
		heap.indexes[heap.Size()], heap.indexes[1]
	heap.reverse[heap.indexes[1]] = 1
	heap.reverse[heap.indexes[heap.Size()]] = 0
	heap.indexes = heap.indexes[:heap.Size()]
	heap.shiftDown(1)
	return v
}
func (heap *IndexMinHeap) ExtractMinIndex() int {
	ret := heap.indexes[1] - 1
	heap.indexes[1], heap.indexes[heap.Size()] =
		heap.indexes[heap.Size()], heap.indexes[1]
	heap.reverse[heap.indexes[1]] = 1
	heap.reverse[heap.indexes[heap.Size()]] = 0
	heap.indexes = heap.indexes[:heap.Size()]
	heap.shiftDown(1)
	return ret
}
func (heap *IndexMinHeap) Contain(i int) bool {
	i++
	if i >= len(heap.data) {
		return false
	}
	return heap.reverse[i] != 0
}
func (heap *IndexMinHeap) shiftUp(k int) {
	for k > 1 && heap.data[heap.indexes[k/2]] > heap.data[heap.indexes[k]] {
		heap.indexes[k], heap.indexes[k/2] = heap.indexes[k/2], heap.indexes[k]
		heap.reverse[heap.indexes[k]] = k
		heap.reverse[heap.indexes[k/2]] = k / 2
		k /= 2
	}
}
func (heap *IndexMinHeap) GetItem(index int) int {
	return heap.data[index+1]
}

func (heap *IndexMinHeap) GetMin() int {
	return heap.data[heap.indexes[1]]
}

func (heap *IndexMinHeap) shiftDown(k int) {
	for 2*k < len(heap.indexes) {
		j := 2 * k
		if j+1 < len(heap.indexes) &&
			heap.data[heap.indexes[j+1]] < heap.data[heap.indexes[j]] {
			j++
		}
		if heap.data[heap.indexes[k]] <= heap.data[heap.indexes[j]] {
			break
		}
		heap.indexes[k], heap.indexes[j] = heap.indexes[j], heap.indexes[k]
		heap.reverse[heap.indexes[k]] = k
		heap.reverse[heap.indexes[j]] = j
		k = j
	}
}

func (heap *IndexMinHeap) Change(index int, weight int) {

	index++
	heap.data[index] = weight
	heap.shiftUp(heap.reverse[index])

	heap.shiftDown(heap.reverse[index])
}
