package searching

// BinarySearch looks up a search key in the sorted array a and
// returns the index of the key if it is present in the array,
// -1 otherwise.
func BinarySearch(a []int, key int) int {
	left, right := 0, len(a)-1
	middle := left + (right-left)/2

	for left <= right {
		switch {
		case key == a[middle]:
			return middle
		case key > a[middle]:
			left = middle + 1
		case key < a[middle]:
			right = middle - 1
		}

		middle = left + (right-left)/2
	}

	return -1
}
