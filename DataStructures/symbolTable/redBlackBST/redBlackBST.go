package redBlackBST

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
	return new(bst)
}

func (b bst) isRed(x *node) bool {
	if x == nil {
		return false
	}
	return x.color == RED
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

func (b bst) IsEmpty() bool {
	return b.root == nil
}

func (b bst) Get(key int) (int, error) {
	v, err := b.get(b.root, key)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func (b bst) get(x *node, key int) (int, error) {
	for x != nil {
		if key < x.key {
			x = x.left
		} else if key > x.key {
			x = x.right
		} else {
			return x.value, nil
		}
	}
	return 0, errors.New("key is not in the symbol table")
}

func (b bst) Contains(key int) bool {
	_, err := b.Get(key)
	if err != nil {
		return false
	}
	return true
}

func (b *bst) Put(key int, value int) {
	b.root = b.put(b.root, key, value)
	b.root.color = BLACK
}

func (b *bst) put(h *node, key int, value int) *node {
	if h == nil {
		return &node{key, value, nil, nil, RED, 1}
	}
	if key < h.key {
		h.left = b.put(h.left, key, value)
	} else if key > h.key {
		h.right = b.put(h.right, key, value)
	} else {
		h.value = value
	}
	// right is red
	if b.isRed(h.right) && !b.isRed(h.left) {
		h = b.rotateLeft(h)
	}
	// left double red
	if b.isRed(h.left) && b.isRed(h.left.left) {
		h = b.rotateRight(h)
	}
	// left right both red
	if b.isRed(h.left) && b.isRed(h.right) {
		b.flipColors(h)
	}
	h.size = 1 + b.size(h.left) + b.size(h.right)
	return h
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

func (b *bst) deleteMin(h *node) *node {
	if h.left == nil {
		return nil
	}
	if !b.isRed(h.left) && !b.isRed(h.left.left) {
		h = b.moveRedLeft(h)
	}
	h.left = b.deleteMin(h.left)
	return b.balance(h)
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

func (b *bst) deleteMax(h *node) *node {
	if b.isRed(h.left) {
		h = b.rotateRight(h)
	}
	if h.right == nil {
		return nil
	}
	if !b.isRed(h.right) && !b.isRed(h.right.left) {
		h = b.moveRedRight(h)
	}
	h.right = b.deleteMax(h.right)
	return b.balance(h)
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

func (b *bst) delete(h *node, key int) *node {
	if key < h.key {
		if !b.isRed(h.left) && !b.isRed(h.left.left) {
			h = b.moveRedLeft(h)
		}
		h.left = b.delete(h.left, key)
	} else {
		if b.isRed(h.left) {
			h = b.moveRedRight(h)
		}
		if key == h.key && h.right == nil {
			return nil
		}
		if !b.isRed(h.right) && !b.isRed(h.right.left) {
			h = b.moveRedRight(h)
		}
		if key == h.key {
			x := b.min(h.right)
			h.key = x.key
			h.value = x.value
			// h.key = min(h.right).key
			// h.value = get(h.right, min(h.right).key)
			h.right = b.deleteMin(h.right)
		} else {
			h.right = b.delete(h.right, key)
		}
	}
	return b.balance(h)
}

func (b *bst) rotateRight(h *node) *node {
	x := h.left
	h.left = x.right
	x.right = h

	x.color = x.right.color
	x.right.color = RED

	x.size = h.size
	h.size = 1 + b.size(h.left) + b.size(h.right)
	return x
}

func (b *bst) rotateLeft(h *node) *node {
	x := h.right
	h.right = x.left
	x.left = h

	x.color = x.left.color
	x.left.color = RED

	x.size = h.size
	h.size = 1 + b.size(h.left) + b.size(h.right)
	return x
}

func (b bst) flipColors(h *node) {
	h.color = !h.color
	h.left.color = !h.left.color
	h.right.color = !h.right.color
}

func (b *bst) moveRedLeft(h *node) *node {
	b.flipColors(h)
	if b.isRed(h.right.left) {
		h.right = b.rotateRight(h.right)
		h = b.rotateLeft(h)
		b.flipColors(h)
	}
	return h
}

func (b *bst) moveRedRight(h *node) *node {
	b.flipColors(h)
	if b.isRed(h.left.left) {
		h = b.rotateRight(h)
		b.flipColors(h)
	}
	return h
}

func (b *bst) balance(h *node) *node {
	if b.isRed(h.right) {
		h = b.rotateLeft(h)
	}
	if b.isRed(h.left) && b.isRed(h.left.left) {
		h = b.rotateRight(h)
	}
	if b.isRed(h.left) && b.isRed(h.right) {
		b.flipColors(h)
	}

	h.size = b.size(h.left) + b.size(h.right) + 1
	return h
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
	if key > x.key {
		return b.ceiling(x.right, key)
	}
	t := b.ceiling(x.left, key)
	if t != nil {
		return t
	} else {
		return x
	}
}

func (b bst) Select(k int) (int, error) {
	if k < 0 || k >= b.Size() {
		return 0, errors.New("illegal index k")
	}
	return b.selectNode(b.root, k).key, nil
}

func (b bst) selectNode(x *node, k int) *node {
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
	return b.rank(b.root, key)
}

func (b bst) rank(x *node, key int) int {
	if x == nil {
		return 0
	}
	if key < x.key {
		return b.rank(x.left, key)
	} else if key > x.key {
		return 1 + b.size(x.left) + b.rank(x.right, key)
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
