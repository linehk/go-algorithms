package mergeSort

func MergeSort(elements []int) {
	l := len(elements)
	aux := make([]int, l)
	sort(elements, aux, 0, l-1)
}

func sort(elements []int, aux []int, lo, hi int) {
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

func merge(elements []int, aux []int, lo, mid, hi int) {
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
