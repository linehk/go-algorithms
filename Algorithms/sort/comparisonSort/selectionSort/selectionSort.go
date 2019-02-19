package selectionSort

func SelectionSort(nums []int) {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
}
