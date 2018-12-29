package depthFirstSearch

import (
	"errors"
	"github.com/linehk/GoAlgorithms/DataStructures/graph/adjacencyMatrix"
)

type dfs struct {
	marked []bool
	count  int
}

func DepthFirstSearch(g adjacencyMatrix.Graph, s int) error {
	d := new(dfs)
	d.marked = make([]bool, g.V())
	d.count = 0
	if err := d.dfs(g, s); err != nil {
		return err
	}
	return nil
}

func (d *dfs) dfs(g adjacencyMatrix.Graph, v int) error {
	d.count++
	d.marked[v] = true
	adj, err := g.Adj(v)
	if err != nil {
		return err
	}
	for _, w := range adj {
		if !d.marked[w] {
			if err := d.dfs(g, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (d dfs) Marked(v int) (bool, error) {
	if err := d.validateVertex(v); err != nil {
		return false, err
	}
	return d.marked[v], nil
}

func (d dfs) Count() int {
	return d.count
}

func (d dfs) validateVertex(v int) error {
	V := len(d.marked)
	if v < 0 || v >= V {
		return errors.New("illegal vertex")
	}
	return nil
}

func NonrecursiveDFS(g adjacencyMatrix.Graph, s int) {
}
