// Package arrayQueue implements queue by slice.
package arrayQueue

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
	return q.cap
}

func (q queue) Len() int {
	return q.len
}

// New creates a queue by cap.
func New(cap int) *queue {
	q := new(queue)
	q.cap = cap
	q.len = 0
	q.front = 0
	q.rear = 0
	q.elements = make([]interface{}, q.cap)
	return q
}

// Enqueue inserts v in the queue.
func (q *queue) Enqueue(v interface{}) error {
	if q.rear == q.cap && q.front == 0 {
		return errors.New("queue is full")
	}

	// rearrange array
	if q.rear == q.cap && q.front != 0 {
		n := q.rear - q.front
		for i := 0; i < n; i++ {
			q.elements[i] = q.elements[q.cap-n+i]
			q.elements[q.cap-n+i] = nil
		}
		q.front = 0
		q.rear = n
	}
	q.elements[q.rear] = v
	q.rear++
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
	q.front++
	q.len--
	return v, nil
}
