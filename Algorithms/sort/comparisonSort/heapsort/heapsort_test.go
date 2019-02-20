package heapsort

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
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
		HeapSort(tt.input)
		if got := isSorted(tt.input); !got {
			t.Errorf("%v. got %v, want %v", i, got, true)
		}
	}

}

func isSorted(elements []int) bool {
	l := len(elements)
	for i := 0; i < l-1; i++ {
		if elements[i] > elements[i+1] {
			return false
		}
	}
	return true
}
