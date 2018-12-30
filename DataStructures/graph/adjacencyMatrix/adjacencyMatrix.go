package adjacencyMatrix

type UndirectedGraph struct {
	v   int
	e   int
	adj [][]int
}

func NewGraph(v int) *UndirectedGraph {
	g := new(UndirectedGraph)
	g.v = v
	g.e = 0
	g.adj = make([][]int, g.v)
	for i := range g.adj {
		g.adj[i] = make([]int, g.v)
	}
	return g
}

type Edge struct {
	v int
	w int
}

// assume vertex >= 0 and < graph vertex
func NewEdge(v, w int) *Edge {
	e := new(Edge)
	e.v = v
	e.w = w
	return e
}

func (u UndirectedGraph) V() int {
	return u.v
}

func (u UndirectedGraph) E() int {
	return u.e
}

func (u *UndirectedGraph) AddEdge(e Edge) {
	u.adj[e.v][e.w] = 1
	u.adj[e.w][e.v] = 1
	u.e++
}

func (u UndirectedGraph) Degree(v int) int {
	degree := 0
	for _, isConnected := range u.adj[v] {
		if isConnected == 1 {
			degree++
		}
	}
	return degree
}

func (u UndirectedGraph) Adj(v int) []int {
	return nil
}

type DirectedGraph struct {
	v        int
	e        int
	adj      [][]int
	inDegree []int
}

func NewDigraph(v int) *DirectedGraph {
	d := new(DirectedGraph)
	d.v = v
	d.e = 0
	d.adj = make([][]int, d.v)
	for i := range d.adj {
		d.adj[i] = make([]int, d.v)
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
	d.adj[e.v][e.w] = 1
	d.inDegree[e.w]++
	d.e++
}

func (d DirectedGraph) Adj(v int) []int {
	return nil
}

func (d DirectedGraph) OutDegree(v int) int {
	outDegree := 0
	for _, isConnected := range d.adj[v] {
		if isConnected == 1 {
			outDegree++
		}
	}
	return outDegree
}

func (d DirectedGraph) InDegree(v int) int {
	return d.inDegree[v]
}

func (d DirectedGraph) Reverse() *DirectedGraph {
	return nil
}
