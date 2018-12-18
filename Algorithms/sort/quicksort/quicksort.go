package quicksort

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

func insertionSort(elements []int, lo, hi int) {
	for i := lo + 1; i < hi; i++ {
		for j := i; j > lo; j-- {
			if elements[j] < elements[j-1] {
				elements[j], elements[j-1] = elements[j-1], elements[j]
			}
		}
	}
}

func partition(elements []int, lo, hi int) int {
	i := lo
	j := hi + 1

	// pivotIndex, pivot := getPivotMid(elements, lo, hi)
	// pivotIndex, pivot := getPivotMedian(elements, lo, hi)
	pivotIndex, pivot := getPivot(elements, lo, hi)
	elements[lo], elements[pivotIndex] = elements[pivotIndex], elements[lo]
	pivotIndex = lo

	for {
		for {
			i++
			if elements[i] < pivot {
				if i == hi {
					break
				}
			} else {
				break
			}
		}

		for {
			j--
			if elements[j] > pivot {
				if j == lo {
					break
				}
			} else {
				break
			}
		}

		if i >= j {
			break
		}
		elements[i], elements[j] = elements[j], elements[i]
	}
	elements[pivotIndex], elements[j] = elements[j], elements[pivotIndex]
	return j
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
