package mergeSort

func MergeSort(elements []int) {
	l := len(elements)
	aux := make([]int, l)
	sort(elements, aux, 0, l-1)
}

func sort(elements, aux []int, lo, hi int) {
	if lo >= hi {
		return
	}
	// mid := int(uint(lo+hi) >> 1)
	// mid is left part last element index
	mid := lo + (hi-lo)/2
	sort(elements, aux, lo, mid)
	sort(elements, aux, mid+1, hi)
	merge(elements, aux, lo, mid, hi)
}

func merge(elements, aux []int, lo, mid, hi int) {
	for k := lo; k <= hi; k++ {
		aux[k] = elements[k]
	}

	i := lo
	j := mid + 1 // right part begin index

	for k := lo; k <= hi; k++ {
		if i > mid {
			elements[k] = aux[j]
			j++
		} else if j > hi {
			elements[k] = aux[i]
			i++
		} else if aux[i] < aux[j] {
			elements[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			elements[k] = aux[j]
			j++
		}
	}
}

func IndexMergeSort(elements []int) []int {
	l := len(elements)
	index := make([]int, l)
	for i := 0; i < l; i++ {
		index[i] = i
	}
	aux := make([]int, l)
	indexSort(elements, index, aux, 0, l-1)
	return index
}

func indexSort(elements, index, aux []int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := lo + (hi-lo)/2
	indexSort(elements, index, aux, lo, mid)
	indexSort(elements, index, aux, mid+1, hi)
	indexMerge(elements, index, aux, lo, mid, hi)
}

func indexMerge(elements, index, aux []int, lo, mid, hi int) {
	for k := lo; k <= hi; k++ {
		aux[k] = index[k]
	}

	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		if i > mid {
			index[k] = aux[j]
			j++
		} else if j > hi {
			index[k] = aux[i]
			i++
		} else if elements[aux[i]] < elements[aux[j]] {
			index[k] = aux[i]
			i++
		} else if elements[aux[j]] < elements[aux[i]] {
			index[k] = aux[j]
			j++
		}
	}
}

func BottomUpMergeSort(elements []int) {
	l := len(elements)
	aux := make([]int, l)
	for length := 1; length < l; length *= 2 {
		for lo := 0; lo < l-length; lo += 2 * length {
			mid := lo + length - 1
			hi := min(lo+2*length-1, l-1)
			merge(elements, aux, lo, mid, hi)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
