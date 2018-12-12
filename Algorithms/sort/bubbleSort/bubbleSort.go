package bubbleSort

func BubbleSort(elements []int) {
	l := len(elements)
	for i := 0; i < l; i++ {
		flag := true
		for j := i + 1; j < l; j++ {
			if elements[i] > elements[j] {
				elements[i], elements[j] = elements[j], elements[i]
				flag = false
			}
		}
		// break if didn't swap element
		if flag {
			break
		}
	}
}
