package search

type MiniHeap struct {
	data []Edge
}

func NewMiniHeap(cap int) *MiniHeap {
	heap := &MiniHeap{data: make([]Edge, 0, cap+1)}
	heap.data = append(heap.data, *NewEdge(0, 0, 0))
	return heap
}
func (heap *MiniHeap) ShiftUp(k int) {
	for k > 1 && heap.data[k/2].Greater(heap.data[k]) {
		heap.data[k], heap.data[k/2] = heap.data[k/2], heap.data[k]
		k /= 2
	}
}

func (heap *MiniHeap) ShiftDown(k int) {
	l := len(heap.data)
	for 2*k < l {
		j := 2 * k
		if j+1 < l && heap.data[j+1].Less(heap.data[j]) {
			j++
		}
		if heap.data[k].LE(heap.data[j]) {
			break
		}
		heap.data[k], heap.data[j] = heap.data[j], heap.data[k]
		k = j
	}
}

func (heap *MiniHeap) Insert(item Edge) {
	heap.data = append(heap.data, item)
	heap.ShiftUp(len(heap.data) - 1)
}

func (heap *MiniHeap) ExtractMin() Edge {
	e := heap.data[1]
	heap.data[1], heap.data[len(heap.data)-1] =
		heap.data[len(heap.data)-1], heap.data[1]
	heap.data = heap.data[:len(heap.data)-1]
	heap.ShiftDown(1)
	return e
}

func (heap *MiniHeap) GetMin() Edge {
	return heap.data[1]
}

func (heap *MiniHeap) IsEmpty() bool {
	return len(heap.data) == 1
}

func (heap *MiniHeap) Size() int {
	return len(heap.data) - 1
}
