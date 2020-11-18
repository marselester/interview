package primitive

// ParityNaive returns 1 if the number of ones in x is odd.
// Parity checks are used to detect single bit errors in storages and communication.
// Bits count is odd when its least significant bit is set.
func ParityNaive(x int) int {
	count := 0
	for x > 0 {
		if x&1 != 0 {
			count++
		}
		x = x >> 1
	}
	return count & 1
}

// ParityNaiveToggle computes parity of x without counting set bits.
// It starts with zero parity, checks if LSB is 1, and flips parity to one.
// Next time parity will flip to zero and so on.
func ParityNaiveToggle(x int) int {
	p := 0
	for x > 0 {
		if x&1 != 0 {
			p = p ^ 1
		}
		x = x >> 1
	}
	return p
}

// ParityKernighan computes parity of x in O(k) time where k is the number of bits set.
func ParityKernighan(x int) int {
	count := 0
	for x > 0 {
		x = x & (x - 1) // Clear the least significant bit set.
		count++
	}
	return count & 1
}

// uintSize is the size of a uint in bits (32 or 64),
// see https://golang.org/src/math/bits/bits.go.
const uintSize = 32 << (^uint(0) >> 32 & 1)

// ParityXor computes parity of x in O(log n) time where n is a word size.
// XOR of a group of bits is its parity, e.g., 11010111 parity equals to (1101 XOR 0111) parity.
func ParityXor(x int) int {
	for shift := uintSize / 2; shift > 0; shift = shift / 2 {
		x = x ^ x>>shift
	}
	return x & 1
}

const (
	// subwordSize is the size of uint subword in bits.
	// For example, 64-bit word has four 16-bit subwords when subwordSize is 16.
	// Note, bigger the size, larger the parityTable, e.g., 16-bit subword requires array size of 65536 (2^16).
	subwordSize = 16
	// subwordMask is a mask of a subword, e.g., 0b11 if size is 2.
	subwordMask = 0xFFFF
)

var parityTable [1 << subwordSize]bool

// init fills the lookup-table with parities of each array index.
func init() {
	for i := range parityTable {
		parityTable[i] = ParityKernighan(i) == 1
	}
}

// ParityLookupTable computes parity using a lookup-table ([subword]=parity).
// It starts from the lowest subword and shifts a higher subword to the right on each iteration.
func ParityLookupTable(x int) int {
	p := 0
	for shift := 0; shift < uintSize; shift += subwordSize {
		if parityTable[x>>shift&subwordMask] {
			p = p ^ 1
		}
	}
	return p
}
