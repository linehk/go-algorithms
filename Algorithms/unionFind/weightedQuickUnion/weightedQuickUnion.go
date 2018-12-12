package weightedQuickUnion

type uf struct {
	id    []int
	scale []int
	count int
}

func New(n int) *uf {
	u := new(uf)
	u.id = make([]int, n)
	u.scale = make([]int, n)
	for i := 0; i < n; i++ {
		u.id[i] = i
		u.scale[i] = 1
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

	if u.scale[pRoot] < u.scale[qRoot] {
		u.id[pRoot] = qRoot
		u.scale[qRoot] += u.scale[pRoot]
	} else {
		u.id[qRoot] = pRoot
		u.scale[pRoot] += u.scale[qRoot]
	}
	u.count--
}

func (u uf) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}
