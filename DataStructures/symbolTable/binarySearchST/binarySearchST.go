package binarySearchST

import (
	"errors"
)

type st struct {
	keys   []int
	values []int
	len    int
	cap    int
}

func New(cap int) *st {
	s := new(st)
	s.keys = make([]int, cap)
	s.values = make([]int, cap)
	s.len = 0
	s.cap = cap
	return s
}

func (s *st) resize(cap int) {
	newKeys := make([]int, cap)
	newValues := make([]int, cap)
	for i := 0; i < s.len; i++ {
		newKeys[i] = s.keys[i]
		newValues[i] = s.values[i]
	}
	s.keys = newKeys
	s.values = newValues
	s.cap = cap
}

func (s st) Contains(key int) bool {
	_, err := s.Get(key)
	if err != nil {
		return false
	}
	return true
}

func (s st) Get(key int) (int, error) {
	if s.len == 0 {
		return 0, errors.New("symbol table is empty")
	}
	i := s.Rank(key)
	if i < s.len && key == s.keys[i] {
		return s.values[i], nil
	}
	return 0, errors.New("key is not in the st")
}

func (s st) Rank(key int) int {
	lo := 0
	hi := s.len - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if key < s.keys[mid] {
			hi = mid - 1
		} else if key > s.keys[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return lo
}

func (s *st) Put(key int, value int) {
	i := s.Rank(key)
	if i < s.len && key == s.keys[i] {
		s.values[i] = value
		return
	}

	if s.len == s.cap {
		s.resize(2 * s.cap)
	}

	for j := s.len; j > i; j-- {
		s.keys[j] = s.keys[j-1]
		s.values[j] = s.values[j-1]
	}
	s.keys[i] = key
	s.values[i] = value
	s.len++
}

func (s *st) Delete(key int) {
	if s.len == 0 {
		return
	}
	i := s.Rank(key)
	if i == s.len || key != s.keys[i] {
		return
	}

	for j := i; j < s.len-1; j++ {
		s.keys[j] = s.keys[j+1]
		s.values[j] = s.keys[j+1]
	}

	s.len--
	// s.keys[s.len] = 0
	// s.values[s.len] = 0

	if s.len > 0 && s.len == s.cap/4 {
		s.resize(s.cap / 2)
	}
}

func (s *st) DeleteMin() error {
	min, err := s.Min()
	if err != nil {
		return err
	}
	s.Delete(min)
	return nil
}

func (s *st) DeleteMax() error {
	max, err := s.Max()
	if err != nil {
		return err
	}
	s.Delete(max)
	return nil
}

func (s st) Min() (int, error) {
	if s.len == 0 {
		return 0, errors.New("symbol table is empty")
	}
	return s.keys[0], nil
}

func (s st) Max() (int, error) {
	if s.len == 0 {
		return 0, errors.New("symbol table is empty")
	}
	return s.keys[s.len-1], nil
}

func (s st) SelectKth(k int) (int, error) {
	if k < 0 || k >= s.len {
		return 0, errors.New("illegal index k")
	}

	return s.keys[k], nil
}

func (s st) Floor(key int) (int, error) {
	i := s.Rank(key)
	if i < s.len && key == s.keys[i] {
		return s.keys[i], nil
	}
	if i == 0 {
		return 0, errors.New("key is not in the table")
	} else {
		return s.keys[i-1], nil
	}
}

func (s st) Ceiling(key int) (int, error) {
	i := s.Rank(key)
	if i == s.len {
		return 0, errors.New("key is not in the table")
	} else {
		return s.keys[i], nil
	}
}

func (s st) Size(lo, hi int) int {
	if lo > hi {
		return 0
	}
	if s.Contains(hi) {
		return s.Rank(hi) - s.Rank(lo) + 1
	} else {
		return s.Rank(hi) - s.Rank(lo)
	}
}
