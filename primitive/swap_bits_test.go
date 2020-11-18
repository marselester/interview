package primitive

import "testing"

func TestSwapBitsNaive(t *testing.T) {
	tt := map[string]struct {
		x    int
		i    int
		j    int
		want int
	}{
		"0000 0001": {
			x:    1,
			i:    0,
			j:    7,
			want: 128, // 0b1000_0000
		},
		"0000 0011": {
			x:    3,
			i:    0,
			j:    1,
			want: 3,
		},
		"0000 1001": {
			x:    9,
			i:    3,
			j:    3,
			want: 9,
		},
		"0001 0111 1101 0110 1110 0011": {
			x:    1562339,
			i:    3,
			j:    20,
			want: 513771, // 0b0000_0111_1101_0110_1110_1011
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := SwapBitsNaive(tc.x, tc.i, tc.j)
			if tc.want != got {
				t.Errorf("expected %b got %b", tc.want, got)
			}
		})
	}
}

func TestSwapBitsXor(t *testing.T) {
	tt := map[string]struct {
		x    int
		i    int
		j    int
		want int
	}{
		"0000 0001": {
			x:    1,
			i:    0,
			j:    7,
			want: 128, // 0b1000_0000
		},
		"0000 0011": {
			x:    3,
			i:    0,
			j:    1,
			want: 3,
		},
		"0000 1001": {
			x:    9,
			i:    3,
			j:    3,
			want: 9,
		},
		"0001 0111 1101 0110 1110 0011": {
			x:    1562339,
			i:    3,
			j:    20,
			want: 513771, // 0b0000_0111_1101_0110_1110_1011
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := SwapBitsXor(tc.x, tc.i, tc.j)
			if tc.want != got {
				t.Errorf("expected %b got %b", tc.want, got)
			}
		})
	}
}

func BenchmarkSwapBitsNaive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SwapBitsNaive(1562339, 3, 20)
	}
}

func BenchmarkSwapBitsXor(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SwapBitsXor(1562339, 3, 20)
	}
}
