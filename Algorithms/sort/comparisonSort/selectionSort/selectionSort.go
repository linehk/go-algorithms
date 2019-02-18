package selectionSort

import (
	"github.com/linehk/GoAlgorithms/Algorithms/sort/utils"
)

func SelectionSort(elements []int) {
	for i := 0; i < len(elements); i++ {
		minIndex := i
		for j := i + 1; j < len(elements); j++ {
			if elements[j] < elements[minIndex] {
				minIndex = j
			}
		}
		utils.Swap(elements, i, minIndex)
	}
}
