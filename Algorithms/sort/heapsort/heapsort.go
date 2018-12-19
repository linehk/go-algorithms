package heapsort

func HeapSort(elements []int) {
	l := len(elements)
	for k := l / 2; k >= 1; k-- {
		sink(elements, k, l)
	}
	for l > 1 {
		exch(elements, 1, l)
		l--
		sink(elements, 1, l)
	}
}

func sink(elements []int, k, l int) {
	for 2*k <= l {
		j := 2 * k
		if j < l && less(elements, j, j+1) {
			j++
		}
		if !less(elements, k, j) {
			break
		}
		exch(elements, k, j)
		k = j
	}
}

func less(elements []int, i, j int) bool {
	if elements[i-1] < elements[j-1] {
		return true
	}
	return false
}

func exch(elements []int, i, j int) {
	elements[i-1], elements[j-1] = elements[j-1], elements[i-1]
}
