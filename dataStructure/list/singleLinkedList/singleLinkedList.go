// Package singleLinkedList implements single linked-list with a head node.
package singleLinkedList

import (
	"errors"
	"fmt"
)

type node struct {
	value interface{}
	next  *node
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
	l.head.next = nil
	l.head.list = l
	return l
}

// NewNode creates a node by v.
func NewNode(v interface{}) *node {
	n := new(node)
	n.value = v
	n.next = nil
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

	pre := l.head
	for pre.next != n {
		pre = pre.next
	}

	pre.next = n.next

	n.value = nil
	n.next = nil
	n.list = nil
	n = nil

	return nil
}

// Show returns list friendly express.
func (l list) Show() string {
	format := ""
	cur := l.head
	for cur.next != nil {
		cur = cur.next
		format += fmt.Sprintf("%v", cur.value)
		if cur.next != nil {
			format += fmt.Sprint("->")
		}
	}
	return format
}
