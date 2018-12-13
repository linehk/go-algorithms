package insertionSort

func InsertionSort(elements []int) {
	l := len(elements)
	for i := 1; i < l; i++ {
		for j := i; j > 0; j-- {
			if elements[j] < elements[j-1] {
				elements[j], elements[j-1] = elements[j-1], elements[j]
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
		}
		exchanges++
	}
	if exchanges == 0 {
		return
	}

	// insertion sort with half-exchanges
	for i := 2; i < l; i++ {
		j := i
		v := elements[i]
		for ; v < elements[j-1]; j-- {
			elements[j] = elements[j-1]
		}
		elements[j] = v
	}
}
