package sequentialSearch

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
	s.first = nil
	s.len = 0
	return s
}

func (s st) Contains(key int) bool {
	_, err := s.Get(key)
	return err == nil
}

func (s st) Get(key int) (int, error) {
	if s.len == 0 {
		return 0, errors.New("symbol table is empty")
	}

	for cur := s.first; cur != nil; cur = cur.next {
		if cur.key == key {
			return cur.value, nil
		}
	}

	return 0, errors.New("key is not exist")
}

func (s *st) Put(key, value int) {
	cur := s.first
	pre := s.first
	for ; cur != nil; cur = cur.next {
		if cur.key == key {
			cur.value = value
			return
		}
		pre = cur
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
	s.first = s.deleteNode(s.first, key)
}

func (s *st) deleteNode(n *node, key int) *node {
	if n == nil {
		return nil
	}

	if n.key == key {
		s.len--
		return n.next
	}

	n.next = s.deleteNode(n.next, key)
	return n
}
