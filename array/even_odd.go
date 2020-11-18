package array

// EvenOdd reorders array a so that even integers appear first.
func EvenOdd(a []int) {
	even, odd := 0, len(a)-1

	for even < odd {
		if a[even]%2 == 0 {
			even++
			continue
		}
		a[even], a[odd] = a[odd], a[even]
		odd--
	}
}
