package separateChainingHash

import (
	"errors"
)

type st struct {
	len  int
	cap  int
	list []list
}

type list struct {
	first *node
}

type node struct {
	key   string
	value int
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
	return err != nil
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

func New(cap int) *st {
	s := new(st)
	s.len = 0
	s.cap = cap
	s.list = make([]list, s.cap)
	return s
}

func (s st) Size() int {
	return s.len
}

func (s st) IsEmpty() bool {
	return s.Size() == 0
}

func (s st) Contains(key string) bool {
	_, err := s.Get(key)
	return err == nil
}

func (s st) Get(key string) (int, error) {
	i := s.hash(key)
	v, err := s.list[i].get(key)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func (s *st) resize(cap int) {
	newST := New(cap)
	for i := 0; i < s.cap; i++ {
		for cur := s.list[i].first; cur != nil; cur = cur.next {
			newST.Put(cur.key, cur.value)
		}
	}
	s.len = newST.len
	s.cap = newST.cap
	s.list = newST.list
}

func (s st) hash(key string) int {
	return (s.hashCode(key) & 0x7fffffff) % s.cap
}

func (s st) hashCode(key string) int {
	hash := 0
	l := len(key)
	for i := 0; i < l; i++ {
		hash = 31*hash + int(key[i])
	}
	return hash
}

func (s *st) Put(key string, value int) {
	if s.len >= (s.cap * 10) {
		s.resize(s.cap * 2)
	}

	i := s.hash(key)
	if !s.list[i].contains(key) {
		s.len++
	}
	s.list[i].put(key, value)
}

func (s *st) Delete(key string) {
	i := s.hash(key)
	if s.list[i].contains(key) {
		s.len--
	}
	s.list[i].delete(key)

	if (s.cap > 4) && (s.len <= (s.cap * 2)) {
		s.resize(s.cap / 2)
	}
}
