package search

type Iterator interface {
	Begin() *Edge
	Next() *Edge
	End() bool
}

//
// type Component struct {
// 	g       Graph
// 	visited []bool
// 	ccount  int
// 	ids     []int
// }
//
// func (c *Component) Count() int {
// 	return c.ccount
// }
// func NewComponent(g Graph) *Component {
// 	c := &Component{g: g, visited: make([]bool, g.V()), ids: make([]int, g.V())}
// 	for i, _ := range c.ids {
// 		c.ids[i] = -1
// 	}
// 	for i := 0; i < g.V(); i++ {
// 		if !c.visited[i] {
// 			c.DFS(i)
// 			c.ccount++
// 		}
// 	}
// 	return c
// }
//
// func (c *Component) DFS(v int) {
// 	c.visited[v] = true
// 	c.ids[v] = c.ccount
// 	it := c.g.Iterator(v)
// 	for w := it.Begin(); !it.End(); w = it.Next() {
// 		if !c.visited[w] {
// 			c.DFS(w)
// 		}
// 	}
// }
//
// func (c *Component) IsConnected(v, w int) bool {
// 	return c.ids[v] == c.ids[w]
// }
