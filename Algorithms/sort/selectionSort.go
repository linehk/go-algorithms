package sort

func SelectionSort(array []int) {
	for i := 0; i < len(array); i++ {
		minIndex := i
		for j := i; j < len(array); j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}
		Swap(array, i, minIndex)
	}
}
