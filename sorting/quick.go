package sorting

// Quicksort sorts a in descending order using quicksort algorithm
// (not stable, in-place, extra space log n due to recursion).
// Running time is n log n on average case (random order, no duplicates),
// and n² in worst (already sorted or has many duplicates).
// Note, 3-way quicksort (partition by <, =, >) works well with duplicates (Dutch flag problem).
//
// Quicksort is the fastest general-purpose sort when space is tight:
// it has only a few instructions in its inner loop,
// it does well with cache memories because it most often references data sequentially.
func Quicksort(a []float64) {
	left, right := 0, len(a)-1
	qsort(a, left, right)
}

// qsort recursively sorts a fragment defined by left and right indices.
func qsort(a []float64, left, right int) {
	// There is nothing to sort if the array fragment is just a single element.
	if left >= right {
		return
	}
	// Partition the array fragment so that elements >= pivot were moved to the left side of the pivot.
	// The pivot's index is returned to recursively sort left and right sides.
	// The pivot element itself doesn't need to move anymore.
	pivot := qpartition(a, left, right)
	qsort(a, left, pivot-1)
	qsort(a, pivot+1, right)
}

/*
qpartition uses rightmost element as a pivot to separate the elements into two groups.
Index j examines elements from start of the fragment till the pivot (excluding).
Elements between left and i should be >= pivot.
Elements between i+1 and j should be < pivot.

	  9 4 8 3 7
	i j       pivot

	0) 9 >= 7, swap(a[0], a[0]) => 9 4 8 3 7
	1) 4 < 7
	2) 8 >= 7, swap(a[1], a[2]) => 9 8 4 3 7
	3) 3 < 7
	4) move pivot, swap(a[2], a[4]) => 9 8 7 3 4

Another way to pick a pivot is to randomly choose 3 elements and take a median.
This should decrease probability of picking a bad pivot.
*/
func qpartition(a []float64, left, right int) int {
	pivot := right

	i := left - 1
	j := left
	for ; j < pivot; j++ {
		if a[j] >= a[pivot] {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	// Put the pivot between two groups.
	i++
	a[pivot], a[i] = a[i], a[pivot]

	pivot = i
	return pivot
}
