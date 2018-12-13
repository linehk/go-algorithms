package mergeSort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
	}
	for i, tt := range tests {
		MergeSort(tt.input)
		if got := isSorted(tt.input); !got {
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

func TestMerge(t *testing.T) {
	input := []int{3, 4, 0, 1}
	aux := make([]int, len(input))
	merge(input, aux, 0, 2, 4)
	t.Log(input)
}
