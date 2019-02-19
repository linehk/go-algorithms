package shellsort

func ShellSort(nums []int) {
	l := len(nums)
	h := 1
	for h < l/3 {
		h = 3*h + 1 // 1, 4, 13, 40, ...
	}
	for h >= 1 {
		for i := h; i < l; i++ {
			for j := i; j >= h; j -= h {
				if nums[j-h] > nums[j] {
					nums[j-h], nums[j] = nums[j], nums[j-h]
				}
			}
		}
		h /= 3
	}
}
