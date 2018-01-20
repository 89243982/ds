package search

import "fmt"

type DenseGraph struct {
	n, m     int
	directed bool
	g        []([]*Edge)
}

func NewDenseGraph(n int, directed bool) *DenseGraph {
	dg := &DenseGraph{n: n, directed: directed, g: make([]([]*Edge), n)}
	for i, _ := range dg.g {
		dg.g[i] = make([]*Edge, n)
	}
	return dg
}

func (dg *DenseGraph) V() int {
	return dg.n
}

func (dg *DenseGraph) E() int {
	return dg.m
}

func (dg *DenseGraph) AddEdge(v int, w int, weight int) {
	if dg.HasEdge(v, w) {
		dg.m--
	}
	dg.g[v][w] = &Edge{a: v, b: w, weight: weight}
	if v != w && !dg.directed {
		dg.g[w][v] = &Edge{a: w, b: v, weight: weight}
	}
	dg.m++
}

func (dg *DenseGraph) HasEdge(v int, w int) bool {
	return dg.g[v][w] != nil
}

func (dg *DenseGraph) Show() {
	for _, vvs := range dg.g {
		for _, v := range vvs {
			if v != nil {
				fmt.Printf("%d\t", v.Wt())
			} else {
				fmt.Printf("%s\t", "Nil")
			}
		}
		fmt.Println()
	}
}

func (dg *DenseGraph) Iterator(v int) Iterator {
	return NewDenseGraphAdjIterator(dg, v)
}

type DenseGraphAdjIterator struct {
	dg    *DenseGraph
	v     int
	vvs   []*Edge
	index int
}

func NewDenseGraphAdjIterator(dg *DenseGraph, v int) *DenseGraphAdjIterator {
	return &DenseGraphAdjIterator{dg: dg, v: v, vvs: dg.g[v]}
}

func (it *DenseGraphAdjIterator) Begin() *Edge {
	it.index = -1
	return it.Next()
}

func (it *DenseGraphAdjIterator) Next() *Edge {
	for it.index += 1; it.index < it.dg.V(); it.index++ {
		if it.vvs[it.index] != nil {
			return it.vvs[it.index]
		}
	}
	return nil
}

func (it *DenseGraphAdjIterator) End() bool {
	return it.index >= it.dg.V()
}
