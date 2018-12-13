package bubbleSort

func BubbleSort(elements []int) {
	l := len(elements)
	for i := 0; i < l-1; i++ {
		changed := false
		for j := 0; j < l-1-i; j++ {
			if elements[j] > elements[j+1] {
				elements[j], elements[j+1] = elements[j+1], elements[j]
				changed = true
			}
		}
		// break if didn't swap element
		if !changed {
			break
		}
	}
}
