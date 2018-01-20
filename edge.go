package search

type Edge struct {
	a, b   int
	weight int
}

func NewEdge(a, b int, weight int) *Edge {
	return &Edge{a: a, b: b, weight: weight}
}

func (e *Edge) V() int {
	return e.a
}

func (e *Edge) W() int {
	return e.b
}

func (e *Edge) Wt() int {
	return e.weight
}

func (e *Edge) Other(x int) int {
	if x == e.a {
		return e.b
	} else {
		return e.a
	}
}

func (e Edge) Less(other Edge) bool {
	return e.Wt() < other.Wt()
}

func (e Edge) LE(other Edge) bool {
	return e.Wt() <= other.Wt()
}
func (e Edge) Greater(other Edge) bool {
	return e.Wt() > other.Wt()
}
func (e Edge) GE(other Edge) bool {
	return e.Wt() >= other.Wt()
}

func (e Edge) Equal(other Edge) bool {
	return e.Wt() == other.Wt()
}
