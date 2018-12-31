package adjacencyMatrix

import (
	"testing"
)

func TestGraph(t *testing.T) {
	v := 10
	u := NewGraph(v)
	e1 := NewEdge(0, 1, 1)
	e2 := NewEdge(1, 2, 1)
	e3 := NewEdge(2, 0, 1)
	u.AddEdge(*e1)
	u.AddEdge(*e2)
	u.AddEdge(*e3)

	if u.V() != v {
		t.Errorf("got %v, want %v", u.V(), v)
	}
	if u.E() != 3 {
		t.Errorf("got %v, want %v", u.E(), 3)
	}

	if u.Degree(0) != 2 {
		t.Errorf("got %v, want %v", u.Degree(0), 2)
	}

	wantAdj := make([]Edge, 2)
	wantAdj[0] = *NewEdge(0, 1, 1)
	wantAdj[1] = *NewEdge(2, 0, 1)
	gotAdj := u.Adj(0)
	for i, e := range gotAdj {
		if !isSameEdge(e, wantAdj[i]) {
			t.Errorf("got %v, want %v", gotAdj, wantAdj)
		}
	}
}

func TestDirectedGraph(t *testing.T) {
	v := 10
	u := NewDigraph(v)
	e1 := NewEdge(0, 1, 1)
	e2 := NewEdge(1, 2, 1)
	e3 := NewEdge(2, 0, 1)
	u.AddEdge(*e1)
	u.AddEdge(*e2)
	u.AddEdge(*e3)

	if u.V() != v {
		t.Errorf("got %v, want %v", u.V(), v)
	}
	if u.E() != 3 {
		t.Errorf("got %v, want %v", u.E(), 3)
	}

	if u.OutDegree(0) != 1 {
		t.Errorf("got %v, want %v", u.OutDegree(0), 1)
	}

	if u.InDegree(0) != 1 {
		t.Errorf("got %v, want %v", u.InDegree(0), 1)
	}

	wantAdj := make([]Edge, 1)
	wantAdj[0] = *NewEdge(0, 1, 1)
	gotAdj := u.Adj(0)
	for i, e := range gotAdj {
		if !isSameEdge(e, wantAdj[i]) {
			t.Errorf("got %v, want %v", gotAdj, wantAdj)
		}
	}

	reverseGraph := u.Reverse()
	wantAdj[0] = *NewEdge(1, 0, 1)
	gotAdj = reverseGraph.Adj(1)
	for i, e := range gotAdj {
		if !isSameEdge(e, wantAdj[i]) {
			t.Errorf("got %v, want %v", gotAdj, wantAdj)
		}
	}
}

func isSameEdge(a, b Edge) bool {
	if a.weight != b.weight {
		return false
	}
	if a.v != b.v {
		return false
	}
	if a.w != b.w {
		return false
	}
	return true
}
