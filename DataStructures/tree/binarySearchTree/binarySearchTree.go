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
	b.root = nil
	return b
}

func (b bst) IsEmpty() bool {
	return b.Size() == 0
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

func (b bst) Contains(key int) bool {
	_, err := b.Get(key)
	if err != nil {
		return false
	}
	return true
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

	//for n != nil {
	//	if key < n.key {
	//	        n = n.left
	//	} else if key > n.key {
	//		n = n.right
	//	} else {
	//		return n, nil
	//	}
	//}
	//return nil, errors.New("key is not exist")
}

//func (b bst) Get(key int) (int, error) {
//	v, err := b.get(b.root, key)
//	if err != nil {
//		return 0, err
//	}
//	return v, nil
//}
//
//func (b bst) get(x *node, key int) (int, error) {
//	if x == nil {
//		return 0, errors.New("key is not in the symbol table")
//	}
//	if key < x.key {
//		return b.get(x.left, key)
//	} else if key > x.key {
//		return b.get(x.right, key)
//	} else {
//		return x.value, nil
//	}
//}

func (b *bst) Put(key, value int) {
	b.root = b.put(b.root, key, value)
}

func (b *bst) put(n *node, key, value int) *node {
	if n == nil {
		return &node{key, value, nil, nil, 1}
	}

	if key < n.key {
		n.left = b.put(n.left, key, value)
	} else if key > n.key {
		n.right = b.put(n.right, key, value)
	} else {
		n.value = value
	}

	n.size = 1 + b.size(n.left) + b.size(n.right)
	return n
}

func (b *bst) DeleteMin() error {
	if b.IsEmpty() {
		return errors.New("symbol table is empty")
	}
	b.root = b.deleteMin(b.root)
	return nil
}

func (b *bst) deleteMin(n *node) *node {
	if n.left == nil {
		return n.right
	}
	n.left = b.deleteMin(n.left)
	n.size = b.size(n.left) + b.size(n.right) + 1
	return n
}

func (b *bst) DeleteMax() error {
	if b.IsEmpty() {
		return errors.New("symbol table is empty")
	}
	b.root = b.deleteMax(b.root)
	return nil
}

func (b *bst) deleteMax(n *node) *node {
	if n.right == nil {
		return n.left
	}
	n.right = b.deleteMax(n.right)
	n.size = b.size(n.left) + b.size(n.right) + 1
	return n
}

func (b *bst) Delete(key int) {
	b.root = b.delete(b.root, key)
}

func (b *bst) delete(n *node, key int) *node {
	if n == nil {
		return nil
	}

	if key < n.key {
		n.left = b.delete(n.left, key)
	} else if key > n.key {
		n.right = b.delete(n.right, key)
	} else {
		if n.right == nil {
			return n.left
		} else if n.left == nil {
			return n.right
		} else {
			t := n
			n = b.min(t.right)
			n.right = b.deleteMin(t.right)
			n.left = t.left
		}
	}

	n.size = 1 + b.size(n.left) + b.size(n.right)
	return n
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

// Floor returns a value less or equal key but is the largest
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

func (b bst) floor(n *node, key int) *node {
	if n == nil {
		return nil
	}

	if key == n.key {
		return n
	}

	if key < n.key {
		return b.floor(n.left, key)
	}
	t := b.floor(n.right, key)
	if t != nil {
		return t
	} else {
		return n
	}
}

// Ceiling returns a value greater or equal key but is the smallest
func (b bst) Ceiling(key int) (int, error) {
	if b.IsEmpty() {
		return 0, errors.New("symbol table is empty")
	}

	n := b.ceiling(b.root, key)
	if n == nil {
		return 0, errors.New("no key is greater than the argument")
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
	// if key < n.key {
	// 	t := b.ceiling(n.left, key)
	// 	if t != nil {
	// 		return t
	// 	} else {
	// 		return n
	// 	}
	// }
	// return b.ceiling(n.right, key)
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

func (b bst) Height() int {
	return b.height(b.root)
}

func (b bst) height(n *node) int {
	if n == nil {
		return -1
	}

	maxHeight := b.height(n.left)
	rightHeight := b.height(n.right)
	if maxHeight < rightHeight {
		maxHeight = rightHeight
	}
	return 1 + maxHeight
}
