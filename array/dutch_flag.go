package array

// DutchFlag reorders array a so that integers less than a[p] (pivot) appear first,
// followed by integers equal to pivot, and greater than pivot.
func DutchFlag(a []int, p int) {
	pivot := a[p]
	lt, gt := 0, len(a)-1

	// Move all ints smaller than pivot at the beginning of the array.
	for i := 0; i < len(a); i++ {
		if a[i] < pivot {
			a[i], a[lt] = a[lt], a[i]
			lt++
		}
	}

	// Move all ints larger than pivot at the end of the array.
	for i := len(a) - 1; i > 0; i-- {
		if a[i] > pivot {
			a[i], a[gt] = a[gt], a[i]
			gt--
		}
	}
}

// DutchFlagOnePass reorders array a as DutchFlag but in a single pass.
func DutchFlagOnePass(a []int, p int) {
	pivot := a[p]
	lt, eq, gt := 0, 0, len(a)

	for eq < gt {
		switch {
		case a[eq] < pivot:
			a[lt], a[eq] = a[eq], a[lt]
			lt++
			eq++
		case a[eq] == pivot:
			eq++
		case a[eq] > pivot:
			gt--
			a[eq], a[gt] = a[gt], a[eq]
		}
	}
}
