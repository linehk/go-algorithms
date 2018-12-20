package maxPriorityQueue

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
	h.cap = cap + 1
	return h
}

func NewByArray(elements []int) *heap {
	l := len(elements)
	h := new(heap)
	h.elements = make([]int, l+1)
	h.len = l
	h.cap = l + 1
	for i := 0; i < l; i++ {
		h.elements[i+1] = elements[i]
	}
	for k := l / 2; k >= 1; k-- {
		h.sink(k)
	}
	return h
}

func (h heap) Max() (int, error) {
	if h.len == 0 {
		return 0, errors.New("priority queue is empty")
	}
	return h.elements[1], nil
}

func (h *heap) DelMax() (int, error) {
	if h.len == 0 {
		return 0, errors.New("priority queue is empty")
	}
	max := h.elements[1]
	h.elements[1], h.elements[h.len] = h.elements[h.len], h.elements[1]
	h.len--
	h.sink(1)
	if (h.len > 0) && (h.len == ((h.cap - 1) / 4)) {
		h.resize(h.cap / 2)
	}
	return max, nil
}

func (h *heap) Insert(v int) {
	// if h.len+1 > h.cap {
	// 	return errors.New("priority queue is full")
	// }
	if h.len == h.cap-1 {
		h.resize(2 * h.cap)
	}
	h.len++
	h.elements[h.len] = v
	h.swim(h.len)
	// return nil
}

func (h *heap) swim(k int) {
	for k > 1 && h.elements[k] > h.elements[k/2] {
		h.elements[k], h.elements[k/2] = h.elements[k/2], h.elements[k]
		k = k / 2
	}
}

func (h *heap) sink(k int) {
	for 2*k <= h.len {
		j := 2 * k
		if j < h.len && h.elements[j] < h.elements[j+1] {
			j++
		}
		if h.elements[j] < h.elements[k] {
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

func (h heap) isMaxHeap() bool {
	return h.isKMaxHeap(1)
}

func (h heap) isKMaxHeap(k int) bool {
	if k > h.len {
		return true
	}
	left := 2 * k
	right := 2*k + 1
	if left <= h.len && h.elements[k] < h.elements[left] {
		return false
	}
	if right <= h.len && h.elements[k] < h.elements[right] {
		return false
	}
	return h.isKMaxHeap(left) && h.isKMaxHeap(right)
}
