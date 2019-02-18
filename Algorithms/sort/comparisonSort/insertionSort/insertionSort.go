package insertionSort

import (
	"github.com/linehk/GoAlgorithms/Algorithms/sort/utils"
)

func InsertionSort(elements []int) {
	for i := 1; i < len(elements); i++ {
		for j := i; j > 0; j-- {
			if elements[j] < elements[j-1] {
				utils.Swap(elements, j, j-1)
			}
		}
	}
}

func InsertSortX(elements []int) {
	l := len(elements)

	// put smallest element in position to serve as sentinel
	exchanges := 0
	for i := l - 1; i > 0; i-- {
		if elements[i] < elements[i-1] {
			elements[i], elements[i-1] = elements[i-1], elements[i]
			exchanges++
		}
	}
	if exchanges == 0 {
		return
	}

	// insertion sort with half-exchanges
	for i := 2; i < l; i++ {
		v := elements[i]
		j := i
		for ; v < elements[j-1]; j-- {
			elements[j] = elements[j-1]
		}
		elements[j] = v
	}
}

func BinaryInsertionSort(elements []int) {
	l := len(elements)
	for i := 1; i < l; i++ {
		// lo = should insert index
		v := elements[i]
		lo := 0
		hi := i
		for lo < hi {
			mid := lo + (hi-lo)/2
			if v < elements[mid] {
				hi = mid
			} else {
				lo = mid + 1
			}
		}

		// move to right
		for j := i; j > lo; j-- {
			elements[j] = elements[j-1]
		}
		elements[lo] = v
	}
}
