package quicksort

import (
	"github.com/linehk/GoAlgorithms/Algorithms/sort/utils"
)

var cutoff = 20

func QuickSort(elements []int) {
	sort(elements, 0, len(elements)-1)
}

func sort(elements []int, lo, hi int) {
	if lo >= hi {
		return
	}
	// n := hi - lo + 1
	// if n <= cutoff {
	// 	insertionSort(elements, lo, hi+1)
	// 	return
	// }
	p := partition(elements, lo, hi)
	sort(elements, lo, p-1)
	sort(elements, p+1, hi)
}

func partition(elements []int, lo, hi int) int {
	i, j := lo, hi
	pivotIndex, pivot := getPivot(elements, lo, hi)
	for i < j {
		for i < j && elements[j] > pivot {
			j--
		}
		for i < j && elements[i] <= pivot {
			i++
		}
		if i < j {
			utils.Swap(elements, i, j)
		}
	}
	utils.Swap(elements, i, pivotIndex)
	return i
}

func insertionSort(elements []int, lo, hi int) {
	for i := lo + 1; i < hi; i++ {
		for j := i; j > lo; j-- {
			if elements[j] < elements[j-1] {
				elements[j], elements[j-1] = elements[j-1], elements[j]
			}
		}
	}
}

func getPivot(elements []int, lo, hi int) (int, int) {
	return lo, elements[lo]
}

func getPivotMid(elements []int, lo, hi int) (int, int) {
	mid := lo + (hi-lo)/2
	return mid, elements[mid]
}

func getPivotMedian(elements []int, lo, hi int) (int, int) {
	mid := lo + (hi-lo)/2
	if elements[lo] < elements[mid] {
		elements[lo], elements[mid] = elements[mid], elements[lo]
	}
	if elements[hi] < elements[lo] {
		elements[hi], elements[lo] = elements[lo], elements[hi]
		if elements[lo] < elements[mid] {
			elements[lo], elements[mid] = elements[mid], elements[lo]
		}
	}
	return mid, elements[mid]
}

func Quick3way(elements []int) {
	sort3way(elements, 0, len(elements)-1)
}

func sort3way(elements []int, lo, hi int) {
	if lo >= hi {
		return
	}
	lt := lo
	gt := hi
	pivot := elements[lo]
	i := lo + 1
	for i <= gt {
		if elements[i] < pivot {
			elements[lt], elements[i] = elements[i], elements[lt]
			lt++
			i++
		} else if elements[i] > pivot {
			elements[i], elements[gt] = elements[gt], elements[i]
			gt--
		} else {
			i++
		}
	}
	sort3way(elements, lo, lt-1)
	sort3way(elements, gt+1, hi)
}
