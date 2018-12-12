// Package doublyCycleLinkedList implements doubly cycle linked-list with
// a head node.
package doublyCycleLinkedList

import (
	"errors"
	"fmt"
)

type node struct {
	value interface{}
	next  *node
	prev  *node
	list  *list
}

type list struct {
	head *node
}

// New creates a list.
func New() *list {
	l := new(list)
	l.head = new(node)
	l.head.value = nil
	l.head.next = l.head
	l.head.prev = l.head
	l.head.list = l
	return l
}

// NewNode creates a node by v.
func NewNode(v interface{}) *node {
	n := new(node)
	n.value = v
	n.next = nil
	n.prev = nil
	n.list = nil
	return n
}

// Value returns the node value.
func (n node) Value() interface{} {
	return n.value
}

// Next returns node the next node.
func (n node) Next() *node {
	return n.next
}

// Prev returns node the prev node.
func (n node) Prev() *node {
	return n.prev
}

// List returns the list belong to node.
func (n node) List() *list {
	return n.list
}

// Head returns the head node belong to the list.
func (l list) Head() *node {
	return l.head
}

// Insert insert n after mark.
func (l *list) Insert(n, mark *node) error {
	if mark.list != l {
		return errors.New("node not in the list")
	}

	n.next = mark.next
	n.prev = mark
	mark.next.prev = n
	mark.next = n
	n.list = l

	return nil
}

// Delete deletes n from list.
func (l *list) Delete(n *node) error {
	if n.list != l {
		return errors.New("node not in the list")
	}

	if n == l.head {
		return errors.New("can't delete head node")
	}

	n.prev.next = n.next
	n.next.prev = n.prev

	n.value = nil
	n.next = nil
	n.prev = nil
	n.list = nil
	n = nil

	return nil
}

// Show returns list friendly express.
func (l list) Show() string {
	format := ""
	cur := l.head
	for cur.next != l.head {
		cur = cur.next
		format += fmt.Sprintf("%v", cur.value)
		if cur.next != l.head {
			format += fmt.Sprint("->")
		}
	}
	return format
}
