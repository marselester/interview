package primitive

// SwapBitsNaive swaps i-th and j-th bits of x.
func SwapBitsNaive(x, i, j int) int {
	// Store bits found in i-th and j-th positiions.
	iv := x >> i & 1
	jv := x >> j & 1

	// Clear i-th, j-th bits and set the stored values.
	x = x&^(1<<i) | (jv << i)
	x = x&^(1<<j) | (iv << j)
	return x
}

// SwapBitsXor swaps i-th and j-th bits of x using XOR.
func SwapBitsXor(x, i, j int) int {
	if (x >> i & 1) != (x >> j & 1) {
		mask := 1<<i | 1<<j
		x = x ^ mask
	}
	return x
}
