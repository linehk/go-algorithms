package shellsort

func ShellSort(nums []int) {
	h := 1
	for h < len(nums)/3 {
		h = 3*h + 1
	}

	for ; h >= 1; h /= 3 {
		for i := h; i < len(nums); i++ {
			for j := i; j >= h; j -= h {
				if nums[j-h] > nums[j] {
					nums[j-h], nums[j] = nums[j], nums[j-h]
				}
			}
		}
	}
}
