package bubbleSort

func BubbleSort(nums []int) {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		isSorted := false
		for j := 0; j < l-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isSorted = true
			}
		}
		// break if didn't swap element
		if !isSorted {
			break
		}
	}
}
