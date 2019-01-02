// Package array implements DataStructure array can dynamically grow and shark.
package array

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
// Assume:
// 0 <= i < a.len
func (a array) Get(i int) interface{} {
	return a.elements[i]
}

// Set assign v to array index i.
// Assume:
// 0 <= i < a.len
func (a *array) Set(i int, v interface{}) {
	a.elements[i] = v
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
// Assume:
// 0 <= i < a.len
func (a *array) Insert(i int, v interface{}) {
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
// Assume:
// 0 <= i <= a.len
func (a *array) Delete(i int) {
	if a.len-1 < a.cap/4 {
		a.resize(a.cap / 2)
	}

	for cur := i; cur < a.len-1; cur++ {
		a.elements[cur] = a.elements[cur+1]
	}

	a.elements[a.len-1] = nil
	a.len--
}

func (a *array) resize(cap int) {
	a.cap = cap
	newElements := make([]interface{}, a.cap)
	for i := 0; i < a.len; i++ {
		newElements[i] = a.elements[i]
	}
	a.elements = newElements
}
