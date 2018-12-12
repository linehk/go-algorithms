// Package array implements DataStructure array can dynamically grow and shark.
package array

import (
	"errors"
)

type array struct {
	len      int
	cap      int
	elements []interface{}
}

// New creates a array with cap.
func New(cap int) *array {
	a := new(array)
	a.len = 0
	a.cap = cap
	a.elements = make([]interface{}, a.cap)
	return a
}

// Get returns the value at i.
func (a array) Get(i int) (interface{}, error) {
	if i < 0 || i > a.len-1 {
		return nil, errors.New("illegal index")
	}

	return a.elements[i], nil
}

// Set assign v to array index i.
func (a *array) Set(i int, v interface{}) error {
	if i < 0 || i > a.len-1 {
		return errors.New("illegal index")
	}

	a.elements[i] = v
	return nil
}

// Len returns array available element counts.
func (a array) Len() int {
	return a.len
}

// Cap returns array capacity.
func (a array) Cap() int {
	return a.cap
}

// Append inserts v at array last.
func (a *array) Append(v interface{}) {
	if a.len+1 > a.cap {
		a.resize(a.cap * 2)
	}

	a.elements[a.len] = v
	a.len++
}

// Insert inserts v before i.
// Can't inserts to a empty array.
func (a *array) Insert(i int, v interface{}) error {
	if i < 0 || i > a.len-1 {
		return errors.New("illegal index")
	}

	if a.len+1 > a.cap {
		a.resize(a.cap * 2)
	}

	for last := a.len - 1; last >= i; last-- {
		a.elements[last+1] = a.elements[last]
	}

	a.elements[i] = v
	a.len++
	return nil
}

// Delete deletes the element at i.
func (a *array) Delete(i int) error {
	if i < 0 || i >= a.len {
		return errors.New("illegal index")
	}

	if a.len-1 < a.cap/4 {
		a.resize(a.cap / 2)
	}

	for cur := i; cur < a.len-1; cur++ {
		a.elements[cur] = a.elements[cur+1]
	}

	a.elements[a.len-1] = nil
	a.len--
	return nil
}

func (a *array) resize(cap int) {
	a.cap = cap
	newElements := make([]interface{}, a.cap)
	for i := 0; i < a.len; i++ {
		newElements[i] = a.elements[i]
	}
	a.elements = newElements
}
