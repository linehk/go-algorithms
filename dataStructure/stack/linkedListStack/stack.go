package linkedListStack

import (
	"errors"
)

type stack struct {
	top *node
}

type node struct {
	value interface{}
	next  *node
}

// New creates a stack.
func New() *stack {
	s := new(stack)
	return s
}

// Push inserts v in the stack.
func (s *stack) Push(v interface{}) {
	n := new(node)
	n.value = v
	n.next = nil

	// empty stack wouldn't execute this
	if s.top != nil {
		n.next = s.top
	}
	s.top = n
}

// Pop deletes the element from the stack and return it.
func (s *stack) Pop() (interface{}, error) {
	if s.top == nil {
		return nil, errors.New("stack is empty")
	}

	v := s.top.value
	s.top = s.top.next
	return v, nil
}
