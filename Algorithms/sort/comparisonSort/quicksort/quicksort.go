package quicksort

var cutoff = 20

func QuickSort(nums []int) {
	sort(nums, 0, len(nums)-1)
}

func sort(nums []int, lo, hi int) {
	if hi <= lo {
		return
	}
	p := partition(nums, lo, hi)
	sort(nums, lo, p-1)
	sort(nums, p+1, hi)
}

func partition(nums []int, lo, hi int) int {
	i, j := lo, hi
	pivotIndex, pivot := lo, nums[lo]
	for i < j {
		for i < j && nums[j] > pivot {
			j--
		}
		for i < j && nums[i] <= pivot {
			i++
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[pivotIndex] = nums[pivotIndex], nums[i]
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
