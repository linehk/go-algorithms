package quickUnion

import (
	"testing"
)

func TestFind(t *testing.T) {
	n := 10
	uf := New(n)
	for i := 0; i < n; i++ {
		if got := uf.Find(i); got != i {
			t.Errorf("%v. got %v, want %v", i, got, i)
		}
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		n int
		p int
		q int
	}{
		{10, 0, 1},
	}
	for i, tt := range tests {
		uf := New(tt.n)
		uf.Union(tt.p, tt.q)
		if got := uf.Connected(tt.p, tt.q); !got {
			t.Errorf("%v. got %v, want %v", i, got, true)
		}
	}
}
