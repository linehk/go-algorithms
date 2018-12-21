package binarySearchTree

import (
	"errors"
)

type bst struct {
	root *node
}

type node struct {
	key   int
	value int
	left  *node
	right *node
	size  int
}

func New() *bst {
	b := new(bst)
	return b
}

func (b bst) IsEmpty() bool {
	return b.Size() == 0
}

func (b bst) Size() int {
	return b.size(b.root)
}

func (b bst) size(x *node) int {
	if x == nil {
		return 0
	} else {
		return x.size
	}
}

func (b bst) Contains(key int) bool {
	_, err := b.Get(key)
	if err != nil {
		return false
	}
	return true
}

func (b bst) Get(key int) (int, error) {
	v, err := b.get(b.root, key)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func (b bst) get(x *node, key int) (int, error) {
	if x == nil {
		return 0, errors.New("node is nil")
	}
	if key < x.key {
		return b.get(x.left, key)
	} else if key > x.key {
		return b.get(x.right, key)
	} else {
		return x.value, nil
	}
}

func (b *bst) Put(key, value int) {
	b.root = b.put(b.root, key, value)
}

func (b *bst) put(x *node, key, value int) *node {
	if x == nil {
		return &node{key, value, nil, nil, 1}
	}
	if key < x.key {
		x.left = b.put(x.left, key, value)
	} else if key > x.key {
		x.right = b.put(x.right, key, value)
	} else {
		x.value = value
	}
	x.size = 1 + b.size(x.left) + b.size(x.right)
	return x
}

func (b *bst) DeleteMin() error {
	if b.IsEmpty() {
		return errors.New("symbol table is empty")
	}
	b.root = b.deleteMin(b.root)
	return nil
}

func (b *bst) deleteMin(x *node) *node {
	if x.left == nil {
		return x.right
	}
	x.left = b.deleteMin(x.left)
	x.size = b.size(x.left) + b.size(x.right) + 1
	return x
}

func (b *bst) DeleteMax() error {
	if b.IsEmpty() {
		return errors.New("symbol table is empty")
	}
	b.root = b.deleteMax(b.root)
	return nil
}

func (b *bst) deleteMax(x *node) *node {
	if x.right == nil {
		return x.left
	}
	x.right = b.deleteMax(x.right)
	x.size = b.size(x.left) + b.size(x.right) + 1
	return x
}

func (b *bst) Delete(key int) {
	b.root = b.delete(b.root, key)
}

func (b *bst) delete(x *node, key int) *node {
	if x == nil {
		return nil
	}
	if key < x.key {
		x.left = b.delete(x.left, key)
	} else if key > x.key {
		x.right = b.delete(x.right, key)
	} else {
		if x.right == nil {
			return x.left
		}
		if x.left == nil {
			return x.right
		}
		t := x
		x = b.min(t.right)
		x.right = b.deleteMin(t.right)
		x.left = t.left
	}
	x.size = b.size(x.left) + b.size(x.right) + 1
	return x
}

func (b bst) Min() (int, error) {
	if b.IsEmpty() {
		return 0, errors.New("symbol table is empty")
	}
	return b.min(b.root).key, nil
}

func (b bst) min(x *node) *node {
	if x.left == nil {
		return x
	} else {
		return b.min(x.left)
	}
}

func (b bst) Max() (int, error) {
	if b.IsEmpty() {
		return 0, errors.New("symbol table is empty")
	}
	return b.max(b.root).key, nil
}

func (b bst) max(x *node) *node {
	if x.right == nil {
		return x
	} else {
		return b.max(x.right)
	}
}

// Floor returns a value less or equal key but is the largest
func (b bst) Floor(key int) (int, error) {
	if b.IsEmpty() {
		return 0, errors.New("symbol table is empty")
	}
	x := b.floor(b.root, key)
	if x == nil {
		return 0, errors.New("key is not in the symbol table")
	} else {
		return x.key, nil
	}
}

func (b bst) floor(x *node, key int) *node {
	if x == nil {
		return nil
	}
	if key == x.key {
		return x
	}
	if key < x.key {
		return b.floor(x.left, key)
	}
	t := b.floor(x.right, key)
	if t != nil {
		return t
	} else {
		return x
	}
}

// Ceiling returns a value greater or equal key but is the smallest
func (b bst) Ceiling(key int) (int, error) {
	if b.IsEmpty() {
		return 0, errors.New("symbol table is empty")
	}
	x := b.ceiling(b.root, key)
	if x == nil {
		return 0, errors.New("key is not in the symbol table")
	} else {
		return x.key, nil
	}
}

func (b bst) ceiling(x *node, key int) *node {
	if x == nil {
		return nil
	}
	if key == x.key {
		return x
	}
	if key < x.key {
		t := b.ceiling(x.left, key)
		if t != nil {
			return t
		} else {
			return x
		}
	}
	return b.ceiling(x.right, key)
}

func (b bst) Select(k int) (int, error) {
	if k < 0 || k >= b.Size() {
		return 0, errors.New("illegal index k")
	}
	return b.selectNode(b.root, k).key, nil
}

func (b bst) selectNode(x *node, k int) *node {
	if x == nil {
		return nil
	}
	t := b.size(x.left)
	if t > k {
		return b.selectNode(x.left, k)
	} else if t < k {
		return b.selectNode(x.right, k-t-1)
	} else {
		return x
	}
}

func (b bst) Rank(key int) int {
	return b.rank(key, b.root)
}

func (b bst) rank(key int, x *node) int {
	if x == nil {
		return 0
	}
	if key < x.key {
		return b.rank(key, x.left)
	} else if key > x.key {
		return 1 + b.size(x.left) + b.rank(key, x.right)
	} else {
		return b.size(x.left)
	}
}

func (b bst) SizeByRange(lo, hi int) int {
	if lo > hi {
		return 0
	}
	if b.Contains(hi) {
		return b.Rank(hi) - b.Rank(lo) + 1
	} else {
		return b.Rank(hi) - b.Rank(lo)
	}
}

func (b bst) Height() int {
	return b.height(b.root)
}

func (b bst) height(x *node) int {
	if x == nil {
		return -1
	}
	maxHeight := b.height(x.left)
	rightHeight := b.height(x.right)
	if maxHeight < rightHeight {
		maxHeight = rightHeight
	}
	return 1 + maxHeight
}
