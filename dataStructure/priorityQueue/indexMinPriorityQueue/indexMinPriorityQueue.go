package indexMinPriorityQueue

import (
	"errors"
)

type heap struct {
	index    []int
	elements []int
	keys     []int
	len      int
	cap      int
}

func (h heap) Len() int {
	return h.len
}

func (h heap) Cap() int {
	return h.cap
}

func New(cap int) *heap {
	h := new(heap)
	h.index = make([]int, cap+1)
	h.elements = make([]int, cap+1)
	h.keys = make([]int, cap+1)
	h.len = 0
	h.cap = cap
	for i := 0; i <= cap; i++ {
		h.elements[i] = -1
	}
	return h
}

func (h heap) Contains(i int) (bool, error) {
	if i < 0 || i >= h.cap {
		return false, errors.New("illegal elements i")
	}
	return h.elements[i] != -1, nil
}

func (h *heap) Insert(i, k int) error {
	if i < 0 || i >= h.cap {
		return errors.New("illegal index i")
	}
	isContains, err := h.Contains(i)
	if err != nil {
		return err
	}
	if isContains {
		return errors.New("index is already in the priority queue")
	}
	h.len++
	h.elements[i] = h.len
	h.index[h.len] = i
	h.keys[i] = k
	h.swim(h.len)
	return nil
}

func (h heap) MinIndex() (int, error) {
	if h.len == 0 {
		return 0, errors.New("priority queue is empty")
	}
	return h.index[1], nil
}

func (h heap) MinKey() (int, error) {
	if h.len == 0 {
		return 0, errors.New("priority queue is empty")
	}
	return h.keys[h.index[1]], nil
}

func (h *heap) DelMin() (int, error) {
	if h.len == 0 {
		return 0, errors.New("priority queue is empty")
	}
	min := h.index[1]
	h.exch(1, h.len)
	h.len--
	h.sink(1)
	h.elements[min] = -1
	// h.keys[min] = nil
	h.keys[min] = 0
	// h.index[h.len+1] = -1
	return min, nil
}

func (h heap) KeyOf(i int) (int, error) {
	if i < 0 || i >= h.cap {
		return 0, errors.New("illegal index i")
	}
	isContains, err := h.Contains(i)
	if err != nil {
		return 0, err
	}
	if !isContains {
		return 0, errors.New("index is not in the priority queue")
	}
	return h.keys[i], nil
}

func (h *heap) ChangeKey(i, k int) error {
	if i < 0 || i >= h.cap {
		return errors.New("illegal index i")
	}
	isContains, err := h.Contains(i)
	if err != nil {
		return err
	}
	if !isContains {
		return errors.New("index is not in the priority queue")
	}
	h.keys[i] = k
	h.swim(h.elements[i])
	h.sink(h.elements[i])
	return nil
}

func (h *heap) DecreaseKey(i, k int) error {
	if i < 0 || i >= h.cap {
		return errors.New("illegal index i")
	}
	isContains, err := h.Contains(i)
	if err != nil {
		return err
	}
	if !isContains {
		return errors.New("index is not in the priority queue")
	}
	// need h.keys[i] > key
	if h.keys[i] <= k {
		return errors.New("key isn't small that keys[i]")
	}
	h.keys[i] = k
	h.swim(h.elements[i])
	return nil
}

func (h *heap) IncreaseKey(i, k int) error {
	if i < 0 || i >= h.cap {
		return errors.New("illegal index i")
	}
	isContains, err := h.Contains(i)
	if err != nil {
		return err
	}
	if !isContains {
		return errors.New("index is not in the priority queue")
	}
	// need h.keys[i] < key
	if h.keys[i] >= k {
		return errors.New("key isn't big that keys[i]")
	}
	h.keys[i] = k
	h.sink(h.elements[i])
	return nil
}

func (h *heap) Delete(i int) error {
	if i < 0 || i >= h.cap {
		return errors.New("illegal index i")
	}
	isContains, err := h.Contains(i)
	if err != nil {
		return err
	}
	if !isContains {
		return errors.New("index is not in the priority queue")
	}
	index := h.elements[i]
	h.exch(index, h.len)
	h.len--
	h.swim(index)
	h.sink(index)
	// h.keys[i] = nil
	h.keys[i] = 0
	h.elements[i] = -1
	return nil
}

func (h heap) greater(i, j int) bool {
	return h.keys[h.index[i]] > h.keys[h.index[j]]
}

func (h *heap) exch(i, j int) {
	h.index[i], h.index[j] = h.index[j], h.index[i]
	h.elements[h.index[i]] = i
	h.elements[h.index[j]] = j
}

func (h *heap) swim(k int) {
	for k > 1 && h.greater(k/2, k) {
		h.exch(k, k/2)
		k = k / 2
	}
}

func (h *heap) sink(k int) {
	for 2*k <= h.len {
		j := 2 * k
		if j < h.len && h.greater(j, j+1) {
			j++
		}
		if !h.greater(k, j) {
			break
		}
		h.exch(k, j)
		k = j
	}
}
