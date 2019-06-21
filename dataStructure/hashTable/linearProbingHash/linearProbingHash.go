package linearProbingHash

import (
	"errors"
)

type st struct {
	keys   []string
	values []int
	len    int
	cap    int
}

const NULL = ""

func New(cap int) *st {
	s := new(st)
	s.keys = make([]string, cap)
	s.values = make([]int, cap)
	s.len = 0
	s.cap = cap
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
	for i := s.hash(key); s.keys[i] != NULL; i = (i + 1) % s.cap {
		if key == s.keys[i] {
			return s.values[i], nil
		}
	}
	return 0, errors.New("key is not exist")
}

func (s st) hash(key string) int {
	return (s.hashCode(key) & 0x7fffffff) % s.cap
}

func (s st) hashCode(key string) int {
	l := len(key)
	hash := 0
	for i := 0; i < l; i++ {
		hash = 31*hash + int(key[i])
	}
	return hash
}

func (s *st) resize(cap int) {
	newST := New(cap)
	for i := 0; i < s.cap; i++ {
		if s.keys[i] != NULL {
			newST.Put(s.keys[i], s.values[i])
		}
	}
	s.keys = newST.keys
	s.values = newST.values
	s.cap = newST.cap
}

func (s *st) Put(key string, value int) {
	if s.len >= (s.cap / 2) {
		s.resize(s.cap * 2)
	}

	i := s.hash(key)
	for ; s.keys[i] != NULL; i = (i + 1) % s.cap {
		if key == s.keys[i] {
			s.values[i] = value
			return
		}
	}

	s.keys[i] = key
	s.values[i] = value
	s.len++
}

func (s *st) Delete(key string) {
	if !s.Contains(key) {
		return
	}

	i := s.hash(key)
	for ; key != s.keys[i]; i = (i + 1) % s.cap {
	}

	s.keys[i] = NULL
	s.values[i] = 0

	i = (i + 1) % s.cap
	for ; s.keys[i] != NULL; i = (i + 1) % s.cap {
		newKey := s.keys[i]
		newVal := s.values[i]

		s.keys[i] = NULL
		s.values[i] = 0

		s.len--
		s.Put(newKey, newVal)
	}

	s.len--

	if (s.len > 0) && (s.len <= (s.cap / 8)) {
		s.resize(s.cap / 2)
	}
}
