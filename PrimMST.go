package search

type PrimMST struct {
	g         Graph
	ipq       IndexMinHeap
	marked    []bool
	mst       []Edge
	edgeTo    []*Edge
	mstWeight int
}

func NewPrimMST(g Graph) *PrimMST {
	p := &PrimMST{g, *NewIndexMinHeap(g.V()), make([]bool, g.V()),
		make([]Edge, 0, g.V()-1), make([]*Edge, g.V()), 0}
	p.visit(0)
	for !p.ipq.IsEmpty() {
		v := p.ipq.ExtractMinIndex()
		p.mst = append(p.mst, *p.edgeTo[v])
		p.visit(v)
	}

	for _, v := range p.mst {
		p.mstWeight += v.Wt()
	}
	return p
}
func (p *PrimMST) visit(v int) {
	p.marked[v] = true
	it := p.g.Iterator(v)
	for e := it.Begin(); !it.End(); e = it.Next() {
		w := e.Other(v)
		if !p.marked[w] {
			if p.edgeTo[w] == nil {
				p.edgeTo[w] = e
				p.ipq.Insert(w, e.Wt())
			} else if e.Wt() < p.edgeTo[w].Wt() {
				p.edgeTo[w] = e
				p.ipq.Change(w, e.Wt())
			}
		}
	}
}

func (p *PrimMST) Result() int {
	return p.mstWeight
}

func (p *PrimMST) MstEdges() []Edge {
	return p.mst
}
