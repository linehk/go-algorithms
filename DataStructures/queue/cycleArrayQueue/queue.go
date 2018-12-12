// Package cycleArrayQueue implements queue by slice.
// elements[] in memory doesn't in order.
package cycleArrayQueue

import (
	"errors"
)

type queue struct {
	cap      int
	len      int
	front    int
	rear     int
	elements []interface{}
}

func (q queue) Cap() int {
	return q.cap - 1
}

func (q queue) Len() int {
	return q.len
}

// New creates a queue by cap.
func New(cap int) *queue {
	q := new(queue)
	// extra one node
	q.cap = cap + 1
	q.len = 0
	q.front = 0
	q.rear = 0
	q.elements = make([]interface{}, q.cap)
	return q
}

// Enqueue inserts v in the queue.
func (q *queue) Enqueue(v interface{}) error {
	if (q.rear+1)%q.cap == q.front {
		return errors.New("queue is full")
	}

	q.elements[q.rear] = v
	q.rear = (q.rear + 1) % q.cap
	q.len++
	return nil
}

// Dequeue deletes v in the queue and return it.
func (q *queue) Dequeue() (interface{}, error) {
	if q.front == q.rear {
		return nil, errors.New("queue is empty")
	}

	v := q.elements[q.front]
	q.elements[q.front] = nil
	q.front = (q.front + 1) % q.cap
	q.len--
	return v, nil
}
