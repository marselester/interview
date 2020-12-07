package sorting

// Mergesort sorts a in descending order using mergersort algorithm (stable, not in-place, extra space n).
// Running time is linearithmic (n log n).
func Mergesort(a []float64) {
	// copy is an auxiliary array allocated once and passed in every call to msort/merge.
	copy := make([]float64, len(a))
	left, right := 0, len(a)-1
	msort(a, copy, left, right)
}

/*
msort recursively sorts fragments (halves) defined by left and right indices.
It needs to descend to the bottom of the tree starting from the left branch
so it can merge pairs (2), then quartets (4), then octets (8)... on the way up.

	  4 9 8 3  4 7 7 0
	         /\
	4 9  8 3    4 7  7 0
	   /\          /\
	4 9  8 3    4 7  7 0

	1) merge([4 9]) => 9 4
	2) merge([8 3])
	3) merge([9 4  8 3]) => 9 8 4 3
	4) merge([4 7]) => 7 4
	5) merge([7 0])
	6) merge([7 4  7 0]) => 7 7 4 0
	7) merge([9 8 4 3  7 7 4 0]) => 9 8 7 7 4 4 3 0
*/
func msort(a, copy []float64, left, right int) {
	// There is nothing to sort/merge if the array fragment is just a single element.
	if left >= right {
		return
	}
	middle := left + (right-left)/2
	// Sort left half (left branch), then right half (right branch).
	msort(a, copy, left, middle)
	msort(a, copy, middle+1, right)

	// Skipping merge when array is ordered yields linear running time (optional optimization).
	// For example, in [4 3  2 1] the smallest element on left side (3) is greater than
	// the biggest element on the right side (2).
	// Since both sides are sorted, then merge call is redundant (concatenation of both sides is already fully sorted).
	if a[middle] >= a[middle+1] {
		return
	}
	// Merge of already sorted left and right halves.
	merge(a, copy, left, middle, right)
}

/*
merge merges left and right sorted halves into one sorted array.
It copies the fragment defined by left/right indices into the auxiliary array.
Then it makes a single pass through the copy and puts elements in order in the original array
by comparing i-th and j-th elements, then advancing the index whose sorting condition was met.
Index i points to left sorted half [left ... middle].
Index j points to right sorted half (middle ... right].

	9 4  8 3
	i    j

	1) 9 => i=1
	2) 8 => j=3
	3) 4 => i=2 (out of range)
	4) 3 => j=4 (out of range)
*/
func merge(a, copy []float64, left, middle, right int) {
	for k := left; k <= right; k++ {
		copy[k] = a[k]
	}

	i := left       // [left ... middle]
	j := middle + 1 // (middle ... right]
	for k := left; k <= right; k++ {
		switch {
		// Left half is exhausted because i is out of range,
		// so we have to take the rest of elements from right side by advancing j.
		case i > middle:
			a[k] = copy[j]
			j++
		// Right half is exhausted because j is out of range,
		// so we have to take the rest of elements from left side by advancing i.
		case j > right:
			a[k] = copy[i]
			i++
		// Item on left side (i-th) is greater than item on right side (j-th).
		// Take the i-th item and advance i.
		case copy[i] > copy[j]:
			a[k] = copy[i]
			i++
		// Item on right side (j-th) is greater or equal to the left side (i-th).
		// Take item from the right side and advance j.
		default:
			a[k] = copy[j]
			j++
		}
	}
}
