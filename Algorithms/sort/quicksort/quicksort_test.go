package quicksort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{[]int{0}},
		{[]int{0, 1}},
		{[]int{1, 0}},
		{[]int{2, 1, 0}},
		{[]int{3, 2, 1, 0}},
		{[]int{4, 3, 2, 1, 0}},
		{[]int{4, 3, 2, 1, 0, 4, 3, 2, 1, 0}},
	}
	for i, tt := range tests {
		QuickSort(tt.input)
		if got := isSorted(tt.input); !got {
			t.Errorf("%v. got %v, want %v", i, got, true)
		}
	}
}

func TestQuick3way(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{[]int{0}},
		{[]int{0, 1}},
		{[]int{1, 0}},
		{[]int{2, 1, 0}},
		{[]int{3, 2, 1, 0}},
		{[]int{4, 3, 2, 1, 0}},
		{[]int{4, 3, 2, 1, 0, 4, 3, 2, 1, 0}},
	}
	for i, tt := range tests {
		Quick3way(tt.input)
		if got := isSorted(tt.input); !got {
			t.Errorf("%v. got %v, want %v", i, got, true)
		}
	}
}

func TestPartition(t *testing.T) {
	input := []int{0, 1}
	j := partition(input, 0, len(input)-1)
	t.Log(input)
	t.Log(j)
}

func isSorted(elements []int) bool {
	l := len(elements)
	for i := 1; i < l; i++ {
		if elements[i] < elements[i-1] {
			return false
		}
	}
	return true
}
