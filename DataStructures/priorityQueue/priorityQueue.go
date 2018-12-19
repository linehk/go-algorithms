package priorityQueue

import (
	"errors"
)

type priorityQueue struct {
	elements []int
	len      int
	cap      int
}

func (p priorityQueue) Len() int {
	return p.len
}

func (p priorityQueue) Cap() int {
	return p.cap
}

func New(cap int) *priorityQueue {
	p := new(priorityQueue)
	p.elements = make([]int, cap+1)
	p.len = 0
	p.cap = cap
	return p
}

func NewByArray(elements []int) *priorityQueue {
	l := len(elements)
	p := new(priorityQueue)
	p.elements = make([]int, l+1)
	p.len = l
	p.cap = l
	for i := 0; i < l; i++ {
		p.elements[i+1] = elements[i]
	}
	for k := l / 2; k >= 1; k-- {
		p.sink(k)
	}
	return p
}

func (p priorityQueue) Max() (int, error) {
	if p.len == 0 {
		return 0, errors.New("priority queue is empty")
	}
	return p.elements[1], nil
}

func (p *priorityQueue) Insert(key int) error {
	if p.len+1 > p.cap {
		return errors.New("priority queue is full")
	}
	if p.len == p.cap {
		p.resize(2 * p.cap)
	}
	p.len++
	p.elements[p.len] = key
	p.swim(p.len)
	return nil
}

func (p *priorityQueue) DelMax() (int, error) {
	if p.len == 0 {
		return 0, errors.New("priority queue is empty")
	}
	max := p.elements[1]
	p.elements[1], p.elements[p.len] = p.elements[p.len], p.elements[1]
	p.len--
	p.sink(1)
	if p.len > 0 && (p.len == (p.cap / 4)) {
		p.resize(p.cap / 2)
	}
	return max, nil
}

func (p *priorityQueue) swim(k int) {
	for k > 1 && p.elements[k/2] < p.elements[k] {
		p.elements[k], p.elements[k/2] = p.elements[k/2], p.elements[k]
		k = k / 2
	}
}

func (p *priorityQueue) sink(k int) {
	for 2*k <= p.len {
		j := 2 * k
		if j < p.len && p.elements[j] < p.elements[j+1] {
			j++
		}
		if p.elements[j] < p.elements[k] {
			break
		}
		p.elements[k], p.elements[j] = p.elements[j], p.elements[k]
		k = j
	}
}

func (p priorityQueue) isMaxHeap() bool {
	return p.isSubMaxHeap(1)
}

func (p priorityQueue) isSubMaxHeap(k int) bool {
	if k > p.len {
		return true
	}

	left := 2 * k
	right := 2*k + 1
	if left <= p.len && p.elements[k] < p.elements[left] {
		return false
	}
	if right <= p.len && p.elements[k] < p.elements[right] {
		return false
	}
	return p.isSubMaxHeap(left) && p.isSubMaxHeap(right)
}

func (p *priorityQueue) resize(cap int) {
	newElements := make([]int, cap)
	for i := 1; i <= p.len; i++ {
		newElements[i] = p.elements[i]
	}
	p.elements = newElements
}
