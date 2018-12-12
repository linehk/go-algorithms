package sort

func BubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		flag := false
		for j := len(array) - 1; j > i; j-- {
			if array[j] < array[j-1] {
				Swap(array, j, j-1)
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
