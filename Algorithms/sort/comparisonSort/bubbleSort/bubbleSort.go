package bubbleSort

func BubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		changed := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				changed = true
			}
		}
		if !changed {
			break
		}
	}
}
