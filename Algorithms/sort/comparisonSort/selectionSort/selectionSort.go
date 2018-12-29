package selectionSort

func SelectionSort(elements []int) {
	l := len(elements)
	for i := 0; i < l; i++ {
		minIndex := i
		for j := i + 1; j < l; j++ {
			if elements[j] < elements[minIndex] {
				minIndex = j
			}
		}
		elements[i], elements[minIndex] = elements[minIndex], elements[i]
	}
}
