package primitive

// CountBitsNaive returns the number of bits that are set to 1 in a non-negative integer x.
// It tests bits one at a time starting from least significat bit.
// The x is shifted right on each iteration so the loop ends when x is zero.
// This allows to work with an integer without knowing its size upfront (32 or 64 bits).
// Time complexity is O(n) where n is a number of bits.
// See https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetNaive.
func CountBitsNaive(x int) (count int) {
	for x > 0 {
		if x&1 != 0 {
			count++
		}
		x = x >> 1
	}
	return count
}

// CountBitsKernighan counts bits Brian Kernighan's way by erasing the lowest set bit in a loop.
// Brian Kernighan's method goes through as many iterations as there are set bits.
// So if we have a 32-bit word with only the high bit set,
// then it will only go once through the loop.
// See https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetKernighan.
func CountBitsKernighan(x int) (count int) {
	for x > 0 {
		count++
		x = x & (x - 1)
	}
	return count
}
