// Package quickFind implements Disjoint-set.
package quickFind

type uf struct {
	id    []int
	count int
}

// New init and returns a new uf.
func New(n int) *uf {
	u := new(uf)
	u.id = make([]int, n)
	for i := 0; i < n; i++ {
		u.id[i] = i
	}
	u.count = n
	return u
}

// Find returns the set of p.
func (u uf) Find(p int) int {
	return u.id[p]
}

// Union joints p and q set.
func (u *uf) Union(p, q int) {
	pid := u.Find(p)
	qid := u.Find(q)

	if pid == qid {
		return
	}

	l := len(u.id)
	for i := 0; i < l; i++ {
		if u.id[i] == pid {
			u.id[i] = qid
		}
	}
	u.count--
}

// Connected returns true if p and q in the same sets.
func (u uf) Connected(p, q int) bool {
	return u.id[p] == u.id[q]
}
