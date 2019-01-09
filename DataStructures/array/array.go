// Package Array implements DataStructure Array can dynamically grow and shark.
package array

import (
	"fmt"
)

type Array struct {
	len      int
	cap      int
	elements []interface{}
}

func New(cap int) *Array {
	return new(Array).Init(cap)
}

func (a *Array) Init(cap int) *Array {
	a.len = 0
	a.cap = cap
	a.elements = make([]interface{}, cap)
	return a
}

// Get returns the value at i.
func (a *Array) Get(i int) interface{} {
	if i < 0 || i >= a.len {
		panic(fmt.Sprintf("illegal index %d", i))
	}

	return a.elements[i]
}

// Set assign v to Array index i.
func (a *Array) Set(i int, v interface{}) {
	if i < 0 || i >= a.len {
		panic(fmt.Sprintf("illegal index %d", i))
	}

	a.elements[i] = v
}

// Len returns Array available element counts.
func (a *Array) Len() int {
	return a.len
}

// Cap returns Array capacity.
func (a *Array) Cap() int {
	return a.cap
}

// Append inserts v at Array last.
func (a *Array) Append(v interface{}) {
	if a.len+1 > a.cap {
		a.resize(a.cap * 2)
	}

	a.elements[a.len] = v
	a.len++
}

// Insert inserts v before i.
// Can't inserts to a empty Array.
func (a *Array) Insert(i int, v interface{}) {
	if i < 0 || i >= a.len {
		panic(fmt.Sprintf("illegal index %d", i))
	}

	if a.len+1 > a.cap {
		a.resize(a.cap * 2)
	}

	for last := a.len - 1; last >= i; last-- {
		a.elements[last+1] = a.elements[last]
	}

	a.elements[i] = v
	a.len++
}

// Delete deletes the element at i.
func (a *Array) Delete(i int) {
	if i < 0 || i >= a.len {
		panic(fmt.Sprintf("illegal index %d", i))
	}

	if a.len-1 < a.cap/4 {
		a.resize(a.cap / 2)
	}

	for cur := i; cur < a.len-1; cur++ {
		a.elements[cur] = a.elements[cur+1]
	}

	a.elements[a.len-1] = nil
	a.len--
}

func (a *Array) resize(cap int) {
	a.cap = cap
	newElements := make([]interface{}, a.cap)
	for i := 0; i < a.len; i++ {
		newElements[i] = a.elements[i]
	}
	a.elements = newElements
}
