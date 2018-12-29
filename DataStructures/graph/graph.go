package graph

import (
	"errors"
)

type Graph struct {
	v   int
	e   int
	adj [][]int
}

func New(v int) *Graph {
	g := new(Graph)
	g.v = v
	g.e = 0
	g.adj = make([][]int, g.v)
	return g
}

func (g Graph) V() int {
	return g.v
}

func (g Graph) E() int {
	return g.e
}

func (g Graph) validateVertex(v int) error {
	if v < 0 || v >= g.v {
		return errors.New("illegal v")
	}
	return nil
}

func (g *Graph) AddEdge(v, w int) {
	g.e++
	g.adj[v][w] = 1
	g.adj[w][v] = 1
}

func (g Graph) Degree(v int) int {
	return len(g.adj[v])
}

func (g Graph) Adj(v int) []int {
	return g.adj[v]
}
