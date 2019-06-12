package arrayStack

import (
	"errors"
)

type stack struct {
	cap      int
	top      int
	elements []interface{}
}

func (s stack) Cap() int {
	return s.cap
}

// New creates a stack by cap.
func New(cap int) *stack {
	s := new(stack)
	s.cap = cap
	s.top = -1
	s.elements = make([]interface{}, s.cap)
	return s
}

// Push inserts v in the stack.
func (s *stack) Push(v interface{}) error {
	if s.top+1 >= s.cap {
		return errors.New("stack is full")
	}

	s.top++
	s.elements[s.top] = v
	return nil
}

// Pop deletes the element from the stack and return it.
func (s *stack) Pop() (interface{}, error) {
	if s.top < 0 {
		return nil, errors.New("stack is empty")
	}

	v := s.elements[s.top]
	s.elements[s.top] = nil
	s.top--
	return v, nil
}
