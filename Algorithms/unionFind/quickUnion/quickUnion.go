package quickUnion

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
	for p != u.id[p] {
		p = u.id[p]
	}
	return p
}

func (u *uf) Union(p, q int) {
	pRoot := u.Find(p)
	qRoot := u.Find(q)

	if pRoot == qRoot {
		return
	}

	u.id[pRoot] = qRoot
	u.count--
}

func (u uf) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}
