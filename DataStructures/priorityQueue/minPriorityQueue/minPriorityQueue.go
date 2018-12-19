package minPriorityQueue

import (
	"errors"
)

type heap struct {
	elements []int
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
	h.elements = make([]int, cap+1)
	h.len = 0
	h.cap = cap
	return h
}

func NewByArray(elements []int) *heap {
	l := len(elements)
	h := new(heap)
	h.elements = make([]int, l+1)
	h.len = l
	h.cap = l
	for i := 0; i < l; i++ {
		h.elements[i+1] = elements[i]
	}
	for k := l / 2; k >= 1; k-- {
		h.sink(k)
	}
	return h
}

func (h heap) Min() (int, error) {
	if h.len == 0 {
		return 0, errors.New("priority queue is empty")
	}
	return h.elements[1], nil
}

func (h *heap) Insert(v int) {
	if h.len == h.cap {
		h.resize(2 * (h.cap + 1))
	}
	h.len++
	h.elements[h.len] = v
	h.swim(h.len)
}

func (h *heap) DelMin() (int, error) {
	if h.len == 0 {
		return 0, errors.New("priority queue is empty")
	}
	min := h.elements[1]
	h.elements[1], h.elements[h.len] = h.elements[h.len], h.elements[1]
	h.len--
	h.sink(1)
	if h.len > 0 && (h.len == (h.cap / 4)) {
		h.resize((h.cap + 1) / 2)
	}
	return min, nil
}

func (h *heap) swim(k int) {
	for k > 1 && h.elements[k] < h.elements[k/2] {
		h.elements[k], h.elements[k/2] = h.elements[k/2], h.elements[k]
		k = k / 2
	}
}

func (h *heap) sink(k int) {
	for 2*k <= h.len {
		j := 2 * k
		if j < h.len && h.elements[j+1] < h.elements[j] {
			j++
		}
		if h.elements[j] > h.elements[k] {
			break
		}
		h.elements[j], h.elements[k] = h.elements[k], h.elements[j]
		k = j
	}
}

func (h *heap) resize(cap int) {
	newElements := make([]int, cap)
	for i := 1; i <= h.len; i++ {
		newElements[i] = h.elements[i]
	}
	h.elements = newElements
}

func (h heap) isMinHeap() bool {
	return h.isKMinHeap(1)
}

func (h heap) isKMinHeap(k int) bool {
	if k > h.len {
		return true
	}
	left := 2 * k
	right := 2*k + 1
	if left <= h.len && h.elements[k] > h.elements[left] {
		return false
	}
	if right <= h.len && h.elements[k] > h.elements[right] {
		return false
	}
	return h.isKMinHeap(left) && h.isKMinHeap(right)
}
