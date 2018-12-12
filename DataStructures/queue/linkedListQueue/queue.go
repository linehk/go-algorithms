package linkedListQueue

import (
	"errors"
)

type node struct {
	value interface{}
	next  *node
}

type queue struct {
	len   int
	front *node
	rear  *node
}

func (q queue) Len() int {
	return q.len
}

// New creates a queue.
func New() *queue {
	q := new(queue)
	q.len = 0
	q.front = new(node)
	q.front.next = nil
	q.front.value = nil
	q.rear = q.front
	return q
}

// Enqueue inserts v in the queue.
func (q *queue) Enqueue(v interface{}) {
	q.rear.value = v

	n := new(node)
	n.value = nil
	n.next = nil

	q.rear.next = n
	q.rear = n
	q.len++
}

// Dequeue deletes v in the queue and return it.
func (q *queue) Dequeue() (interface{}, error) {
	if q.front == q.rear {
		return nil, errors.New("queue is empty")
	}

	v := q.front.value
	q.front.value = nil

	next := q.front.next
	q.front.next = nil
	q.front = next

	q.len--
	return v, nil
}
