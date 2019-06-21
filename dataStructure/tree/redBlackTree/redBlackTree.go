package redBlackTree

import (
	"errors"
)

const RED = true
const BLACK = false

type bst struct {
	root *node
}

type node struct {
	key   int
	value int
	left  *node
	right *node
	color bool
	size  int
}

func New() *bst {
	b := new(bst)
	b.root = nil
	return b
}

func (b bst) isRed(n *node) bool {
	if n == nil {
		return false
	}
	// return n.color == RED
	return n.color
}

func (b bst) Size() int {
	return b.size(b.root)
}

func (b bst) size(n *node) int {
	if n == nil {
		return 0
	} else {
		return n.size
	}
}

func (b bst) IsEmpty() bool {
	return b.Size() == 0
}

func (b bst) Get(key int) (int, error) {
	if b.IsEmpty() {
		return 0, errors.New("symbol table is empty")
	}

	n, err := b.get(b.root, key)
	if err != nil {
		return 0, err
	}
	return n.value, nil
}

func (b bst) get(n *node, key int) (*node, error) {
	if n == nil {
		return nil, errors.New("key is not exist")
	}

	if key < n.key {
		return b.get(n.left, key)
	} else if key > n.key {
		return b.get(n.right, key)
	} else {
		return n, nil
	}
}

func (b bst) Contains(key int) bool {
	_, err := b.Get(key)
	return err == nil
}

func (b *bst) Put(key int, value int) {
	b.root = b.put(b.root, key, value)
	b.root.color = BLACK
}

func (b *bst) put(n *node, key int, value int) *node {
	if n == nil {
		return &node{key, value, nil, nil, RED, 1}
	}

	if key < n.key {
		n.left = b.put(n.left, key, value)
	} else if key > n.key {
		n.right = b.put(n.right, key, value)
	} else {
		n.value = value
	}

	// right is red
	if b.isRed(n.right) && !b.isRed(n.left) {
		n = b.rotateLeft(n)
	}

	// double left is red
	if b.isRed(n.left) && b.isRed(n.left.left) {
		n = b.rotateRight(n)
	}

	// left right both red
	if b.isRed(n.left) && b.isRed(n.right) {
		b.flipColors(n)
	}

	n.size = 1 + b.size(n.left) + b.size(n.right)
	return n
}

func (b *bst) DeleteMin() error {
	if b.IsEmpty() {
		return errors.New("symbol table is empty")
	}

	if !b.isRed(b.root.left) && !b.isRed(b.root.right) {
		b.root.color = RED
	}

	b.root = b.deleteMin(b.root)

	if !b.IsEmpty() {
		b.root.color = BLACK
	}
	return nil
}

func (b *bst) deleteMin(n *node) *node {
	if n.left == nil {
		return nil
	}

	if !b.isRed(n.left) && !b.isRed(n.left.left) {
		n = b.moveRedLeft(n)
	}

	n.left = b.deleteMin(n.left)

	return b.balance(n)
}

func (b *bst) DeleteMax() error {
	if b.IsEmpty() {
		return errors.New("symbol table is empty")
	}

	if !b.isRed(b.root.left) && !b.isRed(b.root.right) {
		b.root.color = RED
	}

	b.root = b.deleteMax(b.root)

	if !b.IsEmpty() {
		b.root.color = BLACK
	}
	return nil
}

func (b *bst) deleteMax(n *node) *node {
	if b.isRed(n.left) {
		n = b.rotateRight(n)
	}

	if n.right == nil {
		return nil
	}

	if !b.isRed(n.right) && !b.isRed(n.right.left) {
		n = b.moveRedRight(n)
	}

	n.right = b.deleteMax(n.right)
	return b.balance(n)
}

func (b *bst) Delete(key int) {
	if !b.Contains(key) {
		return
	}

	if !b.isRed(b.root.left) && !b.isRed(b.root.right) {
		b.root.color = RED
	}

	b.root = b.delete(b.root, key)

	if !b.IsEmpty() {
		b.root.color = BLACK
	}
}

func (b *bst) delete(n *node, key int) *node {
	if key < n.key {
		if !b.isRed(n.left) && !b.isRed(n.left.left) {
			n = b.moveRedLeft(n)
		}
		n.left = b.delete(n.left, key)
	} else {
		if b.isRed(n.left) {
			n = b.moveRedRight(n)
		}
		if key == n.key && n.right == nil {
			return nil
		}
		if !b.isRed(n.right) && !b.isRed(n.right.left) {
			n = b.moveRedRight(n)
		}
		if key == n.key {
			t := b.min(n.right)
			n.key = t.key
			n.value = t.value
			n.right = b.deleteMin(n.right)
		} else {
			n.right = b.delete(n.right, key)
		}
	}
	return b.balance(n)
}

func (b *bst) rotateRight(n *node) *node {
	t := n.left
	n.left = t.right
	t.right = n

	t.color = t.right.color
	t.right.color = RED

	t.size = n.size
	n.size = 1 + b.size(n.left) + b.size(n.right)
	return t
}

func (b *bst) rotateLeft(n *node) *node {
	t := n.right
	n.right = t.left
	t.left = n

	t.color = t.left.color
	t.left.color = RED

	t.size = n.size
	n.size = 1 + b.size(n.left) + b.size(n.right)
	return t
}

func (b bst) flipColors(n *node) {
	n.color = !n.color
	n.left.color = !n.left.color
	n.right.color = !n.right.color
}

func (b *bst) moveRedLeft(n *node) *node {
	b.flipColors(n)
	if b.isRed(n.right.left) {
		n.right = b.rotateRight(n.right)
		n = b.rotateLeft(n)
		b.flipColors(n)
	}
	return n
}

func (b *bst) moveRedRight(n *node) *node {
	b.flipColors(n)
	if b.isRed(n.left.left) {
		n = b.rotateRight(n)
		b.flipColors(n)
	}
	return n
}

func (b *bst) balance(n *node) *node {
	if b.isRed(n.right) {
		n = b.rotateLeft(n)
	}

	if b.isRed(n.left) && b.isRed(n.left.left) {
		n = b.rotateRight(n)
	}

	if b.isRed(n.left) && b.isRed(n.right) {
		b.flipColors(n)
	}

	n.size = b.size(n.left) + b.size(n.right) + 1
	return n
}

func (b bst) Height() int {
	return b.height(b.root)
}

func (b bst) height(n *node) int {
	if n == nil {
		return -1
	}

	maxHeight := b.height(n.left)
	rightHeight := b.height(n.right)
	if rightHeight > maxHeight {
		maxHeight = rightHeight
	}
	return 1 + maxHeight
}

func (b bst) Min() (int, error) {
	if b.IsEmpty() {
		return 0, errors.New("symbol table is empty")
	}
	return b.min(b.root).key, nil
}

func (b bst) min(n *node) *node {
	if n.left == nil {
		return n
	} else {
		return b.min(n.left)
	}
}

func (b bst) Max() (int, error) {
	if b.IsEmpty() {
		return 0, errors.New("symbol table is empty")
	}
	return b.max(b.root).key, nil
}

func (b bst) max(n *node) *node {
	if n.right == nil {
		return n
	} else {
		return b.max(n.right)
	}
}

func (b bst) Floor(key int) (int, error) {
	if b.IsEmpty() {
		return 0, errors.New("symbol table is empty")
	}

	n := b.floor(b.root, key)
	if n == nil {
		return 0, errors.New("no key is less than the argument")
	} else {
		return n.key, nil
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

func (b bst) Ceiling(key int) (int, error) {
	if b.IsEmpty() {
		return 0, errors.New("symbol table is empty")
	}

	n := b.ceiling(b.root, key)
	if n == nil {
		return 0, errors.New("key is not in the symbol table")
	} else {
		return n.key, nil
	}
}

func (b bst) ceiling(n *node, key int) *node {
	if n == nil {
		return nil
	}

	if key == n.key {
		return n
	}

	if key > n.key {
		return b.ceiling(n.right, key)
	}
	t := b.ceiling(n.left, key)
	if t != nil {
		return t
	} else {
		return n
	}
}

func (b bst) Select(k int) (int, error) {
	if k < 0 || k >= b.Size() {
		return 0, errors.New("illegal index")
	}
	return b.selectNode(b.root, k).key, nil
}

func (b bst) selectNode(n *node, k int) *node {
	if n == nil {
		return nil
	}

	t := b.size(n.left)
	if t > k {
		return b.selectNode(n.left, k)
	} else if t < k {
		return b.selectNode(n.right, k-t-1)
	} else {
		return n
	}
}

func (b bst) Rank(key int) int {
	return b.rank(b.root, key)
}

func (b bst) rank(n *node, key int) int {
	if n == nil {
		return 0
	}

	if key < n.key {
		return b.rank(n.left, key)
	} else if key > n.key {
		return 1 + b.size(n.left) + b.rank(n.right, key)
	} else {
		return b.size(n.left)
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
