package shellsort

func ShellSort(nums []int) {
	h := 1
	for h < len(nums)/3 {
		h = h*3 + 1
	}

	for h >= 1 {
		for i := 0; i < len(nums); i++ {
			for j := i; j >= h; j -= h {
				if nums[j-h] > nums[j] {
					nums[j-h], nums[j] = nums[j], nums[j-h]
				}
			}
		}
		h /= 3
	}
}
