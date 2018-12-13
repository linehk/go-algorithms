package mergeSort

func MergeSort(elements []int) {
}

func sort() {
}

func merge(elements []int, aux []int, lo, mid, hi int) {
	for k := lo; k <= hi; k++ {
		aux[k] = elements[k]
	}

	i := lo
	j := mid + 1

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
