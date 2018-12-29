package adjacencyMatrix

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

func (g *Graph) AddEdge(v, w int) error {
	if err := g.validateVertex(v); err != nil {
		return err
	}
	if err := g.validateVertex(w); err != nil {
		return err
	}

	g.e++
	g.adj[v][w] = 1
	g.adj[w][v] = 1

	return nil
}

func (g Graph) Degree(v int) (int, error) {
	if err := g.validateVertex(v); err != nil {
		return 0, err
	}
	return len(g.adj[v]), nil
}

func (g Graph) Adj(v int) ([]int, error) {
	if err := g.validateVertex(v); err != nil {
		return nil, err
	}
	return g.adj[v], nil
}

type Digraph struct {
	v        int
	e        int
	adj      [][]int
	inDegree []int
}

func NewDigraph(v int) *Digraph {
	d := new(Digraph)
	d.v = v
	d.e = 0
	d.inDegree = make([]int, d.v)
	d.adj = make([][]int, d.e)
	return d
}

func (d Digraph) V() int {
	return d.v
}

func (d Digraph) E() int {
	return d.e
}

func (d *Digraph) AddEdge(v, w int) error {
	if err := d.validateVertex(v); err != nil {
		return err
	}
	if err := d.validateVertex(w); err != nil {
		return err
	}

	d.adj[v][w] = 1
	d.inDegree[w]++
	d.e++

	return nil
}

func (d Digraph) validateVertex(v int) error {
	if v < 0 || v >= d.v {
		return errors.New("illegal vertex")
	}
	return nil
}

func (d Digraph) Adj(v int) ([]int, error) {
	if err := d.validateVertex(v); err != nil {
		return nil, err
	}
	return d.adj[v], nil
}

func (d Digraph) OutDegree(v int) (int, error) {
	if err := d.validateVertex(v); err != nil {
		return 0, err
	}
	return len(d.adj[v]), nil
}

func (d Digraph) InDegree(v int) (int, error) {
	if err := d.validateVertex(v); err != nil {
		return 0, err
	}
	return d.inDegree[v], nil
}

func (d Digraph) Reverse() (*Digraph, error) {
	reverse := NewDigraph(d.v)
	for v := 0; v < d.v; v++ {
		for _, w := range d.adj[v] {
			if err := reverse.AddEdge(w, v); err != nil {
				return nil, err
			}
		}
	}
	return reverse, nil
}
