package binarySearch

func BinarySearch(elements []int, key int) bool {
	lo := 0
	hi := len(elements) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if key > elements[mid] {
			lo = mid + 1
		} else if key < elements[mid] {
			hi = mid - 1
		} else {
			return true
		}
	}
	return false
}

func binarySearch(elements []int, key int) int {
	lo := 0
	hi := len(elements) - 1
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if key > elements[mid] {
			lo = mid + 1
		} else if key < elements[mid] {
			hi = mid - 1
		} else if key == elements[mid] {
			return mid
		}
	}
	return lo
}

func binarySearchFirst(elements []int, key int) int {
	lo := 0
	hi := len(elements) - 1
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if key > elements[mid] {
			lo = mid + 1
		} else if key <= elements[mid] {
			hi = mid - 1
		}
	}
	return lo
}

func RecursiveBinarySearch(array []int, key int) bool {
	return rank(array, 0, len(array)-1, key)
}

func rank(array []int, lo, hi int, key int) bool {
	if lo > hi {
		return false
	}
	mid := lo + (hi-lo)/2
	if key > array[mid] {
		return rank(array, mid+1, hi, key)
	}
	if key < array[mid] {
		return rank(array, lo, mid-1, key)
	}
	if key == array[mid] {
		return true
	}
	return false
}
