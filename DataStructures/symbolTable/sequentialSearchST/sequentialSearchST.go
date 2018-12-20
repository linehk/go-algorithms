package sequentialSearchST

import (
	"errors"
)

type st struct {
	first *node
	len   int
}

type node struct {
	key   int
	value int
	next  *node
}

func New() *st {
	s := new(st)
	s.len = 0
	return s
}

func (s *st) Contains(key int) (bool, error) {
	_, err := s.Get(key)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s st) Get(key int) (int, error) {
	for x := s.first; x != nil; x = x.next {
		if x.key == key {
			return x.value, nil
		}
	}
	return 0, errors.New("key is not in the st")
}

func (s *st) Put(key, value int) {
	x := s.first
	pre := s.first
	for ; x != nil; x = x.next {
		if x.key == key {
			x.value = value
			return
		}
		pre = x
	}
	newNode := &node{key, value, nil}
	if s.first == nil {
		s.first = newNode
	} else {
		pre.next = newNode
	}
	s.len++
}

func (s *st) Delete(key int) {
	sts := st{}
	s.first = sts.deleteNode(s.first, key)
}

func (s *st) deleteNode(x *node, key int) *node {
	if x == nil {
		return nil
	}
	if x.key == key {
		s.len--
		return x.next
	}
	sts := st{}
	x.next = sts.deleteNode(x.next, key)
	return x
}
