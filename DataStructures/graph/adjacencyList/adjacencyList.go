package adjacencyList

import (
	"errors"
)

type list struct {
	first *node
}

type node struct {
        edge Edge
	next  *node
}

func (l list) get(key string) (int, error) {
	for cur := l.first; cur != nil; cur = cur.next {
		if key == cur.key {
			return cur.value, nil
		}
	}
	return 0, errors.New("key is not exist")
}

func (l list) contains(key string) bool {
	_, err := l.get(key)
	if err != nil {
		return false
	}
	return true
}

func (l *list) put(key string, value int) {
	n := new(node)
	n.key = key
	n.value = value
	n.next = nil

	if l.first == nil {
		l.first = n
	} else {
		n.next = l.first.next
		l.first.next = n
	}
}

func (l *list) delete(key string) {
	cur := l.first
	pre := l.first
	for ; key != cur.key; cur = cur.next {
		pre = cur
	}

	if (cur == l.first) && (pre == l.first) {
		l.first = nil
	} else {
		pre.next = cur.next
	}
}

type UndirectedGraph struct {
	v   int
	e   int
	adj []list
}

func NewUndirectedGraph(v int) *UndirectedGraph {
	u := new(UndirectedGraph)
	u.v = v
	u.e = 0
	u.adj = make([]list, v)
	return u
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
}

func (u UndirectedGraph) Degree() int {
	return 0
}

func (u UndirectedGraph) Adj(v int) []Edge {
	return nil
}
