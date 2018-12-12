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
