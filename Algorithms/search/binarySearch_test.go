package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	keys := []int{1, 10}
	for _, key := range keys {
		if !BinarySearch(array, key) {
			t.Errorf("%d is fail", key)
		}
	}
}

func TestRecursiveBinarySearch(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	keys := []int{1, 10}
	for _, key := range keys {
		if !RecursiveBinarySearch(array, key) {
			t.Errorf("%d is fail", key)
		}
	}
}
