package bubbleSort

import (
	"github.com/linehk/GoAlgorithms/Algorithms/sort/utils"
)

func BubbleSort(elements []int) {
	for i := 0; i < len(elements)-1; i++ {
		changed := false
		for j := 0; j < len(elements)-1-i; j++ {
			if elements[j] > elements[j+1] {
				utils.Swap(elements, j, j+1)
				changed = true
			}
		}
		// break if didn't swap element
		if !changed {
			break
		}
	}
}
