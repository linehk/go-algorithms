package shellsort

func ShellSort(elements []int) {
	l := len(elements)
	h := 1
	for h < l/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < l; i++ {
			for j := i; j >= h; j -= h {
				if elements[j] < elements[j-h] {
					elements[j], elements[j-h] = elements[j-h], elements[j]
				}
			}
		}
		h /= 3
	}
}
