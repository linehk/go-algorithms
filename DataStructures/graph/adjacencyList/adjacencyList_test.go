package adjacencyList

import (
	"testing"
)

func TestNewEdge(t *testing.T) {
	u := NewUndirectedGraph(10)
	u.adj[0].put("aa", 1)
	v, err := u.adj[0].get("aa")
	if err != nil {
		t.Error(err)
	}
	t.Log(v)
}
