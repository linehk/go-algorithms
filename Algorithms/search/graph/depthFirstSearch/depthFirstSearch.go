package depthFirstSearch

import (
	"github.com/linehk/GoAlgorithms/DataStructures/graph"
)

type dfs struct {
	marked []bool
	count  int
}

func (d *dfs) DepthFirstSearch(g graph.Graph, s int) {
	d.marked = make([]bool, g.V())
	d.dfs(g, s)
}

func (d *dfs) dfs(g graph.Graph, v int) {
	d.count++
	d.marked[v] = true
	for _, w := range g.Adj(v) {
		if !d.marked[w] {
			d.dfs(g, w)
		}
	}
}

func (d dfs) Marked(v int) bool {
	return d.marked[v]
}

func (d dfs) Count() int {
	return d.count
}
