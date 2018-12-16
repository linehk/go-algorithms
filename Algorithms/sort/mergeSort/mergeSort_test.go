package mergeSort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{[]int{0}},
		{[]int{0, 1}},
		{[]int{1, 0}},
		{[]int{2, 1, 0}},
		{[]int{3, 2, 1, 0}},
		{[]int{4, 3, 2, 1, 0}},
	}
	for i, tt := range tests {
		MergeSort(tt.input)
		if got := isSorted(tt.input); !got {
			t.Errorf("%v. got %v, want %v", i, got, true)
		}
	}
}

func TestIndexMergeSort(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{[]int{0}},
		{[]int{0, 1}},
		{[]int{1, 0}},
		{[]int{2, 1, 0}},
		{[]int{3, 2, 1, 0}},
		{[]int{4, 3, 2, 1, 0}},
	}
	for i, tt := range tests {
		index := IndexMergeSort(tt.input)
		elements := make([]int, len(tt.input))
		for k, v := range index {
			elements[k] = tt.input[v]
		}
		if got := isSorted(elements); !got {
			t.Errorf("%v. got %v, want %v", i, got, true)
		}
	}
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
