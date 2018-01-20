package search

type LazyPrimMST struct {
	g         Graph
	pq        MiniHeap
	marked    []bool
	mst       []Edge
	mstWeight int
}

func NewLazyPrimMST(g Graph) *LazyPrimMST {
	lpm := &LazyPrimMST{g: g, pq: *NewMiniHeap(g.E()),
		marked: make([]bool, g.V()), mst: make([]Edge, 0, g.V()-1)}
	lpm.visit(0)
	for !lpm.pq.IsEmpty() {
		e := lpm.pq.ExtractMin()
		if lpm.marked[e.V()] == lpm.marked[e.W()] {
			continue
		}
		lpm.mst = append(lpm.mst, e)
		if lpm.marked[e.V()] {
			lpm.visit(e.W())
		} else {
			lpm.visit(e.V())
		}
	}
	for _, v := range lpm.mst {
		lpm.mstWeight += v.Wt()
	}
	return lpm
}
func (lpm *LazyPrimMST) visit(v int) {
	lpm.marked[v] = true
	it := lpm.g.Iterator(v)
	for e := it.Begin(); !it.End(); e = it.Next() {
		if !lpm.marked[e.Other(v)] {
			lpm.pq.Insert(*e)
		}
	}
}

func (lpm *LazyPrimMST) Result() int {
	return lpm.mstWeight
}

func (lpm *LazyPrimMST) MstEdges() []Edge {
	return lpm.mst
}
