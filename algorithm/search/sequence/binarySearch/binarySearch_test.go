package binarySearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		input []int
		key   int
		want  bool
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 100, false},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 200, false},
	}
	for i, tt := range tests {
		if got := BinarySearch(tt.input, tt.key); got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}

func TestBinarySearchIndex(t *testing.T) {
	tests := []struct {
		input []int
		key   int
		want  int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 100, 10},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -100, 0},
	}
	for i, tt := range tests {
		if got := BinarySearchIndex(tt.input, tt.key); got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}

func TestBinarySearchFirst(t *testing.T) {
	tests := []struct {
		input []int
		key   int
		want  int
	}{
		{[]int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2}, 1, 0},
		{[]int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2}, 2, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 100, 10},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -100, 0},
	}
	for i, tt := range tests {
		if got := BinarySearchFirst(tt.input, tt.key); got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}

func TestRecursiveBinarySearch(t *testing.T) {
	tests := []struct {
		input []int
		key   int
		want  bool
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 100, false},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 200, false},
	}
	for i, tt := range tests {
		if got := BinarySearch(tt.input, tt.key); got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}
