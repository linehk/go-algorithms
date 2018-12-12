package sort

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	array := []int{2, 3, 1, 5, 8, 4, 9, 7, 0, 6, 10}
	SelectionSort(array)
	t.Log(array)
}
