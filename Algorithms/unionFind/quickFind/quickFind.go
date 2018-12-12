package quickFind

type uf struct {
	id    []int
	count int
}

func New(n int) *uf {
	u := new(uf)
	u.id = make([]int, n)
	for i := 0; i < n; i++ {
		u.id[i] = i
	}
	u.count = n
	return u
}

func (u uf) Find(p int) int {
	return u.id[p]
}

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

func (u uf) Connected(p, q int) bool {
	return u.id[p] == u.id[q]
}
