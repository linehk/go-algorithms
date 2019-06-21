package heapsort

func HeapSort(nums []int) {
	l := len(nums)
	for k := l / 2; k >= 1; k-- {
		sink(nums, k, l)
	}
	for l > 1 {
		swap(nums, 1, l)
		l--
		sink(nums, 1, l)
	}
}

func sink(nums []int, k, l int) {
	for 2*k <= l {
		j := 2 * k
		if j < l && less(nums, j, j+1) {
			j++
		}
		if !less(nums, k, j) {
			break
		}
		swap(nums, k, j)
		k = j
	}
}

func less(nums []int, i, j int) bool {
	return nums[i-1] < nums[j-1]
}

func swap(nums []int, i, j int) {
	nums[i-1], nums[j-1] = nums[j-1], nums[i-1]
}
