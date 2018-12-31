package adjacencyMatrix

type UndirectedGraph struct {
	v   int
	e   int
	adj [][]Edge
}

func NewGraph(v int) *UndirectedGraph {
	g := new(UndirectedGraph)
	g.v = v
	g.e = 0
	g.adj = make([][]Edge, g.v)
	for i := range g.adj {
		g.adj[i] = make([]Edge, g.v)
	}
	return g
}

type Edge struct {
	v      int
	w      int
	weight int
}

// assume vertex >= 0 and < graph vertex
func NewEdge(v, w, weight int) *Edge {
	e := new(Edge)
	e.v = v
	e.w = w
	e.weight = weight
	return e
}

func (u UndirectedGraph) V() int {
	return u.v
}

func (u UndirectedGraph) E() int {
	return u.e
}

func (u *UndirectedGraph) AddEdge(e Edge) {
	u.adj[e.v][e.w] = e
	u.adj[e.w][e.v] = e
	u.e++
}

func (u UndirectedGraph) Degree(v int) int {
	degree := 0
	for _, edge := range u.adj[v] {
		if edge.weight != 0 {
			degree++
		}
	}
	return degree
}

func (u UndirectedGraph) Adj(v int) []Edge {
	edges := make([]Edge, 0)
	for _, edge := range u.adj[v] {
		if edge.weight != 0 {
			edges = append(edges, edge)
		}
	}
	return edges
}

type DirectedGraph struct {
	v        int
	e        int
	adj      [][]Edge
	inDegree []int
}

func NewDigraph(v int) *DirectedGraph {
	d := new(DirectedGraph)
	d.v = v
	d.e = 0
	d.adj = make([][]Edge, d.v)
	for i := range d.adj {
		d.adj[i] = make([]Edge, d.v)
	}
	d.inDegree = make([]int, d.v)
	return d
}

func (d DirectedGraph) V() int {
	return d.v
}

func (d DirectedGraph) E() int {
	return d.e
}

func (d *DirectedGraph) AddEdge(e Edge) {
	d.adj[e.v][e.w] = e
	d.inDegree[e.w]++
	d.e++
}

func (d DirectedGraph) Adj(v int) []Edge {
	edges := make([]Edge, 0)
	for _, edge := range d.adj[v] {
		if edge.weight != 0 {
			edges = append(edges, edge)
		}
	}
	return edges
}

func (d DirectedGraph) OutDegree(v int) int {
	outDegree := 0
	for _, edge := range d.adj[v] {
		if edge.weight != 0 {
			outDegree++
		}
	}
	return outDegree
}

func (d DirectedGraph) InDegree(v int) int {
	return d.inDegree[v]
}

func (d DirectedGraph) Reverse() *DirectedGraph {
	reverseGraph := NewDigraph(d.v)
	for v := 0; v < d.v; v++ {
		for _, edge := range d.adj[v] {
			if edge.weight != 0 {
				reverseGraph.AddEdge(Edge{edge.w, edge.v, edge.weight})
			}
		}
	}
	return reverseGraph
}
