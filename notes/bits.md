# Bitwise operation

References:

- [Bit Hacking with Go](https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827)
- [Bit Twiddling Hacks](https://graphics.stanford.edu/~seander/bithacks.html)

Expressions:

- `mask = 1 << 2` mask `0000 0100`
- `a | mask` set 3rd bit
- `a & mask != 0` test 3rd bit
- `a &^ mask` erase 3rd bit
- `a &^ mask | b<<2` modify 3rd bit with value `b` (zero or one)
- `a & (a - 1)` erase lowest set bit
- `a & -a` erase all but lowest set bit
- `a | (a - 1)` set bits to 1 after lowest set bit
- `a & (b - 1)` compute `a mod b` where `b` is 1, 2, 4, ...
- `a & (a - 1) == 0` is `a` power of two (1, 2, 4, ...)?

Use the AND operator to clear the last 4 least significant bits (LSB).

```go
var x uint8 = 0xAC      // 10101100
var y uint8 = 0xF0      // 11110000
fmt.Printf("%b", x & y) // 10100000
```

A number is odd when its least significant bit is set (equal 1).
Apply AND operator to integer `1`.
If the result is `1`, then the original number is odd.

```go
var x uint8 = 3
fmt.Printf("is %v odd: %v", x & 0x1 == 1)
```

Use the OR operator to set the 3rd, 7th, and 8th bit to 1 (from least to most significant bits).

```go
var x uint8 = 0x1       // 00000001
var y uint8 = 0xC4      // 11000100
fmt.Printf("%b", x | y) // 11000101
```

Bit masking as config.

```go
const (
	UPPER = 1 << iota // 00000001
	LOWER             // 00000010
	CAP               // 00000100
	REV               // 00001000
)
var conf byte = LOWER | REV | CAP    // 00001110
if (conf & LOWER) != 0 {
	fmt.Printf("%08b", conf & LOWER) // 00000010
}
```

Bitwise complement unary operator (NOT) `^`, i.e., inversion of bits.
In Python it's `~`, but in Go `~` is used as NOT and XOR depending on context (one variable or two).

```go
var a byte = 0xF
fmt.Printf("%08b\n", a)  // 00001111
fmt.Printf("%08b\n", ^a) // 11110000
```

Use XOR to toggle bits.

```go
var a uint8
for i := 0; i < 4; i++ {
	fmt.Printf("%08b\n", a)
	a = a ^ 1
}
// 00000000
// 00000001
// 00000000
// 00000001
```

Two integers `a`, `b` have:

- the same signs when `(a ^ b) ≥ 0`
- opposite signs when `(a ^ b) < 0`

```go
var (
	a int8 = 3            // 00000011
	b int8 = -10          // -0001010
)
fmt.Printf("%08b", a ^ b) // -0001011
```

The AND NOT operator clears the bits in `a` where `b`'s bits are set to 1.
Clear the last four LSBs in variable `a` from `1010 1011` to `1010 0000`.

```go
var a byte = 0xAB            // 10101011
var b byte = 0xF             // 00001111
fmt.Printf("%08b\n", a &^ b) // 10100000
```

Shift bits to the left n times.

```go
var a byte = 3               // 00000011
fmt.Printf("%08b\n", a << 1) // 00000110
fmt.Printf("%08b\n", a << 2) // 00001100
fmt.Printf("%08b\n", a << 3) // 00011000
```

The left and right shift operators are the multiplication and division
where each shift position represents a power of two.

```go
a := 200
for a > 0 {
	fmt.Printf("%08[1]b %[1]d\n", a)
	a = a >> 1
}
// 11001000 200
// 01100100 100
// 00110010 50
// 00011001 25
// 00001100 12
// 00000110 6
// 00000011 3
// 00000001 1
```

Erase the lowest set bit `a & (a - 1)`.

```go
var a uint8 = 3                 // 00000011
fmt.Printf("%08b", a & (a - 1)) // 00000010
```

Erase all bits except the lowest set bit `a & -a` or `a & ^(a - 1)`.

```go
var a uint8 = 22           // 00010110
fmt.Printf("%08b", a & -a) // 00000010
```

Set bits to 1 after lowest set bit `a | (a - 1)`.

```go
var a uint8 = 0b01010000        // 01010000
fmt.Printf("%08b", a | (a - 1)) // 01011111
```

Compute `a mod b` where `b` is one of 1, 2, 4, 8, 16, 32, ...

```go
a, b := 77, 64
fmt.Println(a & (b - 1)) // 13
```
