package AVLTree

import (
	"errors"
)

type bst struct {
	root *node
}

type node struct {
	key    int
	value  int
	height int
	size   int
	left   *node
	right  *node
}

func newNode(key, value, size, height int) *node {
	n := new(node)
	n.key = key
	n.value = value
	n.size = size
	n.height = height
	n.left = nil
	n.right = nil
	return n
}

func New() *bst {
	return new(bst)
}

func (b bst) IsEmpty() bool {
	return b.root == nil
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

// func (b bst) Get(key int) (int, error) {
//         x, err := b.get(b.root, key)
//         if err != nil {
//                 return 0, err
//         } else {
//                 return x.value, nil
//         }
// }

// func (b bst) get(x *node, key int) (*node, error) {
//         if x == nil {
//                 return nil, errors.New("key is not in the symbol table")
//         }
//
//         if key < x.key {
//                 return b.get(x.left, key)
//         } else if key > x.key {
//                 return b.get(x.right, key)
//         } else {
//                 return x, nil
//         }
// }

func (b bst) Get(key int) (int, error) {
	v, err := b.get(b.root, key)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func (b bst) get(x *node, key int) (int, error) {
	if x == nil {
		return 0, errors.New("key is not in the symbol table")
	}
	if key < x.key {
		return b.get(x.left, key)
	} else if key > x.key {
		return b.get(x.right, key)
	} else {
		return x.value, nil
	}
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
}

func (b *bst) put(x *node, key int, value int) *node {
	if x == nil {
		return newNode(key, value, 0, 1)
	}
	if key < x.key {
		x.left = b.put(x.left, key, value)
	} else if key > x.key {
		x.right = b.put(x.right, key, value)
	} else {
		x.value = value
		return x
	}

	x.size = 1 + b.size(x.left) + b.size(x.right)
	x.height = 1 + max(b.height(x.left), b.height(x.right))

	return b.balance(x)
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

	x.size = 1 + b.size(x.left) + b.size(x.right)
	x.height = 1 + max(b.height(x.left), b.height(x.right))

	return b.balance(x)
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

	x.size = 1 + b.size(x.left) + b.size(x.right)
	x.height = 1 + max(b.height(x.left), b.height(x.right))

	return b.balance(x)
}

func (b *bst) Delete(key int) {
	if !b.Contains(key) {
		return
	}

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
		if x.left == nil {
			return x.right
		} else if x.right == nil {
			return x.left
		} else {
			y := x
			x = b.min(y.right)
			x.right = b.deleteMin(y.right)
			x.left = y.left
		}
	}

	x.size = 1 + b.size(x.left) + b.size(x.right)
	x.height = 1 + max(b.height(x.left), b.height(x.right))

	return b.balance(x)
}

func (b *bst) rotateRight(x *node) *node {
	y := x.left
	x.left = y.right
	y.right = x

	y.size = x.size
	x.size = 1 + b.size(x.left) + b.size(x.right)

	y.height = 1 + max(b.height(y.left), b.height(y.right))
	x.height = 1 + max(b.height(x.left), b.height(x.right))

	return y
}

func (b *bst) rotateLeft(x *node) *node {
	y := x.right
	x.right = y.left
	y.left = x

	y.size = x.size
	x.size = 1 + b.size(x.left) + b.size(x.right)

	y.height = 1 + max(b.height(y.left), b.height(y.right))
	x.height = 1 + max(b.height(x.left), b.height(x.right))
	return y
}

func (b *bst) balance(x *node) *node {
	if b.balanceFactor(x) < -1 {
		if b.balanceFactor(x.right) > 0 {
			x.right = b.rotateRight(x.right)
		}
		x = b.rotateLeft(x)
	} else if b.balanceFactor(x) > 1 {
		if b.balanceFactor(x.left) < 0 {
			x.left = b.rotateLeft(x.left)
		}
		x = b.rotateRight(x)
	}
	return x
}

func (b *bst) balanceFactor(x *node) int {
	return b.height(x.left) - b.height(x.right)
}

func (b bst) Height() int {
	return b.height(b.root)
}

func (b bst) height(x *node) int {
	if x == nil {
		return -1
	} else {
		return x.height
	}
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
		return 0, errors.New("not need")
	} else {
		return x.value, nil
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
	y := b.floor(x.right, key)
	if y != nil {
		return y
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
		return 0, errors.New("not need")
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
	y := b.ceiling(x.left, key)
	if y != nil {
		return y
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

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
